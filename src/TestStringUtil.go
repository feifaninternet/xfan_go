package main

/*
@Time : 2019/1/22 18:01 
@Author : xfan
@Desc:
*/

import (
	"fmt"

	"./util"
)

func main()  {
	fmt.Println(util.Reverse("!oG,olleH"))

	fmt.Println(util.MiddleChar("123f567"))
	fmt.Println(util.MiddleChar("123f5678"))
}
