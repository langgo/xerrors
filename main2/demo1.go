package main

import (
	"encoding/json"
	"fmt"
)

type User1 struct {
	Name  string
	Value interface{}
}

type Int32Value1 struct {
	Value int32 `json:"value"`
}

func demo1() {
	var bs []byte

	fmt.Println("序列化")
	{ // 序列化
		u := &User1{
			Name:  "Int32Value1",
			Value: &Int32Value1{Value: 123},
		}

		var err error
		bs, err = json.Marshal(u)
		if err != nil {
			panic(err)
		}

		fmt.Printf("marshal string: %q\n", string(bs))
	}

	fmt.Println("\n反序列化")
	{ // 反序列化 1
		var u *User1
		if err := json.Unmarshal(bs, &u); err != nil {
			panic(err)
		}

		fmt.Printf("%v %s %T\n", u, u.Name, u.Value)
	}

	fmt.Println("\n 解决方案 1")
	{ // 解决方案 1

		var u *User1
		if err := json.Unmarshal(bs, &u); err != nil {
			panic(err)
		}

		fmt.Printf("%v %s %T\n", u, u.Name, u.Value)

		v3 := &Int32Value1{
			Value: int32(u.Value.(map[string]interface{})["value"].(float64)),
		}

		fmt.Printf("%v\n", v3)

		// 手工进行赋值
		// 1. 极易出错
		// 2. 很机械化
	}

	fmt.Println("\n 解决方案 2")
	{ // 解决方案 2
		type User1_ struct {
			Name  string
			Value json.RawMessage // type Any []byte
		}

		var u_ *User1_
		if err := json.Unmarshal(bs, &u_); err != nil {
			panic(err)
		}

		fmt.Printf("%v %s %s %T\n", u_, u_.Name, u_.Value, u_.Value)

		switch u_.Name {
		case "Int32Value1":
			var v *Int32Value1
			if err := json.Unmarshal(u_.Value, &v); err != nil {
				panic(err)
			}

			fmt.Printf("%v %T\n", v, v)
		}

		// 1. 必须保证两个结构体（User1_ 和 User1）对应（靠人去保证一致性，是不理智的行为）
		// 2. 人工维护，类型和Name的映射
		// 3. 保证Name的唯一性
		// 4. 业务代码要手工控制各个Name相关的反序列化
	}
}
