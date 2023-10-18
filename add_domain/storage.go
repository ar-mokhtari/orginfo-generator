package adddomain

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ar-mokhtari/orginfo-generator/entity"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"io/ioutil"
	"os"
	"strings"
	temp "text/template"
)

func DomainStorageGenerator(domain entity.Domain) error {
	storagePath := env.MainPath + "adapter/storage/"
	//make storage directory if not exist
	if mkdirErr := os.MkdirAll(storagePath+"models", os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	//add setup.go in adapter/storage/ if not exist
	storeFiles := map[uint8][]string{
		0: {"setup", "package storage\n\nimport (\n\tstorage \"" + env.MainRepositoryPath + "/adapter/storage/models\"\n\t\"gorm.io/driver/mysql\"\n\t\"gorm.io/gorm\"\n)\n\ntype DBMS struct {\n\tdb *gorm.DB\n}\n\nfunc GormConnect(dsn string) (DBMS, error) {\n\tgorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})\n\tif err != nil {\n\t\treturn DBMS{}, err\n\t}\n\treturn DBMS{db: gorm}, nil\n}\n\nfunc GormAutoMigrate(db gorm.DB) error {\n\t// Migrate the schema\n\terr := db.AutoMigrate()\n\tif err != nil {\n\t\tpanic(\"Storage auto migrate has error\")\n\t}\n\treturn err\n}\n"},
		1: {"init", "package storage\n\nimport (\n\t\"fmt\"\n)\n\nvar DB DBMS\n\nfunc Init() {\n\tvar eErr error\n\t//I. Define a data source name (DSN)\n\tdsn := \"user:g#Y!298mKwz85@tcp(127.0.0.1:3306)/orginfo?charset=utf8mb4&parseTime=True&loc=Local\"\n\t//II. Try to connect to dsn address\n\tDB, eErr = GormConnect(dsn)\n\t//III. If gorm connect successfully then try to migrate database\n\tif eErr == nil {\n\t\terr := GormAutoMigrate(*DB.gorm)\n\t\tif err != nil {\n\t\t\tfmt.Errorf(\"some error .... %#v\", err)\n\t\t}\n\t}\n}\n"},
	}
	for _, part := range storeFiles {
		if _, existErr := os.Stat(storagePath + part[0] + ".go"); errors.Is(existErr, os.ErrNotExist) {
			setupFile, setupCreateErr := os.Create(storagePath + part[0] + ".go")
			if setupCreateErr != nil {
				return setupCreateErr
			}
			fmt.Fprintf(setupFile, "%s ", part[1])
			defer setupFile.Close()
		}
	}
	// create and exe temp for model
	//domainTemp, domainTempErr := temp.ParseFiles("../temps/storage/models/temp.temp")
	domainTemp, domainTempErr := temp.ParseFiles(env.MainRepository + "add_domain/temps/storage/models/temp.temp")
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
	//actionsTemp, actionsTempErr := temp.ParseFiles("../temps/storage/action_temp.temp")
	actionsTemp, actionsTempErr := temp.ParseFiles(env.MainRepository + "add_domain/temps/storage/action_temp.temp")
	if actionsTempErr != nil {
		return actionsTempErr
	}
	domainActionFile, actionErr := os.Create(storagePath + "action_" + domain.SnakeName + ".go")
	if actionErr != nil {
		return actionErr
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
