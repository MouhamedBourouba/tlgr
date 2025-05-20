package cli

import (
	"flag"
	"fmt"
	"strings"

	p "github.com/mouhamedbourouba/tlgr/config"
)

var (
	platform      p.PlatformType = p.GetDefaultPlatform()
	commandString string         = ""

	parsed bool = false

	listFlag       *bool = flag.Bool("list", false, "List all commands in the cache")
	updateFlag     *bool = flag.Bool("update", false, "Update the local cache")
	helpFlag       *bool = flag.Bool("help", false, "Print help")
	clearCacheFlag *bool = flag.Bool("clear-cache", false, "Clears Local Cache")
	versionFlag    *bool = flag.Bool("version", false, "Print the version")

	platformString = flag.String(
		"platform",
		p.GetDefaultPlatform().ToString(),
		"Override the operating system [possible values: linux, macos, windows, android, freebsd, netbsd]",
	)
)

func init() {
	flag.Usage = func() {
		fmt.Print("TLGR A fast tldr client written in go\n")
		fmt.Print("Auther: Mouhamed Redha bourouba <mouhamedmobiledev@gmail.com>\n\n")
		fmt.Print("Usage: tlgr [OPTIONS] [COMMAND]...", "\n\n")
		fmt.Print("Arguments:\n  [COMMAND]...  The command to show (e.g. `git commit` or `awk`)", "\n\n")
		fmt.Print("Options:\n")
		flag.PrintDefaults()
	}
}

func Parse() error {
	flag.Parse()

	parsedPlatform, err := p.ParsePlatform(*platformString)
	if err != nil {
		return err
	}
	platform = parsedPlatform

	if flag.NArg() >= 1 {
		commandString = strings.ToLower(strings.Join(flag.Args(), "-"))
	}

	parsed = true

	return nil
}

func PrintHelp() {
	flag.Usage()
}

func assertParsed() {
	if !parsed {
		panic("Call Parse Before calling this function")
	}
}

func GetCommandString() string {
	assertParsed()
	return commandString
}

func GetPlatform() p.PlatformType {
	assertParsed()
	return platform
}

func GetListFlag() bool {
	assertParsed()
	return *listFlag
}

func GetUpdateFlag() bool {
	assertParsed()
	return *updateFlag
}

func GetHelpFlag() bool {
	assertParsed()
	return *helpFlag
}

func GetVersionFlag() bool {
	assertParsed()
	return *versionFlag
}

func GetClearCacheFlag() bool {
	assertParsed()
	return *clearCacheFlag
}
