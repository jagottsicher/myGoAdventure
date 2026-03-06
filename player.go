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
	relX   float64 // position as fraction of screen (0.0–1.0), converted at render time
	relY   float64
	stepX  int // movement step in terminal columns/rows (converted to relative at move time)
	stepY  int
	width  int
	height int
	style  tcell.Style
	shape  []*cell
}

func initPlayer() {
	width, height := screen.Size()

	player = &object{
		relX:   float64(width/2) / float64(width),
		relY:   float64(height/3*2) / float64(height),
		width:  2,
		height: 1,
		stepX:  2,
		stepY:  1,
		style:  tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorPurple),
		shape:  playerGfx,
	}

	allObjects = append(allObjects, player)


}
