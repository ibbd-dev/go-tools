package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"testing"
)

const KEY = "def983c33bab071c8170683747be2e13"

func TestCrypt(t *testing.T) {
	crypto := &Crypt{}
	crypto.Init(KEY)
	fmt.Println("==========>")
	hello := "this is chinese!sldfkjsldfjsdklf sldkfjskdjflsdkjfls dfsdfjskdjf sjflskdjeworuiwekslfnsakjeoqiriflaksdnhgjrihginak"
	res, err := crypto.EncryptBase64([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println("---------->")

	res2, err := crypto.DecryptBase64(res)
	if err != nil {
		t.Fatal(err)
	}

	res3 := string(res2)
	if res3 != hello {
		t.Fatal(errors.New("error not eq"))
	}
}

func TestCrypt2(t *testing.T) {
	crypto := &Crypt{}
	crypto.Init(KEY)
	fmt.Println("==========>")
	hello := "=====2(*^98934(*&9))this is chinese!sldfkjsldfjsdklf sldkfjskdjflsdkjfls dfsdfjskdjf sjflskdjeworuiwekslfnsakjeoqiriflaksdnhgjrihginak"
	res, err := crypto.EncryptBase64([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println("---------->")

	res2, err := crypto.DecryptBase64(res)
	if err != nil {
		t.Fatal(err)
	}

	res3 := string(res2)
	if res3 != hello {
		t.Fatal(errors.New("error not eq"))
	}
}

/*
  示例 1
  明文:Hello World
  十六进制密钥:
  6F46756B794C5535777A3534494150326F72503155325177644E447267494843
  十六进制密文:
  2A0956BB2CF0AE98A10F0CA0DFDF396E51FF819D039C338A6F8185AB9209D9BF

  示例 2
  明文:5.5
  十六进制密钥:
  70465A507A31376F564E70344745303336625A30696B56794F7536316B565848
  十六进制密文:
  1CC76C4E999E87CFC06EB425D1672C591F28C4C9659C510A24B9BC147A37FB1A

  示例 3
  明文:www.sohu.com
  十六进制密钥:
  574A435543677446484A4F4C624263617063585A4D536E557A55516D4A675746
  十六进制密文:
  447DD39BEB3604ACB521D48DEDE870A01577DFC8D293DCC1B351EC72617E54AF
*/
func _TestCrypt3(t *testing.T) {
	tests := []struct {
		key          string
		plain_text   string
		encrypt_text string
	}{
		{
			key:          "6F46756B794C5535777A3534494150326F72503155325177644E447267494843",
			plain_text:   "Hello World",
			encrypt_text: "2A0956BB2CF0AE98A10F0CA0DFDF396E51FF819D039C338A6F8185AB9209D9BF",
		},
		{
			key:          "70465A507A31376F564E70344745303336625A30696B56794F7536316B565848",
			plain_text:   "5.5",
			encrypt_text: "1CC76C4E999E87CFC06EB425D1672C591F28C4C9659C510A24B9BC147A37FB1A",
		},
		{
			key:          "574A435543677446484A4F4C624263617063585A4D536E557A55516D4A675746",
			plain_text:   "www.sohu.com",
			encrypt_text: "447DD39BEB3604ACB521D48DEDE870A01577DFC8D293DCC1B351EC72617E54AF",
		},
	}

	for _, test := range tests {
		// 将16进制秘钥转化为字符数组
		bytes, err := hex.DecodeString(test.key)
		if err != nil {
			t.Fatal(err)
		}
		bytes_cal := fmt.Sprintf("%x", bytes)
		if bytes_cal != strings.ToLower(test.key) {
			t.Fatal(errors.New("bytes != test.key"))
		}

		crypto := &Crypt{}
		crypto.SetKey(bytes)

		// 将16进制密文转化为字符数组
		crypt_bytes, err := hex.DecodeString(test.encrypt_text)
		if err != nil {
			t.Fatal(err)
		}
		crypt_bytes_cal := fmt.Sprintf("%x", crypt_bytes)
		println(len(crypt_bytes))
		if crypt_bytes_cal != strings.ToLower(test.encrypt_text) {
			t.Fatal(errors.New("crypt_bytes != test.encrypt_text"))
		}

		// 解密
		block, err := aes.NewCipher(bytes)
		if err != nil {
			panic("aes.NewCipher: " + err.Error())
		}

		iv := crypt_bytes[:aes.BlockSize]
		decryptText := crypt_bytes[aes.BlockSize:]
		fmt.Printf("blocksize: %d\n", aes.BlockSize)

		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(decryptText, decryptText)

		fmt.Println(decryptText)
		fmt.Printf("%x\n", decryptText)
		fmt.Printf("%s\n", string(decryptText))

		// 解密
		//cal_plain_text, err := crypto.Decrypt(crypt_bytes)
		//if err != nil {
		//t.Fatal(err)
		//}

		//fmt.Println(cal_plain_text)
		//fmt.Println(test.plain_text)
	}

}
