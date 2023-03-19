package tetrimino

import (
	"math/rand"
	"time"
)

type Tetrimino struct {
	Shape   [][]bool
	X, Y    int
}

type Cell struct {
	X, Y int
}

func NewTetrimino(shape [][]bool) *Tetrimino {
	return &Tetrimino{
		Shape: shape,
		X:     3,
		Y:     0,
	}
}

func NewRandomTetrimino() *Tetrimino {
	rand.Seed(time.Now().UnixNano())
	shapes := [][][]bool{
		Shapes.I, Shapes.O, Shapes.T, Shapes.S, Shapes.Z, Shapes.J, Shapes.L,
	}
	return NewTetrimino(shapes[rand.Intn(len(shapes))])
}

func (t *Tetrimino) Cells() []Cell {
	cells := []Cell{}

	for y, row := range t.Shape {
		for x, cell := range row {
			if cell {
				cells = append(cells, Cell{X: t.X + x, Y: t.Y + y})
			}
		}
	}

	return cells
}

func (t *Tetrimino) MoveLeft() {
	t.X--
}

func (t *Tetrimino) MoveRight() {
	t.X++
}

func (t *Tetrimino) MoveDown() {
	t.Y++
}

func (t *Tetrimino) Rotate() {
	newShape := make([][]bool, len(t.Shape[0]))
	for i := range newShape {
		newShape[i] = make([]bool, len(t.Shape))
	}

	for y, row := range t.Shape {
		for x, cell := range row {
			newShape[x][len(t.Shape)-1-y] = cell
		}
	}

	t.Shape = newShape
}

func (t *Tetrimino) Clone() *Tetrimino {
	shape := make([][]bool, len(t.Shape))
	for i := range t.Shape {
		shape[i] = make([]bool, len(t.Shape[i]))
		copy(shape[i], t.Shape[i])
	}

	return &Tetrimino{
		Shape: shape,
		X:     t.X,
		Y:     t.Y,
	}
}

var Shapes = struct {
	I, O, T, S, Z, J, L [][]bool
}{
	I: [][]bool{
		{true, true, true, true},
	},
	O: [][]bool{
		{true, true},
		{true, true},
	},
	T: [][]bool{
		{false, true, false},
		{true, true, true},
	},
	S: [][]bool{
		{false, true, true},
		{true, true, false},
	},
	Z: [][]bool{
		{true, true, false},
		{false, true, true},
	},
	J: [][]bool{
		{true, false, false},
		{true, true, true},
	},
	L: [][]bool{
		{false, false, true},
		{true, true, true},
	},
}