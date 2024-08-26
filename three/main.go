package three

import "github.com/davecgh/go-spew/spew"

func Three(a, b int) int {
	r := a + b
	// release v1.0.0
	// release v1.1.0
	// release v1.2.0
	spew.Dump(r)
	return r
}
