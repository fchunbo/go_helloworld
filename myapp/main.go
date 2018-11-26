package main

import (
	"fmt"
	"helloworld/test" // 自定义包是相对于src的
)

func main() {
	fmt.Println("Hello World!!")

	result := test.TestSayHello("Rob")
	fmt.Println(result)
}
