package main

import (
	"fmt"
	subcomponent "github.com/n3wscott/releases-demo/subcomponent/v2"
	v2 "github.com/n3wscott/releases-demo/v2"
)

func main() {
	fmt.Println("-- v2 Bull --")
	b1 := &v2.Bull{}
	fmt.Println(b1.Foo())
	b1.Bar("that")

	fmt.Println("-- v2 Bull --")
	b2 := &subcomponent.Bull{}
	fmt.Println(b2.Foo("this"))
	b2.Bar(5, "that")

	fmt.Println("-- v2 Bull via subcomponent Bull --")
	fmt.Println(b2.Baz().Foo())
	b2.Baz().Bar("another thing")
}
