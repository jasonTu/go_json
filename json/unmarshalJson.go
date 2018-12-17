package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string "json:username"
	NickName string "json:nickname"
	Age int
	Birthday string
	Sex string
	Email string
	Phone string
}

func testStruct() (ret string, err error) {
	user1 := &User{
		UserName: "user1",
		NickName: "haha",
		Age: 18,
		Birthday: "2008/8/8",
		Sex: "Male",
		Email: "woods2001@126.com",
		Phone: "13770844271",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		err = fmt.Errorf("json.marshal failed, err:", err)
		return
	}
	ret = string(data)
	return
}

func testMap() (ret string, err error) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "male"
	var sub_m map[string]interface{}
	sub_m = make(map[string]interface{})
	sub_m["un"] = "user2"
	m["struct"] = sub_m
	
	data, err := json.Marshal(m)
	if err != nil {
		err = fmt.Errorf("json.marshal failed, err:", err)
		return
	}

	ret = string(data)
	return
}

func test2() {
	data, err := testMap()
	if err != nil {
		fmt.Println("test map failed")
	}

	var m map[string]interface{}
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Unmarshal failed, ", err)
		return
	}
	fmt.Println(m["struct"])
	fmt.Println(m)

	/*
	data, err = testStruct()
	if err != nil {
		fmt.Println("test struct failed")
	}
	var ms map[string]interface{}
	err = json.Unmarshal([]byte(data), &ms)
	if err != nil {
		fmt.Println("Unmarshal failed, ", err)
		return
	}
	fmt.Println(ms)
	*/
}

func main() {
	test2()
}
