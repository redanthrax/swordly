package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func colliding() bool {
	if rl.CheckCollisionRecs(plr.dst, rect1) {
		fmt.Println("Collided")
		return true
	}

	fmt.Println("NOT")

	return false
}
