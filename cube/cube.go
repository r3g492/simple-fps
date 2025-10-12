package cube

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cube struct {
	Position rl.Vector3
	Width    float32
	Height   float32
	Length   float32
	Color    rl.Color
	Forward  rl.Vector3
	Roll     float32
}

func (c *Cube) Draw() {
	rl.PushMatrix()
	rl.Translatef(c.Position.X, c.Position.Y, c.Position.Z)

	axis, deg := axisAngleFromForward(c.Forward)
	rl.Rotatef(deg, axis.X, axis.Y, axis.Z)

	rl.Rotatef(c.Roll, 0, 0, 1)

	rl.DrawCube(rl.Vector3{}, c.Width, c.Height, c.Length, c.Color)
	rl.PopMatrix()
}

func axisAngleFromForward(f rl.Vector3) (rl.Vector3, float32) {
	d := rl.Vector3{Z: 1}
	up := rl.Vector3{Y: 1}

	f = rl.Vector3Normalize(f)
	dot := rl.Vector3DotProduct(d, f)

	if dot > 0.9999 {
		return up, 0
	}
	if dot < -0.9999 {
		axis := rl.Vector3Normalize(rl.Vector3CrossProduct(up, d))
		return axis, 180
	}
	axis := rl.Vector3Normalize(rl.Vector3CrossProduct(d, f))
	angle := float32(math.Acos(float64(dot)) * 180 / math.Pi)
	return axis, angle
}
