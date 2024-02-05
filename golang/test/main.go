package main

import "fmt"

//声明全局变量，方法一二三均可

// fang'fa
func main() {
	////方法一：声明一个变量 默认的值是0
	//var a int
	//fmt.Println("a = ", a)
	//fmt.Printf("tyepe of a = %T\n", a)
	//
	////方法二： 声明一个变量，初始化一个值
	//var b int = 100
	//fmt.Println("b = ", b)
	//fmt.Printf("type of b = %T\n", b)
	//
	//var bb string = "abcd"
	//fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)
	//
	////方法三：在初始化的时候，可以省区数据类型，通过值自动匹配当前的变量数据类型
	//var c = 100
	//fmt.Println("c = ", c)
	//fmt.Printf("type of c = %T\n", c)
	//
	//var cc = "abcd"
	//fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)
	//
	////方法四：（常用的方法）省去var关键字，直接自动匹配
	//e := 100
	//fmt.Println("e = ", e)
	//fmt.Printf("type of e = %T\n", e)
	//var (
	//	ua int = 10
	//	ub     = "io"
	//)
	//fmt.Println(ua)
	//fmt.Println(ub)
	//const (
	//	au = iota * 100
	//	bu
	//	bi = iota * 1000
	//)
	//slice1 := []int{0, 1, 2, 3, 4}
	//slice2 := slice1[:]
	//slice3 := slice2[:3]
	//for k, v := range slice3 {
	//	fmt.Printf("k = %d, v = %d\n", k, v)
	//}
	//mp1 := make(map[string]int)
	//mp1["u"] = 1
	//fmt.Println(mp1["u"], "\n")
	//str := "abc"
	//for k := range str {
	//	fmt.Printf("type of k: %T\n", k)
	//	//fmt.Printf("type of v: %T\n", v)
	//}
	//fmt.Printf("type of str %T\n", str[1])
	//
	//s := "ADOBECODEBANC"
	//t := "ABC"g
	s := "abcd"
	//t := "aa"
	//start, left, right := 0, 0, 0
	fmt.Printf("s[1:4]%s\n", s[1:4])
	fmt.Printf("s[0:4]%s\n", s[0:4])
	fmt.Printf("s[0:3]%s\n", s[0:3])
	//mp1 := make(map[uint8]int)
	//for i := 0; i < len(t); i++ {
	//	mp1[t[i]]++
	//}
	//ls := len(s)
	//lm := len(mp1)
	//mp2 := make(map[uint8]int)
	//valid := 0
	//length := 100000
	//for right < ls {
	//	c := s[right]
	//	if _, ok := mp1[c]; ok {
	//		mp2[c]++
	//		if mp2[c] == mp1[c] {
	//			valid++
	//		}
	//	}
	//
	//	for valid == lm {
	//		if right-left < length {
	//			length = right - left
	//			start = left
	//		}
	//		c = s[left]
	//		left++
	//		if _, ok := mp2[c]; ok {
	//			if mp2[c] == mp1[c] {
	//				valid--
	//			}
	//			mp2[c]--
	//		}
	//	}
	//	right++
	//}
	//fmt.Print(s[start : start+length+1])

}
