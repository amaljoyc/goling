package main

import "fmt"

func main() {
	helloWorld := getHelloWorld()
	fmt.Println(helloWorld)

	greetings := []string{"hello", getHelloWorld()}
	greetings = append(greetings, "hi")
	for i, greeting := range greetings {
		fmt.Println(i, greeting)
	}
}

func getHelloWorld() string {
	return "Hello World!!"
}
