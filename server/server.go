package main

import (
	"fmt"
	"runtime"

	"../lib/flags"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	flags.Parse()

	fmt.Println("port has value ", *flags.Port)
	fmt.Println("server has value ", *flags.Server)
}
