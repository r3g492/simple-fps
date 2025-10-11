package enemy

import (
	"simple-fps/animation"
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
	AnimationType     animation.Type
	LastAnimation     time.Time
	Cubes             []cube.Cube
}

func CreateEnemy(position rl.Vector3, normalizedForward rl.Vector3) {
	Enemies = append(
		Enemies, Enemy{
			CreatedAt:         time.Now(),
			Position:          position,
			NormalizedForward: normalizedForward,
			MovementSpeed:     10,
			AnimationType:     animation.Idle,
			LastAnimation:     time.Now(),
		},
	)
}

func UpdateEnemyLogic(now time.Time) {
	// TODO: implement
	for i := range Enemies {
		Enemies[i].CreatedAt = time.Now()
	}
}

func UpdateEnemyAnimation(now time.Time) {
	for i := range Enemies {
		isAnimationDone := false
		if Enemies[i].AnimationType == animation.Move {
			Enemies[i].Cubes, isAnimationDone = animation.EnemyMove(now, Enemies[i].LastAnimation, Enemies[i].Position)
		} else if Enemies[i].AnimationType == animation.Attack {
			Enemies[i].Cubes, isAnimationDone = animation.EnemyAttack(now, Enemies[i].LastAnimation, Enemies[i].Position)
		} else {
			Enemies[i].Cubes, isAnimationDone = animation.EnemyIdle(now, Enemies[i].LastAnimation, Enemies[i].Position)
		}

		if isAnimationDone {
			Enemies[i].LastAnimation = now
		}
	}
}

func DrawEnemies() {
	for _, e := range Enemies {
		for _, c := range e.Cubes {
			c.Draw()
		}
	}
}

func (e *Enemy) IsHit(
	bulletStartingPoint rl.Vector3,
	bulletEndPoint rl.Vector3,
) bool {
	return false
}
