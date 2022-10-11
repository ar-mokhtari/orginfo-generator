package generator

import (
	"fmt"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"os"
	"strings"
)

func New() {
	//-------find project path---------
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	spilitedPwd := strings.Split(pwd, "/")
	mainPathSlice := spilitedPwd[:len(spilitedPwd)-2]
	env.MainPath = strings.Join(mainPathSlice, "/") + "/"
	//-------find project path---------

	//prepare command(s) with flag(s)
	flagErr := rootCommand(os.Args[1:])
	if flagErr != nil {
		println(flagErr.Error())
		os.Exit(1)
	}
}
