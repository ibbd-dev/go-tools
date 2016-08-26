package main

import (
	"fmt"
)

func main() {
	foo1 := Foo{Bar: "hello", Baz: 10.3}
	//foo2 := Foo{Bar: "world", Baz: 10.3}

	fmt.Printf("foo1: %v\n", foo1)
	//fmt.Printf("foo2: %v\n", foo2)

	// Here, we append two messages
	// to the same slice.
	data, _ := foo1.MarshalMsg(nil)
	//data, _ = foo2.MarshalMsg(data)

	// at this point, len(data) should be 0
	fmt.Println("len(data) =", len(data))

	// Now we'll just decode them
	// in reverse:
	//data, _ = foo2.UnmarshalMsg(data)
	data, _ = foo1.UnmarshalMsg(data)

	fmt.Printf("foo1: %v", foo1)
	//fmt.Printf("foo2: %v", foo2)
}
