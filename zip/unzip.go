package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func unzip(target, output string) error {
	archive, err := zip.OpenReader(target)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, file := range archive.File {
		filePath := filepath.Join(output, file.Name)
		fmt.Println("unzipping file:", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(output)+string(os.PathSeparator)) {
			fmt.Println("invalid file path:", filePath)
			return fmt.Errorf("invalid file path: %s", filePath)
		}

		if file.FileInfo().IsDir() {
			fmt.Println("creating directory:", file)
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := file.Open()
		if err != nil {
			return err
		}
		if _, err := io.Copy(outFile, fileInArchive); err != nil {
			return err
		}

		outFile.Close()
		fileInArchive.Close()
	}

	return nil
}

func main() {
	target := flag.String("t", "", "target name")
	output := flag.String("o", "", "output directory")
	flag.Parse()
	fmt.Println("arg", flag.Args())

	if err := unzip(*target, *output); err != nil {
		fmt.Println(err)
	}
}
