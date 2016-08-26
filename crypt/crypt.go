/*
AES加解密
@author Alex
*/
package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

var (
	key       []byte // 加密秘钥
	block     cipher.Block
	blockSize int
)

// 初始化
func Init(key_str string) {
	var err error
	key = []byte(key_str)
	block, err = aes.NewCipher(key)
	if err != nil {
		panic(errors.New("in crypt in init"))
	}

	blockSize = block.BlockSize()
}

// 将字符串加密，并转化为base64编码
func EncryptBase64(str []byte) (string, error) {
	result, err := Encrypt(str)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(result), nil
}

// 将base64编码的字符串解密
func DecryptBase64(str string) ([]byte, error) {
	result, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return Decrypt(result)
}

func Encrypt(origData []byte) ([]byte, error) {
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func Decrypt(crypted []byte) ([]byte, error) {
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
