package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

type zipArchiver struct{}

func (z *zipArchiver) Archive(zipFileName string, files []string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return fmt.Errorf("can't create zip file %s: %w", zipFileName, err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		inputFile, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("can't open file %s: %w", file, err)
		}
		defer inputFile.Close()

		stats, err := inputFile.Stat()
		if err != nil {
			return fmt.Errorf("can't stat file %s: %w", file, err)
		}

		if stats.IsDir() {
			continue
		}

		zipEntry, err := zipWriter.Create(file)
		if err != nil {
			return fmt.Errorf("can't add file %s to zip: %w", file, err)
		}

		_, err = io.Copy(zipEntry, inputFile)
		if err != nil {
			return fmt.Errorf("can't copy file %s to zip archive: %w", file, err)
		}
	}

	return nil
}

func (z *zipArchiver) Unarchive(filePath string) error {
	archive, err := zip.OpenReader(filePath)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer archive.Close()
	err = os.MkdirAll(filePath+".package", 0777)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	for _, f := range archive.File {
		n := filepath.Join(filePath+".package", f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(n, 0777); err != nil {
				slog.Error(err.Error())
				return err
			}
			continue
		}

		dst, err := os.Create(n)
		if err != nil {
			slog.Error(err.Error())
			return err
		}
		fileReader, err := f.Open()
		if err != nil {
			return err
		}
		if _, err := io.Copy(dst, fileReader); err != nil {
			slog.Error(err.Error())
			return err
		}
	}

	return nil
}

func NewZipArchiver() Archiver {
	return &zipArchiver{}
}
