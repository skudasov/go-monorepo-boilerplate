package one

import "github.com/davecgh/go-spew/spew"

func One(a, b int) int {
	r := a + b
	// release #1 tag v1.0.0
	spew.Dump(r)
	return r
}
