package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var width int32 = 1600
	var height int32 = 900

	rl.InitWindow(width, height, "ammo")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	playerPosition := rl.Vector3{
		X: 0,
		Y: 1,
		Z: 0,
	}
	playerForward := rl.Vector3{
		X: 0.1,
		Y: 0,
		Z: 0,
	}
	upVector := rl.Vector3{
		X: 0,
		Y: 1,
		Z: 0,
	}
	playerMoveSpeed := float32(0.1)
	playerNormalizedForward := rl.Vector3Normalize(playerForward)
	playerNormalizedRight := rl.Vector3Normalize(rl.Vector3CrossProduct(playerNormalizedForward, upVector))

	logEvery := 1000 * time.Millisecond
	lastLog := time.Now()
	rl.DisableCursor()

	for !rl.WindowShouldClose() {
		camera := updatePlayerCamera(playerPosition, playerForward, upVector)

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

		if time.Since(lastLog) >= logEvery {
			mouseDelta := rl.GetMouseDelta()
			fmt.Println("mouseDelta:", mouseDelta)
			lastLog = time.Now()
		}

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
