package generator

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ar-mokhtari/orginfo-generator/config/cli/generator/entity"
	"github.com/ar-mokhtari/orginfo-generator/config/env"
	"os"
	"strings"
)

type (
	InputData struct {
		DomainName string         `json:"domain_name"`
		Fields     []entity.Field `json:"fields"`
	}
)

type newCommand struct {
	Fs                  *flag.FlagSet
	FlagDomainNameValue string
	FlagFieldsValue     string
	FlagFromFile        bool
}

func (nc *newCommand) Name() string { return nc.Fs.Name() }

func (nc *newCommand) Init(args []string) error {
	return nc.Fs.Parse(args)
}

func (nc *newCommand) Run() (err error) {
	switch nc.FlagFromFile {
	case false:
		//first, prepare domain_name
		res, prepareErr := preparename.PrepareDomainName(nc.FlagDomainNameValue)
		if prepareErr != nil {
			return prepareErr
		}
		//next prepare fields
		fields := strings.Split(nc.FlagFieldsValue, ",")
		var finalFields []entity.Field
		if len(fields) != 0 {
			for _, field := range fields {
				splitField := strings.Split(field, "-")
				sliceCount := len(splitField)
				if sliceCount != 3 {
					return fmt.Errorf("fields have to content 3 parts: name-type-jsonKey, your input have %d", sliceCount)
				}
				resFields, prepareFieldErr := preparename.PrepareFieldsName(splitField[0], splitField[1], splitField[2])
				if prepareFieldErr != nil {
					return prepareFieldErr
				}
				finalFields = append(finalFields, resFields)
			}
		}
		res.Fields = finalFields
		addInit(res)
	case true:
		allDomain, parseErr := parseDomainFromFile()
		if parseErr != nil {
			return parseErr
		}
		for _, domain := range allDomain {
			res, prepareErr := preparename.PrepareDomainName(domain.DomainName)
			if prepareErr != nil {
				return prepareErr
			}
			for index, field := range domain.Fields {
				domain.Fields[index], err = preparename.PrepareFieldsName(field.Name, field.Type, field.JsonName)
				if err != nil {
					return err
				}
			}
			res.Fields = domain.Fields
			addInit(res)
		}
	}
	return nil
}

type helpCommand struct {
	Fs            *flag.FlagSet
	FlagHelpValue string
}

func (nc *helpCommand) Name() string { return nc.Fs.Name() }

func (nc *helpCommand) Init(args []string) error {
	return nc.Fs.Parse(args)
}

func (nc *helpCommand) Run() (err error) {
	fmt.Println("----------- help menu -----------")
	fmt.Print("Domain generator\nto create new domain with CRUD operation and other prerequisite: run, build and execute this file:\ndomain generator and follow command(s) ... like run this from '/config/cli/generator/bin/':\n\n cd go/src/orginfo/config/cli/generator/bin/\npattern #1\ninput from command and flag one by one:\n./generator sub-command -domain_name=\"DOMAIN NAME\" -fields=\"field1-string-field1_1,field2-uint-field1_2,...\"\nrun this for add a new domain:\n\n./generator new -domain_name=\"DOMAIN_NAME\" -fields=\"codeType-uint-code_type,code-uint-code\"\nAlso, to remove a domain run this:\n\n./generator delete -domain_name=\"DOMAIN_NAME\"\npattern #2\nimport domain(s) from json type file ...\n./generator new -from_file\nremove domain(s) from json type file ...\n./generator delete -from_file\nthe file name must be: input.json. in this address: /config/cli/generator/\nyou can find default of file in this address: sample structure\n[\n  {\n    \"domain_name\": \"string\",\n    \"fields\": [\n      {\n        \"name\": \"string\",\n        \"type\": \"string\",\n        \"json_name\": \"string\"\n      }\n    ]\n  }\n]\n")
	fmt.Println("----------- help menu -----------")
	return nil
}

type deleteCommand struct {
	Fs           *flag.FlagSet
	FlagValue    string
	FlagFromFile bool
}

func (dc *deleteCommand) Name() string { return dc.Fs.Name() }

func (dc *deleteCommand) Init(args []string) error {
	return dc.Fs.Parse(args)
}

func (dc *deleteCommand) Run() (err error) {
	switch dc.FlagFromFile {
	case false:
		res, prepareErr := preparename.PrepareDomainName(dc.FlagValue)
		if prepareErr != nil {
			return prepareErr
		}
		deleteInit(res)
	case true:
		allDomain, parseErr := parseDomainFromFile()
		if parseErr != nil {
			return parseErr
		}
		for _, domain := range allDomain {
			res, prepareErr := preparename.PrepareDomainName(domain.DomainName)
			if prepareErr != nil {
				return prepareErr
			}
			deleteInit(res)
		}
	}
	return nil
}

func parseDomainFromFile() ([]InputData, error) {
	//load and read input.json and convert json to struct (entity.Domain)
	inputJson, jsonErr := os.ReadFile(env.MainPath + "config/cli/generator/" + "input.json")
	if jsonErr != nil {
		return nil, jsonErr
	}
	var allDomain []InputData
	unmarshalErr := json.Unmarshal(inputJson, &allDomain)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return allDomain, nil
}
