package one

import "github.com/davecgh/go-spew/spew"

func One(a, b, c int) int {
	r := a + b + c
	// release #1 tag v1.0.0
	// release v1.1.0
	// release v1.2.0
	// release v1.3.0
	spew.Dump(r)
	return r
}
