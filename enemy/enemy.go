package enemy

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Enemies []Enemy

type Enemy struct {
	CreatedAt         time.Time
	Position          rl.Vector3
	NormalizedForward rl.Vector3
	MovementSpeed     float32
}

func CreateEnemy(position rl.Vector3, normalizedForward rl.Vector3) {
	Enemies = append(Enemies, Enemy{time.Now(), position, normalizedForward, 10})
}

func UpdateEnemy(dt float32) {

}

func DrawEnemies() {
	for _, i := range Enemies {
		rl.DrawSphere(i.Position, 1, rl.Red)
	}
}

func (*Enemy) IsHit(bulletStartingPoint rl.Vector3, bulletEndPoint rl.Vector3) bool {
	return false
}
