package world

import "github.com/gdamore/tcell/v2"

// Room flags matching the original Atari 2600 room properties.
const (
	RoomFlagNone        = 0x00
	RoomFlagMirror      = 0x01 // right half is mirror of left half (else reversed)
	RoomFlagLeftWall    = 0x02 // thin left boundary wall
	RoomFlagRightWall   = 0x04 // thin right boundary wall
)

// room holds all data needed to render and navigate a single screen.
type room struct {
	id         int
	walls      [7][40]bool   // decoded playfield: walls[row][col], row 0 = top
	color      tcell.Color   // foreground/wall colour
	flags      uint8
	upID       int           // room index reached by walking off the top edge
	downID     int           // room index reached by walking off the bottom edge
	leftID     int           // room index reached by walking off the left edge
	rightID    int           // room index reached by walking off the right edge
}

// buildWalls converts a decoded [7]string room graphic into a boolean wall map.
func buildWalls(rows [7]string) [7][40]bool {
	var walls [7][40]bool
	for r, row := range rows {
		for c, ch := range row {
			walls[r][c] = ch == 'X'
		}
	}
	return walls
}

// isWall returns true if the given grid position contains a wall.
// Uses integer grid coordinates to avoid floating-point edge cases.
func (r *room) isWall(col, row int) bool {
	if row < 0 || row >= 7 || col < 0 || col >= 40 {
		return true // out-of-bounds counts as wall (edge of world)
	}
	return r.walls[row][col]
}

// The 30 rooms of Adventure, indexed as in the original C++ source (0x00–0x1E).
// Connections use game-2 (medium difficulty) settings.
var rooms [31]*room

// roomLevelDiffs resolves the level-dependent room connections encoded as 0x80+index.
// For game 2 (gameLevel = 1, zero-indexed).
// Index mapping (each group of 3 is levels 1/2/3):
//   0-2:  down from room 0x01
//   3-5:  down from room 0x02
//   6-8:  down from room 0x03
//   9-11: all four dirs from room 0x1B
//   12-14: down from room 0x1C
//   15-17: up from room 0x1D
var roomLevelDiffs = [18]int{
	0x10, 0x0F, 0x0F, // down from room 0x01
	0x05, 0x11, 0x11, // down from room 0x02
	0x1D, 0x0A, 0x0A, // down from room 0x03
	0x1C, 0x16, 0x16, // u/l/r/d from room 0x1B
	0x1B, 0x0C, 0x0C, // down from room 0x1C
	0x03, 0x0C, 0x0C, // up from room 0x1D
}

// adjustRoom resolves 0x80-flagged room indices for game level 2 (medium).
func adjustRoom(id int) int {
	const gameLevel = 1 // 0=easy, 1=medium, 2=hard
	if id&0x80 != 0 {
		idx := (id & ^0x80) + gameLevel
		return roomLevelDiffs[idx]
	}
	return id
}

func initRooms() {
	// Room 0x00 – Number Room (purple, all exits looping)
	rooms[0x00] = &room{
		id:    0x00,
		walls: buildWalls(decodeRoom(rawNumberRoom, false)),
		color: tcell.ColorPurple,
		flags: RoomFlagNone,
		upID: 0x00, downID: 0x00, leftID: 0x00, rightID: 0x00,
	}

	// Room 0x01 – Top Access (olive green, left thin wall)
	rooms[0x01] = &room{
		id:    0x01,
		walls: buildWalls(decodeRoom(rawBelowYellowCastle, false)),
		color: tcell.ColorOlive,
		flags: RoomFlagLeftWall,
		upID: 0x08, rightID: 0x02,
		downID:  adjustRoom(0x80), // level-dependent: 0x0F for game 2
		leftID:  0x03,
	}

	// Room 0x02 – Top Access (lime green)
	rooms[0x02] = &room{
		id:    0x02,
		walls: buildWalls(decodeRoom(rawBelowYellowCastle, false)),
		color: tcell.ColorLimeGreen,
		flags: RoomFlagNone,
		upID: 0x11, rightID: 0x03,
		downID:  adjustRoom(0x83), // level-dependent: 0x11 for game 2
		leftID:  0x01,
	}

	// Room 0x03 – Left of Name (tan, right thin wall)
	rooms[0x03] = &room{
		id:    0x03,
		walls: buildWalls(decodeRoom(rawLeftOfName, false)),
		color: tcell.ColorTan,
		flags: RoomFlagRightWall,
		upID: 0x06, rightID: 0x01,
		downID:  adjustRoom(0x86), // level-dependent: 0x0A for game 2
		leftID:  0x02,
	}

	// Room 0x04 – Top of Blue Maze (blue)
	rooms[0x04] = &room{
		id:    0x04,
		walls: buildWalls(decodeRoom(rawBlueMazeTop, false)),
		color: tcell.ColorBlue,
		flags: RoomFlagNone,
		upID: 0x10, rightID: 0x05, downID: 0x07, leftID: 0x06,
	}

	// Room 0x05 – Blue Maze #1 (blue)
	rooms[0x05] = &room{
		id:    0x05,
		walls: buildWalls(decodeRoom(rawBlueMaze1, false)),
		color: tcell.ColorBlue,
		flags: RoomFlagNone,
		upID: 0x1D, rightID: 0x06, downID: 0x08, leftID: 0x04,
	}

	// Room 0x06 – Bottom of Blue Maze (blue)
	rooms[0x06] = &room{
		id:    0x06,
		walls: buildWalls(decodeRoom(rawBlueMazeBottom, false)),
		color: tcell.ColorBlue,
		flags: RoomFlagNone,
		upID: 0x07, rightID: 0x04, downID: 0x03, leftID: 0x05,
	}

	// Room 0x07 – Center of Blue Maze (blue)
	rooms[0x07] = &room{
		id:    0x07,
		walls: buildWalls(decodeRoom(rawBlueMazeCenter, false)),
		color: tcell.ColorBlue,
		flags: RoomFlagNone,
		upID: 0x04, rightID: 0x08, downID: 0x06, leftID: 0x08,
	}

	// Room 0x08 – Blue Maze Entry (blue)
	rooms[0x08] = &room{
		id:    0x08,
		walls: buildWalls(decodeRoom(rawBlueMazeEntry, false)),
		color: tcell.ColorBlue,
		flags: RoomFlagNone,
		upID: 0x05, rightID: 0x07, downID: 0x01, leftID: 0x07,
	}

	// Room 0x09 – Maze Middle (light gray)
	rooms[0x09] = &room{
		id:    0x09,
		walls: buildWalls(decodeRoom(rawMazeMiddle, false)),
		color: tcell.ColorSilver,
		flags: RoomFlagNone,
		upID: 0x0A, rightID: 0x0A, downID: 0x0B, leftID: 0x0A,
	}

	// Room 0x0A – Maze Entry (light gray)
	rooms[0x0A] = &room{
		id:    0x0A,
		walls: buildWalls(decodeRoom(rawMazeEntry, false)),
		color: tcell.ColorSilver,
		flags: RoomFlagNone,
		upID: 0x03, rightID: 0x09, downID: 0x09, leftID: 0x09,
	}

	// Room 0x0B – Maze Side (light gray)
	rooms[0x0B] = &room{
		id:    0x0B,
		walls: buildWalls(decodeRoom(rawMazeSide, false)),
		color: tcell.ColorSilver,
		flags: RoomFlagNone,
		upID: 0x09, rightID: 0x0C, downID: 0x1C, leftID: 0x0D,
	}

	// Room 0x0C – Side Corridor (light cyan, right thin wall)
	rooms[0x0C] = &room{
		id:    0x0C,
		walls: buildWalls(decodeRoom(rawSideCorridor, false)),
		color: tcell.ColorLightCyan,
		flags: RoomFlagRightWall,
		upID: 0x1C, rightID: 0x0D, downID: 0x1D, leftID: 0x0B,
	}

	// Room 0x0D – Side Corridor (dark green, left thin wall)
	rooms[0x0D] = &room{
		id:    0x0D,
		walls: buildWalls(decodeRoom(rawSideCorridor, false)),
		color: tcell.ColorDarkGreen,
		flags: RoomFlagLeftWall,
		upID: 0x0F, rightID: 0x0B, downID: 0x0E, leftID: 0x0C,
	}

	// Room 0x0E – Top Entry Room (cyan)
	rooms[0x0E] = &room{
		id:    0x0E,
		walls: buildWalls(decodeRoom(rawTopEntryRoom, false)),
		color: tcell.ColorTeal,
		flags: RoomFlagNone,
		upID: 0x0D, rightID: 0x10, downID: 0x0F, leftID: 0x10,
	}

	// Room 0x0F – White Castle (white)
	rooms[0x0F] = &room{
		id:    0x0F,
		walls: buildWalls(decodeRoom(rawCastle, false)),
		color: tcell.ColorWhite,
		flags: RoomFlagNone,
		upID: 0x0E, rightID: 0x0F, downID: 0x0D, leftID: 0x0F,
	}

	// Room 0x10 – Black Castle (dark/black)
	rooms[0x10] = &room{
		id:    0x10,
		walls: buildWalls(decodeRoom(rawCastle, false)),
		color: tcell.ColorDimGray,
		flags: RoomFlagNone,
		upID: 0x01, rightID: 0x1C, downID: 0x04, leftID: 0x1C,
	}

	// Room 0x11 – Yellow Castle (yellow)
	rooms[0x11] = &room{
		id:    0x11,
		walls: buildWalls(decodeRoom(rawCastle, false)),
		color: tcell.ColorYellow,
		flags: RoomFlagNone,
		upID: 0x06, rightID: 0x03, downID: 0x02, leftID: 0x01,
	}

	// Room 0x12 – Yellow Castle Entry (yellow, all exits stay in this room)
	rooms[0x12] = &room{
		id:    0x12,
		walls: buildWalls(decodeRoom(rawNumberRoom, false)),
		color: tcell.ColorYellow,
		flags: RoomFlagNone,
		upID: 0x12, rightID: 0x12, downID: 0x12, leftID: 0x12,
	}

	// Room 0x13 – Black Maze #1 (light gray)
	rooms[0x13] = &room{
		id:    0x13,
		walls: buildWalls(decodeRoom(rawBlackMaze1, false)),
		color: tcell.ColorSilver,
		flags: RoomFlagNone,
		upID: 0x15, rightID: 0x14, downID: 0x15, leftID: 0x16,
	}

	// Room 0x14 – Black Maze #2 (light gray, mirrored)
	rooms[0x14] = &room{
		id:    0x14,
		walls: buildWalls(decodeRoom(rawBlackMaze2, true)),
		color: tcell.ColorSilver,
		flags: RoomFlagMirror,
		upID: 0x16, rightID: 0x15, downID: 0x16, leftID: 0x13,
	}

	// Room 0x15 – Black Maze #3 (light gray, mirrored)
	rooms[0x15] = &room{
		id:    0x15,
		walls: buildWalls(decodeRoom(rawBlackMaze3, true)),
		color: tcell.ColorSilver,
		flags: RoomFlagMirror,
		upID: 0x13, rightID: 0x16, downID: 0x13, leftID: 0x14,
	}

	// Room 0x16 – Black Maze Entry (light gray)
	rooms[0x16] = &room{
		id:    0x16,
		walls: buildWalls(decodeRoom(rawBlackMazeEntry, false)),
		color: tcell.ColorSilver,
		flags: RoomFlagNone,
		upID: 0x14, rightID: 0x13, downID: 0x1B, leftID: 0x15,
	}

	// Room 0x17 – Red Maze #1 (red)
	rooms[0x17] = &room{
		id:    0x17,
		walls: buildWalls(decodeRoom(rawRedMaze1, false)),
		color: tcell.ColorRed,
		flags: RoomFlagNone,
		upID: 0x19, rightID: 0x18, downID: 0x19, leftID: 0x18,
	}

	// Room 0x18 – Top of Red Maze (red)
	rooms[0x18] = &room{
		id:    0x18,
		walls: buildWalls(decodeRoom(rawRedMazeTop, false)),
		color: tcell.ColorRed,
		flags: RoomFlagNone,
		upID: 0x1A, rightID: 0x17, downID: 0x1A, leftID: 0x17,
	}

	// Room 0x19 – Bottom of Red Maze (red)
	rooms[0x19] = &room{
		id:    0x19,
		walls: buildWalls(decodeRoom(rawRedMazeBottom, false)),
		color: tcell.ColorRed,
		flags: RoomFlagNone,
		upID: 0x17, rightID: 0x1A, downID: 0x17, leftID: 0x1A,
	}

	// Room 0x1A – White Castle Entry (red)
	rooms[0x1A] = &room{
		id:    0x1A,
		walls: buildWalls(decodeRoom(rawWhiteCastleEntry, false)),
		color: tcell.ColorRed,
		flags: RoomFlagNone,
		upID: 0x18, rightID: 0x19, downID: 0x18, leftID: 0x19,
	}

	// Room 0x1B – Black Castle Entry (red, level-dependent exits)
	adj1B := adjustRoom(0x89)
	rooms[0x1B] = &room{
		id:    0x1B,
		walls: buildWalls(decodeRoom(rawTwoExitRoom, false)),
		color: tcell.ColorRed,
		flags: RoomFlagNone,
		upID: adj1B, rightID: adj1B, downID: adj1B, leftID: adj1B,
	}

	// Room 0x1C – Other Purple Room (purple)
	rooms[0x1C] = &room{
		id:    0x1C,
		walls: buildWalls(decodeRoom(rawNumberRoom, false)),
		color: tcell.ColorPurple,
		flags: RoomFlagNone,
		upID: 0x1D, rightID: 0x07, downID: adjustRoom(0x8C), leftID: 0x08,
	}

	// Room 0x1D – Top Entry Room (red)
	rooms[0x1D] = &room{
		id:    0x1D,
		walls: buildWalls(decodeRoom(rawTopEntryRoom, false)),
		color: tcell.ColorRed,
		flags: RoomFlagNone,
		upID: adjustRoom(0x8F), rightID: 0x01, downID: 0x10, leftID: 0x03,
	}

	// Room 0x1E – Name Room (purple) — the secret easter egg room
	rooms[0x1E] = &room{
		id:    0x1E,
		walls: buildWalls(decodeRoom(rawBelowYellowCastle, false)),
		color: tcell.ColorPurple,
		flags: RoomFlagNone,
		upID: 0x06, rightID: 0x01, downID: 0x06, leftID: 0x03,
	}
}

// roomByID returns the room pointer for the given room index.
// Returns nil if the index is out of range.
func roomByID(id int) *room {
	if id < 0 || id >= len(rooms) {
		return nil
	}
	return rooms[id]
}
