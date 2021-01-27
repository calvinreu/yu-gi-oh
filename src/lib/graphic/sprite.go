//Package graphic is using the sdl2 go interface from (c)https://github.com/veandco/go-sdl2/ under the BSD 3 License
package graphic

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//Instance position angle and the center of an instance of a sprite
type Instance struct {
	destRect     sdl.FRect
	angle        float64
	center       sdl.FPoint
	parentSprite *Sprite
}

//Sprite contains the texture a list of instances and a srcRect
type Sprite struct {
	texture   *sdl.Texture
	instances List
	srcRect   sdl.Rect
}

//NewSprite creates a sprite based on a renderer, the image path and a src rectangle
func NewSprite(renderer *sdl.Renderer, imgPath string, srcRect sdl.Rect) (Sprite, error) {
	var sprite Sprite
	var err error
	sprite.texture, err = img.LoadTexture(renderer, imgPath)

	if err != nil {
		return sprite, err
	}

	sprite.srcRect = srcRect

	return sprite, err
}

//NewInstance adds a instance to the sprite and initializes the width and height of the dest rectangle with the src rectangle
func (sprite *Sprite) NewInstance(angle float64, center sdl.FPoint) *Instance {
	var instance Instance

	instance.NewPosition(center)
	instance.destRect.W = (float32)(sprite.srcRect.W)
	instance.destRect.H = (float32)(sprite.srcRect.H)
	instance.angle = angle
	instance.parentSprite = sprite
	sprite.instances.Push(&instance)

	return &instance
}

//NewPosition sets the position of this instance center is the center of the instances new position
func (instance *Instance) NewPosition(center sdl.FPoint) {
	instance.center = sdl.FPoint{instance.destRect.W / 2, instance.destRect.H / 2}
	instance.destRect.X = center.X - (instance.destRect.W / 2)
	instance.destRect.Y = center.Y - (instance.destRect.H / 2)
}

//NewPositionCorner sets the top left corner of the instance to corner
func (instance *Instance) NewPositionCorner(corner sdl.FPoint) {
	instance.center = sdl.FPoint{instance.destRect.W / 2, instance.destRect.H / 2}
	instance.destRect.X = corner.X
	instance.destRect.Y = corner.Y
}

//SetAngle setter for instance.angle
func (instance *Instance) SetAngle(angle float64) {
	instance.angle = angle
}

//ShowInstance makes instance visible
func (sprite *Sprite) ShowInstance(instance *Instance) {
	instance.destRect.W, instance.destRect.H = (float32)(sprite.srcRect.W), (float32)(sprite.srcRect.H)
}

//Hide makes the instance invisible until Show
func (instance *Instance) Hide() {
	instance.destRect.W, instance.destRect.H = 0, 0
}

//Show makes the instance visible
func (instance *Instance) Show() {
	instance.parentSprite.ShowInstance(instance)
}

//Zoom changes the texture size by textureSize*multiplier does not work if the object is hidden could have weird artifacts at small texture sizes
func (instance *Instance) Zoom(multiplier float32) {
	instance.destRect.W, instance.destRect.H = multiplier*instance.destRect.W, multiplier*instance.destRect.H
	instance.NewPosition(sdl.FPoint{instance.center.X + instance.destRect.X, instance.center.Y + instance.destRect.Y})
}

//IsZoomed returns if an instance is zommed in
func (instance *Instance) IsZoomed() bool {
	return (instance.destRect.W != (float32)(instance.parentSprite.srcRect.W) || instance.destRect.H != (float32)(instance.parentSprite.srcRect.H)) && (!instance.IsHidden())
}

//ResetZoom dRect.W, dRect.H is reset to default values
func (instance *Instance) ResetZoom() {
	if instance.IsZoomed() {
		instance.destRect.W, instance.destRect.H = (float32)(instance.parentSprite.srcRect.W), (float32)(instance.parentSprite.srcRect.H)
	}
}

//IsHidden returns if an instance is hidden
func (instance *Instance) IsHidden() bool {
	return instance.destRect.W == 0 || instance.destRect.H == 0
}

//GetBaseWitdth returns the width of an unzommed instance
func (instance Instance) GetBaseWitdth() float32 {
	return (float32)(instance.parentSprite.srcRect.W)
}

//GetBaseHeight returns the height of an unzommed instance
func (instance Instance) GetBaseHeight() float32 {
	return (float32)(instance.parentSprite.srcRect.H)
}
