package bullet

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var PlayerBullets []PlayerBullet

type PlayerBullet struct {
	CreatedAt         time.Time
	Position          rl.Vector3
	NormalizedForward rl.Vector3
	MovementSpeed     float32
}

func CreatePlayerBullet(playerPosition rl.Vector3, playerNormalizedForward rl.Vector3) {
	PlayerBullets = append(PlayerBullets, PlayerBullet{time.Now(), playerPosition, playerNormalizedForward, 1000})
}

func UpdatePlayerBullets(dt float32) {
	now := time.Now()
	dst := PlayerBullets[:0]
	for i := range PlayerBullets {
		if now.Sub(PlayerBullets[i].CreatedAt) <= 3*time.Second {
			PlayerBullets[i].Position = rl.Vector3Add(
				PlayerBullets[i].Position,
				rl.Vector3Scale(PlayerBullets[i].NormalizedForward, PlayerBullets[i].MovementSpeed*dt),
			)
			dst = append(dst, PlayerBullets[i])
		}
	}
	PlayerBullets = dst
}

func DrawPlayerBullets() {
	for _, b := range PlayerBullets {
		rl.DrawSphere(b.Position, 0.15, rl.Yellow)
	}
}
