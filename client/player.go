package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type player struct {
	sprite rl.Texture2D
	src rl.Rectangle
	dst rl.Rectangle
	name string
	hp int
	speed float32
	stamina int
	damage int
	defense int
	moving bool
	direction int 
	up, down, left, right bool
}

func newPlayer(name string, choice int) *player {
	p := player {
		name: name,
		hp: 20,
		speed: 100.0,
		stamina: 5,
		damage: 5,
		defense: 0,
	}

	switch choice {
		case 1:
			p.src = rl.NewRectangle(0, 0, 16, 32)
			p.dst = rl.NewRectangle(0, 0, 16, 32)
			p.sprite = rl.LoadTexture("assets/sprites/female.png")
		case 2:
			p.src = rl.NewRectangle(0, 0, 16, 32)
			p.dst = rl.NewRectangle(0, 0, 16, 32)
			p.sprite = rl.LoadTexture("assets/sprites/male.png")
		case 3:
			p.src = rl.NewRectangle(8 * 16, 2 * 32, 16, 32)
			p.dst = rl.NewRectangle(0, 0, 60, 120)
		case 4:
			p.src = rl.NewRectangle(8 * 16, 3 * 32, 16, 32)
			p.dst = rl.NewRectangle(0, 0, 60, 120)
		case 5:
			p.src = rl.NewRectangle(8 * 16, 4 * 32, 16, 32)
			p.dst = rl.NewRectangle(0, 0, 60, 120)
		case 6:
			p.src = rl.NewRectangle(8 * 16, 5 * 32, 16, 32)
			p.dst = rl.NewRectangle(0, 0, 60, 120)
		case 7:
			p.src = rl.NewRectangle(8 * 16, 6 * 32, 16, 32)
			p.dst = rl.NewRectangle(0, 0, 60, 120)
		default:
			p.src = rl.NewRectangle(8 * 16, 3 * 32, 16, 32)
			p.dst = rl.NewRectangle(0, 0, 60, 120)
	}

	return &p
}
