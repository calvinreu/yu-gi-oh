package main

import (
	"fmt"
	"time"

	"./lib/graphic"
	"github.com/veandco/go-sdl2/sdl"
)

func gameloop(board GameField, window *graphic.Graphic) {
	//if()

	fmt.Scanln()
	window.Render()
	window.Print()
	time.Sleep(1 * time.Second)
	board[0].Deck.Cards[0].FaceUp(board[0].Deck.Center, false)

	window.Render()
	window.Print()
	time.Sleep(1 * time.Second)

	board[0].Deck.Cards[0].FaceDown(board[0].Deck.Center, true)

	window.Render()
	window.Print()
	time.Sleep(1 * time.Second)

	board[0].Deck.Cards[0].FaceUp(board[0].Deck.Center, true)

	window.Render()
	window.Print()
	time.Sleep(1 * time.Second)
}

/*func refresh(board *GameField, window *graphic.Graphic) {
	for _, i := range *board {
		if IsInRect(i.Deck.Place, )
	}
}*/

//IsInRect returns true if position is in rect
func IsInRect(rect sdl.Rect, position sdl.Point) bool {
	if position.X < rect.X+rect.W && position.X > rect.X && position.Y < rect.Y+rect.H && position.Y > rect.Y {
		return true
	}
	return false
}
