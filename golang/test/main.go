package main

import "fmt"

//声明全局变量，方法一二三均可

// fang'fa
func main() {
	//方法一：声明一个变量 默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("tyepe of a = %T\n", a)

	//方法二： 声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法三：在初始化的时候，可以省区数据类型，通过值自动匹配当前的变量数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四：（常用的方法）省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)
	var (
		ua int = 10
		ub     = "io"
	)
	fmt.Println(ua)
	fmt.Println(ub)
	const (
		au = iota * 100
		bu
		bi = iota * 1000
	)
}
