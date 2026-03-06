package world

import "github.com/gdamore/tcell/v2"

var RoomSplashScreen = Room{
	RoomData:   RoomGfxCastle,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.ColorAntiqueWhite,
}

var RoomYellowCastle = Room{
	RoomData:   RoomGfxCastle,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.ColorYellow,
}

var RoomAboveYellowCastle = Room{
	RoomData:   RoomAboveYellowCastleGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.ColorYellow,
}

var RoomBelowYellowCastle = Room{
	RoomData:   RoomBelowYellowCastleGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xa1, 0xb0, 0x34), // COLOR_OLIVEGREEN from original
}

// Room 2 in C++ — "Top Access", left of RoomBelowYellowCastle
var RoomTopAccessRight = Room{
	RoomData:   RoomBelowYellowCastleGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x86, 0xd9, 0x22), // COLOR_LIMEGREEN from original
}

// Room 3 in C++ — "Left of Name", right of RoomBelowYellowCastle
var RoomLeftOfName = Room{
	RoomData:   RoomLeftOfNameGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd5, 0xb5, 0x43), // COLOR_TAN from original
}

func InitDirections() {
	RoomYellowCastle.Up = &RoomAboveYellowCastle
	RoomYellowCastle.Down = &RoomBelowYellowCastle
	RoomYellowCastle.Left = nil
	RoomYellowCastle.Right = nil

	RoomAboveYellowCastle.Up = nil
	RoomAboveYellowCastle.Down = &RoomYellowCastle
	RoomAboveYellowCastle.Left = nil
	RoomAboveYellowCastle.Right = nil

	RoomBelowYellowCastle.Up = &RoomYellowCastle
	RoomBelowYellowCastle.Down = nil
	RoomBelowYellowCastle.Left = &RoomTopAccessRight
	RoomBelowYellowCastle.Right = &RoomLeftOfName

	RoomTopAccessRight.Up = nil
	RoomTopAccessRight.Down = nil
	RoomTopAccessRight.Left = nil
	RoomTopAccessRight.Right = &RoomBelowYellowCastle

	RoomLeftOfName.Up = nil
	RoomLeftOfName.Down = nil
	RoomLeftOfName.Left = &RoomBelowYellowCastle
	RoomLeftOfName.Right = nil
}
