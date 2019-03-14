package main

import (
	"fmt"

	a "github.com/kalbasit/dep-a/v2"
	b "github.com/kalbasit/dep-b/v2"
)

func main() {
	v := b.B{
		A: a.A{
			Second: true,
		},
	}

	fmt.Printf("Second: %t\n", v.A.Second)
}
