package main

import "fmt"

/*
@Time : 2019/1/28 15:16 
@Author : xfan
@Desc: 切片
*/

//将新元素附加到切片 append函数

func main()  {
	var s [] int
	printSlice(s)

	//append 适用于零片
	s = append(s, 0)
	printSlice(s)

	//切片根据需要增长
	s = append(s, 1)
	printSlice(s)

	//可以一次增加多个元素
	s = append(s, 2,3,4)
	printSlice(s)	
	
}

func printSlice(s [] int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


