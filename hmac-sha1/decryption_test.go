package hmacSHA1

import (
	"encoding/base64"
	"errors"
	"fmt"
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
		plain_text   string
		encrypt_text string
	}{
		{
			plain_text:   "23489",
			ekey:         "54MPDbYaVicfYPdjQOKsXfoq4mqVmKxS",
			ikey:         "czyr0wPXEEBT2ORprTjoNo7ZYqxkJiA4",
			encrypt_text: "mrvD2lYBAABjMC4gJjNNVBVwlIbbjjpAgyIudg",
		},
		{
			plain_text:   "22222",
			ekey:         "54MPDbYaVicfYPdjQOKsXfoq4mqVmKxS",
			ikey:         "czyr0wPXEEBT2ORprTjoNo7ZYqxkJiA4",
			encrypt_text: "OUXiC1YBAABTM3Vdehl8FfbevGnCPceKe1G07A",
		},
	}

	for _, test := range tests {
		hmac_sha1(t, test.plain_text, test.ekey, test.ikey, test.encrypt_text)
	}
}

func hmac_sha1(t *testing.T, plain_text, ekey, ikey, encrypt_text string) {
	crypto := &Crypto{
		EKey: []byte(ekey),
		IKey: []byte(ikey),
	}

	fmt.Printf("plain_text = %s\n", plain_text)
	unbase64_text, err := base64.RawStdEncoding.DecodeString(encrypt_text)
	fmt.Printf("encrypt_text = %s len = %d\n", encrypt_text, len(encrypt_text))
	fmt.Printf("unbase64_encrypt_text = %s len = %d\n", unbase64_text, len(unbase64_text))
	if err != nil {
		t.Fatal(err)
	}
	decrypt_text, err := crypto.Decrypt(unbase64_text)
	if err != nil {
		fmt.Println(decrypt_text)
		t.Fatal(err)
	}

	decrypt_text_str := string(decrypt_text)
	fmt.Printf("decrypt_text_str = %s\n", decrypt_text_str)
	if decrypt_text_str != plain_text {
		fmt.Println(decrypt_text)
		t.Fatal(errors.New("Decrypt text not eq plain text!"))
	}
}
