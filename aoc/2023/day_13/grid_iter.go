package main

func NewGridIter[T any](grid [][]T) *GridIter[T] {
	// pre allocate the memory
	col := make([]T, len(grid))
	row := make([]T, len(grid[0]))
	return &GridIter[T]{
		elements: grid,
		col:      col,
		row:      row,
	}
}

type GridIter[T any] struct {
	elements [][]T
	col      []T
	row      []T
}

func (g *GridIter[T]) Row(i int) []T {
	for j := 0; j < len(g.row); j++ {
		g.row[j] = g.elements[i][j]
	}
	return g.row
}
func (g *GridIter[T]) Col(i int) []T {
	for j := 0; j < len(g.col); j++ {
		g.col[j] = g.elements[j][i]
	}
	return g.col
}
