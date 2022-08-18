package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	deltaTime float32
)

func update() {
	deltaTime = rl.GetFrameTime()
	if plr.moving {
		if plr.direction == 1 {
			plr.dst.Y -= plr.speed * deltaTime
		}
	}

	plr.moving = false
}
