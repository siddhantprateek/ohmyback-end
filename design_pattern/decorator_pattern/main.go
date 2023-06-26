package main

import (
	"fmt"
	"time"
)

func mainFun() {
	fmt.Println("main func")
	time.Sleep(1 * time.Second)
}

func additionalFun(a func()) {
	fmt.Printf("Starting function exection %s\n", time.Now())
	a()
	fmt.Printf("Ending function exection %s\n", time.Now())

}

func main() {
	additionalFun(mainFun)
}
