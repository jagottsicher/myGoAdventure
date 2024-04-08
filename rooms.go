package main

import (
	"github.com/gdamore/tcell/v2"
)

// Scaling the room
var defaultXFactor, defaultYFactor int

type room struct {
	compressedRoomData   *[]string
	uncompressedRoomData []*cell
	background           tcell.Color
	foreground           tcell.Color
	barLeft              bool
	barRight             bool
	up                   *room
	down                 *room
	left                 *room
	right                *room
}

func (r room) UncompressedRoomDat() []*cell {
	return r.uncompressedRoomData
}

var rooms = []room{
	roomYellowCastle,
	// roomBlackCastle,
	// roomWhiteCastle,
	// roomStartRoomTopEntryRoom,
	// roomPurpleEasterEggTopEntryRoom,
	// roomNumberRoom,
	// roomOnTopOfYellowCastle,
	// roomLeftOfStartRoom,
	// roomTwoBelowWhiteCastleRoomTopEntryRoom,
	// roomLowerRedRoomTopEntryRoom,
	// roomLeftOfNameRoom,
	// roomOnTopOfBlackCastle,
	// roomBelowWhiteCastle,
	// roomBelowNumberRoom,
	// roomBlueMazeEntry,
	// roomBlueMazeCenter,
	// roomBlueMazeBottom,
	// roomBlueMazeLeft,
	// roomBlueMazeTop,
	// roomMazeMiddle,
	// roomMazeSide,
	// roomMazeEntry,
	// roomRedMazeTopLeft,
	// roomRedMazeBottomLeft,
	// roomRedMazeTopRight,
	// roomRedMazeEntryBottomRight,
	// roomBlackMazeTopLeft,
	// roomBlackMazeBottomLeft,
	// roomBlackMazeTopRight,
	// roomBlackMazeEntryBottomRight,
}

func uncompressRooms() {

	// screensize feststellen
	screenWidth, screenHeight := screen.Size()
	// 	var uncompressedCell cell

	// 	defaultXFactor = int(math.Round(float64(screenWidth) / 40))
	// 	defaultYFactor = 2

	// 	var percentageY float64
	// 	var rowValue int
	// 	MaxRowValue := 12
	// 	var percentageX float64
	// 	var columnValue int
	// 	MaxColumnValue := 40
	// 	var theRow string
	// 	var theSpot rune

	for _, room := range rooms {
		room.uncompressedRoomData = []*cell{
			{1, 1, 'X'},
			{screenWidth, screenHeight, 'Y'},
		}
	}

	// for _, room := range rooms {
	// 		room.uncompressedRoomData = nil

	// 		// fill all rooms with uncompressed screendata
	// 		for y := 0; y < screenHeight; y++ {
	// 			percentageY = float64((y * 100) / (screenHeight - 1))
	// 			rowValue = int(12 * int(percentageY) / 100)
	// 			if rowValue == MaxRowValue {
	// 				rowValue = MaxRowValue - 1
	// 			}

	// 			theRow = room.compressedRoomData[rowValue].(*strings)

	// 			for x := 0; x < screenWidth; x++ {
	// 				percentageX = float64((x * 100) / (screenWidth - 1))
	// 				columnValue = int(40 * int(percentageX) / 100)
	// 				if columnValue == MaxColumnValue {
	// 					columnValue = MaxColumnValue - 1
	// 				}

	// 				theSpot = rune(theRow[columnValue])

	//				uncompressedCell.x = x
	//				uncompressedCell.y = y
	//				uncompressedCell.symbol = theSpot
	//				room.uncompressedRoomData = append(room.uncompressedRoomData, &uncompressedCell)
	//			}
	//			var tempdata = &cell{x: uncompressedCell.x, y: uncompressedCell.y, symbol: 'X'}
	//			room.uncompressedRoomData = append(room.uncompressedRoomData, tempdata)
	//		}
	// }
}

var roomYellowCastle = room{
	compressedRoomData:   roomGfxCastle,
	uncompressedRoomData: nil,
	background:           tcell.ColorDarkGray,
	foreground:           tcell.ColorYellow,
	up:                   nil,
	down:                 nil,
	left:                 nil,
	right:                nil,
}

// var roomBlackCastle = room{
// 	compressedRoomData:   roomGfxCastle,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorBlack),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomWhiteCastle = room{
// 	compressedRoomData:   roomGfxCastle,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorWhite),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomStartRoomTopEntryRoom = room{
// 	compressedRoomData:   roomGfxBelowYellowCastle,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorGreen),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomPurpleEasterEggTopEntryRoom = room{
// 	compressedRoomData:   roomGfxTopEntryRoom,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorPurple),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomNumberRoom = room{
// 	compressedRoomData:   roomGfxNumberRoom,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorPurple),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomOnTopOfYellowCastle = room{
// 	compressedRoomData:   roomGfxNumberRoom,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorYellow),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomLeftOfStartRoom = room{
// 	compressedRoomData:   roomGfxBelowYellowCastle,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorDarkGreen),
// 	barLeft: true,
// 	up:      nil,
// 	down:    nil,
// 	left:    nil,
// 	right:   nil,
// }

// var roomTwoBelowWhiteCastleRoomTopEntryRoom = room{
// 	compressedRoomData:   roomGfxTopEntryRoom,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorDarkGreen),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomLowerRedRoomTopEntryRoom = room{
// 	compressedRoomData:   roomGfxTopEntryRoom,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorRed),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomLeftOfNameRoom = room{
// 	compressedRoomData:   roomGfxLeftOfName,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorGreenYellow),
// 	barRight: true,
// 	up:       nil,
// 	down:     nil,
// 	left:     nil,
// 	right:    nil,
// }

// var roomOnTopOfBlackCastle = room{
// 	compressedRoomData:   roomGfxSideCorridor,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorRed),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomBelowWhiteCastle = room{
// 	compressedRoomData:   roomGfxSideCorridor,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorGreen),
// 	barLeft: true,
// 	up:      nil,
// 	down:    nil,
// 	left:    nil,
// 	right:   nil,
// }

// var roomBelowNumberRoom = room{
// 	compressedRoomData:   roomGfxSideCorridor,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorLightBlue),
// 	barRight: true,
// 	up:       nil,
// 	down:     nil,
// 	left:     nil,
// 	right:    nil,
// }

// var roomBlueMazeEntry = room{
// 	compressedRoomData:   roomGfxBlueMazeEntry,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorBlue),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomBlueMazeCenter = room{
// 	compressedRoomData:   roomGfxBlueMazeCenter,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorBlue),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomBlueMazeBottom = room{
// 	compressedRoomData:   roomGfxBlueMazeBottom,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorBlue),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomBlueMazeLeft = room{
// 	compressedRoomData:   roomGfxBlueMaze1,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorBlue),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomBlueMazeTop = room{
// 	compressedRoomData:   roomGfxBlueMazeTop,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorBlue),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomMazeMiddle = room{
// 	compressedRoomData:   roomGfxMazeMiddle,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorDarkGray),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomMazeSide = room{
// 	compressedRoomData:   roomGfxMazeSide,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorDarkGray),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomMazeEntry = room{
// 	compressedRoomData:   roomGfxMazeEntry,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorDarkGray),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomRedMazeTopLeft = room{
// 	compressedRoomData:   roomGfxRedMaze1,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorRed),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomRedMazeBottomLeft = room{
// 	compressedRoomData:   roomGfxRedMazeBottom,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorRed),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomRedMazeTopRight = room{
// 	compressedRoomData:   roomGfxRedMazeTop,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorRed),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomRedMazeEntryBottomRight = room{
// 	compressedRoomData:   roomGfxWhiteCastleEntry,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorRed),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomBlackMazeTopLeft = room{
// 	compressedRoomData:   roomGfxBlackMaze1,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorOrange),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomBlackMazeBottomLeft = room{
// 	compressedRoomData:   roomGfxBlackMaze3,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorOrange),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomBlackMazeTopRight = room{
// 	compressedRoomData:   roomGfxBlackMaze2,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorOrange),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// var roomBlackMazeEntryBottomRight = room{
// 	compressedRoomData:   roomGfxBlackMazeEntry,
// 	uncompressedRoomData: nil,
// 	roomStyle: tcell.StyleDefault.
// 		Background(tcell.ColorDarkGray).
// 		Foreground(tcell.ColorOrange),
// 	up:    nil,
// 	down:  nil,
// 	left:  nil,
// 	right: nil,
// }

// this function is needed to avoid errors in assignment order
func initDirections() {
	// roomYellowCastle.up = &roomOnTopOfYellowCastle
	// roomYellowCastle.down = &roomStartRoomTopEntryRoom
	roomYellowCastle.up = nil
	roomYellowCastle.down = nil
	roomYellowCastle.left = nil
	roomYellowCastle.right = nil

	// roomBlackCastle.up = &roomOnTopOfBlackCastle
	// roomBlackCastle.down = &roomBlueMazeTop
	// roomBlackCastle.left = nil
	// roomBlackCastle.right = nil

	// roomWhiteCastle.up = &roomRedMazeEntryBottomRight
	// roomWhiteCastle.down = &roomBelowWhiteCastle
	// roomWhiteCastle.left = nil
	// roomWhiteCastle.right = nil

	// roomStartRoomTopEntryRoom.up = &roomYellowCastle
	// roomStartRoomTopEntryRoom.down = nil
	// roomStartRoomTopEntryRoom.left = &roomLeftOfStartRoom
	// roomStartRoomTopEntryRoom.right = &roomLeftOfNameRoom

	// roomPurpleEasterEggTopEntryRoom.up = nil
	// roomPurpleEasterEggTopEntryRoom.down = nil
	// roomPurpleEasterEggTopEntryRoom.left = &roomLeftOfNameRoom
	// roomPurpleEasterEggTopEntryRoom.right = &roomPurpleEasterEggTopEntryRoom

	// roomNumberRoom.up = nil
	// roomNumberRoom.down = &roomBelowNumberRoom
	// roomNumberRoom.left = nil
	// roomNumberRoom.right = nil

	// roomOnTopOfYellowCastle.up = nil
	// roomOnTopOfYellowCastle.down = &roomYellowCastle
	// roomOnTopOfYellowCastle.left = nil
	// roomOnTopOfYellowCastle.right = nil

	// roomLeftOfStartRoom.up = &roomBlueMazeEntry
	// roomLeftOfStartRoom.down = nil
	// roomLeftOfStartRoom.left = nil
	// roomLeftOfStartRoom.right = &roomStartRoomTopEntryRoom

	// roomTwoBelowWhiteCastleRoomTopEntryRoom.up = &roomBelowWhiteCastle
	// roomTwoBelowWhiteCastleRoomTopEntryRoom.down = nil
	// roomTwoBelowWhiteCastleRoomTopEntryRoom.left = nil
	// roomTwoBelowWhiteCastleRoomTopEntryRoom.right = nil

	// roomLowerRedRoomTopEntryRoom.up = &roomBelowNumberRoom
	// roomLowerRedRoomTopEntryRoom.down = nil
	// roomLowerRedRoomTopEntryRoom.left = nil
	// roomLowerRedRoomTopEntryRoom.right = nil

	// roomLeftOfNameRoom.up = nil
	// roomLeftOfNameRoom.down = &roomMazeEntry
	// roomLeftOfNameRoom.left = &roomStartRoomTopEntryRoom
	// roomLeftOfNameRoom.right = &roomPurpleEasterEggTopEntryRoom

	// roomOnTopOfBlackCastle.up = &roomBlackMazeEntryBottomRight
	// roomOnTopOfBlackCastle.down = &roomBlackCastle
	// roomOnTopOfBlackCastle.left = nil
	// roomOnTopOfBlackCastle.right = nil

	// roomBelowWhiteCastle.up = &roomWhiteCastle
	// roomBelowWhiteCastle.down = &roomTwoBelowWhiteCastleRoomTopEntryRoom
	// roomBelowWhiteCastle.left = nil
	// roomBelowWhiteCastle.right = &roomMazeSide

	// roomBelowNumberRoom.up = &roomNumberRoom
	// roomBelowNumberRoom.down = &roomLowerRedRoomTopEntryRoom
	// roomBelowNumberRoom.left = &roomMazeSide
	// roomBelowNumberRoom.right = nil

	// roomBlueMazeEntry.up = &roomBlueMazeLeft
	// roomBlueMazeEntry.down = &roomLeftOfStartRoom
	// roomBlueMazeEntry.left = &roomBlueMazeCenter
	// roomBlueMazeEntry.right = &roomBlueMazeCenter

	// roomBlueMazeCenter.up = &roomBlueMazeTop
	// roomBlueMazeCenter.down = &roomBlueMazeBottom
	// roomBlueMazeCenter.left = &roomBlueMazeEntry
	// roomBlueMazeCenter.right = &roomBlueMazeEntry

	// roomBlueMazeBottom.up = &roomBlueMazeCenter
	// roomBlueMazeBottom.down = nil
	// roomBlueMazeBottom.left = &roomBlueMazeLeft
	// roomBlueMazeBottom.right = &roomBlueMazeTop

	// roomBlueMazeLeft.up = nil
	// roomBlueMazeLeft.down = &roomBlueMazeEntry
	// roomBlueMazeLeft.left = &roomBlueMazeTop
	// roomBlueMazeLeft.right = &roomBlueMazeBottom

	// roomBlueMazeTop.up = &roomBlackCastle
	// roomBlueMazeTop.down = &roomBlueMazeCenter
	// roomBlueMazeTop.left = &roomBlueMazeEntry
	// roomBlueMazeTop.right = &roomBlueMazeLeft

	// roomMazeMiddle.up = &roomMazeEntry
	// roomMazeMiddle.down = &roomMazeSide
	// roomMazeMiddle.left = &roomMazeEntry
	// roomMazeMiddle.right = &roomMazeEntry

	// roomMazeSide.up = &roomMazeMiddle
	// roomMazeSide.down = nil
	// roomMazeSide.left = &roomBelowWhiteCastle
	// roomMazeSide.right = &roomBelowNumberRoom

	// roomMazeEntry.up = &roomLeftOfNameRoom
	// roomMazeEntry.down = &roomMazeMiddle
	// roomMazeEntry.left = &roomMazeMiddle
	// roomMazeEntry.right = &roomMazeMiddle

	// roomRedMazeTopLeft.up = nil
	// roomRedMazeTopLeft.down = &roomRedMazeBottomLeft
	// roomRedMazeTopLeft.left = &roomRedMazeTopRight
	// roomRedMazeTopLeft.right = &roomRedMazeTopRight

	// roomRedMazeBottomLeft.up = &roomRedMazeTopLeft
	// roomRedMazeBottomLeft.down = nil
	// roomRedMazeBottomLeft.left = &roomRedMazeEntryBottomRight
	// roomRedMazeBottomLeft.right = &roomRedMazeEntryBottomRight

	// roomRedMazeTopRight.up = nil
	// roomRedMazeTopRight.down = &roomRedMazeEntryBottomRight
	// roomRedMazeTopRight.left = &roomRedMazeTopLeft
	// roomRedMazeTopRight.right = &roomRedMazeTopLeft

	// roomRedMazeEntryBottomRight.up = &roomRedMazeTopRight
	// roomRedMazeEntryBottomRight.down = &roomWhiteCastle
	// roomRedMazeEntryBottomRight.left = &roomRedMazeBottomLeft
	// roomRedMazeEntryBottomRight.right = &roomRedMazeBottomLeft

	// roomBlackMazeTopLeft.up = &roomBlackMazeBottomLeft
	// roomBlackMazeTopLeft.down = &roomBlackMazeBottomLeft
	// roomBlackMazeTopLeft.left = &roomBlackMazeEntryBottomRight
	// roomBlackMazeTopLeft.right = &roomBlackMazeTopRight

	// roomBlackMazeBottomLeft.up = &roomBlackMazeTopLeft
	// roomBlackMazeBottomLeft.down = &roomBlackMazeTopLeft
	// roomBlackMazeBottomLeft.left = &roomBlackMazeTopRight
	// roomBlackMazeBottomLeft.right = &roomBlackMazeEntryBottomRight

	// roomBlackMazeTopRight.up = nil
	// roomBlackMazeTopRight.down = &roomBlackMazeEntryBottomRight
	// roomBlackMazeTopRight.left = &roomBlackMazeTopLeft
	// roomBlackMazeTopRight.right = &roomBlackMazeBottomLeft

	// roomBlackMazeEntryBottomRight.up = &roomBlackMazeTopRight
	// roomBlackMazeEntryBottomRight.down = &roomBlackCastle
	// roomBlackMazeEntryBottomRight.left = &roomBlackMazeBottomLeft
	// roomBlackMazeEntryBottomRight.right = &roomBlackMazeTopLeft

}
