package main

import "fmt"

func main()  {
	fmt.Println("hello world!") // hello
	var a = 1
	if a == 1 {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	myMap := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(myMap)
}
