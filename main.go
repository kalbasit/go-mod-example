package main

import (
	"fmt"

	a "github.com/kalbasit/dep-a"
	b "github.com/kalbasit/dep-b"
)

func main() {
	v := b.B{
		A: a.A{
			First: true,
		},
	}

	fmt.Printf("First: %t\n", v.A.First)
}
