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
	err := termbox.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize termbox: %v\n", err)
		os.Exit(1)
	}
	defer termbox.Close()

	gameBoard := board.NewBoard(10, 20)
	render := renderer.NewRenderer()
	currentTetrimino := tetrimino.NewRandomTetrimino()

	// Game loop
	gameLoop:
		for {
			render.DrawBoard(gameBoard)
			render.DrawTetrimino(currentTetrimino)

			termbox.Flush()

			select {
			case event := <-eventQueue():
				switch event.Type {
				case termbox.EventKey:
					if event.Key == termbox.KeyCtrlC || event.Ch == 'q' {
						break gameLoop
					}
					handleKeyEvent(event.Key, gameBoard, currentTetrimino)
				}
			case <-time.After(500 * time.Millisecond):
				if gameBoard.CanMoveDown(currentTetrimino) {
					currentTetrimino.MoveDown()
				} else {
					gameBoard.Merge(currentTetrimino)
					currentTetrimino = tetrimino.NewRandomTetrimino()
					if !gameBoard.CanPlace(currentTetrimino) {
						break gameLoop
					}
				}
			}
		}

	fmt.Println("Game Over!")
}

func eventQueue() chan termbox.Event {
	eventChan := make(chan termbox.Event)
	go func() {
		for {
			eventChan <- termbox.PollEvent()
		}
	}()
	return eventChan
}

func handleKeyEvent(key termbox.Key, gameBoard *board.Board, tetrimino *tetrimino.Tetrimino) {
	switch key {
	case termbox.KeyArrowLeft:
		if gameBoard.CanMoveLeft(tetrimino) {
			tetrimino.MoveLeft()
		}
	case termbox.KeyArrowRight:
		if gameBoard.CanMoveRight(tetrimino) {
			tetrimino.MoveRight()
		}
	case termbox.KeyArrowDown:
		if gameBoard.CanMoveDown(tetrimino) {
			tetrimino.MoveDown()
		}
	case termbox.KeySpace:
		gameBoard.Drop(tetrimino)
	case termbox.KeyArrowUp:
		if gameBoard.CanRotate(tetrimino) {
			tetrimino.Rotate()
		}
	}
}