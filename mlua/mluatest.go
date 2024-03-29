package mlua

import (
	"fmt"
)

// LuaTest : lua测试
type LuaTest struct {
}

func test3(L *State) int {
	fmt.Print("cccccc\n")
	return 0
}

func test2(L *State) int {
	fmt.Print("ssssssss\n")
	return 0
}

type act struct {
	A int
	B int
	C int
}

// TestGOStructUserData : 测试UserData
func (T *LuaTest) TestGOStructUserData() {
	L := NewState()
	L.OpenLibs()
	defer L.Close()

	//结构及 函数注册测试完成---------------------------
	td := &act{1, 2, 3}

	L.PushUserGoStruct(td)

	var bbb act
	L.ToUserGoStruct(-1, &bbb)

	fmt.Print(td.A, bbb.A)

}

// TestRegisterFunc 测试注册函数
func (T *LuaTest) TestRegisterFunc() {
	L := NewState()
	L.OpenLibs()
	defer L.Close()

	L.Register("test2", test2)
	L.Register("test3", test3)

	L.DoString("test2() test3()")
}

// TestLuaReg : 测试LuaReg
func (T *LuaTest) TestLuaReg() {
	L := NewState()
	L.OpenLibs()
	defer L.Close()

	af := make([]LuaReg, 1, 1)
	af[0].Name = "test2"
	af[0].Func = test2

	L.NewTable()

	L.SetFuncs(af, 0)

	L.SetGlobal("mm")

	L.DoString("print('ooooo') mm.test2()")
}
