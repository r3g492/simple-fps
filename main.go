package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	var width int32 = 800
	var height int32 = 450

	rl.InitWindow(width, height, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var camera = rl.Camera3D{
		Position: rl.Vector3{
			X: 0,
			Y: 1,
			Z: 0,
		},
		Target: rl.Vector3{
			X: 1,
			Y: 0,
			Z: 0,
		},
		Up: rl.Vector3{
			X: 0,
			Y: 1,
			Z: 0,
		},
		Fovy:       90,
		Projection: rl.CameraPerspective,
	}

	for !rl.WindowShouldClose() {
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
