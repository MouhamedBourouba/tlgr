package cli

import (
	"errors"
	"flag"
	"strings"

	p "github.com/mouhamedbourouba/tlgr/config"
)

var (
	platform      p.Platform
	commandString string = ""

	parsed bool = false

	listFlag    *bool = flag.Bool("list", false, "List all commands in the cache")
	updateFlag  *bool = flag.Bool("update", false, "Update the local cache")
	helpFlag    *bool = flag.Bool("help", false, "Print help")
	versionFlag *bool = flag.Bool("version", false, "Print the version")

	platformString = flag.String(
		"platform",
		p.GetDefaultPlatform().ToString(),
		"Override the operating system, can be specified multiple times in order of preference [possible values: linux, macos, windows, android, freebsd, netbsd]",
	)
)

func Parse() error {
	flag.Parse()

	parsedPlatform, err := p.ParsePlatform(*platformString)
	if err != nil {
		return errors.New("Undefined Platform")
	}
	platform = parsedPlatform

	if flag.NArg() >= 1 {
		commandString = strings.Join(flag.Args(), "-")
	}

	parsed = true

	return nil
}

func PrintHelp() {
	flag.PrintDefaults()
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

func GetPlatformString() p.Platform {
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
