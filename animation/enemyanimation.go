package animation

import (
	"math"
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
) ([]cube.Cube, bool) {
	elapsed := now.Sub(LastAnimation)
	done := elapsed >= IdleAnimationLen
	if !done {
		p := elapsed.Seconds() / IdleAnimationLen.Seconds()
		tri := 1 - math.Abs(2*p-1)
		position.Y -= float32(idleAmp * tri)
	}
	return []cube.Cube{
		{
			Position: position,
			Width:    1,
			Height:   1,
			Length:   1,
			Color:    rl.Red,
		},
	}, done
}

func EnemyMove(
	now time.Time,
	LastAnimation time.Time,
	position rl.Vector3,
) ([]cube.Cube, bool) {
	elapsed := now.Sub(LastAnimation)
	done := elapsed >= IdleAnimationLen
	if !done {
		p := elapsed.Seconds() / IdleAnimationLen.Seconds()
		tri := 1 - math.Abs(2*p-1)
		position.Y -= float32(idleAmp * tri)
	}
	return []cube.Cube{
		{
			Position: position,
			Width:    1,
			Height:   1,
			Length:   1,
			Color:    rl.Red,
		},
	}, done
}

func EnemyAttack(
	now time.Time,
	LastAnimation time.Time,
	position rl.Vector3,
) ([]cube.Cube, bool) {
	elapsed := now.Sub(LastAnimation)
	done := elapsed >= IdleAnimationLen
	if !done {
		p := elapsed.Seconds() / IdleAnimationLen.Seconds()
		tri := 1 - math.Abs(2*p-1)
		position.Y -= float32(idleAmp * tri)
	}
	return []cube.Cube{
		{
			Position: position,
			Width:    1,
			Height:   1,
			Length:   1,
			Color:    rl.Red,
		},
	}, done
}
