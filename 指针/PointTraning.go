package main

import "fmt"

/*
@Time : 2019/1/24 18:22 
@Author : xfan
@Desc: Go中的指针
*/
//Go拥有指针，指针保存了值的内存地址，类型*T是指向T类型的指针。其零值为nil

func main()  {
	var p *int
	i := 42
	//%操作符会生成一个指向其操作数的指针
	p = &i
	//*操作符表示指针指向的底层值
	fmt.Println(*p) //通过指针p获取i
	*p = 21 //通过指针p设置i
	fmt.Println(i)

}