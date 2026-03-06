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

// Room 8 in C++ — "Blue Maze Entry", above RoomTopAccessRight
var RoomBlueMazeEntry = Room{
	RoomData:   RoomBlueMazeEntryGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x6b, 0x64, 0xff), // COLOR_BLUE from original
}

// Room 2 in C++ — "Top Access", left of RoomBelowYellowCastle
var RoomTopAccessRight = Room{
	RoomData:   RoomBelowYellowCastleGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x86, 0xd9, 0x22), // COLOR_LIMEGREEN from original
}

// Room 5 in C++ — "Blue Maze #1"
var RoomBlueMaze1 = Room{
	RoomData:   RoomBlueMaze1Gfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x6b, 0x64, 0xff), // COLOR_BLUE from original
}

// Room 7 in C++ — "Center of Blue Maze"
var RoomBlueMazeCenter = Room{
	RoomData:   RoomBlueMazeCenterGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x6b, 0x64, 0xff), // COLOR_BLUE from original
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

	RoomTopAccessRight.Up = &RoomBlueMazeEntry
	RoomTopAccessRight.Down = nil
	RoomTopAccessRight.Left = nil
	RoomTopAccessRight.Right = &RoomBelowYellowCastle

	RoomBlueMazeEntry.Up = &RoomBlueMaze1
	RoomBlueMazeEntry.Down = &RoomTopAccessRight
	RoomBlueMazeEntry.Left = &RoomBlueMazeCenter
	RoomBlueMazeEntry.Right = &RoomBlueMazeCenter

	RoomBlueMaze1.Up = nil
	RoomBlueMaze1.Down = &RoomBlueMazeEntry
	RoomBlueMaze1.Left = nil
	RoomBlueMaze1.Right = nil

	RoomBlueMazeCenter.Up = nil
	RoomBlueMazeCenter.Down = nil
	RoomBlueMazeCenter.Left = &RoomBlueMazeEntry
	RoomBlueMazeCenter.Right = &RoomBlueMazeEntry

	RoomLeftOfName.Up = nil
	RoomLeftOfName.Down = nil
	RoomLeftOfName.Left = &RoomBelowYellowCastle
	RoomLeftOfName.Right = nil
}
