package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	camera rl.Camera2D
)

func newCamera(player *player) *rl.Camera2D {
	c := rl.Camera2D {}
	c.Target = rl.NewVector2(float32(player.dst.X + 8), float32(player.dst.Y + 16))
	c.Offset = rl.NewVector2(float32(screenWidth / 2), float32(screenHeight / 2))
	c.Rotation = 0.0
	c.Zoom = 3.0

	return &c
}

func updateCamera(camera *rl.Camera2D, player *player) {
	camera.Target = rl.NewVector2(float32(player.dst.X + 8), float32(player.dst.Y + 16))
}
