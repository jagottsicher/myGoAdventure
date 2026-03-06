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
var RoomLeftOfNameGfx = &[]string{
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
	"XXXX  XX  XX  XXXXXXXXXXXX  XX  XX  XXXX",
	"      XX  XX                XX  XX      ",
	"XXXXXXXX  XXXXXXXXXXXXXXXXXXXX  XXXXXXXX",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Blue Maze #1 (Room 5 in C++)
var RoomBlueMaze1Gfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"XXXXXXXXXX  XXXXXXXXXXXXXXXX  XXXXXXXXXX",
	"XXXX              XXXX              XXXX",
	"XXXX  XXXXXXXXXX  XXXX  XXXXXXXXXX  XXXX",
	"      XX      XX  XXXX  XX      XX      ",
	"XXXXXXXX  XX  XX  XXXX  XX  XX  XXXXXXXX",
}

// Blue Maze Center (Room 7 in C++)
var RoomBlueMazeCenterGfx = &[]string{
	"XXXX  XX  XXXXXXXX    XXXXXXXX  XX  XXXX",
	"      XX      XXXX    XXXX      XX      ",
	"XXXXXXXXXXXX  XXXX    XXXX  XXXXXXXXXXXX",
	"          XX  XXXX    XXXX  XX          ",
	"XXXX  XX  XX  XXXX    XXXX  XX  XX  XXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXXXXXX  XX  XX        XX  XX  XXXXXXXX",
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
