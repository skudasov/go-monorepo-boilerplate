package three

import "github.com/davecgh/go-spew/spew"

func Three(a, b, c int) int {
	r := a + b + c
	spew.Dump(r)
	return r
}
