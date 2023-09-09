package asteroid

import "math"

func CircleTest(first, second *Sprite) bool {
	firstRect := Vector2D{first.textureRect.Width, first.textureRect.Height}
	secondRect := Vector2D{second.textureRect.Width, second.textureRect.Height}
	radius1 := (firstRect.X + firstRect.Y) / 4
	radius2 := (secondRect.X + secondRect.Y) / 4
	xd := first.GetPosition().X - second.GetPosition().X
	yd := first.GetPosition().Y - second.GetPosition().Y
	return math.Sqrt(xd*xd+yd*yd) <= radius1+radius2
}
