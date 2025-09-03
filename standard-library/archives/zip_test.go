package archives

import (
	"archive/zip"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestAddFileToZip(t *testing.T) {
	file, err := os.OpenFile("./testdata/generate_test.zip", os.O_WRONLY|os.O_CREATE, 0666)
	assert.NoError(t, err)
	defer file.Close()

	paths := []string{
		filepath.Join("testdata", "a.txt"),
		filepath.Join("testdata", "b.txt"),
	}
	zw := zip.NewWriter(file)
	err = AddFileToZip(zw, paths...)
	assert.NoError(t, err)
	defer zw.Close()
}
