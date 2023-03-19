package renderer

import (
	"github.com/nsf/termbox-go"
	"github.com/YuNaga224/GoTetris/board"
	"github.com/YuNaga224/GoTetris/tetrimino"
)

// 描画の際に使用するオフセット座標を持つ構造体
type Renderer struct {
	OffsetX, OffsetY int
}

func NewRenderer() *Renderer {
	return &Renderer{
		OffsetX: 2,
		OffsetY: 1,
	}
}

// ゲームボードを描画する関数
func (r *Renderer) DrawBoard(b *board.Board) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			cell := b.Cells[y][x]
			if cell {
				termbox.SetCell(r.OffsetX+x*2, r.OffsetY+y, '█', termbox.ColorWhite, termbox.ColorWhite)
				termbox.SetCell(r.OffsetX+x*2+1, r.OffsetY+y, '█', termbox.ColorWhite, termbox.ColorWhite)
			} else {
				termbox.SetCell(r.OffsetX+x*2, r.OffsetY+y, '░', termbox.ColorDefault, termbox.ColorDefault)
				termbox.SetCell(r.OffsetX+x*2+1, r.OffsetY+y, '░', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
}

// 現在のテトリミノを描画する関数
func (r *Renderer) DrawTetrimino(t *tetrimino.Tetrimino) {
	for _, cell := range t.Cells() {
		termbox.SetCell(r.OffsetX+cell.X*2, r.OffsetY+cell.Y, '█', termbox.ColorWhite, termbox.ColorWhite)
		termbox.SetCell(r.OffsetX+cell.X*2+1, r.OffsetY+cell.Y, '█', termbox.ColorWhite, termbox.ColorWhite)
	}
}


