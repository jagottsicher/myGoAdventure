package main

import (
	"math"

	"github.com/gdamore/tcell/v2"
)

// cell is a single terminal character at a fixed offset within a sprite.
type cell struct {
	x, y   int
	symbol rune
}

// object is a game entity that can be positioned and rendered.
type object struct {
	posX   float64     // position in room grid coordinates (0..40)
	posY   float64     // position in room grid coordinates (0..7)
	stepX  float64     // movement step in X (grid units per key press)
	stepY  float64     // movement step in Y (grid units per key press)
	width  int         // bounding box width in grid columns
	height int         // bounding box height in grid rows
	roomID int         // index of the room the object is currently in
	style  tcell.Style
	shape  []*cell
}

// The player is the "ball" in Adventure terms.
var player *object

// allObjects holds every object that should be rendered each frame.
var allObjects []*object

// Grid dimensions – must match the playfield decoder output.
const (
	gridCols = 40
	gridRows = 7
)

// startRoomID is the room where the player begins (Yellow Castle, game 2).
const startRoomID = 0x11

// startX and startY are the initial player grid coordinates inside the
// Yellow Castle. Derived from the C++ source:
//   objectBall.x = 0x50*2 = 160  → grid col = 160/8 = 20
//   objectBall.y = 0x20*2 = 64   → ypos = 64/32 = 2, row = 6-2 = 4, display_row = 6-4 = 2
const (
	startX = 20.0
	startY = 2.5
)

func initPlayer() {
	player = &object{
		posX:   startX,
		posY:   startY,
		stepX:  1.0, // 1 grid column per key press
		stepY:  0.5, // 0.5 grid rows per key press
		width:  2,
		height: 1,
		roomID: startRoomID,
		style:  tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack),
		shape:  playerGfx,
	}
	allObjects = append(allObjects, player)
}

// movePlayer attempts to move the player by (dx, dy) in grid coordinates.
// Wall collision is checked; if the new position overlaps a wall the move is
// blocked. Room transitions happen when the player exits an edge.
func movePlayer(dx, dy float64) {
	r := roomByID(player.roomID)
	if r == nil {
		return
	}

	newX := player.posX + dx
	newY := player.posY + dy

	// --- Room transitions (exit through edges) ---
	if newX+float64(player.width) > float64(gridCols) {
		// Exiting right edge
		nextRoom := roomByID(r.rightID)
		if nextRoom != nil {
			player.roomID = nextRoom.id
			player.posX = 0
			player.posY = newY
		}
		return
	}
	if newX < 0 {
		// Exiting left edge
		nextRoom := roomByID(r.leftID)
		if nextRoom != nil {
			player.roomID = nextRoom.id
			player.posX = float64(gridCols - player.width)
			player.posY = newY
		}
		return
	}
	if newY < 0 {
		// Exiting top edge
		nextRoom := roomByID(r.upID)
		if nextRoom != nil {
			player.roomID = nextRoom.id
			player.posX = newX
			player.posY = float64(gridRows - player.height)
		}
		return
	}
	if newY+float64(player.height) > float64(gridRows) {
		// Exiting bottom edge
		nextRoom := roomByID(r.downID)
		if nextRoom != nil {
			player.roomID = nextRoom.id
			player.posX = newX
			player.posY = 0
		}
		return
	}

	// --- Wall collision ---
	if !playerCollidesWall(r, newX, newY) {
		player.posX = newX
		player.posY = newY
	}
}

// playerCollidesWall returns true if a player at grid position (x, y) would
// overlap any wall cell in the given room.
//
// The player bounding box is [x, x+width) × [y, y+height) in continuous grid
// coordinates.  We enumerate every discrete grid cell whose [col, col+1) ×
// [row, row+1) interval overlaps that box.  This handles fractional positions
// correctly (e.g. posY = 2.5 with height = 1 spans rows 2 and 3).
func playerCollidesWall(r *room, x, y float64) bool {
	colMin := int(x)
	colMax := int(math.Ceil(x+float64(player.width))) - 1
	rowMin := int(y)
	rowMax := int(math.Ceil(y+float64(player.height))) - 1

	// Clamp to valid grid range.
	if colMin < 0 {
		colMin = 0
	}
	if colMax >= gridCols {
		colMax = gridCols - 1
	}
	if rowMin < 0 {
		rowMin = 0
	}
	if rowMax >= gridRows {
		rowMax = gridRows - 1
	}

	for row := rowMin; row <= rowMax; row++ {
		for col := colMin; col <= colMax; col++ {
			if r.isWall(col, row) {
				return true
			}
		}
	}
	return false
}
