package main

import (
	"fmt"

	"github.com/JonasHu1997/learn-golang/hello"
)

func init() {
	var (
		fuck int
		shit int
	)
	fmt.Println("main.go => init()", fuck, shit)
}

func main() {
	fmt.Println("main.go => main()")
	hello.SayHello()
}
