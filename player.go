package main

import (
	"github.com/gdamore/tcell"
)

// Player needed to be globally available
var player playBall

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
	player.dimensions.height = 4
	player.dimensions.width = 2
	player.pos_x = 80
	player.pos_y = 30
}

func (player *playBall) display(s tcell.Screen, r *rooms) {

	// player has same color as actual room
	player.style = r.roomStyle

	// print the player
	emitStr(s, player.pos_x, player.pos_y, player.style, "████")
	emitStr(s, player.pos_x, player.pos_y+1, player.style, "████")
}

func (player *playBall) movement(s tcell.Screen, deltaX, deltaY int) {
	w, h := s.Size()

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

	player.pos_x += deltaX
	player.pos_y += deltaY
}
