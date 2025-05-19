package main

import (
	"archive/zip"
	"io"
	"path/filepath"

	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mouhamedbourouba/tlgr/cli"
	"github.com/mouhamedbourouba/tlgr/config"
)

func main() {
	cli.Parse()

	if cli.GetHelpFlag() {
		cli.PrintHelp()
		os.Exit(0)
	}

	if cli.GetClearCacheFlag() {
		fmt.Printf("clearing cache ---\n")
		clearLocalCache()
	}

	if cli.GetVersionFlag() {
		fmt.Printf("TLGR %s\n", getVersion())
		os.Exit(0)
	}

	if cli.GetUpdateFlag() {
		err := updateCache()
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	if cli.GetListFlag() {
		fmt.Printf("Listing ---\n")
		listAllCommands()
		os.Exit(0)
	}

	if cli.GetCommandString() != "" {
		printTldr(cli.GetCommandString())
		os.Exit(0)
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
	archiveUrl := "https://github.com/tldr-pages/tldr/releases/latest/download/tldr.zip"
	dirPath := "tldr"
	archivePath := "tldr.zip"

	// resp, err := http.Get(archiveUrl)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()
	//
	// if resp.StatusCode != 200 {
	// 	fmt.Printf("resp.Status: %v\n", resp.Status)
	// 	return errors.New("Failed to download")
	// }
	//
	// file, err := os.Create(archivePath)
	// if err != nil {
	// 	return err
	// }
	//
	// var writeCounter = NewDownloadProgressCounter(uint64(resp.ContentLength))
	// if _, err := io.Copy(io.MultiWriter(&writeCounter, file), resp.Body); err != nil {
	// 	return err
	// }
	//
	// zipReader, err := zip.OpenReader(archivePath)
	// if err != nil {
	// 	return err
	// }
	// defer zipReader.Close()
	// println("Unzipping Archive in", dirPath, "...")
	//
	// for _, extractedFile := range zipReader.File {
	// 	filepath := filepath.Join(dirPath, extractedFile.Name)
	//
	// 	if extractedFile.FileInfo().IsDir() {
	// 		err := os.MkdirAll(filepath, os.ModePerm)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		continue
	// 	}
	//
	// 	extractedFileReader, err := extractedFile.Open()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer extractedFileReader.Close()
	//
	// 	createdFile, err := os.Create(filepath)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer createdFile.Close()
	//
	// 	if _, err := io.Copy(createdFile, extractedFileReader); err != nil {
	// 		return err
	// 	}
	// }
	// 
	// println("Done !!!")
	// return nil
	return nil
}

func getVersion() string {
	return "0.0.1"
}
