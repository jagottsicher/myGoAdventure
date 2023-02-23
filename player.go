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

	// check for collision
	// above the player (includes left and right corner)
	var wallRune rune
	y := player.pos_y - 1
	checkXWidth := player.pos_x + player.dimensions.width + 1
	for x := player.pos_x - 1; x < checkXWidth; x++ {
		spot, _, _, _ := s.GetContent(x, y)
		s.SetContent(10, 5, spot, nil, player.style)
		s.SetContent(10, 6, wallRune, nil, player.style)

		// s.Sync()
		if spot == wallRune {
			return
		}
	}

	// x := 0
	// // left and right of the player (only the sides)
	// for y = 0; y < player.dimensions.height; y++ {
	// 	x = player.pos_x - 1
	// 	spot, _, _, _ := s.GetContent(x, y)
	// 	if spot == 'X' {
	// 		return
	// 	}
	// 	x = player.pos_x + player.dimensions.width + 1
	// 	spot, _, _, _ = s.GetContent(x, y)
	// 	if spot == 'X' {
	// 		return
	// 	}
	// }

	// //below the player (includes left and right corner)
	// y = player.dimensions.height + 1
	// for x := player.pos_x - 1; x < (x + player.dimensions.width + 1); x++ {
	// 	spot, _, _, _ := s.GetContent(x, y)
	// 	if spot == 'X' {
	// 		return
	// 	}
	// }

	// only move if not a wall
	player.pos_x += deltaX
	player.pos_y += deltaY
}
