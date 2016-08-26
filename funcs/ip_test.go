package funcs

import "fmt"
import "testing"

func TestIpv4(t *testing.T) {
	s := "127.0.0.1"
	ip, err := Ip2uint(s)
	if err != nil {
		t.Fatal("test")
	}
	fmt.Println("%s => %d", s, ip)

	s = "255.255.255.255"
	ip, err = Ip2uint(s)
	if err != nil {
		t.Fatal("test")
	}
	fmt.Println("%s => %d", s, ip)
}
