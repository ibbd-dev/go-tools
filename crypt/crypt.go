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
)

type Crypt struct {
	key       []byte // 加密秘钥
	block     cipher.Block
	blockSize int
}

func New(key string) *Crypt {
	c := &Crypt{}
	c.SetKey([]byte(key))
	return c
}

// 初始化
func (c *Crypt) Init(key string) {
	c.SetKey([]byte(key))
}

// 设置key
func (c *Crypt) SetKey(key []byte) {
	var err error
	c.key = key
	c.block, err = aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	c.blockSize = c.block.BlockSize()
}

// 将字符串加密，并转化为base64编码
func (c *Crypt) EncryptBase64(str []byte) (string, error) {
	result, err := c.Encrypt(str)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(result), nil
}

// 将base64编码的字符串解密
func (c *Crypt) DecryptBase64(str string) ([]byte, error) {
	result, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return c.Decrypt(result)
}

// 加密
func (c *Crypt) Encrypt(origData []byte) ([]byte, error) {
	origData = pkcs5Padding(origData, c.blockSize)
	// origData = ZeroPadding(origData, c.block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(c.block, c.key[:c.blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 解密
func (c *Crypt) Decrypt(crypted []byte) ([]byte, error) {
	//c.blockSize := c.block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(c.block, c.key[:c.blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	//print("pkcs5UnPadding: ")
	//println(length)
	unpadding := int(origData[length-1])
	//println(unpadding)
	//print("=== ")
	return origData[:(length - unpadding)]
}
