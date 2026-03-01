package main

import (
	"github.com/gdamore/tcell/v2"
)

// Scaling the room
var defaultXFactor, defaultYFactor int

type room struct {
	compressedRoomData   *[]string
	uncompressedRoomData *[]string
	background           tcell.Color
	foreground           tcell.Color
	barLeft              bool
	barRight             bool
	up                   *room
	down                 *room
	left                 *room
	right                *room
}

// func (r *room) assignUncompressedRoomData() {
// 	// Assign uncompressed room data based on compressed room data
// 	// Implement your logic here

// 	// screensize feststellen
// 	screenWidth, screenHeight := screen.Size()

// 	r.uncompressedRoomData = []*cell{
// 		{0, 0, 'X'},
// 		{screenWidth - 1, screenHeight - 1, 'Y'},
// 	}
// }

// var Rooms = []room{
// 	roomYellowCastle,
// }

// func uncompressRooms() {

// 	for _, room := range Rooms {
// 		room.assignUncompressedRoomData()
// 	}
// }

var roomGfxYellowCastleUncompressed = &[]string{
	"XXXXXXXXXXXXXXXXXXXX X X X      X X X XXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXX X X X      X X X XXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXX X X X      X X X XXXXXXXXXXX",
	"X                  X X X X      X X X X         X",
	"X                  X X X X      X X X X         X",
	"X                  X X X X      X X X X         X",
	"X                  XXXXXXX      XXXXXXX         X",
	"X                  XXXXXXXXxxxxXXXXXXXX         X",
	"X                  XXXXXXXXxxxxXXXXXXXX         X",
	"X                  XXXXXXXXxxxxXXXXXXXX         X",
	"X                  XXXXXXXXxxxxXXXXXXXX         X",
	"X                  XXXXXXXXxxxxXXXXXXXX         X",
	"X                  XXXXXXXXxxxxXXXXXXXX         X",
	"X                  XXXXXXXXxxxxXXXXXXXX         X",
	"X                    XXXXXXxxxxXXXXXX           X",
	"X                    XXXXXX    XXXXXX           X",
	"X                    XXXXXX    XXXXXX           X",
	"X                    XXXXXX    XXXXXX           X",
	"X                                               X",
	"X                                               X",
	"X                                               X",
	"X                                               X",
	"X                                               X",
	"XXXXXXXXXXXXXXXX                 XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX                 XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX                 XXXXXXXXXXXXXXXX",
}

var roomSplashScreen = room{
	compressedRoomData:   roomGfxCastle,
	uncompressedRoomData: roomGfxYellowCastleUncompressed,
	background:           tcell.ColorDarkGray,
	foreground:           tcell.ColorAntiqueWhite,
	up:                   nil,
	down:                 nil,
	left:                 nil,
	right:                nil,
}

var roomYellowCastle = room{
	compressedRoomData:   roomGfxCastle,
	uncompressedRoomData: roomGfxYellowCastleUncompressed,
	background:           tcell.ColorDarkGray,
	foreground:           tcell.ColorYellow,
	up:                   nil,
	down:                 nil,
	left:                 nil,
	right:                nil,
}

// this function is needed to avoid errors in assignment order
func initDirections() {
	// roomYellowCastle.up = &roomOnTopOfYellowCastle
	// roomYellowCastle.down = &roomStartRoomTopEntryRoom
	roomYellowCastle.up = nil
	roomYellowCastle.down = nil
	roomYellowCastle.left = nil
	roomYellowCastle.right = nil

	// 	// screensize feststellen
	// _, screenHeight := screen.Size()

	// roomGfxYellowCastleUncompressed

	// for height := 0; height < screenHeight; height++ {
	// 	// 	// (*roomYellowCastle.uncompressedRoomData)[height] = (*roomYellowCastle.compressedRoomData)[nextInteger(float64(height/12))]
	// 	emitStr(screen, 1, height, tcell.StyleDefault.
	// 		Background(roomYellowCastle.background).
	// 		Foreground(roomYellowCastle.foreground), "X")
	// }

	// roomGfxYellowCastleUncompressed = roomGfxCastle
}

func nextInteger(f float64) int {
	// Add 0.5 and then truncate to get the next integer
	return int(f + 0.5)
}
