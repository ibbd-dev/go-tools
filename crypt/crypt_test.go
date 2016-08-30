package crypt

import (
	"errors"
	"fmt"
	"testing"
)

const KEY = "def983c33bab071c8170683747be2e13"

func TestCrypt(t *testing.T) {
	Init(KEY)
	fmt.Println("==========>")
	hello := "this is chinese!sldfkjsldfjsdklf sldkfjskdjflsdkjfls dfsdfjskdjf sjflskdjeworuiwekslfnsakjeoqiriflaksdnhgjrihginak"
	res, err := EncryptBase64([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println("---------->")

	res2, err := DecryptBase64(res)
	if err != nil {
		t.Fatal(err)
	}

	res3 := string(res2)
	if res3 != hello {
		t.Fatal(errors.New("error not eq"))
	}
}

func TestCrypt2(t *testing.T) {
	Init(KEY)
	fmt.Println("==========>")
	hello := "=====2(*^98934(*&9))this is chinese!sldfkjsldfjsdklf sldkfjskdjflsdkjfls dfsdfjskdjf sjflskdjeworuiwekslfnsakjeoqiriflaksdnhgjrihginak"
	res, err := EncryptBase64([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println("---------->")

	res2, err := DecryptBase64(res)
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
func TestCrypt3(t *testing.T) {
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
		Init(test.key)
		cal_encrypt_text, err := Encrypt([]byte(test.plain_text))
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(cal_encrypt_text)
		fmt.Println(test.encrypt_text)
	}

}
