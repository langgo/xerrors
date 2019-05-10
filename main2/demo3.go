package main

import (
	"encoding/json"
	"fmt"
)

type User3 struct {
	Name  string
	Value *Any
}

type Int32Value3 struct {
	Value int32 `json:"value"`
}

func demo3() {
	var bs []byte

	fmt.Println("序列化")
	{ // 序列化
		var err error
		any, err := MarshalJSONAny(&Int32Value1{Value: 123})
		if err != nil {
			panic(err)
		}

		u := &User3{
			Name:  "Int32Value1",
			Value: any,
		}

		bs, err = json.Marshal(u)
		if err != nil {
			panic(err)
		}

		fmt.Printf("marshal string: %q\n", string(bs))
	}

	fmt.Println("序列化")
	{ // 反序列化
		var u *User3
		if err := json.Unmarshal(bs, &u); err != nil {
			panic(err)
		}

		fmt.Printf("%v %T\n", u, u.Value)

		var v *Int32Value3
		if err := UnmarshalJSONAny(u.Value, &v); err != nil {
			panic(err)
		}

		fmt.Printf("%v %T\n", v, v)
	}
}
