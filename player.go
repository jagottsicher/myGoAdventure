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
	posX   int
	posY   int
	stepX  int
	stepY  int
	width  int
	height int
	style  tcell.Style
	shape  []*cell
}

func initPlayer() {
	width, height := screen.Size()

	player = &object{
		posY:   height / 3 * 2,
		posX:   width / 2,
		width:  2,
		height: 1,
		stepX:  2,
		stepY:  1,
		style:  tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorPurple),
		shape:  playerGfx,
	}

	if player.posX%2 != 0 {
		player.posX += 1
	}

	allObjects = append(allObjects, player)

}
