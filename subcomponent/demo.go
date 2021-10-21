package main

import (
	"fmt"
	v1 "github.com/n3wscott/releases-demo"
	v2 "github.com/n3wscott/releases-demo/v2"
)

func main() {
	fmt.Println("-- v1 Bull --")
	b1 := &v1.Bull{}
	fmt.Println(b1.Foo())
	b1.Bar("that")

	fmt.Println("-- v2 Bull --")
	b2 := &v2.Bull{}
	fmt.Println(b2.Foo("this"))
	b2.Bar(5, "that")

	fmt.Println("-- v1 Bull via v2 Bull --")
	fmt.Println(b2.Baz().Foo())
	b2.Baz().Bar("another thing")
}
