package funcs

import (
	"testing"
)

func TestMd5(t *testing.T) {
	data := []byte("hello world--------------------------")
	m := Md5(data)
	println(m)
	if len(m) != 32 {
		t.Fatalf("not eq 32")
	}

	data = []byte("hello world-------3-------------------")
	m = Md5(data)
	println(m)
	if len(m) != 32 {
		t.Fatalf("not eq 32")
	}

	data = []byte("hello world-------3-------------------")
	m = Md5L24(data)
	println(m)
	if len(m) != 24 {
		t.Fatalf("not eq 24")
	}

	data = []byte("hello world-------3-------------------")
	m = Md5L16(data)
	println(m)
	if len(m) != 16 {
		t.Fatalf("not eq 16")
	}

	data = []byte("hello world-------3-------------------")
	m = Md5L8(data)
	println(m)
	if len(m) != 8 {
		t.Fatalf("not eq 8")
	}
}
