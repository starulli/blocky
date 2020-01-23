// Package strategy contains optimization algorithms for converting PNG image
// data into SVG elements.
package strategy

import (
	"github.com/tinychameleon/blocky/grid"
	"github.com/tinychameleon/blocky/shape"
)

// Func is the type of any strategy algorithm.
type Func func(grid.Grid) []shape.Interface
