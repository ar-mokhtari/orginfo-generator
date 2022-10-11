package generator

import (
	"fmt"
	"os"
)

func New() {
	//main root
	fmt.Sprintf("root is: %s", os.Args[1])
	//prepare command(s) with flag(s)
	flagErr := rootCommand(os.Args[1:])
	if flagErr != nil {
		println(flagErr.Error())
		os.Exit(1)
	}
}
