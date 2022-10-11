package adddomain

import (
	"github.com/ar-mokhtari/orginfo-generator/entity"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"os"
	temp "text/template"
)

func DomainUsecaseGenerator(domain entity.Domain) (err error) {
	//---------------------------
	//make directory for new domain (if exist return error)
	if mkdirErr := os.Mkdir(env.MainPath+env.Usecase+"/"+domain.SnakeName, os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	//create init file name in first index
	fileNames := []string{"init", "delete", "new", "edit", "find", "get"}
	// Create a new template and parse the temp into it.
	for _, fileName := range fileNames {
		tmpl, tempCreateErr := temp.ParseFiles(env.MainPath + env.Usecase + "/temp/" + fileName + ".temp")
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
	//make directory for new "validation" domain (if exist return error)
	if mkdirErr := os.Mkdir(env.MainPath+env.Usecase+"/validation/"+domain.SnakeName, os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	tempProtocol, protocolErr := temp.ParseFiles(env.MainPath + env.Usecase + "/validation/temp/protocol.temp")
	if protocolErr != nil {
		return protocolErr
	}
	tempValidator, validatorErr := temp.ParseFiles(env.MainPath + env.Usecase + "/validation/temp/temp_validator.temp")
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
