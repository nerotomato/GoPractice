package main

import "fmt"

func main() {
	/**
	  Go编译器可以根据变量的值来自动推断其类型，类似于 Ruby 和 Python 这类动态语言，只不过它们是在运行时进行推断，而Go是在编译时就已经完成推断过程
	*/
	var a = 15
	var b = false
	var str = "hello, variable!"
	var (
		c    = 10
		d    = 20.5
		str2 = "hello, variable2!"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(str)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(str2)

}
