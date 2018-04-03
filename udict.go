package main

import (
	"flag"
	"fmt"
	"os"
	"udict/dictionary"
)

func main() {
	defineCommand := flag.NewFlagSet("define", flag.ExitOnError)

	listTextPtr := defineCommand.String("word", "", "Word to define. (Required)")

	if len(os.Args) < 2 {
		fmt.Println("error: word subcommand is required")
		defineCommand.PrintDefaults()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "define":
		defineCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if defineCommand.Parsed() {
		if *listTextPtr == "" {
			defineCommand.PrintDefaults()
			os.Exit(1)
		}
		definedWord := dictionary.Define(*listTextPtr)
		fmt.Printf("Definitions are: \n")
		for i, elem := range definedWord.Definitions {
			fmt.Printf("%v: %s\n", i+1, elem.Definition)
		}
	}
}
