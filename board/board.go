package board

import (
	"github.com/YuNaga224/GoTetris/tetrimino"
)

type Board struct {
	Width  int
	Height int
	Cells  [][]bool
}

func NewBoard(width, height int) *Board {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}

	return &Board{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

func (b *Board) CanMoveLeft(t *tetrimino.Tetrimino) bool {
	for _, cell := range t.Cells() {
		x, y := cell.X-1, cell.Y
		if x < 0 || b.Cells[y][x] {
			return false
		}
	}
	return true
}

func (b *Board) CanMoveRight(t *tetrimino.Tetrimino) bool {
	for _, cell := range t.Cells() {
		x, y := cell.X+1, cell.Y
		if x >= b.Width || b.Cells[y][x] {
			return false
		}
	}
	return true
}

func (b *Board) CanMoveDown(t *tetrimino.Tetrimino) bool {
	for _, cell := range t.Cells() {
		x, y := cell.X, cell.Y+1
		if y >= b.Height || b.Cells[y][x] {
			return false
		}
	}
	return true
}

func (b *Board) CanRotate(t *tetrimino.Tetrimino) bool {
	rotated := t.Clone()
	rotated.Rotate()

	for _, cell := range rotated.Cells() {
		x, y := cell.X, cell.Y
		if x < 0 || x >= b.Width || y < 0 || y >= b.Height || b.Cells[y][x] {
			return false
		}
	}

	return true
}

func (b *Board) CanPlace(t *tetrimino.Tetrimino) bool {
	for _, cell := range t.Cells() {
		x, y := cell.X, cell.Y
		if x < 0 || x >= b.Width || y < 0 || y >= b.Height || b.Cells[y][x] {
			return false
		}
	}
	return true
}

func (b *Board) Merge(t *tetrimino.Tetrimino) {
	for _, cell := range t.Cells() {
		b.Cells[cell.Y][cell.X] = true
	}
	b.ClearFullRows()
}

func (b *Board) ClearFullRows() {
	for y := b.Height - 1; y >= 0; y-- {
		if isRowFull(b.Cells[y]) {
			b.Cells = append(b.Cells[:y], b.Cells[y+1:]...)
			b.Cells = append([][]bool{make([]bool, b.Width)}, b.Cells...)
		}
	}
}

func isRowFull(row []bool) bool {
	for _, cell := range row {
		if !cell {
			return false
		}
	}
	return true
}

func (b *Board) Drop(t *tetrimino.Tetrimino) {
	for b.CanMoveDown(t) {
		t.MoveDown()
	}
}