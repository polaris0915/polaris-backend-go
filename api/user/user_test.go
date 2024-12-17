package user

import (
	"encoding/json"
	"testing"
)

// 这是js的结构体
type jsTestUserJson struct {
	UserName string
	Password string
}

// 这是go的结构体
type goTestUserJson struct {
	Name     string `json:"UserName"`
	Password string `json:"Password"`
}

/*
	json数据是结构化的字符串
	"{Name: "liuze", Password: "12345678"}"

	validator/v10
*/

func TestStructTag(t *testing.T) {
	// 以下的操作是js那边的操作
	jsUser := jsTestUserJson{
		UserName: "liuze",
		Password: "123456",
	}
	bytes, err := json.Marshal(jsUser)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("string(bytes): %v\n", string(bytes))
	t.Logf("bytes: %v\n", bytes)

	// js向后端发送数据 string(bytes)
	// gin处理数据
	goUser := &goTestUserJson{}
	if err := json.Unmarshal(bytes, goUser); err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("goUser: %+v\n", goUser)
}
