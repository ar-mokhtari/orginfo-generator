package generator

import (
	"errors"
	"fmt"
	"os"
)

func rootCommand(args []string) error {
	//command-flag implement model : https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
	if len(args) < 1 {
		return errors.New("not receipt sub-command(s)")
	}
	commands := []Runner{
		newPattern(),
		deletePattern(),
		helpPattern(),
	}
	subCommand := os.Args[1]
	for _, command := range commands {
		if command.Name() == subCommand {
			if err := command.Init(os.Args[2:]); err != nil {
				return err
			}
			return command.Run()
		}
	}
	return fmt.Errorf("unknown subcommand: %s", subCommand)
}
