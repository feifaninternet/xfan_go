package main

/*
@Time : 2019/1/22 18:01 
@Author : xfan
@Desc:
*/

import (
	"fmt"

	charUtil "./util"
)

func main() {
	//大写开头的可以引用，类似公私有
	fmt.Println(charUtil.Reverse("!oG,olleH"))

	fmt.Println(charUtil.MiddleChar("123f567"))
	fmt.Println(charUtil.MiddleChar("123f5678"))

	//匿名函数，赋值
	f := func(x, y int) int {
		return x + y
	}(2, 3)

	fmt.Println(f)
}
