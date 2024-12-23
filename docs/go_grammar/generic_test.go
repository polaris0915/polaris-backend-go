package go_grammar

import (
	"fmt"
	"testing"
)

// add 只能实现整型的加法
func addInt(a, b int) int {
	return a + b
}

func addFloat(a, b float64) float64 {
	return a + b
}

/*
func add(a, b ?) ?{

}
*/
type MyNumberType interface {
	int | float32 | float64 | int32
}

func add[T MyNumberType](a, b T) T {
	return a + b
}

//func add1(a, b any) any {
//	return a + b
//}

type cc interface {
	//print()
}

type test1 struct {
	A string
}

func (t *test1) print() {}

/*
any == interface{}
compare a b compare

如果一个结构体实现了一个你自己定义的接口，那么我就可以说，这个结构体属于这个接口的类型
*/

func TestRun(t *testing.T) {
	t.Logf("add(1.1, 2.2) = %v\n", add[float64](1.1, 2.2))
}

// 约束之后所有代码沙箱的行为
type ISandBox interface {
	Printf()
}

// 实现默认的代码沙箱
type SandBox struct {
	Msg string
}

func NewSandBox(msg string) ISandBox {
	return &SandBox{
		Msg: msg,
	}
}

func (s *SandBox) Printf() {
	fmt.Println("SandBox: " + s.Msg)
}

type OtherSandBox struct {
	SandBox
}

func NewOtherSandBox(msg string) ISandBox {
	return &OtherSandBox{
		SandBox{Msg: msg},
	}
}

func (s *OtherSandBox) Printf() {
	fmt.Println("OtherSandBox: " + s.Msg)
}

func TestSandBox(t *testing.T) {
	sandbox := NewOtherSandBox("正在执行代码...")
	sandbox.Printf()

	//var sandbox ISandBox
	//
	//// 以下的代码不用变
	//sandbox.Printf()
}
