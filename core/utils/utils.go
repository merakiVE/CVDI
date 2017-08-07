package utils

import (
	"io/ioutil"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
	"path"
	"strings"
)

func ReadBinaryFile(_path string) ([]byte, error) {
	return ioutil.ReadFile(_path)
}

func Exists(path string) (bool) {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func IsEmptyString(str string) (bool) {
	return len(strings.TrimSpace(str)) == 0
}

/*
 * Genarate rsa keys.
 */

func GenerateKeys(path_dir string) {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := key.PublicKey

	SaveGobKey(path.Join(path_dir, "private.key"), key)
	SavePEMKey(path.Join(path_dir, "private.pem"), key)

	SaveGobKey(path.Join(path_dir, "public.key"), publicKey)
	SavePublicPEMKey(path.Join(path_dir, "public.pem"), publicKey)
}

func SaveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func SavePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func SavePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	//asn1Bytes, err := asn1.Marshal(pubkey)
	asn1Bytes, err := x509.MarshalPKIXPublicKey(&pubkey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
