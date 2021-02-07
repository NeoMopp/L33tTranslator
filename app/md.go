package main

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/md4"
)

func md5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	byteStream := hasher.Sum(nil)
	return hex.EncodeToString(byteStream)

}

func md4Hash(input string) string {
	hasher := md4.New()
	hasher.Write([]byte(input))
	byteStream := hasher.Sum(nil)
	return hex.EncodeToString(byteStream)

}
