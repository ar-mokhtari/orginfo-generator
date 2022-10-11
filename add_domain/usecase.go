package adddomain

import (
	"errors"
	"fmt"
	"github.com/ar-mokhtari/orginfo-generator/entity"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"os"
	temp "text/template"
)

func DomainUsecaseGenerator(domain entity.Domain) (err error) {
	//---------------------------
	//make usecase directory if not exist
	if mkdirErr := os.MkdirAll(env.MainPath+"usecase/validation", os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	//make directory for new domain (if exist return error)
	if mkdirErr := os.MkdirAll(env.MainPath+env.Usecase+"/"+domain.SnakeName, os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	//create init file name in first index
	fileNames := []string{"init", "delete", "new", "edit", "find", "get"}
	// Create a new template and parse the temp into it.
	for _, fileName := range fileNames {
		//tmpl, tempCreateErr := temp.ParseFiles("../temps/usecase/services/" + fileName + ".temp")
		tmpl, tempCreateErr := temp.ParseFiles(env.MainRepository + "add_domain/temps/" + env.Usecase + "/services/" + fileName + ".temp")

		if tempCreateErr != nil {
			return tempCreateErr
		}
		//create file for new domain
		tempFile, tempFileCreateErr := os.Create(env.MainPath + env.Usecase + "/" + domain.SnakeName + "/" + fileName + ".go")
		if tempFileCreateErr != nil {
			return tempFileCreateErr
		}
		// Execute the template for each recipient.
		tmpExeErr := tmpl.Execute(tempFile, domain)
		if tmpExeErr != nil {
			return tmpExeErr
		}
		tempFile.Close()
	}
	//---------------------------
	if _, otherValidationErr := os.Stat(env.MainPath + "usecase/validation/others/other_validation.go"); errors.Is(otherValidationErr, os.ErrNotExist) {
		otherValidationFile, createOtherValidationErr := os.Create(env.MainPath + "usecase/validation/others/other_validation.go")
		if createOtherValidationErr != nil {
			return createOtherValidationErr
		}
		fmt.Fprintf(otherValidationFile, "%s: ", "package validationOthers\n\nimport (\n\t\"errors\"\n\t\"fmt\"\n\tvalidation \"github.com/go-ozzo/ozzo-validation\"\n\t\"strconv\"\n)\n\nfunc CheckNationalID(inputChar string) validation.RuleFunc {\n\treturn func(value interface{}) error {\n\t\tvar (\n\t\t\tcounter, uintChar uint\n\t\t\tcontrolNumber, _  = strconv.Atoi(string(inputChar[9]))\n\t\t)\n\t\tfor index := 1; index < len(inputChar)-1; index++ {\n\t\t\tinternalUintChar, _ := strconv.Atoi(string(inputChar[index]))\n\t\t\tuintChar = uint(internalUintChar)\n\t\t\tcounter += uint(10-index) * uintChar\n\t\t}\n\t\tswitch remaining := counter % 11; remaining < 2 {\n\t\tcase true:\n\t\t\tif remaining == uint(controlNumber) {\n\t\t\t\treturn nil\n\t\t\t}\n\t\tcase false:\n\t\t\tif 11-remaining == uint(controlNumber) {\n\t\t\t\treturn nil\n\t\t\t}\n\t\t}\n\t\terr := errors.New(\"wrong format\")\n\t\treturn fmt.Errorf(\"%v\", err)\n\t}\n}\n")
		defer otherValidationFile.Close()
	}
	//make directory for new "validation" domain (if exist return error)
	if mkdirErr := os.MkdirAll(env.MainPath+env.Usecase+"/validation/"+domain.SnakeName, os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	//tempProtocol, protocolErr := temp.ParseFiles("../temps/usecase/validation/temp/protocol.temp")
	tempProtocol, protocolErr := temp.ParseFiles(env.MainRepository + "add_domain/temps/" + env.Usecase + "/validation/temp/protocol.temp")
	if protocolErr != nil {
		return protocolErr
	}
	//tempValidator, validatorErr := temp.ParseFiles("../temps/usecase/validation/temp/temp_validator.temp")
	tempValidator, validatorErr := temp.ParseFiles(env.MainRepository + "add_domain/temps/" + env.Usecase + "/validation/temp/temp_validator.temp")
	if validatorErr != nil {
		return validatorErr
	}
	fileProtocol, fileProtocolErr := os.Create(env.MainPath + env.Usecase + "/validation/" + domain.SnakeName + "/protocol.go")
	if fileProtocolErr != nil {
		return fileProtocolErr
	}
	fileValidator, fileValidatorErr := os.Create(env.MainPath + env.Usecase + "/validation/" + domain.SnakeName + "/temp_validator.go")
	if fileValidatorErr != nil {
		return fileValidatorErr
	}
	protocolExeErr := tempProtocol.Execute(fileProtocol, domain)
	if protocolExeErr != nil {
		return protocolExeErr
	}
	validatorExeErr := tempValidator.Execute(fileValidator, domain)
	if validatorExeErr != nil {
		return validatorExeErr
	}
	return err
}
