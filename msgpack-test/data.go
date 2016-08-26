package main

type Foo struct {
	Bar string  `msg:"a"` // 注意：使用a来做标签能减少序列化后的长度
	Baz float64 `msg:"b"`
}

//go:generate msgp
