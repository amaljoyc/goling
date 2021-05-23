package main

import (
	"fmt"
)

func main()  {
	defer deferMe()

	go myFun()
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

	fmt.Println(p.getName())
	p.editName()
	fmt.Println(p.getName())

	var employee Employee = person {"Amal", 99,}
	fmt.Println(employee)

	generic("example")
	generic(10)

	sentByValue(p)
	fmt.Println(p)
	sentByReference(&p)
	fmt.Println(p)

	channelExample()
	channelBuffered()
	closeChannel()
	channelSelect()
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

func (p person) getName() string {
	return p.name
}

func (p *person) editName()  {
	p.name = "Lama"
}

type Employee interface {
	getName() string
}

func generic(x interface{})  {
	switch x.(type) {
	case string: fmt.Println(x.(string))
	case int: fmt.Println(x)
	}
}

func sentByValue(p person)  {
	p.name = "not changed"
}

func sentByReference(p *person)  {
	p.name = "new name"
}

func myFun() {
	fmt.Println("inside a goroutine")
}

func channelExample() {
	myChannel := make(chan int)
	go func() { myChannel <- 123 }()
	x := <-myChannel
	fmt.Println(x)
}

func channelBuffered() {
	myChannel := make(chan int, 1)
	myChannel <- 456
	fmt.Println(<-myChannel)
}

func closeChannel() {
	myChannel := make(chan int)
	close(myChannel)
	x, isOpen := <-myChannel
	fmt.Println(x, isOpen)
}

func channelSelect() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	exitChannel := make(chan interface{})

	go func() {
		channel1 <- 1
		channel2 <- 2
		exitChannel <- 0
	}()

	for  {
		select {
		case x := <-channel1:
			fmt.Println("channel1 value: ", x)
		case x := <-channel2:
			fmt.Println("channel2 value: ", x)
		case <-exitChannel:
			return
		}
	}
}