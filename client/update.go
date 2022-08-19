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
		if plr.up {
			plr.dst.Y -= plr.speed * deltaTime
		}

		if plr.down {
			plr.dst.Y += plr.speed * deltaTime
		}

		if plr.left {
			plr.dst.X -= plr.speed * deltaTime
		}

		if plr.right {
			plr.dst.X += plr.speed * deltaTime
		}
	}

	updateCamera(cam, plr)
	plr.moving = false
	plr.up, plr.down, plr.left, plr.right = false, false, false, false
}
