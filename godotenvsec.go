package godotenvsec

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// EncryptDecrypt runs a XOR encryption on the input string, encrypting it if it hasn't already been,
// and decrypting it if it has, using the key provided.
func EncryptDecrypt(input, key string) (output string) {
	kL := len(key)
	for i := range input {
		output += string(input[i] ^ key[i%kL])
	}
	return output
}

// RandStr generates random string of specified length
func RandStr(len int) string {
	buff := make([]byte, len)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	// Base 64 can be longer than len
	return str[:len]
}

// EnvEncode encodes an environment file
func EnvEncode() {
	fmt.Println("==================================")
	fmt.Println("== START: Encoding .env file    ==")
	fmt.Println("==================================")

	fmt.Println("1. Reading .env file...")

	v, err := ioutil.ReadFile(".env")

	if err != nil {
		fmt.Println("Could not read .env file")
	}

	fmt.Println("2. Encoding contents with random key ...")

	fileContents := string(v)

	key := RandStr(73) // min 1 byte, max is 256 bytes!
	out := EncryptDecrypt(fileContents, key)

	fmt.Println("3. Writing encoded content to .eenv file...")

	encContents := base64.StdEncoding.EncodeToString([]byte(out))
	encKey := base64.StdEncoding.EncodeToString([]byte(key))

	fileEncContents := encKey + "_" + encContents

	ioutil.WriteFile(".eenv", []byte(fileEncContents), 0664)

	fmt.Println("==================================")
	fmt.Println("== END: Encoding .env file      ==")
	fmt.Println("==================================")
}

// EnvDecode - Decodes an encoded environment file
func EnvDecode() {
	fmt.Println("==================================")
	fmt.Println("== START: Decoding .eenv file   ==")
	fmt.Println("==================================")

	fmt.Println("1. Reading .eenv file...")

	v, err := ioutil.ReadFile(".eenv")

	if err != nil {
		fmt.Println("Could not read .eenv file")
	}

	fmt.Println("2. Decoding contents ...")

	fileContents := string(v)

	splits := strings.Split(fileContents, "_")

	if len(splits) != 2 {
		fmt.Println("Incorrect file format")
	}

	key, err := base64.StdEncoding.DecodeString(splits[0])

	if err != nil {
		fmt.Println("Incorrect key format")
	}

	content, err := base64.StdEncoding.DecodeString(splits[1])

	if err != nil {
		fmt.Println("Incorrect content format")
	}

	fileDcContents := EncryptDecrypt(string(content), string(key))

	fmt.Println("3. Writing decoded content to .denv file...")

	ioutil.WriteFile(".denv", []byte(fileDcContents), 0664)

	fmt.Println("==================================")
	fmt.Println("== END: Decoding .eenv file     ==")
	fmt.Println("==================================")
}

func eenvToStr() string {
	v, err := ioutil.ReadFile(".eenv")

	if err != nil {
		fmt.Println("Could not read .eenv file")
	}

	fileContents := string(v)

	splits := strings.Split(fileContents, "_")

	if len(splits) != 2 {
		fmt.Println("Incorrect file format")
	}

	key, err := base64.StdEncoding.DecodeString(splits[0])

	if err != nil {
		fmt.Println("Incorrect key format")
	}

	content, err := base64.StdEncoding.DecodeString(splits[1])

	if err != nil {
		fmt.Println("Incorrect content format")
	}

	fileDecContents := EncryptDecrypt(string(content), string(key))

	return fileDecContents

}

func Init() {

	envEnc := flag.String("envenc", "no", "Encodes the environment file")
	envDec := flag.String("envdec", "no", "Decodes the environment file")
	flag.Parse()

	if strings.EqualFold(*envEnc, "yes") {
		EnvEncode()
		return
	}

	if strings.EqualFold(*envDec, "yes") {
		EnvDecode()
		return
	}

	env := eenvToStr()

	res, err := godotenv.Unmarshal(env)

	godotenv.Write(res, ".tempenv")

	err = godotenv.Load(".tempenv")
	
	os.Remove(".tempenv")
	
	if err != nil {
	    log.Fatal("Error loading .env file")
	}
}
