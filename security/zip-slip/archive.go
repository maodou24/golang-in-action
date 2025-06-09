package zipslip

import (
	"archive/zip"
	"os"
)

func WriteToZip() error {
	zipName := "test.zip"
	f, err := os.OpenFile(zipName, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	writer := zip.NewWriter(f)
	defer writer.Close()
  
	//writer.AddFS()

	return nil
}
