package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth int32 = 800
	screenHeight int32 = 600

	running = true

	plr *player
	cam *rl.Camera2D

	rect1 rl.Rectangle
)

func init() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(screenWidth, screenHeight, "Swordly")
	initMap()
	plr = newPlayer("Brinna", 1)
	cam = newCamera(plr)

	rect1 = rl.NewRectangle(100, 100, 20, 20)
}

func main() {
	for running {
		input()
		colliding()
		running = !rl.WindowShouldClose()
		update()
		render()
	}

	quit()
}

func quit() {
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
