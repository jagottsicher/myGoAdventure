package world

import "github.com/gdamore/tcell/v2"

// Room 0x1E in C++ — "Name Room" (Easter Egg, zeigt "Warren Robinett" in Originalfarbe FLASH)
// In Go als Splashscreen genutzt — gleiche Rolle: Sonder-Anzeigeraum außerhalb des Spielablaufs
var RoomSplashScreen = Room{
	RoomData:   RoomGfxCastle,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xA2, 0x51, 0xD9), // COLOR_PURPLE
}

// Room 0x11 in C++ — "Yellow Castle"
var RoomYellowCastle = Room{
	RoomData:   RoomGfxCastle,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xFF, 0xD8, 0x4C), // COLOR_YELLOW
}

// Room 0x12 in C++ — "Yellow Castle Entry" (Schloss-Innenraum, alle Verbindungen self-loop solange gesperrt)
var RoomAboveYellowCastle = Room{
	RoomData:   RoomAboveYellowCastleGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xFF, 0xD8, 0x4C), // COLOR_YELLOW
}

// Room 0x02 in C++ — "Top Access" (limegreen), direkt unter Yellow Castle
var RoomBelowYellowCastle = Room{
	RoomData:   RoomBelowYellowCastleGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x86, 0xd9, 0x22), // COLOR_LIMEGREEN
}

// Room 0x00 in C++ — "Number Room" (Game-Select-Screen, zeigt Spielvariante 1/2/3)
var RoomNumberRoom = Room{
	RoomData:   RoomEasterEggGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xA2, 0x51, 0xD9), // COLOR_PURPLE
}

// Room 0x0A in C++ — "Maze Entry" (light gray, unsichtbare Wände — Originalverhalten)
var RoomMazeEntry = Room{
	RoomData:   RoomGfxMazeEntry,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xcd, 0xcd, 0xcd), // COLOR_LTGRAY
}

// Room 0x09 in C++ — "Maze Middle" (light gray, unsichtbare Wände — Originalverhalten)
var RoomMazeMiddle = Room{
	RoomData:   RoomGfxMazeMiddle,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xcd, 0xcd, 0xcd), // COLOR_LTGRAY
}

// Room 0x0B in C++ — "Maze Side" (light gray, unsichtbare Wände — Originalverhalten)
var RoomMazeSide = Room{
	RoomData:   RoomGfxMazeSide,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xcd, 0xcd, 0xcd), // COLOR_LTGRAY
}

// Room 0x0C in C++ — "Side Corridor" (light cyan)
var RoomSideCorridorCyan = Room{
	RoomData:   RoomGfxSideCorridor,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x55, 0xb6, 0xff), // COLOR_LTCYAN
}

// Room 0x0F in C++ — "White Castle"
var RoomWhiteCastle = Room{
	RoomData:   RoomGfxCastle,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xff, 0xff, 0xff), // COLOR_WHITE
}

// Room 0x13 in C++ — "Black Maze #1" (light gray, unsichtbare Wände — Originalverhalten)
var RoomBlackMaze1 = Room{
	RoomData:   RoomBlackMaze1Gfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xcd, 0xcd, 0xcd), // COLOR_LTGRAY
}

// Room 0x14 in C++ — "Black Maze #2" (light gray, unsichtbare Wände — Originalverhalten)
var RoomBlackMaze2 = Room{
	RoomData:   RoomBlackMaze2Gfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xcd, 0xcd, 0xcd), // COLOR_LTGRAY
}

// Room 0x15 in C++ — "Black Maze #3" (light gray, unsichtbare Wände — Originalverhalten)
var RoomBlackMaze3 = Room{
	RoomData:   RoomBlackMaze3Gfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xcd, 0xcd, 0xcd), // COLOR_LTGRAY
}

// Room 0x16 in C++ — "Black Maze Entry" (light gray, unsichtbare Wände — Originalverhalten)
var RoomBlackMazeEntry = Room{
	RoomData:   RoomBlackMazeEntryGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xcd, 0xcd, 0xcd), // COLOR_LTGRAY
}

// Room 0x1B in C++ — "Black Castle Entry"
var RoomBlackCastleEntry = Room{
	RoomData:   RoomTopEntryRoomGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xfa, 0x52, 0x55), // COLOR_RED
}

// Room 0x1C in C++ — "Other Purple Room"
var RoomOtherPurpleRoom = Room{
	RoomData:   RoomGfxNumberRoom,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xA2, 0x51, 0xD9), // COLOR_PURPLE
}

// Room 0x17 in C++ — "Red Maze #1"
var RoomRedMaze1 = Room{
	RoomData:   RoomRedMaze1Gfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xfa, 0x52, 0x55), // COLOR_RED
}

// Room 0x18 in C++ — "Top of Red Maze"
var RoomRedMazeTop = Room{
	RoomData:   RoomRedMazeTopGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xfa, 0x52, 0x55), // COLOR_RED
}

// Room 0x19 in C++ — "Bottom of Red Maze"
var RoomRedMazeBottom = Room{
	RoomData:   RoomRedMazeBottomGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xfa, 0x52, 0x55), // COLOR_RED
}

// Room 0x1A in C++ — "White Castle Entry" (red maze)
var RoomWhiteCastleEntry = Room{
	RoomData:   RoomWhiteCastleEntryGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xfa, 0x52, 0x55), // COLOR_RED
}

// Room 0x1D in C++ — "Top Entry Room" (red), above Black Castle
var RoomBlackCastleTop = Room{
	RoomData:   RoomBlackCastleTopGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xfa, 0x52, 0x55), // COLOR_RED
}

// Room 0x10 in C++ — "Black Castle"
var RoomBlackCastle = Room{
	RoomData:   RoomGfxCastle,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x00, 0x00, 0x00), // COLOR_BLACK
}

// Room 0x08 in C++ — "Blue Maze Entry"
var RoomBlueMazeEntry = Room{
	RoomData:   RoomBlueMazeEntryGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x6b, 0x64, 0xff), // COLOR_BLUE
}

// Room 0x01 in C++ — "Top Access" (olivegreen), links von RoomBelowYellowCastle
var RoomTopAccessRight = Room{
	RoomData:   RoomBelowYellowCastleGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xa1, 0xb0, 0x34), // COLOR_OLIVEGREEN
}

// Room 0x05 in C++ — "Blue Maze #1"
var RoomBlueMaze1 = Room{
	RoomData:   RoomBlueMaze1Gfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x6b, 0x64, 0xff), // COLOR_BLUE
}

// Room 0x06 in C++ — "Bottom of Blue Maze"
var RoomBlueMazeBottom = Room{
	RoomData:   RoomBlueMazeBottomGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x6b, 0x64, 0xff), // COLOR_BLUE
}

// Room 0x04 in C++ — "Top of Blue Maze"
var RoomBlueMazeTop = Room{
	RoomData:   RoomBlueMazeTopGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x6b, 0x64, 0xff), // COLOR_BLUE
}

// Room 0x07 in C++ — "Center of Blue Maze"
var RoomBlueMazeCenter = Room{
	RoomData:   RoomBlueMazeCenterGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x6b, 0x64, 0xff), // COLOR_BLUE
}

// Room 0x03 in C++ — "Left of Name"
var RoomCorridorRight = Room{
	RoomData:   RoomCorridorRightGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0xd5, 0xb5, 0x43), // COLOR_TAN
}

// Room 0x0D in C++ — "Side Corridor" (COLOR_DKGREEN), links von RoomMazeSide, über RoomWhiteCastle
// C++-Verbindungen: up=0x0F(WhiteCastle), right=0x0B(MazeSide), down=0x0E(TopEntryRoom CYAN — fehlt in Go!), left=0x0C(SideCorridorCyan)
var RoomSideCorridorOlive = Room{
	RoomData:   RoomGfxSideCorridor,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x21, 0xd9, 0x1b), // COLOR_DKGREEN
}

// Room 0x0E in C++ — "Top Entry Room" (COLOR_CYAN), direkt unter RoomSideCorridorOlive (0x0D)
// C++-Verbindungen: up=0x0D, right=0x10(BlackCastle), down=0x0F(WhiteCastle), left=0x10(BlackCastle)
var RoomDeadEndCyan = Room{
	RoomData:   RoomTopEntryRoomGfx,
	Background: tcell.NewRGBColor(0xcd, 0xcd, 0xcd),
	Foreground: tcell.NewRGBColor(0x55, 0xb6, 0xff), // COLOR_LTCYAN
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

	RoomNumberRoom.Up = nil
	RoomNumberRoom.Down = nil
	RoomNumberRoom.Left = &RoomCorridorRight
	RoomNumberRoom.Right = nil

	RoomWhiteCastle.Up = &RoomWhiteCastleEntry
	RoomWhiteCastle.Down = &RoomSideCorridorOlive
	RoomWhiteCastle.Left = nil
	RoomWhiteCastle.Right = nil

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

	RoomSideCorridorOlive.Up = &RoomWhiteCastle
	RoomSideCorridorOlive.Down = &RoomDeadEndCyan
	RoomSideCorridorOlive.Left = nil
	RoomSideCorridorOlive.Right = &RoomMazeSide

	RoomDeadEndCyan.Up = &RoomSideCorridorOlive
	RoomDeadEndCyan.Down = nil
	RoomDeadEndCyan.Left = nil
	RoomDeadEndCyan.Right = nil
}
