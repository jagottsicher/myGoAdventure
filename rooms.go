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

var roomAboveYellowCastle = room{
	roomData:   roomAboveYellowCastleGfx,
	background: tcell.ColorDarkGray,
	foreground: tcell.ColorYellow,
}

var roomBelowYellowCastle = room{
	roomData:   roomBelowYellowCastleGfx,
	background: tcell.ColorDarkGray,
	foreground: tcell.NewRGBColor(0xa1, 0xb0, 0x34), // COLOR_OLIVEGREEN from original
}

// Room 2 in C++ — "Top Access", left of roomBelowYellowCastle
var roomTopAccessRight = room{
	roomData:   roomBelowYellowCastleGfx, // same layout as below-yellow-castle
	background: tcell.ColorDarkGray,
	foreground: tcell.NewRGBColor(0x86, 0xd9, 0x22), // COLOR_LIMEGREEN from original
}

// Room 3 in C++ — "Left of Name", right of roomBelowYellowCastle
var roomLeftOfName = room{
	roomData:   roomLeftOfNameGfx,
	background: tcell.ColorDarkGray,
	foreground: tcell.NewRGBColor(0xd5, 0xb5, 0x43), // COLOR_TAN from original
}

func initDirections() {
	roomYellowCastle.up = &roomAboveYellowCastle
	roomYellowCastle.down = &roomBelowYellowCastle
	roomYellowCastle.left = nil
	roomYellowCastle.right = nil

	roomAboveYellowCastle.up = nil
	roomAboveYellowCastle.down = &roomYellowCastle
	roomAboveYellowCastle.left = nil
	roomAboveYellowCastle.right = nil

	roomBelowYellowCastle.up = &roomYellowCastle
	roomBelowYellowCastle.down = nil
	roomBelowYellowCastle.left = &roomTopAccessRight
	roomBelowYellowCastle.right = &roomLeftOfName

	roomTopAccessRight.up = nil
	roomTopAccessRight.down = nil
	roomTopAccessRight.left = nil
	roomTopAccessRight.right = &roomBelowYellowCastle

	roomLeftOfName.up = nil
	roomLeftOfName.down = nil
	roomLeftOfName.left = &roomBelowYellowCastle
	roomLeftOfName.right = nil
}
