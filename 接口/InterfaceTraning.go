package main

import "fmt"

/*
@Time : 2019/1/23 17:35 
@Author : xfan
@Desc: 接口用interface定义，不需要继承，任何对象只要实现了接口中的方法就实现了此接口
*/

type ITeacher interface {
	Read(s string) string
	Write(s string) string
	Close()string
}

type IRead interface {
	Read(s string) string
}

type IWrite interface {
	Write(s string) string
}

type IClose interface {
	Close() string
}

//以上有四个接口，下面定义一个IStudent对象
type IStudent struct {
	
}

func (f *IStudent) Read(s string) string {
	return "read:" + s
}
func (f *IStudent) Write(s string) string{
	return "write:" + s
}

func (f *IStudent) Close() string{
	return "over"
}

func main()  {
	var student1 IRead = new(IStudent)
	fmt.Println(student1.Read("xfan"))

	var student2 IWrite = new(IStudent)
	fmt.Println(student2.Write("xfan"))

	var student3 IClose = new(IStudent)
	fmt.Println(student3.Close())

	var student4 ITeacher = new(IStudent)
	fmt.Println(student4.Write("parent"))
}

//接口组合，可像struct一样组成新接口

type IReadWrite interface {
	IRead
	IWrite
}
//等价于
type IReadWrite2 interface {
	Read(s string) string
	Write(s string) string
}

//空接口可接受任何类型
var v1 interface{} = 1
var v2 interface{}  = "xfan"



