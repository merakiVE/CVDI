package utils

import (
	"io/ioutil"
)

const (
	PRIVATE_KEY_PATH = "keys/private.pem"
	PUBLIC_KEY_PATH  = "keys/public.pem"
)

func ReadSecrectKey() ([]byte, error) {
	return ioutil.ReadFile(PRIVATE_KEY_PATH)
}

func ReadPublicKey() ([]byte, error) {
	return ioutil.ReadFile(PUBLIC_KEY_PATH)
}
