package main

import (
	"time"

	"./lib/graphic"
	"github.com/veandco/go-sdl2/sdl"
)

func gameloop(board GameField, window *graphic.Graphic) {

	var cursor sdl.Point
	var state uint32

	for true {

		start := time.Now()

		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.GetType() {
			case sdl.QUIT:
				return
			}
		}

		cursor.X, cursor.Y, state = sdl.GetMouseState()

		window.Render()

		if state == sdl.MOUSE
		if cards, ok := StackSelected(&board, cursor); ok { //TODO show cards from stack
			//do smth with the stack
		} else if card, ok := CardSelected(&board, cursor); ok {
			//do smth with the stack
		}

		if time.Since(start) < 16 {
			time.Sleep(16 - time.Since(start))
		}
	}
}

/*func refresh(board *GameField, window *graphic.Graphic) {
	for _, i := range *board {
		if IsInRect(i.Deck.Place, )
	}
}*/

//StackSelected return the stack which is selected if one is selected if not the bool is true
func StackSelected(board *GameField, cursor sdl.Point) (*CardStack, bool) {

	for _, i := range *board {
		if IsInRect(i.Deck.Place, cursor) {
			return &i.Deck, true
		}
		if IsInRect(i.ExtraDeck.Place, cursor) {
			return &i.ExtraDeck, true
		}
		if IsInRect(i.Graveyard.Place, cursor) {
			return &i.Graveyard, true
		}
		if IsInRect(i.BannedStack.Place, cursor) {
			return &i.BannedStack, true
		}
	}

	return nil, false
}

//CardSelected return the stack which is selected if one is selected if not the bool is true
func CardSelected(board *GameField, cursor sdl.Point) (*SingleCardSlot, bool) {

	for _, i := range *board {
		if IsInRect(i.FieldSpell.Place, cursor) {
			return &i.FieldSpell, true
		}
		for j := 0; j < len(i.MonsterSlots); j++ {
			if IsInRect(i.MonsterSlots[j].Place, cursor) {
				return &i.MonsterSlots[j], true
			}
		}
		for j := 0; j < len(i.SpellZone); j++ {
			if IsInRect(i.SpellZone[j].Place, cursor) {
				return &i.SpellZone[j], true
			}
		}
	}

	return nil, false
}

//IsInRect returns true if position is in rect
func IsInRect(rect sdl.Rect, position sdl.Point) bool {
	if position.X < rect.X+rect.W && position.X > rect.X && position.Y < rect.Y+rect.H && position.Y > rect.Y {
		return true
	}
	return false
}
