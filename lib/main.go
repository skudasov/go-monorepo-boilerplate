package lib

import "github.com/davecgh/go-spew/spew"

func Sum(a, b int) int {
	r := a + b
	spew.Dump(r)
	return r
}
