package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func main() {
	//hmac ,use sha1
	key := []byte("54MPDbYaVicfYPdjQOKsXfoq4mqVmKxS")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("23489"))
	fmt.Printf("%x\n", mac.Sum(nil))
}
