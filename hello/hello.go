package hello

import "fmt"

func init() {
	fmt.Println("hello/hello.go => init()")
}

func SayHello() (int, error) {
	val, err := fmt.Println("hello/hello.go => sayHello()")
	if err != nil {
		return -1, err
	}
	return val, nil
}
