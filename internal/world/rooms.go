package world

import "github.com/gdamore/tcell/v2"

// Room 0x1E in C++ — "Name Room" (Easter Egg, zeigt "Warren Robinett" in Originalfarbe FLASH)
// In Go als Splashscreen genutzt — gleiche Rolle: Sonder-Anzeigeraum außerhalb des Spielablaufs
// Room 0x1E in C++ — "Name Room" / Easter Egg room. Uses same graphic as BelowYellowCastle.
var RoomSplashScreen = Room{
	RoomData:   RoomBelowYellowCastleGfx,
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

// RoomsByID maps C++ room IDs (0x00–0x1E) to Go room pointers.
// Used for variation-3 random object placement (mirrors C++ roomBoundsData logic).
var RoomsByID = [0x1F]*Room{
	0x00: &RoomNumberRoom,
	0x01: &RoomTopAccessRight,
	0x02: &RoomBelowYellowCastle,
	0x03: &RoomCorridorRight,
	0x04: &RoomBlueMazeTop,
	0x05: &RoomBlueMaze1,
	0x06: &RoomBlueMazeBottom,
	0x07: &RoomBlueMazeCenter,
	0x08: &RoomBlueMazeEntry,
	0x09: &RoomMazeMiddle,
	0x0A: &RoomMazeEntry,
	0x0B: &RoomMazeSide,
	0x0C: &RoomSideCorridorCyan,
	0x0D: &RoomSideCorridorOlive,
	0x0E: &RoomDeadEndCyan,
	0x0F: &RoomWhiteCastle,
	0x10: &RoomBlackCastle,
	0x11: &RoomYellowCastle,
	0x12: &RoomAboveYellowCastle,
	0x13: &RoomBlackMaze1,
	0x14: &RoomBlackMaze2,
	0x15: &RoomBlackMaze3,
	0x16: &RoomBlackMazeEntry,
	0x17: &RoomRedMaze1,
	0x18: &RoomRedMazeTop,
	0x19: &RoomRedMazeBottom,
	0x1A: &RoomWhiteCastleEntry,
	0x1B: &RoomBlackCastleEntry,
	0x1C: &RoomOtherPurpleRoom,
	0x1D: &RoomBlackCastleTop,
	0x1E: &RoomSplashScreen,
}

func InitDirections(gameType int) {
	// --- Originalverbindungen (vor letztem Arbeitsschritt) ---

	RoomYellowCastle.Up = &RoomAboveYellowCastle
	RoomYellowCastle.Down = &RoomBelowYellowCastle
	RoomYellowCastle.Left = &RoomTopAccessRight
	RoomYellowCastle.Right = &RoomCorridorRight

	RoomAboveYellowCastle.Up = nil
	RoomAboveYellowCastle.Down = &RoomYellowCastle
	RoomAboveYellowCastle.Left = nil
	RoomAboveYellowCastle.Right = nil

	RoomBelowYellowCastle.Up = &RoomYellowCastle
	// V1: Down→BlueMaze1 (0x05) / V2+: Down→YellowCastle (0x11)
	if gameType == 1 {
		RoomBelowYellowCastle.Down = &RoomBlueMaze1
	} else {
		RoomBelowYellowCastle.Down = &RoomYellowCastle
	}
	RoomBelowYellowCastle.Left = &RoomTopAccessRight
	RoomBelowYellowCastle.Right = &RoomCorridorRight

	RoomTopAccessRight.Up = &RoomBlueMazeEntry
	// V1: Down→BlackCastle (0x10) / V2+: Down→WhiteCastle (0x0F)
	if gameType == 1 {
		RoomTopAccessRight.Down = &RoomBlackCastle
	} else {
		RoomTopAccessRight.Down = &RoomWhiteCastle
	}
	RoomTopAccessRight.Left = &RoomCorridorRight
	RoomTopAccessRight.Right = &RoomBelowYellowCastle

	// Room 8 — Blue Maze Entry
	RoomBlueMazeEntry.Up = &RoomBlueMaze1
	RoomBlueMazeEntry.Down = &RoomTopAccessRight
	RoomBlueMazeEntry.Left = &RoomBlueMazeCenter
	RoomBlueMazeEntry.Right = &RoomBlueMazeCenter

	// Room 5 — Blue Maze #1
	RoomBlueMaze1.Up = &RoomBlackCastleTop
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
	// UP: V1→CorridorRight (0x03) / V2+: stays as BlackMazeEntry (confirmed correct in V2)
	if gameType == 1 {
		RoomBlackCastleTop.Up = &RoomCorridorRight
	} else {
		RoomBlackCastleTop.Up = &RoomBlackMazeEntry
	}
	RoomBlackCastleTop.Down = &RoomBlackCastle
	RoomBlackCastleTop.Left = &RoomCorridorRight
	RoomBlackCastleTop.Right = &RoomTopAccessRight

	// Room 0x10 — Black Castle
	RoomBlackCastle.Up = &RoomBlackCastleTop
	RoomBlackCastle.Down = &RoomBlueMazeTop
	RoomBlackCastle.Left = &RoomOtherPurpleRoom
	RoomBlackCastle.Right = &RoomOtherPurpleRoom

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

	RoomCorridorRight.Up = &RoomBlueMazeBottom
	// V1: Down→BlackCastleEntry (0x1B) — gameplay: YellowCastle D→R→D = BlackCastleEntry
	// (mirrors V2: WhiteCastle D→R→R→D = BlackCastleEntry)
	// C++ raw 0x86 decodes to BlackCastleTop (0x1D) for V1, overridden by gameplay requirement.
	// V2+: Down→MazeEntry (0x0A)
	if gameType == 1 {
		RoomCorridorRight.Down = &RoomBlackCastleEntry
	} else {
		RoomCorridorRight.Down = &RoomMazeEntry
	}
	RoomCorridorRight.Left = &RoomBelowYellowCastle
	// V1: Right→TopAccessRight (C++ raw 0x01; Easter Egg not used in V1 navigation)
	// V2/V3: Right→SplashScreen (Easter Egg room, accessible from CorridorRight)
	if gameType == 1 {
		RoomCorridorRight.Right = &RoomTopAccessRight
	} else {
		RoomCorridorRight.Right = &RoomSplashScreen
	}

	RoomMazeEntry.Up = &RoomCorridorRight
	RoomMazeEntry.Down = &RoomMazeMiddle
	RoomMazeEntry.Left = &RoomMazeMiddle
	RoomMazeEntry.Right = &RoomMazeMiddle

	RoomMazeMiddle.Up = &RoomMazeEntry
	RoomMazeMiddle.Down = &RoomMazeSide
	RoomMazeMiddle.Left = &RoomMazeEntry
	RoomMazeMiddle.Right = &RoomMazeEntry

	RoomMazeSide.Up = &RoomMazeMiddle
	RoomMazeSide.Down = &RoomOtherPurpleRoom
	RoomMazeSide.Left = &RoomSideCorridorOlive
	RoomMazeSide.Right = &RoomSideCorridorCyan

	RoomSideCorridorCyan.Up = &RoomOtherPurpleRoom
	RoomSideCorridorCyan.Down = &RoomBlackCastleEntry
	RoomSideCorridorCyan.Left = &RoomMazeSide
	RoomSideCorridorCyan.Right = &RoomSideCorridorOlive

	RoomNumberRoom.Up = nil
	RoomNumberRoom.Down = nil
	RoomNumberRoom.Left = &RoomCorridorRight
	RoomNumberRoom.Right = nil

	RoomSplashScreen.Left = &RoomCorridorRight

	RoomWhiteCastle.Up = &RoomWhiteCastleEntry
	RoomWhiteCastle.Down = &RoomSideCorridorOlive
	RoomWhiteCastle.Left = &RoomWhiteCastle
	RoomWhiteCastle.Right = &RoomWhiteCastle

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

	// Room 0x1B — Black Castle Entry
	// All 4 dirs: V1→OtherPurpleRoom (0x1C) / V2+: stays as current (confirmed correct in V2)
	if gameType == 1 {
		// Up→CorridorRight: symmetric return path (CorridorRight.Down→BlackCastleEntry in V1)
		// Down/Left/Right→OtherPurpleRoom per C++ 0x89 group
		RoomBlackCastleEntry.Up = &RoomCorridorRight
		RoomBlackCastleEntry.Down = &RoomOtherPurpleRoom
		RoomBlackCastleEntry.Left = &RoomOtherPurpleRoom
		RoomBlackCastleEntry.Right = &RoomOtherPurpleRoom
	} else {
		RoomBlackCastleEntry.Up = &RoomSideCorridorCyan
		RoomBlackCastleEntry.Down = &RoomBlackMazeEntry
		RoomBlackCastleEntry.Left = &RoomBlackMazeEntry
		RoomBlackCastleEntry.Right = &RoomBlackMazeEntry
	}

	RoomOtherPurpleRoom.Up = &RoomBlackCastleTop
	// V1: Down→BlackCastleEntry (0x1B) / V2+: Down→SideCorridorCyan (0x0C)
	if gameType == 1 {
		RoomOtherPurpleRoom.Down = &RoomBlackCastleEntry
	} else {
		RoomOtherPurpleRoom.Down = &RoomSideCorridorCyan
	}
	RoomOtherPurpleRoom.Left = &RoomBlueMazeEntry
	RoomOtherPurpleRoom.Right = &RoomBlueMazeCenter

	RoomSideCorridorOlive.Up = &RoomWhiteCastle
	RoomSideCorridorOlive.Down = &RoomDeadEndCyan
	RoomSideCorridorOlive.Left = &RoomSideCorridorCyan
	RoomSideCorridorOlive.Right = &RoomMazeSide

	RoomDeadEndCyan.Up = &RoomSideCorridorOlive
	RoomDeadEndCyan.Down = &RoomWhiteCastle
	RoomDeadEndCyan.Left = &RoomBlackCastle
	RoomDeadEndCyan.Right = &RoomBlackCastle
}
