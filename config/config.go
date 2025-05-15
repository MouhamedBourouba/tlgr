package config

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type Platform int

const (
	Linux Platform = iota
	Macos
	Windows
	Android
	Freebsd
	Netbsd
	openbsd
)

func GetDefaultPlatform() Platform {
	switch runtime.GOOS {
	case "android":
		return Android
	case "darwin":
		return Macos
	case "freebsd":
		return Freebsd
	case "linux":
		return Linux
	case "netbsd":
		return Netbsd
	case "windows":
		return Windows
	}
	panic("Error: unsupported paltform")
}

func (platform Platform) ToString() string {
	switch platform {
	case Android:
		return "Android"
	case Freebsd:
		return "Freebsd"
	case Linux:
		return "Linux"
	case Macos:
		return "Macos"
	case Netbsd:
		return "Netbsd"
	case Windows:
		return "Windows"
	case openbsd:
		return "openbsd"
	default:
		panic(fmt.Sprintf("unexpected main.Platform: %#v", platform))
	}
}

func ParsePlatform(platform string) (Platform, error) {
	switch strings.ToLower(platform) {
	case "linux":
		return Linux, nil
	case "macos":
		return Macos, nil
	case "windows":
		return Windows, nil
	case "android":
		return Android, nil
	case "freebsd":
		return Freebsd, nil
	case "netbsd":
		return Netbsd, nil
	default:
		return -1, errors.New("Undefined platform")
	}
}
