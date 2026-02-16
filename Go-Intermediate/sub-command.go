package main

import (
	"flag"
	"fmt"
	"os"
)

func subCommandExample() {
	subCommand1 := flag.NewFlagSet("firstSub" , flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub" , flag.ExitOnError)

	firstFlag := subCommand1.Bool("processing" , false , "Command for processing")
	secondFlag := subCommand2.Bool("processing" , false , "Command for processing")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'firstSub' or 'secondSub' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("First subcommand processing:", *firstFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("Second subcommand processing:", *secondFlag)
	default:
		fmt.Println("Expected 'firstSub' or 'secondSub' subcommands")
		os.Exit(1)
	}
}