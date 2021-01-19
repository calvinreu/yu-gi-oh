//Package graphic is using the sdl2 go interface from (c)https://github.com/veandco/go-sdl2/ under the BSD 3 License
package graphic

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//Instance position angle and the center of an instance of a sprite
type Instance struct {
	destRect     sdl.Rect
	angle        float64
	center       sdl.Point
	parentSprite *Sprite
}

//Sprite contains the texture a list of instances and a srcRect
type Sprite struct {
	texture   *sdl.Texture
	instances []Instance
	srcRect   sdl.Rect
}

//NewSprite creates a sprite based on a renderer, the image path and a src rectangle
func NewSprite(renderer *sdl.Renderer, imgPath string, srcRect sdl.Rect) (Sprite, error) {
	var sprite Sprite
	var err error
	sprite.texture, err = img.LoadTexture(renderer, imgPath)
	fmt.Println("wololo")

	if err != nil {
		return sprite, err
	}

	sprite.srcRect = srcRect
	sprite.instances = make([]Instance, 0)

	return sprite, err
}

//NewInstance adds a instance to the sprite and initializes the width and height of the dest rectangle with the src rectangle
func (sprite *Sprite) NewInstance(angle float64, center sdl.Point) *Instance {
	var instance Instance

	instance.NewPosition(center)
	instance.destRect.W = sprite.srcRect.W
	instance.destRect.H = sprite.srcRect.H
	instance.angle = angle
	instance.parentSprite = sprite
	sprite.instances = append(sprite.instances, instance)

	return &sprite.instances[len(sprite.instances)-1]
}

//NewPosition sets the position of this instance center is the center of the instances new position
func (instance *Instance) NewPosition(center sdl.Point) {
	instance.center = sdl.Point{instance.destRect.W / 2, instance.destRect.H / 2}
	instance.destRect.X = center.X - (instance.destRect.W / 2)
	instance.destRect.Y = center.Y - (instance.destRect.H / 2)
}

//SetAngle setter for instance.angle
func (instance *Instance) SetAngle(angle float64) {
	instance.angle = angle
}

//ShowInstance makes instance visible
func (sprite *Sprite) ShowInstance(instance *Instance) {
	fmt.Println("show instance call from sprite : ", sprite)
	instance.destRect.W, instance.destRect.H = sprite.srcRect.W, sprite.srcRect.H
}

//Hide makes the instance invisible until Show
func (instance *Instance) Hide() {
	instance.destRect.W, instance.destRect.H = 0, 0
}

//Show makes the instance visible
func (instance *Instance) Show() {
	instance.parentSprite.ShowInstance(instance)
}
