package main

import (
	"fmt"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var width int32 = 1600
	var height int32 = 900

	rl.InitWindow(width, height, "ammo")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	playerPosition2 := rl.Vector3{X: 0, Y: 1, Z: 0}
	playerForward2 := rl.Vector3{X: 0.1, Y: 0, Z: 0}
	playerUpVector := rl.Vector3{X: 0, Y: 1, Z: 0}
	playerMoveSpeed2 := float32(0.1)
	playerNormalizedForward2 := rl.Vector3Normalize(playerForward2)
	playerNormalizedRight2 := rl.Vector3Normalize(rl.Vector3CrossProduct(playerNormalizedForward2, playerUpVector))
	playerPosition, playerForward, upVector, playerMoveSpeed, playerNormalizedForward, playerNormalizedRight := playerPosition2, playerForward2, playerUpVector, playerMoveSpeed2, playerNormalizedForward2, playerNormalizedRight2

	sensitivity := float32(0.0025)
	maxPitch := float32(1.553343)
	var yaw, pitch float32

	logEvery := 1000 * time.Millisecond
	lastLog := time.Now()
	rl.DisableCursor()

	for !rl.WindowShouldClose() {

		// update camera by mouse input
		mouseDelta := rl.GetMouseDelta()
		yaw += mouseDelta.X * sensitivity
		pitch += -mouseDelta.Y * sensitivity
		if pitch > maxPitch {
			pitch = maxPitch
		}
		if pitch < -maxPitch {
			pitch = -maxPitch
		}
		fx := float32(math.Cos(float64(pitch)) * math.Cos(float64(yaw)))
		fy := float32(math.Sin(float64(pitch)))
		fz := float32(math.Cos(float64(pitch)) * math.Sin(float64(yaw)))
		playerForward = rl.Vector3Normalize(rl.NewVector3(fx, fy, fz))
		playerNormalizedForward = playerForward
		playerNormalizedRight = rl.Vector3Normalize(rl.Vector3CrossProduct(playerNormalizedForward, upVector))

		camera := updatePlayerCamera(playerPosition, playerForward, upVector)

		// do key inputs
		if rl.IsKeyDown(rl.KeyW) {
			playerPosition = rl.Vector3Add(playerPosition, rl.Vector3Scale(playerNormalizedForward, playerMoveSpeed))
		}
		if rl.IsKeyDown(rl.KeyS) {
			playerPosition = rl.Vector3Subtract(playerPosition, rl.Vector3Scale(playerNormalizedForward, playerMoveSpeed))
		}
		if rl.IsKeyDown(rl.KeyA) {
			playerPosition = rl.Vector3Subtract(playerPosition, rl.Vector3Scale(playerNormalizedRight, playerMoveSpeed))
		}
		if rl.IsKeyDown(rl.KeyD) {
			playerPosition = rl.Vector3Add(playerPosition, rl.Vector3Scale(playerNormalizedRight, playerMoveSpeed))
		}

		// log
		if time.Since(lastLog) >= logEvery {
			fmt.Println("mouseDelta:", mouseDelta)
			lastLog = time.Now()
		}

		// raylib graphics
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.BeginMode3D(camera)
		rl.DrawGrid(
			1024,
			1,
		)
		rl.EndMode3D()
		rl.EndDrawing()
	}
}

func updatePlayerCamera(playerPosition rl.Vector3, playerForward rl.Vector3, upVector rl.Vector3) rl.Camera3D {
	var camera = rl.Camera3D{
		Position:   playerPosition,
		Target:     rl.Vector3Add(playerPosition, playerForward),
		Up:         upVector,
		Fovy:       90,
		Projection: rl.CameraPerspective,
	}
	return camera
}
