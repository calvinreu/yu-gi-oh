package game

import (
	"../graphic"
	"github.com/veandco/go-sdl2/sdl"
)

//GameField contains the playing field
type GameField []PlayerField

//PlayerField contains all the card slots for one player
type PlayerField struct {
	Deck, ExtraDeck, Graveyard, BannedStack CardStack
	MonsterSlots, SpellZone                 [5]SingleCardSlot
	FieldSpell                              SingleCardSlot
}

//CardStack contains info for a stack of cards
type CardStack struct {
	Cards  []Card
	Place  sdl.Rect
	Center sdl.Point
}

//SingleCardSlot contains position and Card
type SingleCardSlot struct {
	Card   Card
	Place  sdl.Rect
	Center sdl.Point
}

//Init slice in card stack
func (cardStack *CardStack) Init() {
	cardStack.Cards = make([]Card, 0)
}

//CreateRenderStack list the cards and returns the card instances to render
func (cardStack *CardStack) CreateRenderStack() []graphic.Instance {
	instances := make([]graphic.Instance, len(cardStack.Cards))

	for n, i := range cardStack.Cards {
		instances[n] = i.GetCardInstance()
		instances[n].NewPositionCorner(sdl.FPoint{(float32)(n) * (i.GetCardInstance().GetBaseWitdth()), 0})
		instances[n].Show()
	}

	return instances
}
