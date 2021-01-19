package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"./lib/graphic"
	"github.com/veandco/go-sdl2/sdl"
)

//ConfigFileData bla
type ConfigFileData struct {
	Players               []PlayerFieldConfig
	Cards                 map[string]string //Cardname filename
	CardWidth, CardHeight int32
	Window                GraphicsConfigFileData
}

//Config data
type Config struct {
	Players               []PlayerFieldConfig
	Cards                 map[string]CardConfig
	CardWidth, CardHeight int32
	Window                graphic.Graphic
}

//PlayerFieldConfig Config file data for Player
type PlayerFieldConfig struct {
	Deck, ExtraDeck                                                                             []string
	LocationDeck, LocationExtraDeck, LocationGraveyard, LocationFieldSpell, LocationBannedStack sdl.Point
	LocationMonsterSlot, LocationSpellZone                                                      [5]sdl.Point
}

//CardConfig Config file data for Card
type CardConfig struct {
	ImgFilename string
	SRect       sdl.Rect
}

//GraphicsConfigFileData Config file data for window and renderer
type GraphicsConfigFileData struct {
	Title                      string
	X, Y, Width, Height        int32
	WindowFlags, RendererFlags uint32
}

//Load config From file
func (config *Config) Load(filename string) {

	var tempCardConfig CardConfig
	var configFileData ConfigFileData

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = json.Unmarshal(data, &configFileData)
	if err != nil {
		fmt.Println("error:", err)
	}

	//Configure Graphics system
	config.Window, err = graphic.New(configFileData.Window.Title, configFileData.Window.X, configFileData.Window.Y, configFileData.Window.Width, configFileData.Window.Height, configFileData.Window.WindowFlags, configFileData.Window.RendererFlags)
	if err != nil {
		fmt.Println("error:", err)
	}

	//Configure Players
	config.Players = configFileData.Players
	config.CardWidth, config.CardHeight = configFileData.CardWidth, configFileData.CardHeight

	//Configure Cards
	config.Cards = make(map[string]CardConfig)

	for key, file := range configFileData.Cards {
		tempCardConfig.Load(file)
		config.Cards[key] = tempCardConfig
	}

	fmt.Println(configFileData)
}

//Load CardConfig from json file
func (cardConfig *CardConfig) Load(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = json.Unmarshal(data, cardConfig)
	if err != nil {
		fmt.Println("error:", err)
	}
}

//Convert convertes a config file to a graphic.Graphic and a GameField
func (config *Config) Convert() (GameField, *graphic.Graphic) {
	cards := make(map[string]uint32)
	var err error
	gameField := make([]PlayerField, len(config.Players))

	for name, card := range config.Cards {
		cards[name], err = config.Window.AddSprite(card.ImgFilename, card.SRect)
		if err != nil {
			fmt.Println(err)
			return nil, &config.Window
		}
		fmt.Println(name, card)
	}

	n := 0 //count players
	var emptyPoint sdl.Point
	for _, i := range config.Players {

		//Create deck
		deckSpace := CreateRect(i.LocationDeck, config.CardWidth, config.CardHeight)
		deck := CardStack{make([]Card, len(i.Deck)), deckSpace, i.LocationDeck}
		for j := 0; j < len(i.Deck)-1; j++ {
			deck.Cards[j].SetCardInstance(config.Window.Sprites[cards[i.Deck[j]]].NewInstance(0, emptyPoint))
			deck.Cards[j].SetFaceDownInstance(config.Window.Sprites[cards["facedown"]].NewInstance(0, emptyPoint))
			deck.Cards[j].Hide()
		}
		lastCard := len(i.Deck) - 1 //last item in deck
		deck.Cards[lastCard].SetCardInstance(config.Window.Sprites[cards[i.Deck[lastCard]]].NewInstance(0, emptyPoint))
		deck.Cards[lastCard].SetFaceDownInstance(config.Window.Sprites[cards["facedown"]].NewInstance(0, emptyPoint))
		deck.Cards[lastCard].FaceDown(i.LocationDeck, false)

		//Create Extra Deck
		extraDeckSpace := CreateRect(i.LocationExtraDeck, config.CardWidth, config.CardHeight)
		extraDeck := CardStack{make([]Card, len(i.ExtraDeck)), extraDeckSpace, i.LocationExtraDeck}
		for j := 0; j < len(i.ExtraDeck)-1; j++ {
			extraDeck.Cards[j].SetCardInstance(config.Window.Sprites[cards[i.ExtraDeck[j]]].NewInstance(0, emptyPoint))
			extraDeck.Cards[j].SetFaceDownInstance(config.Window.Sprites[cards["facedown"]].NewInstance(0, emptyPoint))
			extraDeck.Cards[j].Hide()
		}
		lastCard = len(i.ExtraDeck) - 1 //last item in deck

		if lastCard != -1 {
			extraDeck.Cards[lastCard].SetCardInstance(config.Window.Sprites[cards[i.ExtraDeck[lastCard]]].NewInstance(0, emptyPoint))
			extraDeck.Cards[lastCard].SetFaceDownInstance(config.Window.Sprites[cards["facedown"]].NewInstance(0, emptyPoint))
			extraDeck.Cards[lastCard].FaceDown(i.LocationExtraDeck, false)
		}

		var player PlayerField
		player.Deck, player.ExtraDeck = deck, extraDeck
		player.Graveyard.Init()
		player.BannedStack.Init()
		gameField[n] = player
		n++
	}

	return gameField, &config.Window

}

//CreateRect create a rect with W and H and the center position of rect
func CreateRect(center sdl.Point, W, H int32) sdl.Rect {
	var rect sdl.Rect
	rect.X = center.X - W/2
	rect.Y = center.Y - H/2
	rect.W = W
	rect.H = H
	return rect
}
