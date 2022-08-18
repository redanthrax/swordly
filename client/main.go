package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth int32 = 800
	screenHeight int32 = 600

	running = true

	plr *player
)

func init() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(screenWidth, screenHeight, "Swordly")
	plr = newPlayer("Brinna", 1)
}

func main() {
	for running {
		input()
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
