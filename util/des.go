package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"management-platform/config"
)

// DesEncrypt Des加密
func DesEncrypt(origData []byte) ([]byte, error) {
	key := []byte(config.Config.Des.Key)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypt := make([]byte, len(origData))
	blockMode.CryptBlocks(crypt, origData)
	return crypt, nil
}

// DesDecrypt Des解密
func DesDecrypt(crypt []byte) ([]byte, error) {
	key := []byte(config.Config.Des.Key)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypt))
	blockMode.CryptBlocks(origData, crypt)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
