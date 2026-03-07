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

// Castle
var RoomGfxCastle = &[]string{
	"XXXXXXXXXXX X X X      X X X XXXXXXXXXXX",
	"X         X X X X      X X X X         X",
	"X         XXXXXXX      XXXXXXX         X",
	"X         XXXXXXXXxxxxXXXXXXXX         X",
	"X           XXXXXXxxxxXXXXXX           X",
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
