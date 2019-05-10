package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func demo2() {
	var bs []byte

	fmt.Println("序列化")
	{ // 序列化
		any, err := ptypes.MarshalAny(&Int32Value2{
			Value: 13,
		})
		if err != nil {
			panic(err)
		}

		u := &User2{
			Name:  "Int32Value2",
			Value: any,
		}

		bs, err = proto.Marshal(u)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%q\n", string(bs))
	}

	fmt.Println("\n反序列化")
	{ // 反序列化
		var u = &User2{}
		if err := proto.Unmarshal(bs, u); err != nil {
			panic(err)
		}

		fmt.Printf("%v %T\n", u, u.Value)

		detail := &ptypes.DynamicAny{}
		if err := ptypes.UnmarshalAny(u.Value, detail); err != nil {
			panic(err)
		}

		v := detail.Message.(*Int32Value2)
		fmt.Printf("%v %T\n", v, v)
	}
}
