package animation

import (
	"simple-fps/cube"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var IdleAnimationLen = 3 * time.Second

const idleAmp = 0.4

func EnemyIdle(
	now time.Time,
	LastAnimation time.Time,
	position rl.Vector3,
	normalizedForward rl.Vector3,
) ([]cube.Cube, bool) {

	head := cube.Cube{
		Position: rl.Vector3{
			X: position.X,
			Y: position.Y,
			Z: position.Z,
		},
		Width:   0.5,
		Height:  0.5,
		Length:  0.5,
		Color:   rl.Red,
		Forward: normalizedForward,
		Roll:    30,
	}

	body := cube.Cube{
		Position: rl.Vector3{
			X: position.X,
			Y: position.Y - 0.5,
			Z: position.Z,
		},
		Width:   3,
		Height:  0.5,
		Length:  0.5,
		Color:   rl.Red,
		Forward: normalizedForward,
		Roll:    0,
	}

	elapsed := now.Sub(LastAnimation)
	done := elapsed >= IdleAnimationLen
	if done {
		return []cube.Cube{}, true
	}
	return []cube.Cube{
		head,
		body,
	}, done
}

func EnemyMove(
	now time.Time,
	LastAnimation time.Time,
	position rl.Vector3,
	normalizedForward rl.Vector3,
) ([]cube.Cube, bool) {
	return []cube.Cube{}, false
}

func EnemyAttack(
	now time.Time,
	LastAnimation time.Time,
	position rl.Vector3,
	normalizedForward rl.Vector3,
) ([]cube.Cube, bool) {
	return []cube.Cube{}, false
}
