package main

import (
	"errors"
	"fmt"

	"github.com/ibbd-dev/go-tools/crypt"
)

const KEY = "def983c33bab071c8170683747be2e13"

func main() {
	plain_text := "hello world!"
	fmt.Printf("plain_text: %s\n", plain_text)

	// 加密
	c := crypt.New(KEY)
	crypt_text, err := c.EncryptBase64([]byte(plain_text))
	if err != nil {
		panic(errors.New("encrypt error"))
	}
	fmt.Printf("crypt_text: %s\n", crypt_text)

	// 解密
	c = crypt.New(KEY)
	cal_plain_text_bytes, err := c.DecryptBase64(crypt_text)
	if err != nil {
		panic(errors.New("decrypt error"))
	}
	cal_plain_text := string(cal_plain_text_bytes)
	if cal_plain_text != plain_text {
		panic(errors.New("cal_plain_text != plain_text"))
	}
	fmt.Printf("after plain_text: %s\n", cal_plain_text)
}
