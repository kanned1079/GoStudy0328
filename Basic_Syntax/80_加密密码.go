package main

import (
	"encoding/base64"
	"log"
)

func main() {
	msg := "debYUD#^&DFDYEUgd2"
	encode := base64.StdEncoding.EncodeToString([]byte(msg))
	log.Println(encode)

	decode, _ := base64.StdEncoding.DecodeString(encode)
	log.Println(string(decode))
}
