package main

import (
	"github.com/gdamore/tcell"
)

// Player needed to be globally available
var player playBall

// Player-Wall collision
var playerWallCollsion bool

// type size declared in graphics.go
var defaultPlayerSize = size{
	width:  4,
	height: 2,
}

type playBall struct {
	dimensions size
	pos_x      int
	pos_y      int
	style      tcell.Style
}

func (player *playBall) init() {
	player.dimensions.width = 4
	player.dimensions.height = 2
	player.pos_x = 80
	player.pos_y = 30
}

func (player *playBall) display(s tcell.Screen, r *rooms) {

	// player has same color as actual room
	player.style = r.roomStyle
	// player.style = tcell.StyleDefault.
	// 	Background(tcell.ColorBlack).
	// 	Foreground(tcell.ColorBlack)

	// print the player
	emitStr(s, player.pos_x, player.pos_y, player.style, "████")
	emitStr(s, player.pos_x, player.pos_y+1, player.style, "████")
}

func (player *playBall) movement(s tcell.Screen, deltaX, deltaY int) {

	// assume the way is not blocked
	playerWallCollsion = false

	// Valid for terminal Size, but we move on screens only
	// w, h := s.Size()

	// this is the stagesize
	w := 160
	h := 44

	// turn right
	if player.pos_x+deltaX >= w {
		// move player to the left side
		player.pos_x -= w
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
	}

	// turn left
	if player.pos_x+deltaX < 0 {
		player.pos_x += w
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
	// intendedDirection above = 1
	// intendedDirection above = 2
	// intendedDirection above = 3
	// intendedDirection above = 4

	switch intendedDirection {
	case 1: // check above the player
		for x := 0; x < player.dimensions.width; x++ {
			spot, _, _, _ := s.GetContent(player.pos_x+x, player.pos_y-1)
			if spot == rune('X') || spot == rune('┼') {
				return true
			}
		}
	case 2: // check right of the player
		for y := 0; y < player.dimensions.height; y++ {
			spot, _, _, _ := s.GetContent(player.pos_x+player.dimensions.width+1, player.pos_y+y)
			if spot == rune('X') || spot == rune('┼') {
				return true
			}
		}
	case 3: // check below the player
		for x := 0; x < player.dimensions.width; x++ {
			spot, _, _, _ := s.GetContent(player.pos_x+x, player.pos_y+player.dimensions.height+1)
			if spot == rune('X') || spot == rune('┼') {
				return true
			}
		}
	case 4: // check left of the player
		for y := 0; y < player.dimensions.height; y++ {
			spot, _, _, _ := s.GetContent(player.pos_x-1, player.pos_y+y)
			if spot == rune('X') || spot == rune('┼') {
				return true
			}
		}
	}
	// if all idrections free we you shall pass
	return false
}
