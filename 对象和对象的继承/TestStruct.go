package main

import "fmt"

/*
@Time : 2019/1/23 15:54
@Author : xfan
@Desc: 给类型加上方法便是对象
*/

type Customer int

func (c Customer) Bigger (s Customer) bool{
	return c > s
}

func main()  {
	var a Customer = 5
	fmt.Println(a.Bigger(4))
	fmt.Println(a.Bigger(6))


}

//一般使用struct来定义对象
type Person struct {
	name string
	age int
}

func (p Person)getName() string{
	return p.name
}

func (p *Person)setName(s string)  {
	p.name = s
}