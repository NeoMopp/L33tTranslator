package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func sha1Hash(input string)  string{
	hasher := sha1.New()
	hasher.Write([]byte(input))
	byteStream := hasher.Sum(nil)
	return hex.EncodeToString(byteStream)
}

func sha256Hash(input string) string{
	hasher := sha256.New()
	hasher.Write([]byte(input))
	byteStream := hasher.Sum(nil)
	return hex.EncodeToString(byteStream)
}

func sha384Hash(input string) string{
	hasher := sha512.New384()
	hasher.Write([]byte(input))
	byteStream := hasher.Sum(nil)
	return hex.EncodeToString(byteStream)
}

func sha512Hash(input string) string{
	hasher := sha512.New()
	hasher.Write([]byte(input))
	byteStream := hasher.Sum(nil)
	return hex.EncodeToString(byteStream)
}
