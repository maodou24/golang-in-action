package _archive

import (
	"archive/tar"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func AddFileToTar(tw *tar.Writer, filePaths ...string) error {
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

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(tw, file)
		if err != nil {
			return err
		}
	}

	return tw.Flush()
}

func ExtractTar(tr *tar.Reader, dtsDir string) error {
	_, err := os.Stat(dtsDir)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(dtsDir, 0755)
		if err != nil {
			return err
		}
	}
	for {
		header, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		outPath := filepath.Join(dtsDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(outPath, 0755)
		case tar.TypeReg:
			_ = os.MkdirAll(filepath.Dir(outPath), 0755)
			outFile, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, tr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
