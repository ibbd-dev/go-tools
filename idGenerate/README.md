# 发号器

用于分布式生成唯一ID, 例如竞价ID的生成

多次测试, 并发300w/s, 不会冲突


## Usages

```go
package main

import "git.ibbd.net/dsp/go-tools/idGenerate"

func main() {
    var bid_id string = idGenerate.NextId()
}
```

## 性能测试

```
// 环境: 4C4G ubuntu16.04 PC
// ./benchmark
// 或者 go test -test.bench=".*"
BenchmarkNextId-4               10000000           182 ns/op
BenchmarkNextIdParallel-4       10000000           126 ns/op
```

