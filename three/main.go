package three

import "github.com/davecgh/go-spew/spew"

func Three(a, b int) int {
	r := a + b
	spew.Dump(r)
	return r
}
