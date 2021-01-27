package main

import (
	"time"

	"./lib/game"
	"./lib/graphic"
	"github.com/veandco/go-sdl2/sdl"
)

func gameloop(board game.GameField, window *graphic.Graphic) {

	var cursor sdl.Point
	//var state uint32

	for true {

		start := time.Now()

		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.GetType() {
			case sdl.QUIT:
				return
			}
		}

		cursor.X, cursor.Y, _ = sdl.GetMouseState()

		//if state == sdl.MOUSEBUTTONDOWN {
		if cards, ok := StackSelected(&board, cursor); ok { //TODO show cards from stack
			if len(cards.Cards) > 0 {
				window.RenderStack(cards.CreateRenderStack())
				//do smth with the stack
			}
		} else if card, ok := CardSelected(&board, cursor); ok {
			card.Card.ZoomOnce()
			//do smth with card
		} else {
			window.RenderBoard()
		}
		//}

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
func StackSelected(board *game.GameField, cursor sdl.Point) (*game.CardStack, bool) {

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
func CardSelected(board *game.GameField, cursor sdl.Point) (*game.SingleCardSlot, bool) {

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
