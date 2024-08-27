package two

import "github.com/davecgh/go-spew/spew"

func Two(a, b int) int {
	r := a + b
	spew.Dump(r)
	return r
}
