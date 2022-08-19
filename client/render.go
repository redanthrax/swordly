package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"strconv"

	"github.com/gen2brain/raylib-go/raylib"
)

var (
	bkgColor = rl.NewColor(34, 34, 34, 255)
	tileDest rl.Rectangle
	tileSrc rl.Rectangle
	tileMap []int
	srcMap []string
	mapW, mapH int
	tex rl.Texture2D
	grassSprite rl.Texture2D
)

func render() {
	rl.DrawFPS(0, 0)
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(*cam)
	draw()
	rl.EndMode2D()
	rl.EndDrawing()
}

func initMap() {
	tex = rl.LoadTexture("assets/tilesets/dungeon_parts.png")
	tileSrc = rl.NewRectangle(0, 0, 16, 16)
	tileDest = rl.NewRectangle(0, 0, 16, 16)
	loadMap("maps/arena1.map")
}

func draw() {
	for i := 0; i < len(tileMap); i++ {
		if tileMap[i] != 0 {
			tileSrc.X = tileSrc.Width * float32(
				(tileMap[i] - 1) % int(tex.Width / int32(tileSrc.Width)))
			tileSrc.Y = tileSrc.Height * float32(
				(tileMap[i] - 1) / int(tex.Width / int32(tileSrc.Width)))

			tileDest.X = tileDest.Width * float32(i % mapW)
			tileDest.Y = tileDest.Height * float32(i / mapW)

			rl.DrawTexturePro(
				tex, 
				tileSrc, 
				tileDest, 
				rl.NewVector2(tileDest.Width, tileDest.Height), 
				0, 
				rl.White)
		}
	}

	rl.DrawTexturePro(
		plr.sprite, 
		plr.src, 
		plr.dst, 
		rl.NewVector2(plr.dst.Width, plr.dst.Height), 
		0, 
		rl.White)

	rl.DrawRectangleRec(rect1, rl.Red)
	rl.DrawRectangle(int32(plr.dst.X) - 16, int32(plr.dst.Y) - 32, int32(plr.dst.Width), int32(plr.dst.Height), rl.White)
}

func loadMap(mapFile string) {
	file, err := ioutil.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	remNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(remNewLines, " ")
	mapW = -1
	mapH = -1
	for i := 0; i < len(sliced); i++ {
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		if mapW == -1 {
			mapW = m
		} else if mapH == -1 {
			mapH = m
		} else if i < (mapW * mapH + 2) {
			tileMap = append(tileMap, m)
		} else {
			srcMap = append(srcMap, sliced[i])
		}
	}

	if len(tileMap) > mapW * mapH {
		tileMap = tileMap[:len(tileMap) - 1]
	}
}
