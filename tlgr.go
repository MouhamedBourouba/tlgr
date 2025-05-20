package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mouhamedbourouba/tlgr/cache"
	"github.com/mouhamedbourouba/tlgr/cli"
	"github.com/mouhamedbourouba/tlgr/config"
)

func main() {
	cli.Parse()

	if cli.GetHelpFlag() {
		cli.PrintHelp()
		return
	}

	if cli.GetVersionFlag() {
		fmt.Printf("TLGR %s\n", getVersion())
		return
	}

	archiveUrl := config.GetArchiveUrlPath()
	appCacheDir, err := config.GetAndCreateCacheDir()
	if err != nil {
		panic(err)
	}
	cacheInstance, err := cache.LoadCache(appCacheDir, archiveUrl)

	// Todo: add auto update option
	// if cacheInstance.GetState() == cache.CacheStateEmpty {
	// 	cacheInstance.Update()
	// } else if cacheInstance.GetState() == cache.CacheStateOutdated {
	// 	printOutdatedWarning(cacheInstance.GetCacheTime())
	// }

	if cli.GetListFlag() {
		fmt.Printf("Listing ---\n")
		listAllCommands()
		return
	}

	// fallthrou flags
	if cli.GetClearCacheFlag() {
		cacheInstance.Clear()
		return
	}

	if cli.GetUpdateFlag() {
		err := cacheInstance.Update()

		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			return
		}

		return
	}

	if cli.GetCommandString() != "" {
		printTldr(cli.GetCommandString())
		return
	}

	cli.PrintHelp()
}

func printOutdatedWarning(currentTime time.Time) {
	days := int(time.Since(currentTime).Hours() / 24)
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Printf("%sWarning: Cache is %d day(s) old!%s\n", yellow, days, reset)
}

func printTldr(s string) error {
	panic("unimplemented")
}

func listAllCommands() {
	panic("unimplemented")
}

func getVersion() string {
	return "0.0.1"
}
