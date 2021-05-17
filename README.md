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
- go get example.com/pkg -> downloads the package into global cache
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
```
var a = 1
if a == 1 {
	fmt.Println("if")
} else {
	fmt.Println("else")
}
```

### for loop
```
for i, card := range cards {
	fmt.Println(i, card)
}
```

### array and slice
```
a = [3]int{1,2,3} // array - limited size
a = []int{1,2,3} // slice - no limit of size
```

### map
```
myMap := map[string]int{
	"a": 1,
	"b": 2,
}
```

### struct
