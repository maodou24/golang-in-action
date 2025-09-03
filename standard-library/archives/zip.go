package archives

import (
	"archive/zip"
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
