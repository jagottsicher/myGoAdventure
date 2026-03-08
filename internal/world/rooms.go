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

// Room 0x00 in C++ — "Number Room" (Easter Egg / Purple)
var RoomNumberRoom = Room{
	RoomData:   RoomEasterEggGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xa8, 0x00, 0xff), // COLOR_PURPLE
}

// Room 0x01 in C++ — "Top Access" (olive, LEFTTHINWALL)
var RoomTopAccess1 = Room{
	RoomData:   RoomBelowYellowCastleGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xa1, 0xb0, 0x34), // COLOR_OLIVEGREEN
}

// Room 0x0A in C++ — "Maze Entry" (light gray)
var RoomMazeEntry = Room{
	RoomData:   RoomGfxMazeEntry,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x9a, 0x9a, 0x9a),
}

// Room 0x09 in C++ — "Maze Middle" (light gray)
var RoomMazeMiddle = Room{
	RoomData:   RoomGfxMazeMiddle,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x9a, 0x9a, 0x9a), // COLOR_LTGRAY
}

// Room 0x0B in C++ — "Maze Side" (light gray)
var RoomMazeSide = Room{
	RoomData:   RoomGfxMazeSide,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x9a, 0x9a, 0x9a),
}

// Room 0x0C in C++ — "Side Corridor" (light cyan)
var RoomSideCorridorCyan = Room{
	RoomData:   RoomGfxSideCorridor,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x55, 0xb6, 0xff), // COLOR_LTCYAN
}

// Room 0x0D in C++ — "Side Corridor" (dark green)
var RoomSideCorridorGreen = Room{
	RoomData:   RoomGfxSideCorridor,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x21, 0xd9, 0x1b), // COLOR_DKGREEN
}

// Room 0x0E in C++ — "Top Entry Room" above White Castle (cyan)
var RoomWhiteCastleTop = Room{
	RoomData:   RoomTopEntryRoomGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x61, 0xd0, 0x70), // COLOR_CYAN
}

// Room 0x0F in C++ — "White Castle"
var RoomWhiteCastle = Room{
	RoomData:   RoomGfxCastle,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.ColorWhite,
}

// Room 0x12 in C++ — "Yellow Castle Entry"
var RoomYellowCastleEntry = Room{
	RoomData:   RoomGfxNumberRoom,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.ColorYellow,
}

// Room 0x13 in C++ — "Black Maze #1" (orange per user request)
var RoomBlackMaze1 = Room{
	RoomData:   RoomBlackMaze1Gfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x70, 0x00), // orange
}

// Room 0x14 in C++ — "Black Maze #2" (orange per user request)
var RoomBlackMaze2 = Room{
	RoomData:   RoomBlackMaze2Gfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x70, 0x00),
}

// Room 0x15 in C++ — "Black Maze #3" (orange per user request)
var RoomBlackMaze3 = Room{
	RoomData:   RoomBlackMaze3Gfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x70, 0x00),
}

// Room 0x16 in C++ — "Black Maze Entry" (orange per user request)
var RoomBlackMazeEntry = Room{
	RoomData:   RoomBlackMazeEntryGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x70, 0x00),
}

// Room 0x1B in C++ — "Black Castle Entry"
var RoomBlackCastleEntry = Room{
	RoomData:   RoomTopEntryRoomGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x28, 0x20), // COLOR_RED
}

// Room 0x1C in C++ — "Other Purple Room"
var RoomOtherPurpleRoom = Room{
	RoomData:   RoomGfxNumberRoom,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xa8, 0x00, 0xff), // COLOR_PURPLE
}

// Room 0x1E in C++ — "Name Room" (purple)
var RoomNameRoom = Room{
	RoomData:   RoomBelowYellowCastleGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xa8, 0x00, 0xff),
}

// Room 0x17 in C++ — "Red Maze #1"
var RoomRedMaze1 = Room{
	RoomData:   RoomRedMaze1Gfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x28, 0x20), // COLOR_RED
}

// Room 0x18 in C++ — "Top of Red Maze"
var RoomRedMazeTop = Room{
	RoomData:   RoomRedMazeTopGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x28, 0x20),
}

// Room 0x19 in C++ — "Bottom of Red Maze"
var RoomRedMazeBottom = Room{
	RoomData:   RoomRedMazeBottomGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x28, 0x20),
}

// Room 0x1A in C++ — "White Castle Entry" (red maze)
var RoomWhiteCastleEntry = Room{
	RoomData:   RoomWhiteCastleEntryGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x28, 0x20),
}

// Room 0x1D in C++ — "Top Entry Room" (red), above Black Castle
var RoomBlackCastleTop = Room{
	RoomData:   RoomBlackCastleTopGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd0, 0x28, 0x20), // COLOR_RED from original
}

// Room 0x10 in C++ — "Black Castle"
var RoomBlackCastle = Room{
	RoomData:   RoomGfxCastle,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x18, 0x18, 0x18), // COLOR_BLACK — dark walls on gray bg
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

// Room 6 in C++ — "Bottom of Blue Maze"
var RoomBlueMazeBottom = Room{
	RoomData:   RoomBlueMazeBottomGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x6b, 0x64, 0xff), // COLOR_BLUE from original
}

// Room 4 in C++ — "Top of Blue Maze"
var RoomBlueMazeTop = Room{
	RoomData:   RoomBlueMazeTopGfx,
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
var RoomCorridorRight = Room{
	RoomData:   RoomCorridorRightGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xd5, 0xb5, 0x43), // COLOR_TAN from original
}

// Left of RoomMazeSide — same graphic as SideCorridor, olivegreen (same as RoomBelowYellowCastle)
var RoomSideCorridorOlive = Room{
	RoomData:   RoomGfxSideCorridor,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0xa1, 0xb0, 0x34), // COLOR_OLIVEGREEN
}

// Below RoomSideCorridorOlive — only open upward, dark green
var RoomDeadEndDarkGreen = Room{
	RoomData:   RoomTopEntryRoomGfx,
	Background: tcell.ColorDarkGray,
	Foreground: tcell.NewRGBColor(0x0c, 0x6e, 0x0c), // dark green
}

func InitDirections() {
	// --- Originalverbindungen (vor letztem Arbeitsschritt) ---

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
	RoomBelowYellowCastle.Right = &RoomCorridorRight

	RoomTopAccessRight.Up = &RoomBlueMazeEntry
	RoomTopAccessRight.Down = nil
	RoomTopAccessRight.Left = nil
	RoomTopAccessRight.Right = &RoomBelowYellowCastle

	// Room 8 — Blue Maze Entry
	RoomBlueMazeEntry.Up = &RoomBlueMaze1
	RoomBlueMazeEntry.Down = &RoomTopAccessRight
	RoomBlueMazeEntry.Left = &RoomBlueMazeCenter
	RoomBlueMazeEntry.Right = &RoomBlueMazeCenter

	// Room 5 — Blue Maze #1
	RoomBlueMaze1.Up = nil
	RoomBlueMaze1.Down = &RoomBlueMazeEntry
	RoomBlueMaze1.Left = &RoomBlueMazeTop
	RoomBlueMaze1.Right = &RoomBlueMazeBottom

	// Room 6 — Blue Maze Bottom
	RoomBlueMazeBottom.Up = &RoomBlueMazeCenter
	RoomBlueMazeBottom.Down = &RoomCorridorRight
	RoomBlueMazeBottom.Left = &RoomBlueMaze1
	RoomBlueMazeBottom.Right = &RoomBlueMazeTop

	// Room 0x17 — Red Maze #1
	RoomRedMaze1.Up = &RoomRedMazeBottom
	RoomRedMaze1.Right = &RoomRedMazeTop
	RoomRedMaze1.Down = &RoomRedMazeBottom
	RoomRedMaze1.Left = &RoomRedMazeTop

	// Room 0x18 — Top of Red Maze
	RoomRedMazeTop.Up = &RoomWhiteCastleEntry
	RoomRedMazeTop.Right = &RoomRedMaze1
	RoomRedMazeTop.Down = &RoomWhiteCastleEntry
	RoomRedMazeTop.Left = &RoomRedMaze1

	// Room 0x19 — Bottom of Red Maze
	RoomRedMazeBottom.Up = &RoomRedMaze1
	RoomRedMazeBottom.Right = &RoomWhiteCastleEntry
	RoomRedMazeBottom.Down = &RoomRedMaze1
	RoomRedMazeBottom.Left = &RoomWhiteCastleEntry

	// Room 0x1A — White Castle Entry (Eingang zum Red Maze, über White Castle)
	RoomWhiteCastleEntry.Up = &RoomRedMazeTop
	RoomWhiteCastleEntry.Right = &RoomRedMazeBottom
	RoomWhiteCastleEntry.Down = &RoomWhiteCastle
	RoomWhiteCastleEntry.Left = &RoomRedMazeBottom

	// Room 0x1D — Black Castle Top
	RoomBlackCastleTop.Up = &RoomBlackMazeEntry
	RoomBlackCastleTop.Down = &RoomBlackCastle
	RoomBlackCastleTop.Left = nil
	RoomBlackCastleTop.Right = nil

	// Room 0x10 — Black Castle
	RoomBlackCastle.Up = &RoomBlackCastleTop
	RoomBlackCastle.Down = &RoomBlueMazeTop
	RoomBlackCastle.Left = nil
	RoomBlackCastle.Right = nil

	// Room 4 — Blue Maze Top
	RoomBlueMazeTop.Up = &RoomBlackCastle
	RoomBlueMazeTop.Down = &RoomBlueMazeCenter
	RoomBlueMazeTop.Left = &RoomBlueMazeBottom
	RoomBlueMazeTop.Right = &RoomBlueMaze1

	// Room 7 — Blue Maze Center
	RoomBlueMazeCenter.Up = &RoomBlueMazeTop
	RoomBlueMazeCenter.Down = &RoomBlueMazeBottom
	RoomBlueMazeCenter.Left = &RoomBlueMazeEntry
	RoomBlueMazeCenter.Right = &RoomBlueMazeEntry

	RoomCorridorRight.Up = nil
	RoomCorridorRight.Down = &RoomMazeEntry
	RoomCorridorRight.Left = &RoomBelowYellowCastle
	RoomCorridorRight.Right = &RoomNumberRoom

	// --- Neu hinzugefügte Räume: alle disconnected ---

	RoomTopAccess1.Up = nil
	RoomTopAccess1.Down = nil
	RoomTopAccess1.Left = nil
	RoomTopAccess1.Right = nil

	RoomMazeEntry.Up = &RoomCorridorRight
	RoomMazeEntry.Down = &RoomMazeMiddle
	RoomMazeEntry.Left = &RoomMazeMiddle
	RoomMazeEntry.Right = &RoomMazeMiddle

	RoomMazeMiddle.Up = &RoomMazeEntry
	RoomMazeMiddle.Down = &RoomMazeSide
	RoomMazeMiddle.Left = &RoomMazeEntry
	RoomMazeMiddle.Right = &RoomMazeEntry

	RoomMazeSide.Up = &RoomMazeMiddle
	RoomMazeSide.Down = nil
	RoomMazeSide.Left = &RoomSideCorridorOlive
	RoomMazeSide.Right = &RoomSideCorridorCyan

	RoomSideCorridorCyan.Up = &RoomOtherPurpleRoom
	RoomSideCorridorCyan.Down = &RoomBlackCastleEntry
	RoomSideCorridorCyan.Left = &RoomMazeSide
	RoomSideCorridorCyan.Right = nil

	RoomSideCorridorGreen.Up = nil
	RoomSideCorridorGreen.Down = nil
	RoomSideCorridorGreen.Left = nil
	RoomSideCorridorGreen.Right = nil

	RoomNumberRoom.Up = nil
	RoomNumberRoom.Down = nil
	RoomNumberRoom.Left = &RoomCorridorRight
	RoomNumberRoom.Right = nil

	RoomWhiteCastleTop.Up = nil
	RoomWhiteCastleTop.Down = nil
	RoomWhiteCastleTop.Left = nil
	RoomWhiteCastleTop.Right = nil

	RoomWhiteCastle.Up = &RoomWhiteCastleEntry
	RoomWhiteCastle.Down = &RoomSideCorridorOlive
	RoomWhiteCastle.Left = nil
	RoomWhiteCastle.Right = nil

	RoomYellowCastleEntry.Up = nil
	RoomYellowCastleEntry.Down = nil
	RoomYellowCastleEntry.Left = nil
	RoomYellowCastleEntry.Right = nil

	// Room 0x13 — Black Maze #1
	RoomBlackMaze1.Up = &RoomBlackMaze3
	RoomBlackMaze1.Right = &RoomBlackMaze2
	RoomBlackMaze1.Down = &RoomBlackMaze3
	RoomBlackMaze1.Left = &RoomBlackMazeEntry

	// Room 0x14 — Black Maze #2
	RoomBlackMaze2.Up = &RoomBlackMazeEntry
	RoomBlackMaze2.Right = &RoomBlackMaze3
	RoomBlackMaze2.Down = &RoomBlackMazeEntry
	RoomBlackMaze2.Left = &RoomBlackMaze1

	// Room 0x15 — Black Maze #3
	RoomBlackMaze3.Up = &RoomBlackMaze1
	RoomBlackMaze3.Right = &RoomBlackMazeEntry
	RoomBlackMaze3.Down = &RoomBlackMaze1
	RoomBlackMaze3.Left = &RoomBlackMaze2

	// Room 0x16 — Black Maze Entry
	RoomBlackMazeEntry.Up = &RoomBlackMaze2
	RoomBlackMazeEntry.Right = &RoomBlackMaze1
	RoomBlackMazeEntry.Down = &RoomBlackCastleTop
	RoomBlackMazeEntry.Left = &RoomBlackMaze3

	RoomBlackCastleEntry.Up = &RoomSideCorridorCyan
	RoomBlackCastleEntry.Down = nil
	RoomBlackCastleEntry.Left = nil
	RoomBlackCastleEntry.Right = nil

	RoomOtherPurpleRoom.Up = nil
	RoomOtherPurpleRoom.Down = &RoomSideCorridorCyan
	RoomOtherPurpleRoom.Left = nil
	RoomOtherPurpleRoom.Right = nil

	RoomNameRoom.Up = nil
	RoomNameRoom.Down = nil
	RoomNameRoom.Left = nil
	RoomNameRoom.Right = nil

	RoomSideCorridorOlive.Up = &RoomWhiteCastle
	RoomSideCorridorOlive.Down = &RoomDeadEndDarkGreen
	RoomSideCorridorOlive.Left = nil
	RoomSideCorridorOlive.Right = &RoomMazeSide

	RoomDeadEndDarkGreen.Up = &RoomSideCorridorOlive
	RoomDeadEndDarkGreen.Down = nil
	RoomDeadEndDarkGreen.Left = nil
	RoomDeadEndDarkGreen.Right = nil
}
