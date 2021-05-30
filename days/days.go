package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	format := "2.1.06" // dd.mm.yy
	date1, _ := time.Parse(format, os.Args[1])
	date2, _ := time.Parse(format, os.Args[2])
	days := date2.Sub(date1).Hours() / 24
	fmt.Println("days =", days)
}
