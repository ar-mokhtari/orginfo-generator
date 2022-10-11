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

func DomainStorageGenerator(domain entity.Domain) error {
	storagePath := env.MainPath + "adapter/storage/"
	// create and exe temp for model
	domainTemp, domainTempErr := temp.ParseFiles(storagePath + "models/temp.temp")
	if domainTempErr != nil {
		return domainTempErr
	}
	domainModelFile, fileModelErr := os.Create(storagePath + "models/" + domain.SnakeName + ".go")
	if fileModelErr != nil {
		return fileModelErr
	}
	tempExeErr := domainTemp.Execute(domainModelFile, domain)
	if tempExeErr != nil {
		return tempExeErr
	}
	// create and exe temp for actions
	actionsTemp, actionsTempErr := temp.ParseFiles(storagePath + "action_temp.temp")
	if actionsTempErr != nil {
		return actionsTempErr
	}
	domainActionFile, fileModelErr := os.Create(storagePath + "action_" + domain.SnakeName + ".go")
	if fileModelErr != nil {
		return fileModelErr
	}
	tempActionExeErr := actionsTemp.Execute(domainActionFile, domain)
	if tempActionExeErr != nil {
		return tempActionExeErr
	}
	//change storage setup file, add domain for migrate ...
	input, readErr := ioutil.ReadFile(storagePath + "setup.go")
	if readErr != nil {
		return readErr
	}
	contentText := "storage." + domain.UpperName + "{}"
	textExist := strings.Contains(string(input), contentText)
	if !textExist {
		output := bytes.Replace(input, []byte("db.AutoMigrate("), []byte("db.AutoMigrate("+contentText+", "), -1)
		if err := ioutil.WriteFile(storagePath+"/setup.go", output, 0666); err != nil {
			return err
		}
	}
	return nil
}
