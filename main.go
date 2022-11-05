package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/ar-mokhtari/orginfo-generator/env"
)

func New() {
	//-------find project path---------
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	spilitedPwd := strings.Split(pwd, "/")
	mainPathSlice := spilitedPwd[:6]
	env.MainRepositoryPath = strings.Join(spilitedPwd[5:6], "/")
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
