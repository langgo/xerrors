package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type User struct {
	Name  string
	Value interface{}
}

//go:generate protoc -I . --go_out=plugins=grpc:. user.proto
func main() {
	demo2()
}

func demo1() {
	u := &User{
		Name:  "int32",
		Value: &wrappers.Int32Value{Value: 13},
	}

	bs, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	var u2 *User
	if err := json.Unmarshal(bs, &u2); err != nil {
		panic(err)
	}

	fmt.Printf("%v %T\n", u2, u2.Value)
}

func demo2() {
	value := wrappers.Int32Value{Value: 13}
	any, err := ptypes.MarshalAny(&value)
	if err != nil {
		panic(err)
	}

	u := &User2{
		Name:  "int32",
		Value: any,
	}

	bs, err := proto.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", string(bs))

	var u2 = &User2{}
	if err := proto.Unmarshal(bs, u2); err != nil {
		panic(err)
	}

	fmt.Printf("%#v %T\n", u2, u2.Value)

	detail := &ptypes.DynamicAny{}
	if err := ptypes.UnmarshalAny(u2.Value, detail); err != nil {
		panic(err)
	}

	v := detail.Message.(*wrappers.Int32Value)
	fmt.Printf("%v %T\n", v, v)
}
