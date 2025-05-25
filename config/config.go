package config

import (
	"fmt"
	"os"
	"path/filepath"
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

func (platform PlatformType) String() string {
	switch platform {
	case Android:
		return "android"
	case Freebsd:
		return "freebsd"
	case Linux:
		return "linux"
	case Macos:
		return "macos"
	case Netbsd:
		return "netbsd"
	case Windows:
		return "windows"
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
		return -1, fmt.Errorf("undefined platform `%s` [possible values: linux, macos, windows, android, freebsd, netbsd]", platform)
	}
}

func GetArchiveUrlPath() string {
	return "https://github.com/tldr-pages/tldr/releases/latest/download/tldr.zip"
}

func GetAndCreateCacheDir() (string, error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	userCacheDir = filepath.Join(userCacheDir, "tlgr")

	err = os.MkdirAll(userCacheDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return userCacheDir, nil
}
