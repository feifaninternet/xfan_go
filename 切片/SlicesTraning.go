package main

import "fmt"

/*
@Time : 2019/1/28 15:16 
@Author : xfan
@Desc: 切片
*/

//通过指定由冒号分隔的两个索引形成切片

func testCreate() {
	primes := [5] int{1, 2, 3, 4, 5}
	var s = primes [1 : 4]
	fmt.Println(s)
}

//切片就像对数组的引用，更改切片的元素会修改其基础数组的相应元素
func test() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

//省略下限(默认0)和上限(切片长度)的写法
func test2()  {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

//切片的长度和容量，切片的长度是它包含的元素数，切片的容量是基础数组中的元素数
//只要具有足够的容量，便可以重新切片来延长切片的长度，当然，扩展超出容量则会报错

func test3() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 切片使其长度为0
	s = s[:0]
	printSlice(s)

	// 延长它的长度
	s = s[:4]
	printSlice(s)

	// 删除前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//切片的零值是nil

//
func main()  {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {

	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}



