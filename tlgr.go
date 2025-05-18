package main

import (
	"fmt"
	"os"

	"github.com/mouhamedbourouba/tlgr/cli"
	"github.com/mouhamedbourouba/tlgr/config"
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

type Command struct {
	name       string
	pathToTldr string
	platforms  []config.PlatformType
}

type Commands []Command

func IndexDir(dirPath string) error {
	dir, err := os.ReadDir(dirPath)

	if err != nil {
		return err
	}

	for _, dirEntry := range dir {
		println(dirEntry.Name())
	}

	return nil
}

func printTldr(s string) error {
	const cachePath = "./tldr/pages/"
	
	IndexDir(cachePath)

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
