package main

import (
	"fmt"
	"os"

	"github.com/mouhamedbourouba/tlgr/cache"
	"github.com/mouhamedbourouba/tlgr/cli"
	"github.com/mouhamedbourouba/tlgr/config"
)

func main() {
	if err := cli.Parse(); err != nil {
		fmt.Fprint(os.Stderr, "Failed to parse flags: ", err.Error())
		os.Exit(1)
	}

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

	cacheInstance := cache.LoadCache(appCacheDir, archiveUrl)

	// Todo: add auto update option
	// if cacheInstance.GetState() == cache.CacheStateEmpty {
	// 	cacheInstance.Update()
	// } else if cacheInstance.GetState() == cache.CacheStateOutdated {
	// 	printOutdatedWarning(cacheInstance.GetCacheTime())
	// }

	if cli.GetListFlag() {
		pages, err := cacheInstance.GetCommandListForPlatform(cli.GetPlatform())

		if err != nil {
			fmt.Fprint(os.Stderr, "Error: ", err.Error(), ", please run tlgr -update to download the cache\n")
			os.Exit(1)
		}

		for _, page := range pages {
			fmt.Print(page, "\n")
		}
		return
	}

	// fallthrou flags
	if cli.GetClearCacheFlag() {
		if err = cacheInstance.Clear(); err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
		return
	}

	if cli.GetUpdateFlag() {
		err := cacheInstance.Update()

		if err != nil {
			fmt.Fprint(os.Stderr, "Error: ", err.Error())
			return
		}
		return
	}

	if cli.GetCommandString() != "" {
		_ = printTldr(cli.GetCommandString())
		return
	}

	cli.PrintHelp()
}

// func printOutdatedWarning(currentTime time.Time) {
// 	days := int(time.Since(currentTime).Hours() / 24)
// 	yellow := "\033[33m"
// 	reset := "\033[0m"
// 	fmt.Printf("%sWarning: Cache is %d day(s) old!%s\n", yellow, days, reset)
// }

func printTldr(s string) error {
	panic("unimplemented")
}

func getVersion() string {
	return "0.0.1"
}
