package main

import (
	"github.com/gdamore/tcell/v2"
)

type cell struct {
	x, y   int
	symbol rune
}

type object struct {
	// room   *rooms
	posX   float64
	posY   float64
	stepX  float64
	stepY  float64
	width  int
	height int
	style  tcell.Style
	shape  []*cell
}

func initPlayer() {
	template := *roomYellowCastle.compressedRoomData
	templateH := float64(len(template))
	templateW := float64(len([]rune(template[0])))

	player = &object{
		posX:   templateW / 2.0,
		posY:   templateH / 3.0 * 2.0,
		width:  2,
		height: 1,
		stepX:  2.0,
		stepY:  1.0,
		style:  tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorPurple),
		shape:  playerGfx,
	}

	if int(player.posX)%2 != 0 {
		player.posX += 1.0
	}

	allObjects = append(allObjects, player)
}
