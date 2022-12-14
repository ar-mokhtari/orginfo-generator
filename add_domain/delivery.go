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

func DomainDeliveryGenerator(domain entity.Domain) (err error) {
	deliveryHttpV1Path := env.MainPath + env.Delivery + "/http/V1"
	deliveryInit := deliveryHttpV1Path + "/init.go"
	//make directory for new domain (if exist return error)
	if mkdirErr := os.MkdirAll(deliveryHttpV1Path+"/"+domain.SnakeName, os.ModePerm); mkdirErr != nil {
		return mkdirErr
	}
	//---------------------------
	//add delivery/http/init.go if not exist
	_, errHttpInit := os.Stat(env.MainPath + env.Delivery + "/http/init.go")
	if errors.Is(errHttpInit, os.ErrNotExist) {
		httpInit, httpInitCreateErr := os.Create(env.MainPath + env.Delivery + "/http/init.go")
		if httpInitCreateErr != nil {
			return httpInitCreateErr
		}
		fmt.Fprintf(httpInit, "%s", "package http\n\nimport (\n\t\""+env.MainRepositoryPath+"/delivery/http/V1\"\n\t\"github.com/labstack/echo/v4\"\n)\n\nfunc Init() {\n\t//initialise new Echo (web framework)\n\te := echo.New()\te.Use(middleware.CORSWithConfig(middleware.CORSConfig{\n\t\tAllowOrigins: []string{\"*\"/*, \"https://labstack.net\"*/},\n\t\tAllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},\n\t}))\n\t//handle incoming requests:\n\tV1.Init(e)\n\t//serve with port\n\te.Logger.Fatal(e.Start(\":1315\"))\n}\n")
		defer httpInit.Close()
	}
	//---------------------------
	//create action(s)
	filesNames := []string{"init", "delete", "new", "edit", "find", "get"}
	// Create a new template and parse the temp into it.
	for _, fileName := range filesNames {
		tmpl, tempCreateErr := temp.ParseFiles(env.MainRepository + "add_domain/temps/" + env.Delivery + "/actions/" + fileName + ".temp")
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
	if _, existErr := os.Stat(deliveryInit); errors.Is(existErr, os.ErrNotExist) {
		initFile, initCreateErr := os.Create(deliveryInit)
		if initCreateErr != nil {
			return initCreateErr
		}
		defer initFile.Close()
		fmt.Fprintf(initFile, "%s", "package V1\n\nimport (\n\t\"github.com/labstack/echo/v4\"\n)\n\nfunc Init(echo *echo.Echo) {\n\n}\n")
	}
	input, readErr := ioutil.ReadFile(deliveryInit)
	if readErr != nil {
		return readErr
	}
	contentText := env.Delivery + domain.UpperName + ".Routs(echo)"
	contentImport := "import (\n\t" + env.Delivery + domain.UpperName + ` "` + env.MainRepositoryPath + "/delivery/http/V1/" + domain.SnakeName + `"`
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
