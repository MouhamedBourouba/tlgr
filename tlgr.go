package main

import (
	"fmt"
	"os"

	"github.com/mouhamedbourouba/tlgr/cli"
)

func main() {
	cli.Parse()

	if cli.GetHelpFlag() {
		cli.PrintHelp()
		os.Exit(0)
	}

	if cli.GetVersionFlag() {
		fmt.Printf("TLGR %s\n", getVersion())
		os.Exit(0)
	}

	if cli.GetUpdateFlag() {
		fmt.Printf("Updating cache ---\n")
		updateCache()
		os.Exit(0)
	}

	if cli.GetListFlag() {
		fmt.Printf("Listing ---\n")
		listAllCommands()
		os.Exit(0)
	}
	
	if cli.GetCommandString() != "" {
		fmt.Print("Printing tldr for", cli.GetCommandString())
		printTldr(cli.GetCommandString())
		os.Exit(0)
	}

	cli.PrintHelp()
}

func printTldr(s string) {
	panic("unimplemented")
}

func listAllCommands() {
	panic("unimplemented")
}

func updateCache() {
	panic("unimplemented")
}

func getVersion() string {
	return "0.0.1"
}
