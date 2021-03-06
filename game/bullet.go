package game

import "github.com/veandco/go-sdl2/sdl"

// bulletConfig holds bullet configuration
type bulletConfig struct {
	speed     int32
	direction int32 // -1 up 1 down
	colorR    uint8
	colorG    uint8
	colorB    uint8
}

// bullet holds bullet state information
type bullet struct {
	c *bulletConfig
	r *sdl.Renderer
	x int32
	y int32
	w int32
	h int32
}

// newBullet renerates a new bullet and adds it to the bullet list
func newBullet(r *sdl.Renderer, bl *bulletList, x, y int32, c *bulletConfig) {
	b := &bullet{
		r: r,
		c: c,
		x: x,
		y: y,
		w: 7,
		h: 9,
	}

	*bl = append(*bl, b)

	b.Draw()

}

// Draw an individual bullet
func (b *bullet) Draw() {
	b.r.SetDrawColor(b.c.colorR, b.c.colorG, b.c.colorB, 0xFF)

	b.r.FillRect(
		&sdl.Rect{X: b.x, Y: b.y, W: b.w, H: b.h},
	)
}

// Update updates a bullets position
// This will return false if the bullet is out of bounds
func (b *bullet) Update() bool {
	_, maxY, _ := b.r.GetRendererOutputSize()

	b.y += b.c.direction * b.c.speed

	return !(b.y < 0 || b.y > int32(maxY))
}

// Holds all bullets currently on the screen
type bulletList []*bullet

// Draw renders all existing bullets
func (bl *bulletList) Draw() {
	tmpBl := bulletList{}
	for _, b := range *bl {
		if b.Update() {
			b.Draw()
			tmpBl = append(tmpBl, b)
		}
	}
	*bl = tmpBl
}

// remove removes a bullet from the bullet list
func (bl *bulletList) remove(b *bullet) {
	tmpBl := bulletList{}
	for _, tb := range *bl {
		if tb != b {
			tmpBl = append(tmpBl, tb)
		}
	}

	*bl = tmpBl
}
