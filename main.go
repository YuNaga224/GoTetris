package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/YuNaga224/GoTetris/board"
	"github.com/YuNaga224/GoTetris/render"
	"github.com/YuNaga224/GoTetris/tetrimino"
)

func main() {
	// termboxの初期化
	err := termbox.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize termbox: %v\n", err)
		os.Exit(1)
	}
	// ゲームボード、レンダー、現在のテトリミノを作成
	gameBoard := board.NewBoard(10, 20)
	render := renderer.NewRenderer()
	currentTetrimino := tetrimino.NewRandomTetrimino()
	//ゲーム終了時の処理
	defer fmt.Println("Game Over!\n")
	defer fmt.Println("Score: ", gameBoard.Score, "Point")
	defer termbox.Close()

	// ゲームループ
	gameLoop:
		for {
			//　ゲームボードと現在のテトリミノを描画
			render.DrawBoard(gameBoard)
			render.DrawTetrimino(currentTetrimino)
			// 描画内容を画面に反映
			termbox.Flush()

			//イベント処理
			select {
			case event := <-eventQueue():
				switch event.Type {
				case termbox.EventKey:
					// Ctr+Cかqでゲーム終了
					if event.Key == termbox.KeyCtrlC || event.Ch == 'q' {
						break gameLoop
					}
					// キーイベントを処理
					handleKeyEvent(event.Key, gameBoard, currentTetrimino)
				}
			case <-time.After(500 * time.Millisecond):
				//　テトリミノを下に動かす
				if gameBoard.CanMoveDown(currentTetrimino) {
					currentTetrimino.MoveDown()
				} else {
					// テトリミノが動かせない場合、ボードに固定
					gameBoard.Merge(currentTetrimino)
					// 新しいテトリミノを生成
					currentTetrimino = tetrimino.NewRandomTetrimino()
					// ゲームオーバー判定
					if !gameBoard.CanPlace(currentTetrimino) {
						break gameLoop
					}
				}
			}
		}

	fmt.Println("Game Over!")
}

// イベントキューを作成
func eventQueue() chan termbox.Event {
	eventChan := make(chan termbox.Event)
	go func() {
		for {
			eventChan <- termbox.PollEvent()
		}
	}()
	return eventChan
}

// キーイベントの処理
func handleKeyEvent(key termbox.Key, gameBoard *board.Board, tetrimino *tetrimino.Tetrimino) {
	switch key {
	case termbox.KeyArrowLeft:
		// 左矢印キー：テトリミノを左に移動
		if gameBoard.CanMoveLeft(tetrimino) {
			tetrimino.MoveLeft()
		}
	case termbox.KeyArrowRight:
		// 右矢印キー：テトリミノを右に移動
		if gameBoard.CanMoveRight(tetrimino) {
			tetrimino.MoveRight()
		}
	case termbox.KeyArrowDown:
		// 下矢印キー：テトリミノを下に移動
		if gameBoard.CanMoveDown(tetrimino) {
			tetrimino.MoveDown()
		}
	case termbox.KeySpace:
		// スペースキー：テトリミノを一気に下に落とす
		gameBoard.Drop(tetrimino)
	case termbox.KeyArrowUp:
		// 上矢印キー：テトリミノを回転
		if gameBoard.CanRotate(tetrimino) {
			tetrimino.Rotate()
		}
	}
}