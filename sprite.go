package asteroid

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Sprite struct {
	Transform
	textureRect Rect
	texture     *ebiten.Image
}

func NewSprite(texture *ebiten.Image) *Sprite {

	textureRect := NewRect(float64(texture.Bounds().Min.X),
		float64(texture.Bounds().Min.Y),
		float64(texture.Bounds().Max.X),
		float64(texture.Bounds().Max.Y))

	return &Sprite{textureRect: *textureRect, texture: texture}
}

func (s *Sprite) GetTextureRect() Rect {
	return s.textureRect
}

func (s *Sprite) SetTextureRect(width, height float64) {
	s.textureRect = Rect{Width: width, Height: height}
}

func (s *Sprite) GetTexture() *ebiten.Image {
	return s.texture
}
func (s *Sprite) SetTexture(texture *ebiten.Image) {
	s.texture = texture
}

func (s *Sprite) SetPivotToCenter() {
	rect := s.GetTextureRect()
	x, y := rect.GetCenter()
	s.SetPivot(x, y)
}

func (s *Sprite) GetTransform() *Transform {
	return &s.Transform
}

func (s *Sprite) SetTransform(transform Transform) {
	s.Transform = transform
}

func (s *Sprite) updateGeoM(geom ebiten.GeoM) ebiten.GeoM {
	geom.Translate(-s.pivot.X, -s.pivot.Y)
	geom.Rotate(float64(s.rotation%360) * 2 * math.Pi / 360)
	geom.Translate(s.position.X, s.position.Y)
	return geom
}

func (s *Sprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	local_op := &ebiten.DrawImageOptions{}
	local_op.GeoM = op.GeoM
	local_op.GeoM = s.updateGeoM(op.GeoM)
	s.Transform.SetGeoM(local_op.GeoM)
	if s.texture != nil {
		target.DrawImage(s.texture, local_op)
	}
}
