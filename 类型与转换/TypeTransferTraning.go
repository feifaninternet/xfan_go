package main

import "fmt"

/*
@Time : 2019/1/24 16:15 
@Author : xfan
@Desc: Go中的类型转换
*/

//表达式T(v)将值v转换成类型T
//与C不同的是，Go在不同类型的项之间赋值时需要显式转换，否则会报错: cannot use ** as type ** in argument ...
func main()  {

	i := 40
	f := float64(i)
	u := uint(f)
	fmt.Println(u)

	//类型推导
	//在声明一个变量不指定类型时，变量的类型由右值推导得出
	var p int
	j := p
	fmt.Printf("j is of type %T\n", j)

	//可以修改V的值，查看类型的变化
	v := "xfan"
	//v := 42
	fmt.Printf("v is of type %T\n", v)
}
