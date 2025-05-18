package config

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type PlatformType int

const (
	Linux PlatformType = iota
	Macos
	Windows
	Android
	Freebsd
	Netbsd
	openbsd
	common
)

func GetDefaultPlatform() PlatformType {
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

func (platform PlatformType) ToString() string {
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
	case common:
		return "common"
	default:
		panic(fmt.Sprintf("unexpected main.Platform: %#v", platform))
	}
}

func ParsePlatform(platform string) (PlatformType, error) {
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
	case "common":
		return common, nil
	default:
		return -1, errors.New("Undefined platform")
	}
}
