package main

import "fmt"

/*
@Time : 2019/1/23 16:57 
@Author : xfan
@Desc: 继承效果的实现
*/

type Base struct {
	name string
}

func (b Base) getName() string{
  return "Son have get base`s name"
}

type Son struct {
	Base
}

func main()  {
	var s Son
	fmt.Println(s.getName())
}
