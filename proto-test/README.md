# Proto

## 使用前的准备：	
 - go get -u github.com/golang/protobuf/proto
 - go get -u github.com/golang/protobuf/protoc-gen-go  (看情况)
 + protobuf 软件安装

## .Proto文件中的接口
 - Constants 	常量
 - Enums 		枚举
 - Oneof		 	择一 (单一出现)
 - message 		复合结构体
	- required	必需 
	- optional	可选
	- repeated	数组
	- group		结构体
	- oneof 		单一出现
	
## .pd.go文件的生成

```sh
$protoc --go_out=. *.proto 
```	

## go程序使用:

###引用import .pd.go文件地址
###方法引用：
	+ Marshal 		编码
	+ UnMarshal		解码
	
## 参考文档地址：
	- https://godoc.org/github.com/golang/protobuf/proto
	- http://www.minaandrawos.com/2014/05/27/practical-guide-protocol-buffers-protobuf-go-golang/