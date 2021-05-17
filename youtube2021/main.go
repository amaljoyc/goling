package main

import "fmt"

func main()  {
	defer deferMe()
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

	p := person{
		name: "Amal",
		age: 99,
	}
	fmt.Println("person", p)

	hello, world := helloWorld()
	fmt.Println(hello, world)
}

type person struct {
	name string
	age int
}

func hello() string {
	return "hello"
}

func helloWorld() (string, string)  {
	return "hello", "world"
}

func deferMe()  {
	fmt.Println("defer me!")
}