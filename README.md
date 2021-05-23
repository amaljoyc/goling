# goling

set of go programs

### basic commands
- go build main.go -> will just compile and create executable
- go run main.go -> will compile and run the program rightaway (without creating an exe)
- go fmt -> formats all code in each file of the current dir
- go install -> compiles and installs a package
- go get -> downloads the source code of someone else's package
- go test -> runs all tests in the projects

- go mod init <full path of the go module to init>
    - eg. go mod init amaljoyc.com/my-first-module
    - eg. go mod init github.com/amaljoyc/my-first-module (if I am keeping it in my github repo my-first-module)
- go build (do this inside your project ie. inside youtube2021, which will build all go files in that project)
- GOOS=linux GOARCH=amd64 go build (to build for a specific os or arch, by default it will be built for the os/arch which is used for build)
- go get github.com/google/pprof -> to get the dependency from this repo (it will also add this depen into go.mod file).
- go get example.com/pkg -> downloads the package into global cache (same as example just above)
- go get -u example.com/pkg -> gets and upgrades the package
- go list -u -m all -> view available dependency upgrades
- go test -> to run tests
- go vet -> runs the linter
- go tool -> builtin profiling


### files
- go.mod -> required packages to run the project
- go.sum -> lock versions for reproducible builds

### packages
- package main -> executable package, will give an exe
- package example or package hello -> reuable package (for dependency/lib)

### OOP vs type/receivers
In go, instead of creating classes and objects, we create types and receivers to tie up a type to different functions/methods.

### Special Convention
- Functions and Variables starting with Capital letter are exported (visible from another package)
- DoSomething() vs doSomething()

### if else
```go
var a = 1
if a == 1 {
	fmt.Println("if")
} else {
	fmt.Println("else")
}
```

### for loop
```go
for i, card := range cards {
	fmt.Println(i, card)
}
```

### array and slice
```go
a = [3]int{1,2,3} // array - limited size
a = []int{1,2,3} // slice - no limit of size
```

### map
```go
myMap := map[string]int{
	"a": 1,
	"b": 2,
}
```

### struct
```go
type person struct {
	name string
	age int
}

p := person{
	name: "Amal",
	age: 99,
}
```

### function
```go
func hello() string {
	return "hello"
}

func helloWorld() (string, string)  {
	return "hello", "world"
}
```

### function as first class citizen
- you can pass around function like a variable or another function argument.
```go
hello, world := helloWorld()
fmt.Println(hello, world)
```

### defer a function
- we can defer the execution or call of a function to the last part...
```go
func main()  {
	defer deferMe()
	fmt.Println("hello")
}

func deferMe()  {
	fmt.Println("defer me!")
}	
```
- here, "hello" will be printed first, then "defer me!" will be printed as that method is only executed by the end of main() execution

### function receivers
- receivers are used to attach funtions to a struct.
- by doing this, you can call those functions using the struct object/instance.
- you pass the receiver just before the function name
```go
// pass by value - can only read p's properties, cannot edit it (p is immutable)
func (p person) getName() string {
	return p.name
}

// pass by reference - can update the properties of p (p is mutable)
func (p *person) editName()  {
	p.name = "Lama"
}
```

### interfaces
- used to define functions
- we could save a struct value into an interface variable if that struct has all functions of that interface
```go
type Employee interface {
	getName() string
}

func (p person) getName() string {
	return p.name
}

var employee Employee = person {"Amal", 99,}
```

### generic interface{} variables
- we can use `interface{}` to represent any variable type
- to know/infer the real type of such variables, we can use variableName.(type)
```go
func generic(x interface{})  {
	switch x.(type) {
	case string: fmt.Println(x.(string))
	case int: fmt.Println(x)
	}
}

generic("example")
generic(10)
```

### pointers
- most of the types in go are "sent by value" when passed as a function argument -> so we cannot update them as they are immutable
- in order to update such types, we need to use pointers which are "sent by reference"
- while using pointers, the function arg is represented using `*x` while the variable passed during function call is `&x`
- there are still some types in go which are by default "sent by reference" and hence no need to use pointers for them - eg, slices, maps, channels
```go
func sentByValue(p person)  {
	p.name = "not changed"
}

func sentByReference(p *person)  {
	p.name = "new name"
}

sentByValue(p)
sentByReference(&p)
```

### goroutines
- it is a construct built on top of threads
- it is lot more lightweight than threads (upto 250 times lighter)
- to make any function call asynchronously, just call that function with `go` keyword and it will run as a goroutine
```go
func myFun() {
	fmt.Println("inside a goroutine")
}

go myFun()
```

### channels
- channels allow to pass information in a thread safe way
- so it is ideal to send data in and out of goroutines
- a channel is created using the `make` function
- you can't take something out of a channel until something is sent to a channel, otherwise it will create a deadlock error.
- also you can't put something into a channel until something is taken from it
- essentially this means that, the read and write into a channel happens at the same time (or is synchronised) for it to work 
- channel read and write are both blocking operations
```go
func channelExample() {
	myChannel := make(chan int)
	go func() { myChannel <- 123}()
	x := <-myChannel
	fmt.Println(x)
}
```

### buffered channel
- the regular channels (described above) are called unbuffered channels
- buffered channels creates a buffer for a channel until it is blocked, ie. you can tell how many items the channel can take without blocking. After that limit is reached, it will start blocking.
- in below eg. we are creating a buffer of 1 item, so we can add an int into that channel without blocking (hence no need to wrap it into a goroutine)
```go
func channelBuffered() {
	myChannel := make(chan int, 1)
	myChannel <- 456
	fmt.Println(<-myChannel)
}
```

### close a channel
- to close a channel, we use the `close` function and pass in the name of the channel.
- a normal channel is blocking, but a closed channel is not blocking
- we can still read from a closed channel, but we cannot write to a closed channel
- if we read from a closed channel, we get the default value of the channel type, ie `0` for int, `false` for bool etc.
- as a norm, it is the responsibility of the producer to close the channel, and not the responsibility of the one who listens to the channel.
- you can also take a second argument from a channel while reading a value, this second arg returns if the channel is open or not (true/false)
```go
func closeChannel() {
	myChannel := make(chan int)
	close(myChannel)
	x, isOpen := <-myChannel
	fmt.Println(x, isOpen)
}
```

### select
- can be used if we want to deal with multiple channels at once
- which means we can read or write from multiple channels whichever is available at that point in time.
```go
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
```

### error handling
- there is no default exception or error handling support in go
- rather it is advised to handle errors as return values (or additional return values) from a function.

### panic and recover
- it is not intended to use panic as a form of error/exception mechanism
- rather use `panic` to only indicate "impossible" situations
- `recover()` returns the value which is passed by the panic function
```go
func panicExample()  {
	defer recoverExample()
	fmt.Println("Starting..")
	panic("Panicking..")
}

func recoverExample() {
	panicValue := recover()
	fmt.Println(panicValue)
	fmt.Println("Recovering..")
}
```

### no inheritence, but we have composition support
- we can inherit types not using inheritence but rather with composition
```go
type Address struct {
	 city string
}

type Office struct {
	name string
	Address
}

office := Office{}
office.name = "My Office"
office.city = "London"
fmt.Println(office)
```