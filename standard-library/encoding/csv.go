package encoding

import (
	"encoding/csv"
	"io"
	"os"
)

func ParseCSVFile[T any](filename string, fn func(record []string) (T, error)) ([]T, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	result := make([]T, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		v, err := fn(record)
		if err != nil {
			return nil, err
		}

		result = append(result, v)
	}

	return result, nil
}
