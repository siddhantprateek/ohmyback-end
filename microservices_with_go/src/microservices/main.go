package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %s\nArch: %s\n", runtime.GOOS, runtime.GOARCH)

}
