package main

import (
	"os"

	"github.com/comavius/unfoldcpp"
)

func tooFewArgsError() {
	println("Too few arguments.\nSee unfoldcppcli --help for more information.")
}

func invalidActionError() {
	println("Invalid action.\nSee unfoldcppcli --help for more information.")
}

func showHelp() {
	println("unfoldcppcli is a command line interface for unfoldcpp.\nUsage: unfoldcppcli [action]\nActions:\n\thelp - show this help\n\tunfold [path] - unfold C++ code from [path] and stdout it")
}

func main() {
	// check args
	if len(os.Args) < 2 {
		tooFewArgsError()
		return
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
			return
		}
		// get path
		path := os.Args[2]
		// unfold
		code, err := unfoldcpp.Unfold(path)
		if err != nil {
			println(err.Error())
			return
		}
		// stdout
		println(code)
	default:
		invalidActionError()
	}
}
