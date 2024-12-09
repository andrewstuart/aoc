package ezaoc

import "fmt"

// Print2dGrid simply iterates each item and prints it out in a fmt.Print 2d
// grid. No spacing but newlines.
func Print2dGrid[T any, Ts []T](ts []Ts) {
	for _, row := range ts {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

// Print2dGridWithNumbers prints out a 2d grid with row and column numbers.
func Print2dGridWithNumbers[T any, Ts []T](ts []Ts) {
	for i, row := range ts {
		fmt.Printf("%d: ", i)
		for j, cell := range row {
			fmt.Printf("%d:%v ", j, cell)
		}
		fmt.Println()
	}
}

// Make2DSlice creates a 2d slice of type T and length ixj, and sets the i,jth
// elements of the 2d array to the result of f(i,j). Attempting here to follow
// more of an existing Go idiom (sort.Slice) than something purely generic.
func Make2DSlice[T any](i, j int, f func(i, j int) T) [][]T {
	m := make([][]T, i)
	for i := range m {
		m[i] = make([]T, j)
		for j := range m[i] {
			m[i][j] = f(i, j)
		}
	}
	return m
}

// IsInBounds returns for any 2d slice whether the given ints are in bounds
func IsInBounds[T any, Ts ~[]T](ts []Ts, i, j int) bool {
	gtZero := i >= 0 && j >= 0
	inBounds := i < len(ts) && len(ts) > 0 && j < len(ts[0])
	return gtZero && inBounds
}

// Type Cell is used by many of the 2D slice methods to indicate both value and
// slice indices to the caller/callee.
type Cell[T any] struct {
	I, J  int
	Value T
}

// Set should be used with the orignal slice to avoid panics, and updates the
// in the Cell index to that passed as a parameter.
func (c Cell[T]) Set(ts [][]T, to T) {
	ts[c.I][c.J] = to
}

// Point returns [2]int{i, j}; useful for a comparable map or set key.
func (c Cell[T]) Point() [2]int {
	return [2]int{c.I, c.J}
}

// SliceNeighbors is a utility function to get the elements surrounding a particular 2d index.
func SliceNeighbors[T any](ts [][]T, n, m int) []Cell[T] {
	var out []Cell[T]
	for i := n - 1; i < n+2; i++ {
		for j := m - 1; j < m+2; j++ {
			if IsInBounds(ts, i, j) && !(i == n && j == m) { // You are not your own neighbor
				out = append(out, Cell[T]{I: i, J: j, Value: ts[i][j]})
			}
		}
	}
	return out
}

// NonDiagSliceNeighbors is a utility function to get the elements surrounding a
// particular 2d index, not including diagonally adjacent elements.
func NonDiagSliceNeighbors[T any](ts [][]T, n, m int) []Cell[T] {
	var out []Cell[T]
	for i := n - 1; i < n+2; i++ {
		for j := m - 1; j < m+2; j++ {
			if IsInBounds(ts, i, j) && !(i == n && j == m) && !(i != n && j != m) { // You are not your own neighbor, ignore diags
				out = append(out, Cell[T]{I: i, J: j, Value: ts[i][j]})
			}
		}
	}
	return out
}

// VisitCells calls a function for a Cell of each value in the given 2D array.
func VisitCells[T any](ts [][]T, f func(Cell[T]) error) {
	var c Cell[T]
	for i, row := range ts {
		for j := range row {
			c.I, c.J, c.Value = i, j, ts[i][j]
			if f(c) != nil {
				return
			}
		}
	}
}

// VisitNeighbors iterates over a 2d array, calling a func with each index and
// a list of neighbors.
func VisitNeighbors[T any](ts [][]T, f func(Cell[T], []Cell[T]) error) {
	var c Cell[T]
	for i, row := range ts {
		for j := range row {
			c.I, c.J, c.Value = i, j, ts[i][j]
			if f(c, SliceNeighbors(ts, i, j)) != nil {
				return
			}
		}
	}
}

// VisitNonDiagNeighbors iterates over a 2d array, calling a func with each index and
// a list of neighbors, not including diagonal neighbors.
func VisitNonDiagNeighbors[T any](ts [][]T, f func(Cell[T], []Cell[T]) error) {
	var c Cell[T]
	for i, row := range ts {
		for j := range row {
			c.I, c.J, c.Value = i, j, ts[i][j]
			if f(c, NonDiagSliceNeighbors(ts, i, j)) != nil {
				return
			}
		}
	}
}

// Cols returns a column of ints with index n
func Cols[T any](ts [][]T, n int) []Cell[T] {
	col := make([]Cell[T], len(ts))
	for i, row := range ts {
		col[i].I = i
		col[i].J = n
		col[i].Value = row[n]
	}
	return col
}

// RawCols returns a column of type T in the 2d T array having second-dimension
// index n
func RawCols[T any](ts [][]T, n int) []T {
	col := make([]T, len(ts))
	for i, row := range ts {
		col[i] = row[n]
	}
	return col
}

type Direction int

const (
	Unknown Direction = iota
	Up
	Down
	Left
	Right
	UpLeft
	UpRight
	DownLeft
	DownRight
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	case UpLeft:
		return "UpLeft"
	case UpRight:
		return "UpRight"
	case DownLeft:
		return "DownLeft"
	case DownRight:
		return "DownRight"
	}
	return "Unknown"
}

var AllDirections = []Direction{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}

func GetCellsInDirection[T any](ts [][]T, d Direction, n, m, count int) []Cell[T] {
	var out []Cell[T]
	if count == 0 {
		return out
	}
	if !IsInBounds(ts, n, m) {
		return out
	}
	out = append(out, Cell[T]{I: n, J: m, Value: ts[n][m]})
	switch d {
	case Up:
		return append(out, GetCellsInDirection(ts, d, n-1, m, count-1)...)
	case Down:
		return append(out, GetCellsInDirection(ts, d, n+1, m, count-1)...)
	case Left:
		return append(out, GetCellsInDirection(ts, d, n, m-1, count-1)...)
	case Right:
		return append(out, GetCellsInDirection(ts, d, n, m+1, count-1)...)
	case UpLeft:
		return append(out, GetCellsInDirection(ts, d, n-1, m-1, count-1)...)
	case UpRight:
		return append(out, GetCellsInDirection(ts, d, n-1, m+1, count-1)...)
	case DownLeft:
		return append(out, GetCellsInDirection(ts, d, n+1, m-1, count-1)...)
	case DownRight:
		return append(out, GetCellsInDirection(ts, d, n+1, m+1, count-1)...)
	}
	return out
}
