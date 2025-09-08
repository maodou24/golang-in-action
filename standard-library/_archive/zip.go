package _archive

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func AddFileToZip(zw *zip.Writer, filePaths ...string) error {
	for _, path := range filePaths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			return err
		}

		fileHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		fileHeader.Name = filepath.Base(path)
		fileHeader.Method = zip.Deflate

		writer, err := zw.CreateHeader(fileHeader)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}
	}

	return zw.Flush()
}

func ExtractZip(zw *zip.ReadCloser, dtsDir string) error {
	_, err := os.Stat(dtsDir)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(dtsDir, 0755)
		if err != nil {
			return err
		}
	}

	for _, file := range zw.File {
		outFilePath := filepath.Join(dtsDir, file.Name)
		outFile, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_WRONLY, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		reader, err := file.Open()
		if err != nil {
			return err
		}
		defer reader.Close()

		_, err = io.Copy(outFile, reader)
		if err != nil {
			return err
		}
	}
	return nil
}
