package main

import (
	"github.com/skudasov/go-monorepo-boilerplate/one"
	"github.com/skudasov/go-monorepo-boilerplate/three"
	"github.com/skudasov/go-monorepo-boilerplate/two"
)

func main() {
	one.One(1, 2)
	two.Two(1, 2)
	three.Three(1, 2)
}
