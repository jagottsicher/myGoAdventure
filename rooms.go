// TO-DO
// own struct for sizes
// add all screens
// add and initialize all directions
// monolithisch erstellen

package main

import "github.com/gdamore/tcell"

type compressedRoom []string

type roomDimensions struct {
	dimensions                                          size
	defaultXFactor, defaultYFactor, defaultYInnerFactor int
}

// type size declared in graphics.go
var defaultRoomSize = size{
	width:  160,
	height: 44,
}

var defaultDimensions = roomDimensions{
	dimensions:          defaultRoomSize,
	defaultXFactor:      4,
	defaultYFactor:      2,
	defaultYInnerFactor: 4,
}

type rooms struct {
	allRoomDimensions  roomDimensions
	compressedRoomData compressedRoom
	roomStyle          tcell.Style
	up                 *rooms
	down               *rooms
	left               *rooms
	right              *rooms
}

var roomYellowCastle = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxCastle,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorYellow),
	up:    nil,
	down:  &roomStartRoomTopEntryRoom,
	left:  nil,
	right: nil,
}

var roomBlackCastle = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxCastle,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorBlack),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomWhiteCastle = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxCastle,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorWhite),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomStartRoomTopEntryRoom = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBelowYellowCastle,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorGreen),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomPurpleEasterEggTopEntryRoom = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxTopEntryRoom,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorPurple),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomNumberRoom = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxNumberRoom,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorPurple),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomOnTopOfYellowCastle = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxNumberRoom,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorYellow),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomLeftOfStartRoom = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBelowYellowCastle,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorDarkGreen),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomTwoBelowWhiteCastleRoomTopEntryRoom = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxTopEntryRoom,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorDarkGreen),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomLowerRedRoomTopEntryRoom = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxTopEntryRoom,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomLeftOfNameRoom = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxLeftOfName,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorGreenYellow),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomOnTopOfBlackCastle = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxSideCorridor,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBelowWhiteCastle = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxSideCorridor,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorGreen),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBelowNumberRoom = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxSideCorridor,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorLightBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlueMazeEntry = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBlueMazeEntry,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlueMazeCenter = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBlueMazeCenter,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlueMazeBottom = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBlueMazeBottom,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlueMazeLeft = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBlueMaze1,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlueMazeTop = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBlueMazeTop,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorBlue),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomMazeMiddle = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxMazeMiddle,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorOrange).
		Foreground(tcell.ColorGray),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomMazeSide = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxMazeSide,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorOrange).
		Foreground(tcell.ColorGray),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomMazeEntry = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxMazeEntry,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorOrange).
		Foreground(tcell.ColorGray),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomRedMazeTopLeft = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxRedMaze1,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomRedMazeBottomLeft = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxRedMazeBottom,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomRedMazeTopRight = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxRedMazeTop,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomRedMazeEntryBottomRight = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxWhiteCastleEntry,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorRed),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlackMazeTopLeft = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBlackMaze1,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorOrange),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlackMazeBottomLeft = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBlackMaze3,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorOrange),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlackMazeTopRight = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBlackMaze2,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
		Foreground(tcell.ColorOrange),
	up:    nil,
	down:  nil,
	left:  nil,
	right: nil,
}

var roomBlackMazeEntryBottomRight = rooms{
	allRoomDimensions:  defaultDimensions,
	compressedRoomData: roomGfxBlackMazeEntry,
	roomStyle: tcell.StyleDefault.
		Background(tcell.ColorGray).
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

func display(s tcell.Screen, r *rooms) {

	// set the room colors
	roomStyle := r.roomStyle
	s.SetStyle(roomStyle)

	// line by line
	// top and last lines repeat twice, the inner lines 8 times
	for lines, content := range r.compressedRoomData[0:1] {
		for yScale := 0; yScale < r.allRoomDimensions.defaultYFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.allRoomDimensions.defaultXFactor; xScale++ {
					emitStr(s, ((r.allRoomDimensions.defaultXFactor * columns) + xScale), (lines + yScale), roomStyle, string(char))
				}
			}
		}
	}

	for lines, content := range r.compressedRoomData[1:11] {
		for yScale := 0; yScale < r.allRoomDimensions.defaultYInnerFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.allRoomDimensions.defaultXFactor; xScale++ {
					emitStr(s, ((r.allRoomDimensions.defaultXFactor * columns) + xScale), ((r.allRoomDimensions.defaultYInnerFactor * (lines)) + r.allRoomDimensions.defaultYFactor + yScale), roomStyle, string(char))
				}
			}
		}
	}

	for _, content := range r.compressedRoomData[11:] {
		for yScale := 0; yScale < r.allRoomDimensions.defaultYFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.allRoomDimensions.defaultXFactor; xScale++ {
					emitStr(s, ((r.allRoomDimensions.defaultXFactor * columns) + xScale), ((r.allRoomDimensions.dimensions.height) - r.allRoomDimensions.defaultYFactor + yScale), roomStyle, string(char))
				}
			}
		}
	}
}
