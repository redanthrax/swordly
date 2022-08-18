package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		plr.moving = true
		plr.direction = 1
	}

	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		plr.moving = true
		plr.direction = 2
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		plr.moving = true
		plr.direction = 3
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		plr.moving = true
		plr.direction = 4
	}
}
