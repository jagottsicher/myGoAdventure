package main

import "github.com/gdamore/tcell"

// player needed to be globally available
var player playBall

type playBall struct {
	// room   *rooms
	contentX     [2]rune
	contentStyle [2]tcell.Style
	content_Y    rune
	pos_x        int
	pos_y        int
	width        int
	height       int
	style        tcell.Style
}

func (player *playBall) init() {
	player.style = tcell.StyleDefault.
		Background(tcell.ColorWhite).
		Foreground(tcell.ColorBlack)
	// player.room = &roomStartRoomTopEntryRoom
	player.pos_x = 80
	// we insist on even player position
	if player.pos_x&1 == 1 {
		player.pos_x = player.pos_x - 1
	}
	player.pos_y = (48 / 3) * 2
	player.width = 2
	player.height = 1
}

func (player *playBall) display(s tcell.Screen) {

	for i := 0; i < player.width; i++ {
		player.contentX[i], _, player.contentStyle[i], _ = s.GetContent(player.pos_x+i, player.pos_y)
	}

	for i := 0; i < player.width; i++ {
		s.SetContent(player.pos_x+i, player.pos_y, ' ', nil, player.style)
	}
	// 	emitStr(s, player.pos_x, player.pos_y, player.style, "██")
	// }

	// func (player *playBal, ' ', nil, player.stylel) movement(s tcell.Screen, deltaX, deltaY int) {

	// 	// turn to right
	// 	if player.pos_x+deltaX > stageWidth-player.width {
	// 		// move player to the left side
	// 		player.pos_x = 0 - player.width
	// 		// switch room
	// 		if currentRoom.right != nil {
	// 			currentRoom = currentRoom.right
	// 			display(s, currentRoom)
	// 		}
	// 	}

	// 	// turn to bottom stage
	// 	if player.pos_y+deltaY >= stageHeight {
	// 		player.pos_y -= stageHeight
	// 		if currentRoom.down != nil {
	// 			currentRoom = currentRoom.down
	// 			display(s, currentRoom)
	// 		}
	// 	}

	// 	// turn to left stage
	// 	if player.pos_x+deltaX < 0 {
	// 		if stageWidth&1 == 1 {
	// 			player.pos_x += stageWidth - 1
	// 		} else {
	// 			player.pos_x += stageWidth
	// 		}

	// 		if currentRoom.left != nil {
	// 			currentRoom = currentRoom.left
	// 			display(s, currentRoom)
	// 		}
	// 	}

	// 	// turn to upper room
	// 	if player.pos_y+deltaY < 0 {
	// 		player.pos_y += stageHeight
	// 		if currentRoom.up != nil {
	// 			currentRoom = currentRoom.up
	// 			display(s, currentRoom)
	// 		}
	// 	}

	// 	// remove the player and put what was there before
	// 	s.SetContent(player.pos_x, player.pos_y, playerLeft, nil, playerSpaceStyleLeft)
	// 	s.SetContent(player.pos_x+1, player.pos_y, playerRight, nil, playerSpaceStyleRight)

	// 	player.pos_x += deltaX
	// 	player.pos_y += deltaY

	// // remember content under player
	// playerLeft, _, playerSpaceStyleLeft, _ = s.GetContent(player.pos_x, player.pos_y)
	// playerRight, _, playerSpaceStyleRight, _ = s.GetContent(player.pos_x+1, player.pos_y)
}
