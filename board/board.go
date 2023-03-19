package board

import (
	"github.com/YuNaga224/GoTetris/tetrimino"
)

type Board struct {
	Width  int
	Height int
	Cells  [][]bool
	Score int
}

// 新しいゲームボードを生成する関数
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

// テトリミノを左に移動できるかを判定する関数
func (b *Board) CanMoveLeft(t *tetrimino.Tetrimino) bool {
	for _, cell := range t.Cells() {
		x, y := cell.X-1, cell.Y
		if x < 0 || b.Cells[y][x] {
			return false
		}
	}
	return true
}

// テトリミノを右に移動できるかを判定する関数
func (b *Board) CanMoveRight(t *tetrimino.Tetrimino) bool {
	for _, cell := range t.Cells() {
		x, y := cell.X+1, cell.Y
		if x >= b.Width || b.Cells[y][x] {
			return false
		}
	}
	return true
}

// テトリミノを下に移動できるかを判定する関数
func (b *Board) CanMoveDown(t *tetrimino.Tetrimino) bool {
	for _, cell := range t.Cells() {
		x, y := cell.X, cell.Y+1
		if y >= b.Height || b.Cells[y][x] {
			return false
		}
	}
	return true
}

// テトリミノを回転させられるかを判定する関数
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

// テトリミノがボード上に配置できるかどうかを判定する関数
func (b *Board) CanPlace(t *tetrimino.Tetrimino) bool {
	for _, cell := range t.Cells() {
		x, y := cell.X, cell.Y
		if x < 0 || x >= b.Width || y < 0 || y >= b.Height || b.Cells[y][x] {
			return false
		}
	}
	return true
}

// テトリミノをボードに結合する関数
func (b *Board) Merge(t *tetrimino.Tetrimino) {
	for _, cell := range t.Cells() {
		b.Cells[cell.Y][cell.X] = true
	}
	b.ClearFullRows()
}

// ボード上のすべてのセルが埋まった行を削除する関数
func (b *Board) ClearFullRows() {
	clearedRows := 0
	for y := b.Height - 1; y >= 0; y-- {
		if isRowFull(b.Cells[y]) {
			b.Cells = append(b.Cells[:y], b.Cells[y+1:]...)
			b.Cells = append([][]bool{make([]bool, b.Width)}, b.Cells...)
			clearedRows++
		}
	}
	b.updateScore(clearedRows)
}

// 行が全て埋まっているかどうかを判断する関数
func isRowFull(row []bool) bool {
	for _, cell := range row {
		if !cell {
			return false
		}
	}
	return true
}

// テトリミノを下に落とす関数
func (b *Board) Drop(t *tetrimino.Tetrimino) {
	for b.CanMoveDown(t) {
		t.MoveDown()
	}
}

// スコアを更新する関数
func (b *Board) updateScore(clearedRows int) {
	switch clearedRows {
	case 1:
		b.Score += 100
	case 2:
		b.Score += 300
	case 3:
		b.Score += 500
	case 4:
		b.Score += 800
	}
}