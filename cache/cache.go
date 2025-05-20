package cache

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/mouhamedbourouba/tlgr/config"
	"github.com/schollz/progressbar/v3"
)

const ARCHIVE_FILE_NAME = "tldr.zip"
const ARCHIVE_OUTPUT_DIR = "tldr"

type CacheState int

const (
	CacheStateFresh CacheState = iota
	CacheStateEmpty
	CacheStateOutdated
)

type Cache struct {
	sate             CacheState
	currentCacheTime time.Time
	archiveUrl       string
	cacheDir         string
}

func isCacheExpired(modtime time.Time) bool {
	expirationDate := modtime.AddDate(0, 0, 30)
	return time.Now().After(expirationDate)
}

func LoadCache(cacheDir string, archiveUrl string) (Cache, error) {
	cacheState := CacheStateEmpty
	cacheTime := time.Now()

	tldrCacheDir := filepath.Join(cacheDir, ARCHIVE_OUTPUT_DIR)
	indexJsonFilePath := filepath.Join(tldrCacheDir, "index.json")

	_, err := os.Stat(tldrCacheDir)

	if err == nil {
		fileInfo, err := os.Stat(indexJsonFilePath)
		if err == nil {
			modTime := fileInfo.ModTime()
			cacheTime = modTime

			if !isCacheExpired(modTime) {
				cacheState = CacheStateFresh
			} else {
				cacheState = CacheStateOutdated
			}
		}
	}

	return Cache{
		archiveUrl:       archiveUrl,
		cacheDir:         cacheDir,
		sate:             cacheState,
		currentCacheTime: cacheTime,
	}, nil
}

func (cache Cache) GetState() CacheState {
	return cache.sate
}

func (cache Cache) GetCacheTime() time.Time {
	return cache.currentCacheTime
}

func (cache Cache) Update() error {
	archivePath := filepath.Join(cache.cacheDir, ARCHIVE_FILE_NAME)
	outputDir := filepath.Join(cache.cacheDir, ARCHIVE_OUTPUT_DIR)

	err := downloadTldrArchive(cache.archiveUrl, archivePath)
	if err != nil {
		return err
	}

	println("Unzipping Archive in", outputDir, "...")

	err = unzipArchive(archivePath, outputDir)
	if err != nil {
		return err
	}

	os.Remove(archivePath)

	return nil
}

func (cache Cache) Clear() {
}

func (cache Cache) FindPage(p config.PlatformType) string {
	pagePath := filepath.Join(cache.cacheDir, ARCHIVE_OUTPUT_DIR, p.ToString(), p.ToString()+"."+p.ToString()+".md")
	if _, err := os.Stat(pagePath); err == nil {
		return pagePath
	}

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
	defer file.Close()

	progress := progressbar.DefaultBytes(resp.ContentLength, "Downloading Repository")
	if _, err := io.Copy(io.MultiWriter(progress, file), resp.Body); err != nil {
		return err
	}

	return nil
}

func unzipArchive(archivePath string, outputPath string) error {
	zipReader, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, extractedFile := range zipReader.File {
		filepath := filepath.Join(outputPath, extractedFile.Name)

		if extractedFile.FileInfo().IsDir() {
			err := os.MkdirAll(filepath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		extractedFileReader, err := extractedFile.Open()
		if err != nil {
			return err
		}
		defer extractedFileReader.Close()

		createdFile, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer createdFile.Close()

		if _, err := io.Copy(createdFile, extractedFileReader); err != nil {
			return err
		}
	}

	return nil
}
