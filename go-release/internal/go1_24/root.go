package go1_24

import (
	"io"
	"os"
)

func ReadFile(filename string) error {
	root, err := os.OpenRoot("./")
	if err != nil {
		return err
	}
	defer root.Close()

	file, err := root.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	return err
}
