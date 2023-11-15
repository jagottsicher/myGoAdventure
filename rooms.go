package main

import (
	"github.com/gdamore/tcell"
)

type compressedRoom []string

// Scaling the room
var defaultXFactor, defaultYFactor int

// Internal scaling factor
// const defaultYInnerFactor = 4

var stages []tcell.Screen

func initStages(r []room) {

	// Any maze

	// defaultXFactor = int(math.Round(float64(stageWidth) / 40))
	// defaultYFactor = 2

	// stageYFactor := int(math.Floor(float64(stageHeight) / 12))

	// 	var percentageY float64
	// 	var rowValue int
	// 	MaxRowValue := 12
	// 	var percentageX float64
	// 	var columnValue int
	// 	MaxColumnValue := 40

	// 	var theRow string
	// 	var theSpot rune

	// 	for i := 0; i < 30; i++ {
	// 		// create one screen indexed by id
	// 		stages[i], err = tcell.NewScreen()
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "%v\n", err)
	// 			os.Exit(1)
	// 		}
	// 		if err := stages[i].Init(); err != nil {
	// 			fmt.Fprintf(os.Stderr, "%v\n", err)
	// 			os.Exit(1)
	// 		}

	// 		// fill all screens
	// 		for y := 0; y < stageHeight; y++ {
	// 			percentageY = float64((y * 100) / (stageHeight - 1))
	// 			rowValue = int(12 * int(percentageY) / 100)
	// 			if rowValue == MaxRowValue {
	// 				rowValue = MaxRowValue - 1
	// 			}
	// 			theRow = r[i].compressedRoomData[rowValue]

	// 			for x := 0; x < stageWidth; x++ {
	// 				percentageX = float64((x * 100) / (stageWidth - 1))
	// 				columnValue = int(40 * int(percentageX) / 100)
	// 				if columnValue == MaxColumnValue {
	// 					columnValue = MaxColumnValue - 1
	// 				}

	// 				theSpot = rune(theRow[columnValue])

	//				stages[i].SetContent(x, y, theSpot, nil, tcell.StyleDefault.
	//					Background(tcell.ColorDarkGray).
	//					Foreground(tcell.ColorYellow))
	//			}
	//		}
	//	}
}

type room struct {
	id                 uint8
	compressedRoomData compressedRoom
	uncompressedData   *tcell.Screen
	roomStyle          tcell.Style
	barLeft            bool
	barRight           bool
	up                 *room
	down               *room
	left               *room
	right              *room
}

var rooms []room

var roomYellowCastle = room{
	id:                 0,
	compressedRoomData: roomGfxCastle,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorYellow),
	up:    nil,
	down:  &roomStartRoomTopEntryRoom,
	left:  nil,
	right: nil,
}

var roomBlackCastle = room{
	id:                 1,
	compressedRoomData: roomGfxCastle,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlack),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomWhiteCastle = room{
	id:                 2,
	compressedRoomData: roomGfxCastle,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorWhite),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomStartRoomTopEntryRoom = room{
	id:                 3,
	compressedRoomData: roomGfxBelowYellowCastle,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorGreen),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomPurpleEasterEggTopEntryRoom = room{
	id:                 4,
	compressedRoomData: roomGfxTopEntryRoom,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorPurple),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomNumberRoom = room{
	id:                 5,
	compressedRoomData: roomGfxNumberRoom,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorPurple),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomOnTopOfYellowCastle = room{
	id:                 6,
	compressedRoomData: roomGfxNumberRoom,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorYellow),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomLeftOfStartRoom = room{
	id:                 7,
	compressedRoomData: roomGfxBelowYellowCastle,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorDarkGreen),
	barLeft: true,
	up:      nil,
	down:    nil,
	left:    nil,
	right:   nil,
}

var roomTwoBelowWhiteCastleRoomTopEntryRoom = room{
	id:                 8,
	compressedRoomData: roomGfxTopEntryRoom,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorDarkGreen),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomLowerRedRoomTopEntryRoom = room{
	id:                 9,
	compressedRoomData: roomGfxTopEntryRoom,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomLeftOfNameRoom = room{
	id:                 10,
	compressedRoomData: roomGfxLeftOfName,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorGreenYellow),
	barRight: true,
	up:       nil,
	down:     nil,
	left:     nil,
	right:    nil,
}

var roomOnTopOfBlackCastle = room{
	id:                 11,
	compressedRoomData: roomGfxSideCorridor,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBelowWhiteCastle = room{
	id:                 12,
	compressedRoomData: roomGfxSideCorridor,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorGreen),
	barLeft: true,
	up:      nil,
	down:    nil,
	left:    nil,
	right:   nil,
}

var roomBelowNumberRoom = room{
	id:                 13,
	compressedRoomData: roomGfxSideCorridor,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorLightBlue),
	barRight: true,
	up:       nil,
	down:     nil,
	left:     nil,
	right:    nil,
}

var roomBlueMazeEntry = room{
	id:                 14,
	compressedRoomData: roomGfxBlueMazeEntry,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlueMazeCenter = room{
	id:                 15,
	compressedRoomData: roomGfxBlueMazeCenter,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlueMazeBottom = room{
	id:                 16,
	compressedRoomData: roomGfxBlueMazeBottom,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlueMazeLeft = room{
	id:                 17,
	compressedRoomData: roomGfxBlueMaze1,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlueMazeTop = room{
	id:                 18,
	compressedRoomData: roomGfxBlueMazeTop,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomMazeMiddle = room{
	id:                 19,
	compressedRoomData: roomGfxMazeMiddle,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorDarkGray),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomMazeSide = room{
	id:                 20,
	compressedRoomData: roomGfxMazeSide,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorDarkGray),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomMazeEntry = room{
	id:                 21,
	compressedRoomData: roomGfxMazeEntry,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorDarkGray),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomRedMazeTopLeft = room{
	id:                 22,
	compressedRoomData: roomGfxRedMaze1,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomRedMazeBottomLeft = room{
	id:                 23,
	compressedRoomData: roomGfxRedMazeBottom,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomRedMazeTopRight = room{
	id:                 23,
	compressedRoomData: roomGfxRedMazeTop,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomRedMazeEntryBottomRight = room{
	id:                 25,
	compressedRoomData: roomGfxWhiteCastleEntry,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlackMazeTopLeft = room{
	id:                 26,
	compressedRoomData: roomGfxBlackMaze1,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorOrange),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlackMazeBottomLeft = room{
	id:                 27,
	compressedRoomData: roomGfxBlackMaze3,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorOrange),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlackMazeTopRight = room{
	id:                 28,
	compressedRoomData: roomGfxBlackMaze2,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorOrange),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlackMazeEntryBottomRight = room{
	id:                 29,
	compressedRoomData: roomGfxBlackMazeEntry,
	uncompressedData:   nil,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorOrange),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

// this function is needed to avoid errors in assignment order
func initDirections() {
	roomYellowCastle.up = &roomOnTopOfYellowCastle
	roomYellowCastle.down = &roomStartRoomTopEntryRoom
	roomYellowCastle.left = nil
	roomYellowCastle.right = nil

	roomBlackCastle.up = &roomOnTopOfBlackCastle
	roomBlackCastle.down = &roomBlueMazeTop
	roomBlackCastle.left = nil
	roomBlackCastle.right = nil

	roomWhiteCastle.up = &roomRedMazeEntryBottomRight
	roomWhiteCastle.down = &roomBelowWhiteCastle
	roomWhiteCastle.left = nil
	roomWhiteCastle.right = nil

	roomStartRoomTopEntryRoom.up = &roomYellowCastle
	roomStartRoomTopEntryRoom.down = nil
	roomStartRoomTopEntryRoom.left = &roomLeftOfStartRoom
	roomStartRoomTopEntryRoom.right = &roomLeftOfNameRoom

	roomPurpleEasterEggTopEntryRoom.up = nil
	roomPurpleEasterEggTopEntryRoom.down = nil
	roomPurpleEasterEggTopEntryRoom.left = &roomLeftOfNameRoom
	roomPurpleEasterEggTopEntryRoom.right = &roomPurpleEasterEggTopEntryRoom

	roomNumberRoom.up = nil
	roomNumberRoom.down = &roomBelowNumberRoom
	roomNumberRoom.left = nil
	roomNumberRoom.right = nil

	roomOnTopOfYellowCastle.up = nil
	roomOnTopOfYellowCastle.down = &roomYellowCastle
	roomOnTopOfYellowCastle.left = nil
	roomOnTopOfYellowCastle.right = nil

	roomLeftOfStartRoom.up = &roomBlueMazeEntry
	roomLeftOfStartRoom.down = nil
	roomLeftOfStartRoom.left = nil
	roomLeftOfStartRoom.right = &roomStartRoomTopEntryRoom

	roomTwoBelowWhiteCastleRoomTopEntryRoom.up = &roomBelowWhiteCastle
	roomTwoBelowWhiteCastleRoomTopEntryRoom.down = nil
	roomTwoBelowWhiteCastleRoomTopEntryRoom.left = nil
	roomTwoBelowWhiteCastleRoomTopEntryRoom.right = nil

	roomLowerRedRoomTopEntryRoom.up = &roomBelowNumberRoom
	roomLowerRedRoomTopEntryRoom.down = nil
	roomLowerRedRoomTopEntryRoom.left = nil
	roomLowerRedRoomTopEntryRoom.right = nil

	roomLeftOfNameRoom.up = nil
	roomLeftOfNameRoom.down = &roomMazeEntry
	roomLeftOfNameRoom.left = &roomStartRoomTopEntryRoom
	roomLeftOfNameRoom.right = &roomPurpleEasterEggTopEntryRoom

	roomOnTopOfBlackCastle.up = &roomBlackMazeEntryBottomRight
	roomOnTopOfBlackCastle.down = &roomBlackCastle
	roomOnTopOfBlackCastle.left = nil
	roomOnTopOfBlackCastle.right = nil

	roomBelowWhiteCastle.up = &roomWhiteCastle
	roomBelowWhiteCastle.down = &roomTwoBelowWhiteCastleRoomTopEntryRoom
	roomBelowWhiteCastle.left = nil
	roomBelowWhiteCastle.right = &roomMazeSide

	roomBelowNumberRoom.up = &roomNumberRoom
	roomBelowNumberRoom.down = &roomLowerRedRoomTopEntryRoom
	roomBelowNumberRoom.left = &roomMazeSide
	roomBelowNumberRoom.right = nil

	roomBlueMazeEntry.up = &roomBlueMazeLeft
	roomBlueMazeEntry.down = &roomLeftOfStartRoom
	roomBlueMazeEntry.left = &roomBlueMazeCenter
	roomBlueMazeEntry.right = &roomBlueMazeCenter

	roomBlueMazeCenter.up = &roomBlueMazeTop
	roomBlueMazeCenter.down = &roomBlueMazeBottom
	roomBlueMazeCenter.left = &roomBlueMazeEntry
	roomBlueMazeCenter.right = &roomBlueMazeEntry

	roomBlueMazeBottom.up = &roomBlueMazeCenter
	roomBlueMazeBottom.down = nil
	roomBlueMazeBottom.left = &roomBlueMazeLeft
	roomBlueMazeBottom.right = &roomBlueMazeTop

	roomBlueMazeLeft.up = nil
	roomBlueMazeLeft.down = &roomBlueMazeEntry
	roomBlueMazeLeft.left = &roomBlueMazeTop
	roomBlueMazeLeft.right = &roomBlueMazeBottom

	roomBlueMazeTop.up = &roomBlackCastle
	roomBlueMazeTop.down = &roomBlueMazeCenter
	roomBlueMazeTop.left = &roomBlueMazeEntry
	roomBlueMazeTop.right = &roomBlueMazeLeft

	roomMazeMiddle.up = &roomMazeEntry
	roomMazeMiddle.down = &roomMazeSide
	roomMazeMiddle.left = &roomMazeEntry
	roomMazeMiddle.right = &roomMazeEntry

	roomMazeSide.up = &roomMazeMiddle
	roomMazeSide.down = nil
	roomMazeSide.left = &roomBelowWhiteCastle
	roomMazeSide.right = &roomBelowNumberRoom

	roomMazeEntry.up = &roomLeftOfNameRoom
	roomMazeEntry.down = &roomMazeMiddle
	roomMazeEntry.left = &roomMazeMiddle
	roomMazeEntry.right = &roomMazeMiddle

	roomRedMazeTopLeft.up = nil
	roomRedMazeTopLeft.down = &roomRedMazeBottomLeft
	roomRedMazeTopLeft.left = &roomRedMazeTopRight
	roomRedMazeTopLeft.right = &roomRedMazeTopRight

	roomRedMazeBottomLeft.up = &roomRedMazeTopLeft
	roomRedMazeBottomLeft.down = nil
	roomRedMazeBottomLeft.left = &roomRedMazeEntryBottomRight
	roomRedMazeBottomLeft.right = &roomRedMazeEntryBottomRight

	roomRedMazeTopRight.up = nil
	roomRedMazeTopRight.down = &roomRedMazeEntryBottomRight
	roomRedMazeTopRight.left = &roomRedMazeTopLeft
	roomRedMazeTopRight.right = &roomRedMazeTopLeft

	roomRedMazeEntryBottomRight.up = &roomRedMazeTopRight
	roomRedMazeEntryBottomRight.down = &roomWhiteCastle
	roomRedMazeEntryBottomRight.left = &roomRedMazeBottomLeft
	roomRedMazeEntryBottomRight.right = &roomRedMazeBottomLeft

	roomBlackMazeTopLeft.up = &roomBlackMazeBottomLeft
	roomBlackMazeTopLeft.down = &roomBlackMazeBottomLeft
	roomBlackMazeTopLeft.left = &roomBlackMazeEntryBottomRight
	roomBlackMazeTopLeft.right = &roomBlackMazeTopRight

	roomBlackMazeBottomLeft.up = &roomBlackMazeTopLeft
	roomBlackMazeBottomLeft.down = &roomBlackMazeTopLeft
	roomBlackMazeBottomLeft.left = &roomBlackMazeTopRight
	roomBlackMazeBottomLeft.right = &roomBlackMazeEntryBottomRight

	roomBlackMazeTopRight.up = nil
	roomBlackMazeTopRight.down = &roomBlackMazeEntryBottomRight
	roomBlackMazeTopRight.left = &roomBlackMazeTopLeft
	roomBlackMazeTopRight.right = &roomBlackMazeBottomLeft

	roomBlackMazeEntryBottomRight.up = &roomBlackMazeTopRight
	roomBlackMazeEntryBottomRight.down = &roomBlackCastle
	roomBlackMazeEntryBottomRight.left = &roomBlackMazeBottomLeft
	roomBlackMazeEntryBottomRight.right = &roomBlackMazeTopLeft

}
