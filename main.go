package main

import (
	"fmt"
	"os"

	"github.com/comavius/unfoldcpp"
)

func tooFewArgsError() {
	fmt.Fprintln(os.Stderr, "Too few arguments.\nSee unfoldcppcli --help for more information.")
}

func invalidActionError() {
	fmt.Fprintln(os.Stderr, "Invalid action.\nSee unfoldcppcli --help for more information.")
}

func showHelp() {
	println(os.Stderr, "unfoldcppcli is a command line interface for unfoldcpp.\nUsage: unfoldcppcli [action]\nActions:\n\thelp - show this help\n\tunfold [path] - unfold C++ code from [path] and stdout it")
}

func main() {
	// check args
	if len(os.Args) < 2 {
		tooFewArgsError()
		os.Exit(1)
	}
	// get type of action
	action := os.Args[1]
	// check action
	switch action {
	case "help":
		showHelp()
	case "unfold":
		// check args
		if len(os.Args) < 3 {
			tooFewArgsError()
			os.Exit(1)
		}
		// get path
		path := os.Args[2]
		// unfold
		code, err := unfoldcpp.Unfold(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// stdout
		fmt.Println(code)
	default:
		invalidActionError()
		os.Exit(1)
	}
	os.Exit(0)
}
