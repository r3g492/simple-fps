package cube

import rl "github.com/gen2brain/raylib-go/raylib"

type Cube struct {
	Position rl.Vector3
	Width    float32
	Height   float32
	Length   float32
	Color    rl.Color
}

func (c *Cube) Draw() {
	rl.DrawCube(c.Position, c.Width, c.Height, c.Length, c.Color)
}
