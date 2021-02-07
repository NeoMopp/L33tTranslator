package main

import (
	"flag"
	"fmt"
)

type Operation string

const (
    Leet Operation = "l33t"
    Md5 Operation = "md5"
    Md4 Operation = "md4"
    Sha1 Operation = "sha1"
    Sha256 Operation = "sha256"
    Sha384 Operation = "sha384"
    Sha512 Operation = "sha512"
)

func parseArguments() (word string, operation Operation) {
	wordPtr := flag.String("word", "foo", "the string you want converted")
	operationPtr := flag.String("operation", "default", "the operation you want to perform on the string")
	flag.Parse()
	//Can't ever be null as far as I can tell
	return *wordPtr, Operation(*operationPtr)
}

func main() {
	word, operation := parseArguments()

	var result string

    switch operation {
    case Leet:
		result = leetHash(word)
    case Md5:
		result = md5Hash(word)
    case Md4:
		result = md4Hash(word)
    case Sha1:
		result = sha1Hash(word)
    case Sha256:
		result = sha256Hash(word)
    case Sha384:
		result = sha384Hash(word)
    case Sha512:
		result = sha512Hash(word)
    default:
		result = string("Not a valid Operation: " + operation)
    }

    fmt.Println(result)

}
