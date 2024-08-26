package two

import "github.com/davecgh/go-spew/spew"

func Two(a, b int) int {
	r := a + b
	// release v1.0.0
	spew.Dump(r)
	return r
}
