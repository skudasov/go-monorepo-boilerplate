package two

import "github.com/davecgh/go-spew/spew"

func Two(a, b, c int) int {
	r := a + b + c
	spew.Dump(r)
	return r
}
