package cache

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
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
	tldrRepoDir      string
}

func isCacheExpired(modtime time.Time) bool {
	expirationDate := modtime.AddDate(0, 0, 30)
	return time.Now().After(expirationDate)
}

func LoadCache(cacheDir string, archiveUrl string) Cache {
	cacheState := CacheStateEmpty
	cacheTime := time.Now()

	tldrRepoDir := filepath.Join(cacheDir, ARCHIVE_OUTPUT_DIR)
	indexJsonFilePath := filepath.Join(tldrRepoDir, "index.json")

	_, err := os.Stat(tldrRepoDir)

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
		tldrRepoDir:      tldrRepoDir,
	}
}

func (cache Cache) GetState() CacheState {
	return cache.sate
}

func (cache Cache) GetCacheTime() time.Time {
	return cache.currentCacheTime
}

func (cache *Cache) Update() error {
	archivePath := filepath.Join(cache.cacheDir, ARCHIVE_FILE_NAME)
	outputDir := filepath.Join(cache.cacheDir, ARCHIVE_OUTPUT_DIR)

	err := downloadTldrArchive(cache.archiveUrl, archivePath)
	if err != nil {
		return err
	}

	err = unzipArchive(archivePath, outputDir)
	if err != nil {
		return err
	}

	_ = os.Remove(archivePath)
	cache.sate = CacheStateFresh

	return nil
}

func (cache Cache) GetCommandListForPlatform(platform config.PlatformType) ([]string, error) {
	if cache.sate == CacheStateEmpty {
		return nil, errors.New("page cache not found, try tlgr -update")
	}

	commonPagesPath := filepath.Join(cache.tldrRepoDir, "pages/common")
	platformPagesPth := filepath.Join(cache.tldrRepoDir, "pages", platform.String())

	commonPages, err := os.ReadDir(commonPagesPath)
	if err != nil {
		print(err)
		return nil, errors.New("invalid cache, try tlgr -update")
	}
	platformPages, err := os.ReadDir(platformPagesPth)
	if err != nil {
		return nil, errors.New("invalid cache, try tlgr -update")
	}

	allPages := slices.Concat(commonPages, platformPages)
	allPagesNames := make([]string, 0)

	for _, entry := range allPages {
		if !slices.Contains(allPagesNames, entry.Name()) {
			allPagesNames = append(allPagesNames, entry.Name())
		}
	}

	return allPagesNames, nil
}

func (cache Cache) Clear() error {
	return os.RemoveAll(cache.cacheDir)
}

func (cache Cache) FindPage(page string, p config.PlatformType) (string, error) {
	commonPath := filepath.Join(cache.tldrRepoDir, "pages", "common", (page + ".md"))
	platformPath := filepath.Join(cache.tldrRepoDir, "pages", p.String(), (page + ".md"))

	if _, err := os.Stat(platformPath); err == nil {
		return platformPath, nil
	}

	if _, err := os.Stat(commonPath); err == nil {
		return commonPath, nil
	}

	return "", fmt.Errorf("page `%s` not found in cache, try tlgr -update", page)
}

func downloadTldrArchive(url string, dst string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to download, status code %d", resp.StatusCode)
	}

	file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer file.Close()

	bar := progressbar.NewOptions(
		int(resp.ContentLength),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetDescription("Downloading TLDR repository"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	if _, err := io.Copy(io.MultiWriter(bar, file), resp.Body); err != nil {
		return err
	}

	fmt.Print("\n")

	return nil
}

func unzipArchive(archivePath string, outputPath string) error {
	zipReader, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	bar := progressbar.NewOptions(
		len(zipReader.File) - 300,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetDescription("Unzipping the TLDR archive "),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

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

		_ = bar.Add(1)
	}

	fmt.Println()

	return nil
}
