package main

import (
	"fmt"
)

/*
@Time : 2019/1/23 18:04 
@Author : xfan
@Desc: 循环与遍历
*/

//if，else的普通用法与java语言一样，go还有个特殊用法，可以初始化变量
func testIf(j int)  {
	if i:=j; i > 0 {
		fmt.Println(j)
	}
}

//switch,不需要用break退出一个case，默认退出
//switch可以初始化变量
func testSwitch(j int)  {
	i := j
	switch i {
	case 1,2,3:
		fmt.Println("1")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("0")
	}
}

//for循环，普通用法同C
//遍历列表
func testFor()  {
	aList := [] int{1, 2, 3}
	for index,value := range aList {
		fmt.Println(index)
		fmt.Println(value)
	}

	for _,value := range aList {
		fmt.Println(value)
	}
}

//遍历map
func testMap()  {
	//创建map,声明的是空值，还需要make创建
	var myMap map[string] string
	myMap = make(map[string] string)
	myMap["a"] = "mapA"
	myMap["b"] = "mapB"
	myMap["c"] = "mapC"
	for key,value := range myMap {
		fmt.Println(key)
		fmt.Println(value)
	}
}
func main()  {
	testIf(4)
	testSwitch(5)
	testFor()
	testMap()
}