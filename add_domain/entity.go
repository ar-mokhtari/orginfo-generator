package adddomain

import (
	"github.com/ar-mokhtari/orginfo-generator/entity"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"os"
	temp "text/template"
)

func DomainEntityGenerator(domain entity.Domain) (err error) {
	//---------------------------
	//make directory for new domain (if exist return error)
	if mkdirErr := os.MkdirAll(env.MainPath+"entity/"+domain.SnakeName, os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	// Create a new template and parse the temp into it.
	for _, fileName := range []string{"model", "protocol"} {
		//tmpl, tempCreateErr := temp.ParseFiles("../entity/" + fileName + ".temp")
		tmpl, tempCreateErr := temp.ParseFiles(env.MainRepository + "add_domain/temps/" + env.Entity + "/" + fileName + ".temp")
		if tempCreateErr != nil {
			return tempCreateErr
		}
		//create entity file for new domain
		tempFile, tempFileCreateErr := os.Create(env.MainPath + "entity/" + domain.SnakeName + "/" + fileName + ".go")
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
	return err
}
