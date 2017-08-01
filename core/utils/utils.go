package utils

import (
	"io/ioutil"
)

func ReadBinaryFile(_path string) ([]byte, error) {
	return ioutil.ReadFile(_path)
}
