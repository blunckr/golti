/*
 * Genarate rsa keys.
 * https://gist.github.com/sdorra/1c95de8cb80da31610d2ad767cd6f251
 */

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func MakeKeys() rsa.PublicKey {
	if fileExists("private.pem") && fileExists("public.pem") {
		return getExisting()
	}

	return makeNew()
}

func getExisting() rsa.PublicKey {
	fileContent, _ := ioutil.ReadFile("public.pem")
	block, _ := pem.Decode(fileContent)
	key, _ := x509.ParsePKCS1PublicKey(block.Bytes) //.(rsa.PublicKey, error)
	fmt.Println(key)
	return *key
}

func makeNew() rsa.PublicKey {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := key.PublicKey

	savePEMKey("private.pem", key)

	savePublicPEMKey("public.pem", publicKey)
	return publicKey
}

func fileExists(name string) bool {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
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

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	asn1Bytes, err := asn1.Marshal(pubkey)
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
