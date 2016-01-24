package main

import (
	"fmt"
	"go/build"
	"runtime"
)

func main() {
	fmt.Printf("Hello brave new world! It's working on %v %v/%v!\n", runtime.Version(), build.Default.GOOS, build.Default.GOARCH)
}