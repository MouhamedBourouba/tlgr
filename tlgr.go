package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mouhamedbourouba/tlgr/cli"
)

func main() {
	cli.Parse()

	if cli.GetHelpFlag() {
		cli.PrintHelp()
		os.Exit(0)
	}

	if cli.GetClearCacheFlag() {
		fmt.Printf("clearing cache ---\n")
		clearLocalCache()
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
		printTldr(cli.GetCommandString())
		os.Exit(0)
	}

	cli.PrintHelp()
}

func printTldr(s string) error {
	const cachePath = "./tldr/pages/common"

	dir, err := os.ReadDir(cachePath)
	if err != nil {
		return err
	}
	for _, dirEntry := range dir {
		var commandName = strings.TrimSuffix(dirEntry.Name(), ".md")
		if commandName == s {
			file, err := os.ReadFile(cachePath + "/" + dirEntry.Name())
			if err != nil {
				return err
			}
			print(string(file))
		}
	}
	return nil
}

func clearLocalCache() {
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
