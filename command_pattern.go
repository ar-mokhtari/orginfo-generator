package generator

import (
	"flag"
)

func newPattern() *newCommand {
	res := &newCommand{
		Fs: flag.NewFlagSet("new", flag.ContinueOnError),
	}
	res.Fs.StringVar(&res.FlagDomainNameValue, "domain_name", "", "")
	res.Fs.StringVar(&res.FlagFieldsValue, "fields", "", "")
	res.Fs.BoolVar(&res.FlagFromFile, "from_file", false, "")
	return res
}

func deletePattern() *deleteCommand {
	res := &deleteCommand{
		Fs: flag.NewFlagSet("delete", flag.ContinueOnError),
	}
	res.Fs.StringVar(&res.FlagValue, "domain_name", "", "")
	res.Fs.BoolVar(&res.FlagFromFile, "from_file", false, "")
	return res
}

func helpPattern() *helpCommand {
	res := &helpCommand{
		Fs: flag.NewFlagSet("help", flag.ContinueOnError),
	}
	return res
}
