package adddomain

import (
	"bytes"
	"github.com/ar-mokhtari/orginfo-generator/entity"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"io/ioutil"
	"os"
	"strings"
	temp "text/template"
)

func DomainDeliveryGenerator(domain entity.Domain) (err error) {
	deliveryHttpV1Path := env.MainPath + env.Delivery + "/http/V1"
	//---------------------------
	//make directory for new domain (if exist return error)
	if mkdirErr := os.Mkdir(deliveryHttpV1Path+"/"+domain.SnakeName, os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	//create action(s)
	filesNames := []string{"init", "delete", "new", "edit", "find", "get"}
	// Create a new template and parse the temp into it.
	for _, fileName := range filesNames {
		tmpl, tempCreateErr := temp.ParseFiles(deliveryHttpV1Path + "/temp/" + fileName + ".temp")
		if tempCreateErr != nil {
			return tempCreateErr
		}
		//create file for new domain
		tempFile, tempFileCreateErr := os.Create(deliveryHttpV1Path + "/" + domain.SnakeName + "/" + fileName + ".go")
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
	//add route init in delivery/http/V1/init.go
	input, readErr := ioutil.ReadFile(deliveryHttpV1Path + "/init.go")
	if readErr != nil {
		return readErr
	}
	contentText := env.Delivery + domain.UpperName + ".Routs(echo)"
	contentImport := "import (\n\t" + env.Delivery + domain.UpperName + ` "` + "github.com/ar-mokhtari/orginfo-generator/delivery/http/V1/" + domain.SnakeName + `"`
	textExist := strings.Contains(string(input), contentText)
	if !textExist {
		output := bytes.Replace(input, []byte("}"), []byte("\t"+contentText+"\n}"), -1)
		output = bytes.Replace(output, []byte("import ("), []byte(contentImport), -1)
		if err = ioutil.WriteFile(deliveryHttpV1Path+"/init.go", output, 0666); err != nil {
			return err
		}
	}
	return err
}