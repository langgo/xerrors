package main

//go:generate protoc -I . --go_out=plugins=grpc:. user.proto
//go:generate thrift --gen go base.thrift
func main() {
	demo3()
}

// 编程语言 和 接口定义语言
