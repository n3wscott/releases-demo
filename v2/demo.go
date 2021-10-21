package subcomponent

import (
	"fmt"
)

// Example is an interface.
type Example interface {
	// Foo sometimes you eat the bar.
	Foo() int
	// Bar sometimes the bar eats you.
	Bar(string)
}

// Bull implements the v1 Example interface.
type Bull struct{ Greeting string }

func (b *Bull) greeting() string {
	if len(b.Greeting) > 0 {
		return b.Greeting
	}
	return "How you doing there,"
}

func (b *Bull) Foo() int {
	return len(b.greeting()) + 42
}

func (b *Bull) Bar(s string) {
	fmt.Println(b.greeting(), s, "once")
}

var _ Example = (*Bull)(nil)
