package world

// Player graphics
var PlayerGfx = []*Cell{
	{X: 0, Y: 0, Symbol: 'L'},
	{X: 1, Y: 0, Symbol: 'R'},
}

var PlayerGfxBefore = []*Cell{
	{X: 0, Y: 0, Symbol: 'B'},
	{X: 1, Y: 0, Symbol: '4'},
}

// Yellow Key (8 wide x 2 terminal rows = 3 pixel rows via half-block chars)
// From objectGfxKey[] = { 3, 0x07, 0xFD, 0xA7 } in Adventure.cpp
// ▄▄▄▄▄█▀█
// ▀ ▀  ▀▀▀
var KeyGfx = []*Cell{
	// Row 0: ▄▄▄▄▄█▀█
	{X: 0, Y: 0, Symbol: '▄'},
	{X: 1, Y: 0, Symbol: '▄'},
	{X: 2, Y: 0, Symbol: '▄'},
	{X: 3, Y: 0, Symbol: '▄'},
	{X: 4, Y: 0, Symbol: '▄'},
	{X: 5, Y: 0, Symbol: '█'},
	{X: 6, Y: 0, Symbol: '▀'},
	{X: 7, Y: 0, Symbol: '█'},
	// Row 1: ▀ ▀  ▀▀▀
	{X: 0, Y: 1, Symbol: '▀'},
	{X: 2, Y: 1, Symbol: '▀'},
	{X: 5, Y: 1, Symbol: '▀'},
	{X: 6, Y: 1, Symbol: '▀'},
	{X: 7, Y: 1, Symbol: '▀'},
}

// Dragon (8 wide x 10 tall terminal rows = 20 pixel rows via half-block chars)
// State 0 from objectGfxDrag[] in Adventure.cpp
//     ▄██▄
// ████▄▄█▀
//     ▀█▀
//    ▄▄█▄
//  ▄██████
// ██▀   ██
// ██   ▄██
// ▀▀████▀▀
// ▄   █▄▄▄
// ▀▀█▄▄▄▄█
var DragonGfx = []*Cell{
	// Row 0: ▄██▄  (cols 4-7)
	{X: 4, Y: 0, Symbol: '▄'},
	{X: 5, Y: 0, Symbol: '█'},
	{X: 6, Y: 0, Symbol: '█'},
	{X: 7, Y: 0, Symbol: '▄'},
	// Row 1: ████▄▄█▀
	{X: 0, Y: 1, Symbol: '█'},
	{X: 1, Y: 1, Symbol: '█'},
	{X: 2, Y: 1, Symbol: '█'},
	{X: 3, Y: 1, Symbol: '█'},
	{X: 4, Y: 1, Symbol: '▄'},
	{X: 5, Y: 1, Symbol: '▄'},
	{X: 6, Y: 1, Symbol: '█'},
	{X: 7, Y: 1, Symbol: '▀'},
	// Row 2:     ▀█▀  (cols 4-6)
	{X: 4, Y: 2, Symbol: '▀'},
	{X: 5, Y: 2, Symbol: '█'},
	{X: 6, Y: 2, Symbol: '▀'},
	// Row 3:    ▄▄█▄  (cols 3-6)
	{X: 3, Y: 3, Symbol: '▄'},
	{X: 4, Y: 3, Symbol: '▄'},
	{X: 5, Y: 3, Symbol: '█'},
	{X: 6, Y: 3, Symbol: '▄'},
	// Row 4:  ▄██████  (cols 1-7)
	{X: 1, Y: 4, Symbol: '▄'},
	{X: 2, Y: 4, Symbol: '█'},
	{X: 3, Y: 4, Symbol: '█'},
	{X: 4, Y: 4, Symbol: '█'},
	{X: 5, Y: 4, Symbol: '█'},
	{X: 6, Y: 4, Symbol: '█'},
	{X: 7, Y: 4, Symbol: '█'},
	// Row 5: ██▀   ██
	{X: 0, Y: 5, Symbol: '█'},
	{X: 1, Y: 5, Symbol: '█'},
	{X: 2, Y: 5, Symbol: '▀'},
	{X: 6, Y: 5, Symbol: '█'},
	{X: 7, Y: 5, Symbol: '█'},
	// Row 6: ██   ▄██
	{X: 0, Y: 6, Symbol: '█'},
	{X: 1, Y: 6, Symbol: '█'},
	{X: 5, Y: 6, Symbol: '▄'},
	{X: 6, Y: 6, Symbol: '█'},
	{X: 7, Y: 6, Symbol: '█'},
	// Row 7: ▀▀████▀▀
	{X: 0, Y: 7, Symbol: '▀'},
	{X: 1, Y: 7, Symbol: '▀'},
	{X: 2, Y: 7, Symbol: '█'},
	{X: 3, Y: 7, Symbol: '█'},
	{X: 4, Y: 7, Symbol: '█'},
	{X: 5, Y: 7, Symbol: '█'},
	{X: 6, Y: 7, Symbol: '▀'},
	{X: 7, Y: 7, Symbol: '▀'},
	// Row 8: ▄   █▄▄▄
	{X: 0, Y: 8, Symbol: '▄'},
	{X: 4, Y: 8, Symbol: '█'},
	{X: 5, Y: 8, Symbol: '▄'},
	{X: 6, Y: 8, Symbol: '▄'},
	{X: 7, Y: 8, Symbol: '▄'},
	// Row 9: ▀▀█▄▄▄▄█
	{X: 0, Y: 9, Symbol: '▀'},
	{X: 1, Y: 9, Symbol: '▀'},
	{X: 2, Y: 9, Symbol: '█'},
	{X: 3, Y: 9, Symbol: '▄'},
	{X: 4, Y: 9, Symbol: '▄'},
	{X: 5, Y: 9, Symbol: '▄'},
	{X: 6, Y: 9, Symbol: '▄'},
	{X: 7, Y: 9, Symbol: '█'},
}

// Bat State 03 — wings up (compact), decoded from objectGfxBat[] State 03
// 7 pixel rows → 4 terminal rows (3 pairs + 1 lone top-half row)
//
// █      █
// ██    ██
// ▀█▀██▀█▀
//  ▀▀  ▀▀
var BatGfx = []*Cell{
	// Row 0: 0x81+0x81 → █      █
	{X: 0, Y: 0, Symbol: '█'},
	{X: 7, Y: 0, Symbol: '█'},
	// Row 1: 0xC3+0xC3 → ██    ██
	{X: 0, Y: 1, Symbol: '█'},
	{X: 1, Y: 1, Symbol: '█'},
	{X: 6, Y: 1, Symbol: '█'},
	{X: 7, Y: 1, Symbol: '█'},
	// Row 2: 0xFF+0x5A → ▀█▀██▀█▀
	{X: 0, Y: 2, Symbol: '▀'},
	{X: 1, Y: 2, Symbol: '█'},
	{X: 2, Y: 2, Symbol: '▀'},
	{X: 3, Y: 2, Symbol: '█'},
	{X: 4, Y: 2, Symbol: '█'},
	{X: 5, Y: 2, Symbol: '▀'},
	{X: 6, Y: 2, Symbol: '█'},
	{X: 7, Y: 2, Symbol: '▀'},
	// Row 3: 0x66 lone → ▀▀  ▀▀ (top-half only)
	{X: 1, Y: 3, Symbol: '▀'},
	{X: 2, Y: 3, Symbol: '▀'},
	{X: 5, Y: 3, Symbol: '▀'},
	{X: 6, Y: 3, Symbol: '▀'},
}

// Bat State FF — wings spread/down, decoded from objectGfxBat[] State FF
// 11 pixel rows → 6 terminal rows (5 pairs + 1 lone top-half row)
//
// ▄      ▀
// ▄      ▀
//  ▄▀██▀▄
// ▄█▀  ▀█▄
// █      █
// ▀      ▀
var BatGfxOpen = []*Cell{
	// Row 0: 0x01+0x80 → ▄      ▀
	{X: 0, Y: 0, Symbol: '▄'},
	{X: 7, Y: 0, Symbol: '▀'},
	// Row 1: 0x01+0x80 → ▄      ▀
	{X: 0, Y: 1, Symbol: '▄'},
	{X: 7, Y: 1, Symbol: '▀'},
	// Row 2: 0x3C+0x5A →  ▄▀██▀▄
	{X: 1, Y: 2, Symbol: '▄'},
	{X: 2, Y: 2, Symbol: '▀'},
	{X: 3, Y: 2, Symbol: '█'},
	{X: 4, Y: 2, Symbol: '█'},
	{X: 5, Y: 2, Symbol: '▀'},
	{X: 6, Y: 2, Symbol: '▄'},
	// Row 3: 0x66+0xC3 → ▄█▀  ▀█▄
	{X: 0, Y: 3, Symbol: '▄'},
	{X: 1, Y: 3, Symbol: '█'},
	{X: 2, Y: 3, Symbol: '▀'},
	{X: 5, Y: 3, Symbol: '▀'},
	{X: 6, Y: 3, Symbol: '█'},
	{X: 7, Y: 3, Symbol: '▄'},
	// Row 4: 0x81+0x81 → █      █
	{X: 0, Y: 4, Symbol: '█'},
	{X: 7, Y: 4, Symbol: '█'},
	// Row 5: 0x81 lone → ▀      ▀ (top-half only)
	{X: 0, Y: 5, Symbol: '▀'},
	{X: 7, Y: 5, Symbol: '▀'},
}

// Dragon State 01 — mouth open, decoded from objectGfxDrag[] State 01 in Adventure.cpp
// 22 pixel rows → 11 terminal rows via half-block pairs (bit7=leftmost pixel)
//
// ▀▄
//   ▀▄▄██▄
//     █▄█▀
//   ▄▀▀█▀
// ▄▀  ▄█▄
//   ▄████▄
//  ███████
//  ███████
//   ▀███▀
// ▄▄▄▄█
// █▄▄
var DragonGfxOpen = []*Cell{
	// All Y values shifted -1 so total height matches DragonGfx (10 rows).
	// Row -1 (clipped): pixels 0x80+0x40 → ▀▄  (out of bounds, not drawn)
	{X: 0, Y: -1, Symbol: '▀'},
	{X: 1, Y: -1, Symbol: '▄'},
	// Row 0: pixels 0x26+0x1F → ▀▄▄██▄
	{X: 2, Y: 0, Symbol: '▀'},
	{X: 3, Y: 0, Symbol: '▄'},
	{X: 4, Y: 0, Symbol: '▄'},
	{X: 5, Y: 0, Symbol: '█'},
	{X: 6, Y: 0, Symbol: '█'},
	{X: 7, Y: 0, Symbol: '▄'},
	// Row 1: pixels 0x0B+0x0E → █▄█▀
	{X: 4, Y: 1, Symbol: '█'},
	{X: 5, Y: 1, Symbol: '▄'},
	{X: 6, Y: 1, Symbol: '█'},
	{X: 7, Y: 1, Symbol: '▀'},
	// Row 2: pixels 0x1E+0x24 → ▄▀▀█▀
	{X: 2, Y: 2, Symbol: '▄'},
	{X: 3, Y: 2, Symbol: '▀'},
	{X: 4, Y: 2, Symbol: '▀'},
	{X: 5, Y: 2, Symbol: '█'},
	{X: 6, Y: 2, Symbol: '▀'},
	// Row 3: pixels 0x44+0x8E → ▄▀  ▄█▄
	{X: 0, Y: 3, Symbol: '▄'},
	{X: 1, Y: 3, Symbol: '▀'},
	{X: 4, Y: 3, Symbol: '▄'},
	{X: 5, Y: 3, Symbol: '█'},
	{X: 6, Y: 3, Symbol: '▄'},
	// Rows 4–9: body from DragonGfx (rows 4–9), same Y
	// Row 4:  ▄██████
	{X: 1, Y: 4, Symbol: '▄'},
	{X: 2, Y: 4, Symbol: '█'},
	{X: 3, Y: 4, Symbol: '█'},
	{X: 4, Y: 4, Symbol: '█'},
	{X: 5, Y: 4, Symbol: '█'},
	{X: 6, Y: 4, Symbol: '█'},
	{X: 7, Y: 4, Symbol: '█'},
	// Row 5:  ██▀   ██
	{X: 0, Y: 5, Symbol: '█'},
	{X: 1, Y: 5, Symbol: '█'},
	{X: 2, Y: 5, Symbol: '▀'},
	{X: 6, Y: 5, Symbol: '█'},
	{X: 7, Y: 5, Symbol: '█'},
	// Row 6:  ██   ▄██
	{X: 0, Y: 6, Symbol: '█'},
	{X: 1, Y: 6, Symbol: '█'},
	{X: 5, Y: 6, Symbol: '▄'},
	{X: 6, Y: 6, Symbol: '█'},
	{X: 7, Y: 6, Symbol: '█'},
	// Row 7:  ▀▀████▀▀
	{X: 0, Y: 7, Symbol: '▀'},
	{X: 1, Y: 7, Symbol: '▀'},
	{X: 2, Y: 7, Symbol: '█'},
	{X: 3, Y: 7, Symbol: '█'},
	{X: 4, Y: 7, Symbol: '█'},
	{X: 5, Y: 7, Symbol: '█'},
	{X: 6, Y: 7, Symbol: '▀'},
	{X: 7, Y: 7, Symbol: '▀'},
	// Row 8:  ▄   █▄▄▄
	{X: 0, Y: 8, Symbol: '▄'},
	{X: 4, Y: 8, Symbol: '█'},
	{X: 5, Y: 8, Symbol: '▄'},
	{X: 6, Y: 8, Symbol: '▄'},
	{X: 7, Y: 8, Symbol: '▄'},
	// Row 9:  ▀▀█▄▄▄▄█
	{X: 0, Y: 9, Symbol: '▀'},
	{X: 1, Y: 9, Symbol: '▀'},
	{X: 2, Y: 9, Symbol: '█'},
	{X: 3, Y: 9, Symbol: '▄'},
	{X: 4, Y: 9, Symbol: '▄'},
	{X: 5, Y: 9, Symbol: '▄'},
	{X: 6, Y: 9, Symbol: '▄'},
	{X: 7, Y: 9, Symbol: '█'},
}

// Bridge — objectGfxBridge, 24 pixel rows → 12 terminal rows
// ██    ██  (top caps ×2)
//  █    █   (pillars ×8)
// ██    ██  (bottom caps ×2)
var BridgeGfx = []*Cell{
	// Rows 0-1: 0xC3 pairs → ██    ██
	{X: 0, Y: 0, Symbol: '█'}, {X: 1, Y: 0, Symbol: '█'}, {X: 6, Y: 0, Symbol: '█'}, {X: 7, Y: 0, Symbol: '█'},
	{X: 0, Y: 1, Symbol: '█'}, {X: 1, Y: 1, Symbol: '█'}, {X: 6, Y: 1, Symbol: '█'}, {X: 7, Y: 1, Symbol: '█'},
	// Rows 2-9: 0x42 pairs →  █    █
	{X: 1, Y: 2, Symbol: '█'}, {X: 6, Y: 2, Symbol: '█'},
	{X: 1, Y: 3, Symbol: '█'}, {X: 6, Y: 3, Symbol: '█'},
	{X: 1, Y: 4, Symbol: '█'}, {X: 6, Y: 4, Symbol: '█'},
	{X: 1, Y: 5, Symbol: '█'}, {X: 6, Y: 5, Symbol: '█'},
	{X: 1, Y: 6, Symbol: '█'}, {X: 6, Y: 6, Symbol: '█'},
	{X: 1, Y: 7, Symbol: '█'}, {X: 6, Y: 7, Symbol: '█'},
	{X: 1, Y: 8, Symbol: '█'}, {X: 6, Y: 8, Symbol: '█'},
	{X: 1, Y: 9, Symbol: '█'}, {X: 6, Y: 9, Symbol: '█'},
	// Rows 10-11: 0xC3 pairs → ██    ██
	{X: 0, Y: 10, Symbol: '█'}, {X: 1, Y: 10, Symbol: '█'}, {X: 6, Y: 10, Symbol: '█'}, {X: 7, Y: 10, Symbol: '█'},
	{X: 0, Y: 11, Symbol: '█'}, {X: 1, Y: 11, Symbol: '█'}, {X: 6, Y: 11, Symbol: '█'}, {X: 7, Y: 11, Symbol: '█'},
}

// Sword — objectGfxSword, 5 pixel rows → 3 terminal rows
//  ▄▀
// ▀█▀▀▀▀▀▀
//   ▀
var SwordGfx = []*Cell{
	// Row 0: 0x20+0x40 → col1=▄ col2=▀
	{X: 1, Y: 0, Symbol: '▄'}, {X: 2, Y: 0, Symbol: '▀'},
	// Row 1: 0xFF+0x40 → ▀█▀▀▀▀▀▀
	{X: 0, Y: 1, Symbol: '▀'}, {X: 1, Y: 1, Symbol: '█'},
	{X: 2, Y: 1, Symbol: '▀'}, {X: 3, Y: 1, Symbol: '▀'},
	{X: 4, Y: 1, Symbol: '▀'}, {X: 5, Y: 1, Symbol: '▀'},
	{X: 6, Y: 1, Symbol: '▀'}, {X: 7, Y: 1, Symbol: '▀'},
	// Row 2: 0x20 lone → col2=▀
	{X: 2, Y: 2, Symbol: '▀'},
}

// Chalice — objectGfxChallise, 9 pixel rows → 5 terminal rows
// █      █
// ▀█▄▄▄▄█▀
//  ▀████▀
//    ██
//  ▀▀▀▀▀▀
var ChaliceGfx = []*Cell{
	// Row 0: 0x81+0x81 → █      █
	{X: 0, Y: 0, Symbol: '█'}, {X: 7, Y: 0, Symbol: '█'},
	// Row 1: 0xC3+0x7E → ▀█▄▄▄▄█▀
	{X: 0, Y: 1, Symbol: '▀'}, {X: 1, Y: 1, Symbol: '█'},
	{X: 2, Y: 1, Symbol: '▄'}, {X: 3, Y: 1, Symbol: '▄'},
	{X: 4, Y: 1, Symbol: '▄'}, {X: 5, Y: 1, Symbol: '▄'},
	{X: 6, Y: 1, Symbol: '█'}, {X: 7, Y: 1, Symbol: '▀'},
	// Row 2: 0x7E+0x3C →  ▀████▀
	{X: 1, Y: 2, Symbol: '▀'}, {X: 2, Y: 2, Symbol: '█'},
	{X: 3, Y: 2, Symbol: '█'}, {X: 4, Y: 2, Symbol: '█'},
	{X: 5, Y: 2, Symbol: '█'}, {X: 6, Y: 2, Symbol: '▀'},
	// Row 3: 0x18+0x18 →    ██
	{X: 3, Y: 3, Symbol: '█'}, {X: 4, Y: 3, Symbol: '█'},
	// Row 4: 0x7E lone →  ▀▀▀▀▀▀ (top-half only)
	{X: 1, Y: 4, Symbol: '▀'}, {X: 2, Y: 4, Symbol: '▀'},
	{X: 3, Y: 4, Symbol: '▀'}, {X: 4, Y: 4, Symbol: '▀'},
	{X: 5, Y: 4, Symbol: '▀'}, {X: 6, Y: 4, Symbol: '▀'},
}

// Magnet — objectGfxMagnet, 8 pixel rows → 4 terminal rows
//  ▄████▄
// ██▀  ▀██
// ██    ██
// ██    ██
var MagnetGfx = []*Cell{
	// Row 0: 0x3C+0x7E →  ▄████▄
	{X: 1, Y: 0, Symbol: '▄'}, {X: 2, Y: 0, Symbol: '█'},
	{X: 3, Y: 0, Symbol: '█'}, {X: 4, Y: 0, Symbol: '█'},
	{X: 5, Y: 0, Symbol: '█'}, {X: 6, Y: 0, Symbol: '▄'},
	// Row 1: 0xE7+0xC3 → ██▀  ▀██
	{X: 0, Y: 1, Symbol: '█'}, {X: 1, Y: 1, Symbol: '█'},
	{X: 2, Y: 1, Symbol: '▀'}, {X: 5, Y: 1, Symbol: '▀'},
	{X: 6, Y: 1, Symbol: '█'}, {X: 7, Y: 1, Symbol: '█'},
	// Rows 2-3: 0xC3+0xC3 → ██    ██
	{X: 0, Y: 2, Symbol: '█'}, {X: 1, Y: 2, Symbol: '█'}, {X: 6, Y: 2, Symbol: '█'}, {X: 7, Y: 2, Symbol: '█'},
	{X: 0, Y: 3, Symbol: '█'}, {X: 1, Y: 3, Symbol: '█'}, {X: 6, Y: 3, Symbol: '█'}, {X: 7, Y: 3, Symbol: '█'},
}

// Dot — objectGfxDot, 1 pixel row → 1 terminal row (top-half only)
var DotGfx = []*Cell{
	{X: 0, Y: 0, Symbol: '▀'},
}

// Castle
var RoomGfxCastle = &[]string{
	"XXXXXXXXXXX X X X      X X X XXXXXXXXXXX",
	"X         X X X X      X X X X         X",
	"X         XXXXXXX      XXXXXXX         X",
	"X         XXXXXXXXXXXXXXXXXXXX         X",
	"X           XXXXXXXXXXXXXXXX           X",
	"X           XXXXXX    XXXXXX           X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Left of Name room: solid top wall, open sides, opening at bottom
var RoomCorridorRightGfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Room below yellow castle: opening at top, open sides, solid bottom
var RoomBelowYellowCastleGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Room above yellow castle
var RoomAboveYellowCastleGfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Blue Maze Entry: complex maze walls, opening at bottom center
var RoomBlueMazeEntryGfx = &[]string{
	"XXXXXXXX  XX  XX  XXXX  XX  XX  XXXXXXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXX  XX  XX  XXXXXXXXXXXX  XX  XX  XXXX",
	"XXXX  XX  XX  XXXXXXXXXXXX  XX  XX  XXXX",
	"      XX  XX                XX  XX      ",
	"      XX  XX                XX  XX      ",
	"XXXXXXXX  XXXXXXXXXXXXXXXXXXXX  XXXXXXXX",
	"XXXXXXXX  XXXXXXXXXXXXXXXXXXXX  XXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Blue Maze #1 (Room 5 in C++)
var RoomBlueMaze1Gfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXX  XXXXXXXXXXXXXXXX  XXXXXXXXXX",
	"XXXXXXXXXX  XXXXXXXXXXXXXXXX  XXXXXXXXXX",
	"XXXX              XXXX              XXXX",
	"XXXX              XXXX              XXXX",
	"XXXX  XXXXXXXXXX  XXXX  XXXXXXXXXX  XXXX",
	"XXXX  XXXXXXXXXX  XXXX  XXXXXXXXXX  XXXX",
	"      XX      XX  XXXX  XX      XX      ",
	"      XX      XX  XXXX  XX      XX      ",
	"XXXXXXXX  XX  XX  XXXX  XX  XX  XXXXXXXX",
}

// Red Maze #1 (Room 0x17 in C++)
var RoomRedMaze1Gfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"              XX        XX              ",
	"              XX        XX              ",
	"XXXXXXXXXXXX  XX        XX  XXXXXXXXXXXX",
	"XXXXXXXXXXXX  XX        XX  XXXXXXXXXXXX",
	"XXXX      XX  XX  XXXX  XX  XX      XXXX",
	"XXXX      XX  XX  XXXX  XX  XX      XXXX",
	"XXXX  XX  XXXXXX  XXXX  XXXXXX  XX  XXXX",
}

// Bottom of Red Maze (Room 0x19 in C++)
var RoomRedMazeBottomGfx = &[]string{
	"XXXX  XX  XXXXXX  XXXX  XXXXXX  XX  XXXX",
	"XXXX  XX                        XX  XXXX",
	"XXXX  XX                        XX  XXXX",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"      XX  XX                XX  XX      ",
	"      XX  XX                XX  XX      ",
	"XXXXXXXXXXXX                XXXXXXXXXXXX",
	"XXXXXXXXXXXX                XXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Top of Red Maze (Room 0x18 in C++)
var RoomRedMazeTopGfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                  XXXX                  ",
	"                  XXXX                  ",
	"XXXXXXXXXXXXXXXX  XXXX  XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX  XXXX  XXXXXXXXXXXXXXXX",
	"              XX  XXXX  XX              ",
	"              XX  XXXX  XX              ",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"XXXX  XX  XX                XX  XX  XXXX",
	"XXXX  XX  XX                XX  XX  XXXX",
	"XXXX  XXXXXX  XX        XX  XXXXXX  XXXX",
}

// White Castle Entry — red (Room 0x1A in C++)
var RoomWhiteCastleEntryGfx = &[]string{
	"XXXX  XXXXXX  XX        XX  XXXXXX  XXXX",
	"XXXX          XX        XX          XXXX",
	"XXXX          XX        XX          XXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXX  XX                        XX  XXXX",
	"XXXX  XX                        XX  XXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Top Entry Room — solid bottom (Room 0x0E cyan in C++)
var RoomTopEntryRoomGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Top Entry Room — opening at bottom (Room 0x1D red in C++, above Black Castle)
var RoomBlackCastleTopGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Blue Maze Bottom (Room 6 in C++)
var RoomBlueMazeBottomGfx = &[]string{
	"XXXXXXXX  XX  XX        XX  XX  XXXXXXXX",
	"      XX      XX        XX      XX      ",
	"      XX      XX        XX      XX      ",
	"XXXX  XXXXXXXXXX        XXXXXXXXXX  XXXX",
	"XXXX  XXXXXXXXXX        XXXXXXXXXX  XXXX",
	"XXXX                                XXXX",
	"XXXX                                XXXX",
	"XXXXXXXX                        XXXXXXXX",
	"XXXXXXXX                        XXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Blue Maze Top (Room 4 in C++)
var RoomBlueMazeTopGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"        XX    XX        XX    XX        ",
	"        XX    XX        XX    XX        ",
	"XXXX    XX    XXXX    XXXX    XX    XXXX",
	"XXXX    XX    XXXX    XXXX    XX    XXXX",
	"XXXX    XX                    XX    XXXX",
	"XXXX    XX                    XX    XXXX",
	"XXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXX",
	"      XX        XX    XX        XX      ",
	"      XX        XX    XX        XX      ",
	"XXXX  XX  XXXXXXXX    XXXXXXXX  XX  XXXX",
}

// Blue Maze Center (Room 7 in C++)
var RoomBlueMazeCenterGfx = &[]string{
	"XXXX  XX  XXXXXXXX    XXXXXXXX  XX  XXXX",
	"      XX      XXXX    XXXX      XX      ",
	"      XX      XXXX    XXXX      XX      ",
	"XXXXXXXXXXXX  XXXX    XXXX  XXXXXXXXXXXX",
	"XXXXXXXXXXXX  XXXX    XXXX  XXXXXXXXXXXX",
	"          XX  XXXX    XXXX  XX          ",
	"          XX  XXXX    XXXX  XX          ",
	"XXXX  XX  XX  XXXX    XXXX  XX  XX  XXXX",
	"XXXX  XX  XX  XXXX    XXXX  XX  XX  XXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXXXXXX  XX  XX        XX  XX  XXXXXXXX",
}

// Black Maze #1 (Room 0x13 in C++, ROOMFLAG_NONE: right = reverse(left))
var RoomBlackMaze1Gfx = &[]string{
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"            XX            XX            ",
	"            XX            XX            ",
	"XXXXXXXXXXXXXX            XXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXX            XXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX",
	"XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
}

// Black Maze #2 (Room 0x14 in C++, ROOMFLAG_MIRROR: right = same as left)
var RoomBlackMaze2Gfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                  XX                  XX",
	"                  XX                  XX",
	"XXXXXXXXXXXXXXXX  XXXXXXXXXXXXXXXXXX  XX",
	"XXXXXXXXXXXXXXXX  XXXXXXXXXXXXXXXXXX  XX",
	"              XX                  XX    ",
	"              XX                  XX    ",
	"XXXX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXX",
	"XXXX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXX",
	"        XXXX      XX        XXXX      XX",
	"        XXXX      XX        XXXX      XX",
	"XX  XX  XXXX  XX  XXXX  XX  XXXX  XX  XX",
}

// Black Maze #3 (Room 0x15 in C++, ROOMFLAG_MIRROR: right = same as left)
var RoomBlackMaze3Gfx = &[]string{
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"XX                  XX                  ",
	"XX                  XX                  ",
	"XX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXX",
	"XX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXX",
	"      XX                  XX            ",
	"      XX                  XX            ",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"XX          XX      XX          XX      ",
	"XX          XX      XX          XX      ",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
}

// Black Maze Entry (Room 0x16 in C++, ROOMFLAG_NONE: right = reverse(left))
var RoomBlackMazeEntryGfx = &[]string{
	"XX  XX  XXXX  XX  XXXX  XX  XXXX  XX  XX",
	"    XX        XX  XXXX  XX        XX    ",
	"    XX        XX  XXXX  XX        XX    ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Easter Egg Room (Room 0x00): opening at top, no side walls, solid bottom
var RoomEasterEggGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Number Room (Rooms 0x12, 0x1C in C++)
var RoomGfxNumberRoom = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Side Corridor (Rooms 0x0C, 0x0D in C++)
var RoomGfxSideCorridor = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Maze Entry (Room 0x0A in C++)
var RoomGfxMazeEntry = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXX  XX    XXXXXXXXXXXXXXXX    XX  XXXX",
	"XXXX  XX    XXXXXXXXXXXXXXXX    XX  XXXX",
	"      XX          XXXX          XX      ",
	"      XX          XXXX          XX      ",
	"XXXXXXXX  XX      XXXX      XX  XXXXXXXX",
	"XXXXXXXX  XX      XXXX      XX  XXXXXXXX",
	"          XX      XXXX      XX          ",
	"          XX      XXXX      XX          ",
	"XXXXXXXXXXXX  XX  XXXX  XX  XXXXXXXXXXXX",
}

// Maze Middle (Room 0x09 in C++)
var RoomGfxMazeMiddle = &[]string{
	"XXXXXXXXXXXX  XX  XXXX  XX  XXXXXXXXXXXX",
	"              XX  XXXX  XX              ",
	"              XX  XXXX  XX              ",
	"XXXX      XXXXXX  XXXX  XXXXXX      XXXX",
	"XXXX      XXXXXX  XXXX  XXXXXX      XXXX",
	"          XX                XX          ",
	"          XX                XX          ",
	"XXXXXXXX  XX  XXXXXXXXXXXX  XX  XXXXXXXX",
	"XXXXXXXX  XX  XXXXXXXXXXXX  XX  XXXXXXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXX  XX  XX  XX  XXXX  XX  XX  XX  XXXX",
}

// Maze Side (Room 0x0B in C++)
var RoomGfxMazeSide = &[]string{
	"XXXX  XX  XX  XX  XXXX  XX  XX  XX  XXXX",
	"      XX      XX  XXXX  XX      XX      ",
	"      XX      XX  XXXX  XX      XX      ",
	"      XXXXXXXXXX  XXXX  XXXXXXXXXX      ",
	"      XXXXXXXXXX  XXXX  XXXXXXXXXX      ",
	"                  XXXX                  ",
	"                  XXXX                  ",
	"      XXXXXXXX    XXXX    XXXXXXXX      ",
	"      XXXXXXXX    XXXX    XXXXXXXX      ",
	"      XX          XXXX          XX      ",
	"      XX          XXXX          XX      ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// ConvertToBinary converts a row of room graphics to a bitmask for collision detection.
func ConvertToBinary(data string) int64 {
	binary := int64(0)
	for _, char := range data {
		binary <<= 1
		if char == 'X' {
			binary |= 1
		}
	}
	return binary
}
