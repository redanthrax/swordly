package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	bkgColor = rl.NewColor(34, 34, 34, 255)
)

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)
	draw()
	rl.EndMode2D()
	rl.EndDrawing()
}

func draw() {
	rl.DrawTexturePro(
		plr.sprite, 
		plr.src, 
		plr.dst, 
		rl.NewVector2(plr.dst.Width, plr.dst.Height), 
		0, 
		rl.White)
}
