# 使用hmac-sha1进行加解密

## 使用

```go
package main

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/ibbd-dev/go-tools/hmac-sha1"
)

func main() {
	//hmac ,use sha1
	test := struct {
		ekey         string
		ikey         string
		plain_int    int
		encrypt_text string
	}{
		plain_int:    23489,
		ekey:         "54MPDbYaVicfYPdjQOKsXfoq4mqVmKxS",
		ikey:         "czyr0wPXEEBT2ORprTjoNo7ZYqxkJiA4",
		encrypt_text: "mrvD2lYBAABjMC4gJjNNVBVwlIbbjjpAgyIudg",
	}

	crypto := &hmacSHA1.Crypto{
		EKey: []byte(test.ekey),
		IKey: []byte(test.ikey),
	}
	unbase64_text, err := base64.RawStdEncoding.DecodeString(test.encrypt_text)
	if err != nil {
		//TODO
	}

	decrypt_bytes, err := crypto.Decrypt(unbase64_text)
	if err != nil {
		//TODO
	}

	price_str := fmt.Sprintf("%x", decrypt_bytes)
	price64, err := strconv.ParseInt(price_str, 16, 32)
	if err != nil {
		//TODO
	}

	price := int(price64)
	fmt.Printf("%s => %d\n", price_str, price)
	if price == test.plain_int {
		fmt.Println("====> OK!")
	}
}
```

## TODO

- 加密
