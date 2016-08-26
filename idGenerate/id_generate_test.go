package idGenerate

import "fmt"
import "testing"

func TestNextId(t *testing.T) {
	var id string
	id = NextId()
	if len(id) != 16 {
		t.Fatal("len(id) != 16")
	}
	fmt.Println(id)

	id = NextId()
	if len(id) != 16 {
		t.Fatal("len(id) != 16")
	}
	fmt.Println(id)

	Init("ab")
	id = NextId()
	if len(id) != 16 {
		t.Fatal("len(id) != 16")
	}
	fmt.Println(id)
}

func BenchmarkNextId(b *testing.B) {
	var id, last string

	for i := 0; i < b.N; i++ {
		id = NextId()
		if id == last {
			b.Fatal("NextId return eq!!")
		}
		last = id
	}
}

func BenchmarkNextIdParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var id, last string
		for pb.Next() {
			id = NextId()
			if id == last {
				b.Fatal("NextId return eq!!")
			}
			last = id
		}
	})
}
