package generator

import (
	"os"
)

func New() {
	//prepare command(s) with flag(s)
	flagErr := rootCommand(os.Args[1:])
	if flagErr != nil {
		println(flagErr.Error())
		os.Exit(1)
	}
}
