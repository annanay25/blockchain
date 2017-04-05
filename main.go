package main

import (
//	"crypto"
	"fmt"
	"io/ioutil"
//	"errors"
//	"encoding/pem"
//	"encoding/x509"

	"github.com/spacemonkeygo/openssl"
)

func checkError(err error){
	fmt.Errorf("Error: ", err)
}

func loadPrivateKey(path string) (openssl.PrivateKey, error){
	pri_key, err := ioutil.ReadFile(path)
	checkError(err)

	return openssl.LoadPrivateKeyFromPEM(pri_key)
}

func loadPublicKey(path string) (openssl.PublicKey, error){
	pub_key, err := ioutil.ReadFile(path)
	checkError(err)

	return openssl.LoadPublicKeyFromPEM(pub_key)
}

func main(){
	
	private, err := loadPrivateKey("private.pem")
	checkError(err)

	// fmt.Println("Private Key: ",private)

	public, err := loadPublicKey("public.pem")
	checkError(err)

	// fmt.Println("Public Key: ", public)
}

