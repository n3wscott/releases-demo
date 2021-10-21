package v2

import (
	"fmt"

	v1 "github.com/n3wscott/releases-demo"
)

// Example is an interface at version 2.
type Example interface {
	// Foo sometimes you eat the bar with a name.
	Foo(string) int
	// Bar sometimes the bar eats you with a count.
	Bar(int, string)
	// Baz returns you the original bull.
	Baz() v1.Example
}

// Bull implements the v2 Example interface.
type Bull struct{}

func (b *Bull) Baz() v1.Example {
	return &v1.Bull{Greeting: "whelp"}
}

func (b *Bull) Foo(s string) int {
	return len(s)
}

func (b *Bull) Bar(i int, s string) {
	fmt.Println(s, i, "times")
}

var _ Example = (*Bull)(nil)
