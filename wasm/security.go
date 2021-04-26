package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func EncryptAES_CBC(src, key, iv string) (string, error) {
	data := []byte(src)
	keyByte, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	data = PKCS5Padding(data, block.BlockSize())
	ivByte, err := hex.DecodeString(iv)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, ivByte)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out), nil
}

func DecryptAES_CBC(src, key, iv string) (string, error) {
	keyByte, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	data, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	ivBye, err := hex.DecodeString(iv)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, ivBye)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext), nil
}

func MD5(v string)string{
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}