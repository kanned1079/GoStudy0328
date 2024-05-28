package encrypt

import (
	"encoding/base64"
)

type Encryptor struct {
}

func (e *Encryptor) Encrypt(plainText string) (cipherText string) {
	var temp []byte
	temp = []byte(plainText)
	cipherText = string(base64.StdEncoding.EncodeToString(temp))
	return
}

func (e *Encryptor) Decrypt(cipherText string) (plainText string) {
	var temp []byte
	temp, _ = base64.StdEncoding.DecodeString(cipherText)
	plainText = string(temp)
	return
}
