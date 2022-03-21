package util

import (
	"fmt"
	"io/ioutil"
)

type FileReader interface {
	JSON() ([]byte, error)
}

type fileReader struct {
	file string
}

func NewFileReader(file string) FileReader {
	return &fileReader{file: file}
}

func (r *fileReader) JSON() ([]byte, error) {
	file, err := ioutil.ReadFile(r.file)
	if err != nil {
		return nil, fmt.Errorf("unable to read file %v: %v", r.file, err)
	}

	return file, nil
}
