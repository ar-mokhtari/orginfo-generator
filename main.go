package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ar-mokhtari/orginfo-generator/env"
)

func init() {
	repo, repoErr := ioutil.ReadFile(env.MainPath + "go.mod")
	if repoErr != nil {
		fmt.Println(repoErr)
	}
	repoContent := string(repo)
	lines := strings.Split(repoContent, "\n")
	firstLine := lines[0]
	finalRepo := bytes.Replace([]byte(firstLine), []byte("module "), []byte(""), -1)
	env.MainRepositoryPath = string(finalRepo)
}

func New() {
	//-------find project path---------
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	spilitedPwd := strings.Split(pwd, "/")
	mainPathSlice := spilitedPwd[:6]
	if env.MainRepositoryPath == "" {
		env.MainRepositoryPath = strings.Join(spilitedPwd[5:6], "/")
	}
	env.MainPath = strings.Join(mainPathSlice, "/") + "/"
	//-------find project path---------

	//FOR DEBUG
	fmt.Println(os.Args)

	//prepare command(s) with flag(s)
	flagErr := rootCommand(os.Args[1:])
	if flagErr != nil {
		fmt.Println(flagErr.Error())
		os.Exit(1)
	}
}
