package cache

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CacheState int

const (
	CacheStateFresh CacheState = iota
	CacheStateEmpty
	CacheStateOutdated
)

type Cache struct {
	sate CacheState
}

func LoadCache(cacheDir string) Cache {
	return Cache{}
}

func (cache Cache) GetState() CacheState {
	return cache.sate
}

func (cache Cache) Update() error {
	return nil
}

func (cache Cache) Clear() {
}

func (cache Cache) findPage() string {
	return ""
}

func downloadTldrArchive(url string, dst string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("resp.Status: %v\n", resp.Status)
		return errors.New("Failed to download")
	}

	file, err := os.Create(dst)
	if err != nil {
		return err
	}

	var writeCounter = newDownloadProgressCounter(uint64(resp.ContentLength))
	if _, err := io.Copy(io.MultiWriter(&writeCounter, file), resp.Body); err != nil {
		return err
	}
	return nil
}
