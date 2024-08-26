package two

import "github.com/davecgh/go-spew/spew"

func Two(a, b, c int) int {
	r := a + b + c
	// release v1.0.0
	// release v1.1.0
	// release v1.2.0
	// another release
	spew.Dump(r)
	return r
}
