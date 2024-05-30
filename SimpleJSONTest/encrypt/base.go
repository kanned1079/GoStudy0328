package encrypt

import (
	"encoding/base64"
)

var Base64Encryptor Encryptor

type Encryptor struct {
}

func (e *Encryptor) Base64Encrypt(plainText string) (cipherText string) {
	var temp []byte
	temp = []byte(plainText)
	cipherText = string(base64.StdEncoding.EncodeToString(temp))
	return
}

func (e *Encryptor) Base64Decrypt(cipherText string) (plainText string) {
	var temp []byte
	temp, _ = base64.StdEncoding.DecodeString(cipherText)
	plainText = string(temp)
	return
}
