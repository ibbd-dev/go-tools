package funcs

import "testing"
import "bytes"
import "strings"

func _BenchmarkStringJoin(b *testing.B) {
	var a, c string
	a = "hello"
	for i := 0; i < b.N; i++ {
		c += " " + a
	}
	println(len(c))
}

func BenchmarkStrings(b *testing.B) {
	var a string
	var s []string
	a = "hello"

	for i := 0; i < b.N; i++ {
		//s = []string{a, c}
		s = append(s, a)
	}
	println(len(strings.Join(s, " ")))
}

func BenchmarkStringBuffer(b *testing.B) {
	var a string
	var c bytes.Buffer

	a = "hello"
	for i := 0; i < b.N; i++ {
		c.WriteString(a)
	}
	println(len(c.String()))

	c.Reset()
	c.WriteString(a)
	c.WriteString(a)
	println(c.String())
}
