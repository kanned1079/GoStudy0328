package user

import (
	"encoding/base64"
	"log"
)

type Encryptor struct {
}

// Encrypt 用于对用户密码进行base64加密
func (e *Encryptor) Encrypt(PlainPassword string) (EncryptedPassword string) {
	var temp []byte
	temp = []byte(PlainPassword)
	return base64.StdEncoding.EncodeToString(temp)
}

// Decrypt 用于对用户密码进行base解密获得明文进行比对
func (e *Encryptor) Decrypt(EncryptedPassword string) (PlainPassword string) {
	var temp []byte
	temp, err := base64.StdEncoding.DecodeString(EncryptedPassword)
	if err != nil {
		log.Println("base64 decode error:", err)
		return
	}
	return string(temp)

}
