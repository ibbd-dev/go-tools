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
