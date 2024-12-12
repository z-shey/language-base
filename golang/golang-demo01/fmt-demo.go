package main

import "fmt"

func main() {
	// 输出字符串
	fmt.Println("Hello, World!")
	/*
		fmt中Println()输出的内容会自动换行
	*/

	fmt.Print("Hello, World!")
	/*
		fmt中Print()输出的内容不会自动换行
	*/

	fmt.Printf("Hello, %s!\n", "World")
	/*
		fmt.Printf()可以格式化输出字符串
		%s表示字符串
	*/
}
