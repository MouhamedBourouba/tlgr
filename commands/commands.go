package commands

import (
	"fmt"
	"os"
	c "github.com/mouhamedbourouba/tlgr/config"
)

type Tldr struct {
}

var commands map[string]Tldr = make(map[string]Tldr)

func IndexCache(path string, platform c.PlatformType) error {
	entrs, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, value := range entrs {
		fmt.Printf("value.Name(): %v\n", value.Name())
	}

	return nil
}
