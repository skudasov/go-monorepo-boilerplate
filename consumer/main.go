package main

import (
	"github.com/skudasov/go-monorepo-boilerplate/one"
	"github.com/skudasov/go-monorepo-boilerplate/three"
	"github.com/skudasov/go-monorepo-boilerplate/two/v2"
)

func main() {
	one.One(1, 2, 3)
	two.Two(1, 2)
	three.Three(1, 2)
}
