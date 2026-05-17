package _archive

import (
	"archive/tar"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestAddFileToTar(t *testing.T) {
	file, err := os.OpenFile("./testdata/test.tar", os.O_WRONLY|os.O_CREATE, 0666)
	assert.NoError(t, err)
	defer file.Close()

	paths := []string{
		filepath.Join("testdata", "a.txt"),
		filepath.Join("testdata", "b.txt"),
	}
	tw := tar.NewWriter(file)
	defer tw.Close()
	err = AddFileToTar(tw, paths...)
	assert.NoError(t, err)
}

func TestExtractTar(t *testing.T) {
	file, err := os.Open("./testdata/test.tar")
	assert.NoError(t, err)
	defer file.Close()

	tr := tar.NewReader(file)

	err = ExtractTar(tr, "./testdata/testdir")
	assert.NoError(t, err)
}
