//Package graphic is using the sdl2 go interface from (c)https://github.com/veandco/go-sdl2/ under the BSD 3 License
package graphic

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

//Graphic contains the information required to render a window with diffrent Sprites
type Graphic struct {
	Sprites    []Sprite
	Renderer   *sdl.Renderer
	window     *sdl.Window
	cardSpaces []sdl.Rect
}

//RenderBoard renders the information from the graphic object to the screen
func (graphic *Graphic) RenderBoard() {

	graphic.Renderer.SetDrawColor(10, 10, 10, 1)
	graphic.Renderer.Clear()

	graphic.Renderer.SetDrawColor(0, 100, 200, 1)
	graphic.Renderer.FillRects(graphic.cardSpaces)

	for _, i := range graphic.Sprites {
		for j := i.instances.First(); j.Next() != nil; j = j.Next() {
			err := graphic.Renderer.CopyExF(i.texture, &i.srcRect, &j.Value.destRect, j.Value.angle, &j.Value.center, sdl.FLIP_NONE)
			if err != nil {
				fmt.Println("error in render call : ", err)
			}
		}
	}
	graphic.Renderer.Present()
}

//RenderStack renders one stack of cards
func (graphic *Graphic) RenderStack(instances []Instance) {
	graphic.Renderer.SetDrawColor(10, 10, 10, 1)
	graphic.Renderer.Clear()
	for _, i := range instances {
		err := graphic.Renderer.CopyExF(i.parentSprite.texture, &i.parentSprite.srcRect, &i.destRect, i.angle, &i.center, sdl.FLIP_NONE)
		if err != nil {
			fmt.Println("error in render call : ", err)
		}
	}
	graphic.Renderer.Present()
}

//Print info about every sprite and instance
func (graphic *Graphic) Print() {
	for _, i := range graphic.Sprites {
		fmt.Println("sRect : ", i.srcRect)
		for j := i.instances.First(); j.Next() != nil; j = j.Next() {
			fmt.Println("dRect", j.Value.destRect)
			fmt.Println("Center", j.Value.center)
			fmt.Println("Angle", j.Value.angle)
		}
	}
	fmt.Println("--------------------------")
}

//SetCardSpaces set the rectangle for card spaces
func (graphic *Graphic) SetCardSpaces(rects []sdl.Rect) {
	graphic.cardSpaces = rects
}

//New returns a Graphic object with initialized renderer and window note that Sprites have to be added manual
func New(title string, x, y, width, heigh int32, WindowFlags, RendererFlags uint32) (Graphic, error) {
	var graphic Graphic
	var err error

	graphic.window, err = sdl.CreateWindow(title, x, y, width, heigh, WindowFlags)
	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO)
		return graphic, err
	}

	graphic.Renderer, err = sdl.CreateRenderer(graphic.window, -1, RendererFlags)

	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO)
		return graphic, err
	}

	return graphic, nil
}

//AddSprite adds another sprite which can be used be creating a instance of it see Sprite.NewInstance
func (graphic *Graphic) AddSprite(imgPath string, srcRect sdl.Rect) (uint32, error) {
	var err error
	var sprite Sprite
	retIndex := len(graphic.Sprites)
	sprite, err = NewSprite(graphic.Renderer, imgPath, srcRect)
	if err != nil {
		fmt.Print(err)
		return 0, err
	}
	graphic.Sprites = append(graphic.Sprites, sprite)

	return uint32(retIndex), err
}
