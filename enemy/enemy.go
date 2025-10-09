package enemy

import (
	"simple-fps/cube"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Enemies []Enemy

type Enemy struct {
	CreatedAt         time.Time
	Position          rl.Vector3
	NormalizedForward rl.Vector3
	MovementSpeed     float32
	AnimationType     AnimationType
	Head              cube.Cube
	Body              cube.Cube
}

func CreateEnemy(position rl.Vector3, normalizedForward rl.Vector3) {
	Enemies = append(
		Enemies, Enemy{
			CreatedAt:         time.Now(),
			Position:          position,
			NormalizedForward: normalizedForward,
			MovementSpeed:     10,
			AnimationType:     Idle,
			Head: cube.Cube{
				Position: position,
				Width:    1,
				Height:   1,
				Length:   1,
				Color:    rl.Red,
			},
			Body: cube.Cube{
				Position: position,
				Width:    1,
				Height:   1,
				Length:   1,
				Color:    rl.Red,
			},
		},
	)
}

func UpdateEnemy(
	dt float32,
) {
}

func DrawEnemies() {
	for _, i := range Enemies {
		i.Head.Draw()
		i.Body.Draw()
	}
}

func (e *Enemy) IsHit(
	bulletStartingPoint rl.Vector3,
	bulletEndPoint rl.Vector3,
) bool {

	return false
}
