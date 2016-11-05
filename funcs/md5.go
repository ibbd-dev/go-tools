package funcs

import (
	"crypto/md5"
	"fmt"
)

// Md5 32字节加密字符串
func Md5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// Md5L24 24字节加密字符串
func Md5L24(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))[4:28]
}

// Md5L16 16字节加密字符串
func Md5L16(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))[8:24]
}

// Md5L8 8字节加密字符串
func Md5L8(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))[12:20]
}
