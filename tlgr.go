package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mouhamedbourouba/tlgr/cache"
	"github.com/mouhamedbourouba/tlgr/cli"
	"github.com/mouhamedbourouba/tlgr/config"
)


func printOutdatedWarning(currentTime time.Time) {
	days := int(time.Since(currentTime).Hours() / 24)
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Printf("%sWarning: Cache is %d day(s) old!%s\n", yellow, days, reset)
}

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

	appCacheDir, err := config.GetAndCreateCacheDir()
	if err != nil {
		panic(err)
	}
	archiveUrl := config.GetArchiveUrlPath()

	cacheInstance, err := cache.LoadCache(appCacheDir, archiveUrl)
	if cacheInstance.GetState() == cache.CacheStateEmpty {
		cacheInstance.Update()
	} else if cacheInstance.GetState() == cache.CacheStateOutdated {
		printOutdatedWarning(cacheInstance.GetCacheTime())
	}

	if cli.GetListFlag() {
		fmt.Printf("Listing ---\n")
		listAllCommands()
		return
	}

	// fallthrou flags
	if cli.GetClearCacheFlag() {
		fmt.Printf("clearing cache ---\n")
		clearLocalCache()
	}

	if cli.GetUpdateFlag() {
		err := updateCache()
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			return
		}
	}

	if cli.GetCommandString() != "" {
		printTldr(cli.GetCommandString())
		return
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

func updateCache() error {
	return nil
}

func getVersion() string {
	return "0.0.1"
}
