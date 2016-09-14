package hmacSHA1

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"testing"
)

/*
字符串：22222
加密ekey：54MPDbYaVicfYPdjQOKsXfoq4mqVmKxS
签名ikey：czyr0wPXEEBT2ORprTjoNo7ZYqxkJiA4
加密后的字符串为：OUXiC1YBAABTM3Vdehl8FfbevGnCPceKe1G07A
*/
func TestDecode(t *testing.T) {
	tests := []struct {
		ekey         string
		ikey         string
		plain_int    int
		encrypt_text string
	}{
		{
			plain_int:    23489,
			ekey:         "54MPDbYaVicfYPdjQOKsXfoq4mqVmKxS",
			ikey:         "czyr0wPXEEBT2ORprTjoNo7ZYqxkJiA4",
			encrypt_text: "mrvD2lYBAABjMC4gJjNNVBVwlIbbjjpAgyIudg",
		},
		{
			plain_int:    22222,
			ekey:         "54MPDbYaVicfYPdjQOKsXfoq4mqVmKxS",
			ikey:         "czyr0wPXEEBT2ORprTjoNo7ZYqxkJiA4",
			encrypt_text: "OUXiC1YBAABTM3Vdehl8FfbevGnCPceKe1G07A",
		},
		{
			plain_int:    88888,
			ekey:         "54MPDbYaVicfYPdjQOKsXfoq4mqVmKxS",
			ikey:         "czyr0wPXEEBT2ORprTjoNo7ZYqxkJiA4",
			encrypt_text: "B-98KFcBAABlQTk1bh0WK6AjN3nNlrxZvS9peA",
		},
	}

	for _, test := range tests {
		hmac_sha1(t, test.plain_int, test.ekey, test.ikey, test.encrypt_text)
	}
}

func hmac_sha1(t *testing.T, plain_int int, ekey, ikey, encrypt_text string) {
	crypto := &Crypto{
		EKey: []byte(ekey),
		IKey: []byte(ikey),
	}

	unbase64_text, err := base64.RawURLEncoding.DecodeString(encrypt_text)
	fmt.Printf("encrypt_text = %s len = %d\n", encrypt_text, len(encrypt_text))
	fmt.Printf("unbase64_encrypt_text = %x len = %d\n", unbase64_text, len(unbase64_text))
	if err != nil {
		t.Fatal(err)
	}
	decrypt_text, err := crypto.Decrypt(unbase64_text)
	if err != nil {
		t.Fatal(err)
	}

	price_str := fmt.Sprintf("%x", decrypt_text)
	price64, err := strconv.ParseInt(price_str, 16, 32)
	if err != nil {
		t.Fatal(err)
	}

	price := int(price64)
	fmt.Printf("%s => %d\n", price_str, price)

	if price != plain_int {
		fmt.Println(decrypt_text)
		t.Fatal(errors.New("Decrypt text not eq plain text!"))
	}
}
