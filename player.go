package main

import (
	"github.com/gdamore/tcell"
)

// this is the stagesize
// var w = stageWidth
var h = 44

// player needed to be globally available
var player playBall

// player-Wall collision
var playerWallCollsion bool

// type size declared in graphics.go
var defaultPlayerSize = size{
	width:  2,
	height: 1,
}

type playBall struct {
	dimensions size
	pos_x      int
	pos_y      int
	style      tcell.Style
}

func (player *playBall) init() {
	player.dimensions.width = 2
	player.dimensions.height = 1
	player.pos_x = stageWidth / 2
	// we insist on even player position
	if player.pos_x&1 == 1 {
		player.pos_x = player.pos_x - 1
	}
	player.pos_y = (stageHeight / 3) * 2
}

func (player *playBall) display(s tcell.Screen, r *rooms) {

	// player has same color as actual room
	player.style = r.roomStyle
	// player.style = tcell.StyleDefault.
	// 	Background(tcell.ColorBlack).
	// 	Foreground(tcell.ColorBlack)

	// print the player
	// if player.pos_x&1 == 1 {
	// 	player.pos_x = player.pos_x - 1
	// }
	emitStr(s, player.pos_x, player.pos_y, player.style, "██")
	// emitStr(s, player.pos_x, player.pos_y+1, player.style, "████")
}

func (player *playBall) movement(s tcell.Screen, deltaX, deltaY int) {

	// assume the way is not blocked
	playerWallCollsion = false

	// turn right
	if player.pos_x+deltaX > stageWidth-player.dimensions.width {
		// move player to the left side
		player.pos_x = 0 - player.dimensions.width
		// switch room
		if currentRoom.right != nil {
			currentRoom = currentRoom.right
		}
	}

	// turn down
	if player.pos_y+deltaY >= h {
		player.pos_y -= h
		if currentRoom.down != nil {
			currentRoom = currentRoom.down
		}

		if currentRoom == &roomYellowCastle || currentRoom == &roomWhiteCastle || currentRoom == &roomBlackCastle {
			player.pos_y += 18
			if player.pos_x > (stageWidth / 2) {
				player.pos_x = (stageWidth / 2)
			} else {
				player.pos_x = (stageWidth / 2) - player.dimensions.width
			}

		}

	}

	// turn left
	if player.pos_x+deltaX < 0 {
		if stageWidth&1 == 1 {
			player.pos_x += stageWidth - 1
		} else {
			player.pos_x += stageWidth
		}

		if currentRoom.left != nil {
			currentRoom = currentRoom.left
		}
	}

	// turn up
	if player.pos_y+deltaY < 0 {
		player.pos_y += h
		if currentRoom.up != nil {
			currentRoom = currentRoom.up
		}
	}

	// only move if not a wall
	player.pos_x += deltaX
	player.pos_y += deltaY
}

func checkPlayerWallCollision(s tcell.Screen, intendedDirection uint8) bool {
	// check for collision

	switch intendedDirection {
	case 1: // check above the player
		for x := 0; x < player.dimensions.width; x++ {
			spot, _, _, _ := s.GetContent(player.pos_x+x, player.pos_y-1)

			if spot == rune('x') || spot == rune('┼') {
				if currentRoom == &roomYellowCastle {
					if yellowCastleGate.unlocked == false {
						currentRoom = currentRoom.up
						player.pos_y = h - player.dimensions.height
						return true
					}
				}
				if currentRoom == &roomWhiteCastle {
					if yellowCastleGate.unlocked == false {
						currentRoom = currentRoom.up
						player.pos_y = h - player.dimensions.height
						return true
					}
				}
				if currentRoom == &roomBlackCastle {
					if yellowCastleGate.unlocked == false {
						currentRoom = currentRoom.up
						player.pos_y = h - player.dimensions.height
						return true
					}
				}
			}

			if spot == rune('X') {
				return true
			}
		}
	case 2: // check right of the player
		for y := 0; y < player.dimensions.height; y++ {
			spot, _, _, _ := s.GetContent(player.pos_x+player.dimensions.width, player.pos_y+y)
			if spot == rune('X') {
				return true
			}
		}
	case 3: // check below the player
		for x := 0; x < player.dimensions.width; x++ {
			spot, _, _, _ := s.GetContent(player.pos_x+x, player.pos_y+player.dimensions.height)
			if spot == rune('X') {
				return true
			}
		}
	case 4: // check left of the player
		for y := 0; y < player.dimensions.height; y++ {
			spot, _, _, _ := s.GetContent(player.pos_x-1, player.pos_y+y)
			if spot == rune('X') {
				return true
			}
		}
	}

	// if all idrections free we you shall pass
	return false
}
