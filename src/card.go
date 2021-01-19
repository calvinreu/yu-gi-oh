package main

import (
	"./lib/graphic"
	"github.com/veandco/go-sdl2/sdl"
)

//Card contains the information of every card
type Card struct {
	cardInstance     *graphic.Instance
	faceDownInstance *graphic.Instance
}

//SetCardInstance Setter for cardInstance
func (card *Card) SetCardInstance(cardInstance *graphic.Instance) {
	card.cardInstance = cardInstance
}

//SetFaceDownInstance Setter for facedown instance
func (card *Card) SetFaceDownInstance(faceDownInstance *graphic.Instance) {
	card.faceDownInstance = faceDownInstance
}

//FaceUp shows the card image
func (card *Card) FaceUp(position sdl.Point, turned bool) {

	card.cardInstance.Show()
	card.faceDownInstance.Hide()

	card.cardInstance.NewPosition(position)
	if turned {
		card.cardInstance.SetAngle(90)
	} else {
		card.cardInstance.SetAngle(0)
	}
}

//FaceDown Hides the card image and shows a card backsite image
func (card *Card) FaceDown(position sdl.Point, turned bool) {

	card.faceDownInstance.Show()
	card.cardInstance.Hide()

	card.faceDownInstance.NewPosition(position)
	if turned {
		card.faceDownInstance.SetAngle(90)
	} else {
		card.faceDownInstance.SetAngle(0)
	}

}

//Hide makes the card invisible
func (card *Card) Hide() {

	card.faceDownInstance.Hide()
	card.cardInstance.Hide()

}
