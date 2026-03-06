package main

import (
	"github.com/gdamore/tcell/v2"
)

type room struct {
	roomData   *[]string
	background tcell.Color
	foreground tcell.Color
	barLeft    bool
	barRight   bool
	up         *room
	down       *room
	left       *room
	right      *room
}

var roomSplashScreen = room{
	roomData:   roomGfxCastle,
	background: tcell.ColorDarkGray,
	foreground: tcell.ColorAntiqueWhite,
}

var roomYellowCastle = room{
	roomData:   roomGfxCastle,
	background: tcell.ColorDarkGray,
	foreground: tcell.ColorYellow,
}

func initDirections() {
	roomYellowCastle.up = nil
	roomYellowCastle.down = nil
	roomYellowCastle.left = nil
	roomYellowCastle.right = nil
}
