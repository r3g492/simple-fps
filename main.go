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

	playerPosition := rl.Vector3{X: 0, Y: 1, Z: 0}
	playerUpVector := rl.Vector3{X: 0, Y: 1, Z: 0}
	playerMoveSpeed := float32(8)
	playerNormalizedForward := rl.Vector3Normalize(rl.Vector3{X: 0.1, Y: 0, Z: 0})
	playerUpVelocity := float32(0.0)
	sensitivity := float32(0.0015)
	maxPitch := float32(1.553343)
	var yaw, pitch float32
	jumpPower := float32(6.4)
	gravitationalForce := float32(-9.81)
	lowerestGroundPoint := float32(1)

	blastCooldown := 200 * time.Millisecond
	lastBlast := time.Now()

	logEvery := 1000 * time.Millisecond
	lastLog := time.Now()
	rl.DisableCursor()

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		// player movement start

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
		playerNormalizedForward = rl.Vector3Normalize(rl.NewVector3(fx, fy, fz))

		// do key inputs
		// project forward onto XZ plane
		playerNormalizedRight := rl.Vector3Normalize(rl.Vector3CrossProduct(playerNormalizedForward, playerUpVector))

		// accumulate input
		move := rl.NewVector3(0, 0, 0)
		if rl.IsKeyDown(rl.KeyW) {
			move = rl.Vector3Add(move, rl.Vector3{X: playerNormalizedForward.X, Y: 0, Z: playerNormalizedForward.Z})
		}
		if rl.IsKeyDown(rl.KeyS) {
			move = rl.Vector3Subtract(move, rl.Vector3{X: playerNormalizedForward.X, Y: 0, Z: playerNormalizedForward.Z})
		}
		if rl.IsKeyDown(rl.KeyA) {
			move = rl.Vector3Subtract(move, rl.Vector3{X: playerNormalizedRight.X, Y: 0, Z: playerNormalizedRight.Z})
		}
		if rl.IsKeyDown(rl.KeyD) {
			move = rl.Vector3Add(move, rl.Vector3{X: playerNormalizedRight.X, Y: 0, Z: playerNormalizedRight.Z})
		}
		blast := false
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && time.Since(lastBlast) >= blastCooldown {
			blast = true
			lastBlast = time.Now()
		}

		// normalize so diagonals are not faster
		if rl.Vector3Length(move) > 0 {
			dir := rl.Vector3Normalize(move)
			playerPosition = rl.Vector3Add(playerPosition, rl.Vector3Scale(dir, playerMoveSpeed*dt))
		}

		// player must be at the ground to jump
		if rl.IsKeyPressed(rl.KeySpace) && playerPosition.Y <= 2 {
			playerUpVelocity = jumpPower
		}
		playerUpVelocity += gravitationalForce * dt
		playerPosition = rl.Vector3Add(playerPosition, rl.Vector3Scale(playerUpVector, playerUpVelocity*dt))
		if playerPosition.Y < lowerestGroundPoint {
			playerPosition.Y = lowerestGroundPoint
			playerUpVelocity = 0
		}
		// player movement end

		camera := updatePlayerCamera(playerPosition, playerNormalizedForward, playerUpVector)

		// raylib graphics
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.BeginMode3D(camera)
		rl.DrawGrid(
			1024,
			1,
		)

		// --- start gun ---
		right := rl.Vector3Normalize(rl.Vector3CrossProduct(playerNormalizedForward, playerUpVector))
		camUp := rl.Vector3Normalize(rl.Vector3CrossProduct(right, playerNormalizedForward))

		const offF = float32(0.6)
		const offR = float32(0.40)
		const offU = float32(-0.64)
		const barrelLen = float32(0.5)

		gunStart := rl.Vector3Add(playerPosition,
			rl.Vector3Add(
				rl.Vector3Add(
					rl.Vector3Scale(playerNormalizedForward, offF),
					rl.Vector3Scale(right, offR),
				),
				rl.Vector3Scale(camUp, offU),
			),
		)
		gunEnd := rl.Vector3Add(gunStart, rl.Vector3Scale(playerNormalizedForward, barrelLen))
		rl.DrawLine3D(gunStart, gunEnd, rl.Blue)
		rl.DrawCylinderEx(gunStart, gunEnd, 0.06, 0.03, 32, rl.Black)
		if blast {
			rl.DrawSphere(gunEnd, 0.15, rl.Yellow)
		}
		// --- end gun ---

		rl.EndMode3D()

		rl.DrawCircle(width/2, height/2, 3, rl.DarkGray)
		rl.EndDrawing()

		// log
		if time.Since(lastLog) >= logEvery {
			fmt.Println("mouseDelta:", mouseDelta)
			lastLog = time.Now()
		}
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
