package main

import (
	"fmt"
	"runtime"

	"../lib/flags"
)

func main() {
	flags.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("port has value ", *flags.Port)
	fmt.Println("server has value ", *flags.Server)
}
