// TO-DO
// own struct for sizes
// add all screens
// add and initialize all directions
// monolithisch erstellen

package main

import "github.com/gdamore/tcell"

type compressedRoom []string

type roomID int

type rooms struct {
	width               int
	height              int
	defaultXFactor      int
	defaultYFactor      int
	defaultYInnerFactor int
	compressedRoomData  compressedRoom
	up                  *rooms
	down                *rooms
	left                *rooms
	right               *rooms
}

var YellowCastle = rooms{
	width:               160,
	height:              44,
	defaultXFactor:      4,
	defaultYFactor:      2,
	defaultYInnerFactor: 4,
	compressedRoomData:  roomGfxCastle,
	up:                  nil,
	down:                &topEntryRoom,
	left:                nil,
	right:               nil,
}

var topEntryRoom = rooms{
	width:               160,
	height:              44,
	defaultXFactor:      4,
	defaultYFactor:      2,
	defaultYInnerFactor: 4,
	compressedRoomData:  roomGfxTopEntryRoom,
	up:                  nil,
	down:                nil,
	left:                nil,
	right:               nil,
}

func initDirections(r *rooms) {
	topEntryRoom.up = &YellowCastle
}

func display(s tcell.Screen, r *rooms) {

	// line by line
	// top and last lines repeat twice, the inner lines 8 times
	for lines, content := range r.compressedRoomData[0:1] {
		for yScale := 0; yScale < r.defaultYFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.defaultXFactor; xScale++ {
					emitStr(s, ((r.defaultXFactor * columns) + xScale), (lines + yScale), player.style, string(char))
				}
			}
		}
	}

	for lines, content := range r.compressedRoomData[1:11] {
		for yScale := 0; yScale < r.defaultYInnerFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.defaultXFactor; xScale++ {
					emitStr(s, ((r.defaultXFactor * columns) + xScale), ((r.defaultYInnerFactor * (lines)) + r.defaultYFactor + yScale), player.style, string(char))
				}
			}
		}
	}

	for _, content := range r.compressedRoomData[11:] {
		for yScale := 0; yScale < r.defaultYFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.defaultXFactor; xScale++ {
					emitStr(s, ((r.defaultXFactor * columns) + xScale), ((r.height) - r.defaultYFactor + yScale), player.style, string(char))
					// r.defaultYFactor + ((r.defaultYInnerFactor*r.defaultYFactor)*(r.height-(2*r.defaultYFactor)) + yScale)), player.style, string(char))
				}
			}
		}
	}
}
