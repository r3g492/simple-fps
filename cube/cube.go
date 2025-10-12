package cube

import rl "github.com/gen2brain/raylib-go/raylib"

type Cube struct {
	Position   rl.Vector3
	Width      float32
	Height     float32
	Length     float32
	Color      rl.Color
	RotAxis    rl.Vector3
	RotDegrees float32
}

func (c *Cube) Draw() {
	rl.PushMatrix()
	rl.Translatef(c.Position.X, c.Position.Y, c.Position.Z)
	rl.Rotatef(c.RotDegrees, c.RotAxis.X, c.RotAxis.Y, c.RotAxis.Z)
	rl.DrawCube(rl.Vector3{X: 0, Y: 0, Z: 0}, c.Width, c.Height, c.Length, c.Color)
	rl.PopMatrix()
}
