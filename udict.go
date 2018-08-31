package main

import (
	"flag"
	color "github.com/fatih/color"
	"os"
	"strings"
	"udict/dictionary"
)

func main() {
	defineCommand := flag.NewFlagSet("define", flag.ExitOnError)

	listTextPtr := defineCommand.String("word", "", "Word to define. (Required)")

	if len(os.Args) < 2 {
		color.Red("error: missing the subcommand define")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "define":
		defineCommand.Parse(os.Args[2:])
	default:
		color.Red("error: unknown subcommand")
		os.Exit(1)
	}

	if defineCommand.Parsed() {
		if *listTextPtr == "" {
			defineCommand.PrintDefaults()
			os.Exit(1)
		}
		definedWord := dictionary.Define(*listTextPtr)
		top5Defintions := definedWord.Definitions
		color.Yellow("Definition(s) are: \n")
		for i, elem := range top5Defintions {
			def := strings.Replace(elem.Definition, "[", "", -1)
			def = strings.Replace(def, "]", "", -1)
			color.Cyan("%v: %s\n\n", i+1, def)
		}
	}
}
