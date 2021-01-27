package game

import (
	"../graphic"
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

//GetCardInstance Returns a copy of the card instance of the card
func (card *Card) GetCardInstance() graphic.Instance {
	return *card.cardInstance
}

//SetFaceDownInstance Setter for facedown instance
func (card *Card) SetFaceDownInstance(faceDownInstance *graphic.Instance) {
	card.faceDownInstance = faceDownInstance
}

//FaceUp shows the card image
func (card *Card) FaceUp(position sdl.FPoint, turned bool) {

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
func (card *Card) FaceDown(position sdl.FPoint, turned bool) {

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

//ZoomOnce to card resets zoom of card if it is zommed in then zooms in
func (card *Card) ZoomOnce() {
	if card.cardInstance.IsZoomed() {
		card.cardInstance.ResetZoom()
	}
	if card.faceDownInstance.IsZoomed() {
		card.faceDownInstance.ResetZoom()
	}

	card.faceDownInstance.Zoom(1.3)
	card.cardInstance.Zoom(1.3)
}
