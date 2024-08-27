package one

import "github.com/davecgh/go-spew/spew"

func One(a, b int) int {
	r := a + b
	spew.Dump(r)
	return r
}
