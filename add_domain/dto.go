package adddomain

import (
	"github.com/ar-mokhtari/orginfo-generator/entity"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"os"
	temp "text/template"
)

func DomainDTOGenerator(domain entity.Domain) (err error) {
	//dto
	//---------------------------
	// Create a new template and parse the temp into it.
	tmpl, tempCreateErr := temp.ParseFiles(env.MainPath + "dto/temp/temp.temp")
	if tempCreateErr != nil {
		return tempCreateErr
	}
	//make directory for new domain (if exist return error)
	if mkdirErr := os.MkdirAll(env.MainPath+"dto/"+domain.SnakeName, os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	//create dto file for new domain
	tempFile, tempFileCreateErr := os.Create(env.MainPath + "dto/" + domain.SnakeName + "/" + domain.SnakeName + ".go")
	if tempFileCreateErr != nil {
		return tempFileCreateErr
	}
	// Execute the template for each recipient.
	tmpExeErr := tmpl.Execute(tempFile, domain)
	if tmpExeErr != nil {
		return tmpExeErr
	}
	tempFile.Close()
	return err
}
