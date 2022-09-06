package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/gdamore/tcell/v2"

	"github.com/mattn/go-runewidth"
)

// type spriteSize struct {
// 	width  int16
// 	Height int16
// }

// type position struct {
// 	x int16
// 	y int16
// }

// type player struct {
// 	playerPosition position
// 	size           spriteSize
// 	color          tcell.Style
// 	// bool carryingItem;
// 	// Items * itemCarrying;
// }

type playBall struct {
	width  int
	height int
	pos_x  int
	pos_y  int
	Style  tcell.Style
}

func (player playBall) init() {
	Player.width = 4
	Player.height = 2
	Player.pos_x = 80
	Player.pos_y = 40
	Player.Style = tcell.StyleDefault.Foreground(tcell.ColorDarkBlue.TrueColor()).Background(tcell.ColorYellow)
}

func (Player playBall) display(s tcell.Screen) {
	s.Clear()
	emitStr(s, Player.pos_x, Player.pos_y, Player.Style, "PPPP")
	emitStr(s, Player.pos_x, Player.pos_y+1, Player.Style, fmt.Sprintf("%d/%d", Player.pos_x, Player.pos_y))
	s.Show()
}

// func (Player playBall) movement(s tcell.Screen, deltaXY int8) {

// }

var Player playBall

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

// func displayPlayer(s tcell.Screen) {
// 	_, h := s.Size()

// 	// if positionX > 0 {
// 	// 	positionX = (w/2 - 2) + playerX
// 	// } else {
// 	// 	positionX = w - playerX
// 	// }

// 	positionY := (h / 2) + playerY

// 	s.Clear()
// 	style := tcell.StyleDefault.Foreground(tcell.ColorDarkBlue.TrueColor()).Background(tcell.ColorYellow)

// 	emitStr(s, positionX, positionY, style, "PPPP")
// 	emitStr(s, positionX, positionY+1, style, fmt.Sprintf("%d", positionX))
// 	// emitStr(s, w/2-9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")
// 	s.Show()
//}

// clear the screen depending on your OS
func clearScreen() {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// package main

// import (
// 	// "log"

// 	"github.com/gdamore/tcell/v2"
// )

// // type world struct {
// // 	stage     tcell.Screen
// // 	leftWall  bool
// // 	rightWall bool
// // 	up        tcell.Screen
// // 	down      tcell.Screen
// // 	left      tcell.Screen
// // 	right     tcell.Screen
// // }

// var bgrStyle, YellowDrawStyle, OrangeDrawStyle, BlackDrawStyle, WhiteDrawStyle, DarkGreenDrawStyle tcell.Style

// var character [4][2]rune
// var colorStyle [4][2]tcell.Style

// var castles, bottomOpen []string
// var playerX, playerY int

// func InitPlayer() {
// 	playerX = 80
// 	playerY = 40
// }

// func getUnderPlayer(currentScreen tcell.Screen, playerX, playerY int) {
// 	character[0][0], _, colorStyle[0][0], _ = currentScreen.GetContent(playerX, playerY)
// 	character[1][0], _, colorStyle[1][0], _ = currentScreen.GetContent(playerX+1, playerY)
// 	character[2][0], _, colorStyle[2][0], _ = currentScreen.GetContent(playerX+2, playerY)
// 	character[3][0], _, colorStyle[3][0], _ = currentScreen.GetContent(playerX+3, playerY)
// 	character[0][1], _, colorStyle[0][1], _ = currentScreen.GetContent(playerX, playerY+1)
// 	character[1][1], _, colorStyle[1][1], _ = currentScreen.GetContent(playerX+1, playerY+1)
// 	character[2][1], _, colorStyle[2][1], _ = currentScreen.GetContent(playerX+2, playerY+1)
// 	character[3][1], _, colorStyle[3][1], _ = currentScreen.GetContent(playerX+3, playerY+1)
// }

// func erasePlayer(currentScreen tcell.Screen, playerX, playerY int) {
// 	currentScreen.SetContent(playerX, playerY, character[0][0], nil, colorStyle[0][0])
// 	currentScreen.SetContent(playerX+1, playerY, character[1][0], nil, colorStyle[1][0])
// 	currentScreen.SetContent(playerX+2, playerY, character[2][0], nil, colorStyle[2][0])
// 	currentScreen.SetContent(playerX+3, playerY, character[3][0], nil, colorStyle[3][0])
// 	currentScreen.SetContent(playerX, playerY+1, character[0][1], nil, colorStyle[0][1])
// 	currentScreen.SetContent(playerX+1, playerY+1, character[1][1], nil, colorStyle[1][1])
// 	currentScreen.SetContent(playerX+2, playerY+1, character[2][1], nil, colorStyle[2][1])
// 	currentScreen.SetContent(playerX+3, playerY+1, character[3][1], nil, colorStyle[3][1])
// }

// func drawPlayer(currentScreen tcell.Screen, playerX, playerY int) {
// 	currentScreen.SetContent(playerX, playerY, 'X', nil, YellowDrawStyle)
// 	currentScreen.SetContent(playerX+1, playerY, 'X', nil, YellowDrawStyle)
// 	currentScreen.SetContent(playerX+2, playerY, 'X', nil, YellowDrawStyle)
// 	currentScreen.SetContent(playerX+3, playerY, 'X', nil, YellowDrawStyle)
// 	currentScreen.SetContent(playerX, playerY+1, 'X', nil, YellowDrawStyle)
// 	currentScreen.SetContent(playerX+1, playerY+1, 'X', nil, YellowDrawStyle)
// 	currentScreen.SetContent(playerX+2, playerY+1, 'X', nil, YellowDrawStyle)
// 	currentScreen.SetContent(playerX+3, playerY+1, 'X', nil, YellowDrawStyle)
// }

// func InitColors() {
// 	// define all colors as drawing color
// 	// defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
// 	bgrStyle = tcell.StyleDefault.Foreground(tcell.ColorLightGray).Background(tcell.ColorLightGray)

// 	OrangeDrawStyle = tcell.StyleDefault.Foreground(tcell.ColorOrange).Background(tcell.ColorOrange)
// 	YellowDrawStyle = tcell.StyleDefault.Foreground(tcell.Color226).Background(tcell.Color226)
// 	BlackDrawStyle = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)
// 	// RedDrawStyle = tcell.StyleDefault.Foreground(tcell.ColorIndianRed).Background(tcell.ColorIndianRed)
// 	WhiteDrawStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorWhite)
// 	DarkGreenDrawStyle = tcell.StyleDefault.Foreground(tcell.ColorDarkGreen).Background(tcell.ColorDarkGreen)
// 	// GreenDrawStyle = tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorGreen)
// 	// DarkOliveGreenDrawStyle = tcell.StyleDefault.Foreground(tcell.ColorDarkOliveGreen).Background(tcell.ColorDarkOliveGreen)
// 	// PurpleDrawStyle = tcell.StyleDefault.Foreground(tcell.ColorMediumPurple).Background(tcell.ColorMediumPurple)
// 	// LightBlueDrawStyle = tcell.StyleDefault.Foreground(tcell.ColorDeepSkyBlue).Background(tcell.ColorDeepSkyBlue)
// 	// NeonYellowDrawStyle = tcell.StyleDefault.Foreground(tcell.Color227).Background(tcell.Color227)
// 	// BlueDrawStyle = tcell.StyleDefault.Foreground(tcell.Color21).Background(tcell.Color21)
// 	return
// }

// func switchScreen(screenName tcell.Screen, content []string, drawingColor tcell.Style) {
// 	screenName.SetStyle(bgrStyle)
// 	// s.EnableMouse()
// 	// s.EnablePaste()
// 	screenName.Clear()

// 	var tempString string
// 	// var tempContent = *content

// 	for i, _ := range content {
// 		tempString = content[i]
// 		characters := []rune(tempString)
// 		for j := 0; j < len(tempString); j++ {
// 			if string(characters[j]) == "X" {
// 				col := j * 2
// 				screenName.SetContent(col, i, 'X', nil, drawingColor)
// 				screenName.SetContent(col+1, i, 'X', nil, drawingColor)
// 			}
// 		}
// 	}
// }

// func InitContent() {
// 	// all different kinds of stages
// 	castles = []string{
// 		"XXXXXXXXXXXXXXXX   XXX   XXX   XXX            XXX   XXX   XXX   XXXXXXXXXXXXXXXX",
// 		"XXXXXXXXXXXXXXXX   XXX   XXX   XXX            XXX   XXX   XXX   XXXXXXXXXXXXXXXX",
// 		"XX           XXX   XXX   XXX   XXX            XXX   XXX   XXX   XXX           XX",
// 		"XX           XXX   XXX   XXX   XXX            XXX   XXX   XXX   XXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
// 		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
// 		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// 		"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// 	}

// 	bottomOpen = []string{
// 		"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// 		"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XX                                                                            XX",
// 		"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// 		"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// 	}

// }

// // bottomLeftRightOpen := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // topOpen := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // topLeftRightOpen := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // bottomTopOpen := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XX                                                                            XX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // bottomTopLeftRightOpen := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // orangeTop := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"         XXX                                                        XXX         ",
// // 	"         XXX                                                        XXX         ",
// // 	"         XXX                                                        XXX         ",
// // 	"         XXX                                                        XXX         ",
// // 	"         XXX                                                        XXX         ",
// // 	"         XXX                                                        XXX         ",
// // 	"         XXX                                                        XXX         ",
// // 	"         XXX                                                        XXX         ",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"XXXXX    XXX          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXX    XXXXX",
// // 	"         XXX                        XXXXXXXX                        XXX         ",
// // 	"         XXX                        XXXXXXXX                        XXX         ",
// // 	"         XXX                        XXXXXXXX                        XXX         ",
// // 	"         XXX                        XXXXXXXX                        XXX         ",
// // 	"         XXX                        XXXXXXXX                        XXX         ",
// // 	"         XXX                        XXXXXXXX                        XXX         ",
// // 	"         XXX                        XXXXXXXX                        XXX         ",
// // 	"         XXX                        XXXXXXXX                        XXX         ",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX                XXXXXXXX                XXXX    XXXXXXXXXXXX",
// // 	"                XXXX                XXXXXXXX                XXXX                ",
// // 	"                XXXX                XXXXXXXX                XXXX                ",
// // 	"                XXXX                XXXXXXXX                XXXX                ",
// // 	"                XXXX                XXXXXXXX                XXXX                ",
// // 	"                XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX                ",
// // 	"                XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX                ",
// // 	"                XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX                ",
// // 	"                XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX                ",
// // 	"XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX",
// // }

// // orangeMiddle := []string{
// // 	"XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX",
// // 	"                        XXXXXXXX    XXXXXXXX    XXXXXXXX                        ",
// // 	"                        XXXXXXXX    XXXXXXXX    XXXXXXXX                        ",
// // 	"                        XXXXXXXX    XXXXXXXX    XXXXXXXX                        ",
// // 	"                        XXXXXXXX    XXXXXXXX    XXXXXXXX                        ",
// // 	"                        XXXXXXXX    XXXXXXXX    XXXXXXXX                        ",
// // 	"                        XXXXXXXX    XXXXXXXX    XXXXXXXX                        ",
// // 	"                        XXXXXXXX    XXXXXXXX    XXXXXXXX                        ",
// // 	"                        XXXXXXXX    XXXXXXXX    XXXXXXXX                        ",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"XXX             XXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXX             XXX",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"        XXXX    XXXX    XXXXXXXX                XXXXXXXX    XXXX    XXXX        ",
// // 	"        XXXX    XXXX    XXXXXXXX                XXXXXXXX    XXXX    XXXX        ",
// // 	"        XXXX    XXXX    XXXXXXXX                XXXXXXXX    XXXX    XXXX        ",
// // 	"        XXXX    XXXX    XXXXXXXX                XXXXXXXX    XXXX    XXXX        ",
// // 	"        XXXX    XXXX    XXXXXXXX                XXXXXXXX    XXXX    XXXX        ",
// // 	"        XXXX    XXXX    XXXXXXXX                XXXXXXXX    XXXX    XXXX        ",
// // 	"        XXXX    XXXX    XXXXXXXX                XXXXXXXX    XXXX    XXXX        ",
// // 	"        XXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXX        ",
// // 	"XXXX    XXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXX    XXXX",
// // 	"XXXX    XXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXX    XXXX",
// // }

// // orangeBottom := []string{
// // 	"XXXX    XXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXX    XXXX",
// // 	"XXXX    XXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXX    XXXX",
// // 	"        XXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXX        ",
// // 	"        XXXX            XXXXXXXX    XXXXXXXX    XXXXXXXX            XXXX        ",
// // 	"        XXXX            XXXXXXXX    XXXXXXXX    XXXXXXXX            XXXX        ",
// // 	"        XXXX            XXXXXXXX    XXXXXXXX    XXXXXXXX            XXXX        ",
// // 	"        XXXX            XXXXXXXX    XXXXXXXX    XXXXXXXX            XXXX        ",
// // 	"        XXXX            XXXXXXXX    XXXXXXXX    XXXXXXXX            XXXX        ",
// // 	"        XXXX            XXXXXXXX    XXXXXXXX    XXXXXXXX            XXXX        ",
// // 	"        XXXX            XXXXXXXX    XXXXXXXX    XXXXXXXX            XXXX        ",
// // 	"        XXXX            XXXXXXXX    XXXXXXXX    XXXXXXXX            XXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXX        ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"        XXXXXXXXXXXXXXXXXXX         XXXXXXXX         XXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXX         XXXXXXXX         XXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXX         XXXXXXXX         XXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXX         XXXXXXXX         XXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXX         XXXXXXXX         XXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXX         XXXXXXXX         XXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXX         XXXXXXXX         XXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXXXXXXXXXXXXXXXXX         XXXXXXXX         XXXXXXXXXXXXXXXXXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"        XXXX                        XXXXXXXX                        XXXX        ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// //	}

// // redRightTop := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"                                    XXXXXXXX                                    ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                        XXXX        XXXXXXXX        XXXX                        ",
// // 	"                        XXXX        XXXXXXXX        XXXX                        ",
// // 	"                        XXXX        XXXXXXXX        XXXX                        ",
// // 	"                        XXXX        XXXXXXXX        XXXX                        ",
// // 	"                        XXXX        XXXXXXXX        XXXX                        ",
// // 	"                        XXXX        XXXXXXXX        XXXX                        ",
// // 	"                        XXXX        XXXXXXXX        XXXX                        ",
// // 	"                        XXXX        XXXXXXXX        XXXX                        ",
// // 	"XXXXXX     XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX     XXXXXX",
// // 	"XXXXXX     XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX     XXXXXX",
// // 	"XXXXXX     XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX     XXXXXX",
// // 	"XXXXXX     XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XX    XX                                          XX    XX     XXXXXX",
// // 	"XXXXXX     XXXXXXXX     XXXX                        XXXX     XXXXXXXX     XXXXXX",
// // 	"XXXXXX     XXXXXXXX     XXXX                        XXXX     XXXXXXXX     XXXXXX",
// // 	"XXXXXX     XXXXXXXX     XXXX                        XXXX     XXXXXXXX     XXXXXX",
// // 	"XXXXXX     XXXXXXXX     XXXX                        XXXX     XXXXXXXX     XXXXXX",
// // 	"XXXXXX     XXXXXXXX     XXXX                        XXXX     XXXXXXXX     XXXXXX",
// // }

// // redRightBottom := []string {
// // 	"XXXXXX     XXXXXXXX     XXXX                        XXXX     XXXXXXXX     XXXXXX",
// // 	"XXXXXX     XXXXXXXX     XXXX                        XXXX     XXXXXXXX     XXXXXX",
// // 	"XXXXXX     XXXXXXXX     XXXX                        XXXX     XXXXXXXX     XXXXXX",
// // 	"XXXXXX     XXXXXXXX     XXXX                        XXXX     XXXXXXXX     XXXXXX",
// // 	"XXXXXX                  XXXX                        XXXX                  XXXXXX",
// // 	"XXXXXX                  XXXX                        XXXX                  XXXXXX",
// // 	"XXXXXX                  XXXX                        XXXX                  XXXXXX",
// // 	"XXXXXX                  XXXX                        XXXX                  XXXXXX",
// // 	"XXXXXX                  XXXX                        XXXX                  XXXXXX",
// // 	"XXXXXX                  XXXX                        XXXX                  XXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"XXXXXXXX        XXXX                                        XXXX        XXXXXXXX",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"                XXXX                                        XXXX                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // redLeftTop:= []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                              XX                XX                              ",
// // 	"                              XX                XX                              ",
// // 	"                              XX                XX                              ",
// // 	"                              XX                XX                              ",
// // 	"                              XX                XX                              ",
// // 	"                              XX                XX                              ",
// // 	"                              XX                XX                              ",
// // 	"                              XX                XX                              ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
// // }

// // redLeftBottom := []string {
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX                                                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX                                                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX                                                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX                                                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX                                                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX                                                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX                                                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX                                                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX                                                XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"XXXXXXXXXXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXXXXXXXXXX",
// // 	"                XXXX    XXXX                        XXXX    XXXX                ",
// // 	"                XXXX    XXXX                        XXXX    XXXX                ",
// // 	"                XXXX    XXXX                        XXXX    XXXX                ",
// // 	"                XXXX    XXXX                        XXXX    XXXX                ",
// // 	"                XXXX    XXXX                        XXXX    XXXX                ",
// // 	"                XXXX    XXXX                        XXXX    XXXX                ",
// // 	"                XXXX    XXXX                        XXXX    XXXX                ",
// // 	"                XXXX    XXXX                        XXXX    XXXX                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // blueTop := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"             XXX        XXXX                        XXXX        XXX             ",
// // 	"             XXX        XXXX                        XXXX        XXX             ",
// // 	"             XXX        XXXX                        XXXX        XXX             ",
// // 	"             XXX        XXXX                        XXXX        XXX             ",
// // 	"             XXX        XXXX                        XXXX        XXX             ",
// // 	"             XXX        XXXX                        XXXX        XXX             ",
// // 	"             XXX        XXXX                        XXXX        XXX             ",
// // 	"             XXX        XXXX                        XXXX        XXX             ",
// // 	"XXXXX        XXX        XXXXXXXX                XXXXXXXX        XXX        XXXXX",
// // 	"XXXXX        XXX        XXXXXXXX                XXXXXXXX        XXX        XXXXX",
// // 	"XXXXX        XXX        XXXXXXXX                XXXXXXXX        XXX        XXXXX",
// // 	"XXXXX        XXX        XXXXXXXX                XXXXXXXX        XXX        XXXXX",
// // 	"XXXXX        XXX        XXXXXXXX                XXXXXXXX        XXX        XXXXX",
// // 	"XXXXX        XXX        XXXXXXXX                XXXXXXXX        XXX        XXXXX",
// // 	"XXXXX        XXX        XXXXXXXX                XXXXXXXX        XXX        XXXXX",
// // 	"XXXXX        XXX        XXXXXXXX                XXXXXXXX        XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXX        XXX                                                XXX        XXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"              XXXX            XX                XX            XXXX              ",
// // 	"              XXXX            XX                XX            XXXX              ",
// // 	"              XXXX            XX                XX            XXXX              ",
// // 	"              XXXX            XX                XX            XXXX              ",
// // 	"              XXXX            XX                XX            XXXX              ",
// // 	"              XXXX            XX                XX            XXXX              ",
// // 	"              XXXX            XX                XX            XXXX              ",
// // 	"              XXXX            XX                XX            XXXX              ",
// // 	"XXXXXXXX      XXXX      XXXXXXXX                XXXXXXXX      XXXX      XXXXXXXX",
// // 	"XXXXXXXX      XXXX      XXXXXXXX                XXXXXXXX      XXXX      XXXXXXXX",
// // }

// // blueMiddle := []string{
// // 	"XXXXXXXX      XXXX      XXXXXXXX                XXXXXXXX      XXXX      XXXXXXXX",
// // 	"XXXXXXXX      XXXX      XXXXXXXX                XXXXXXXX      XXXX      XXXXXXXX",
// // 	"              XXXX        XXXXXX                XXXXXX        XXXX              ",
// // 	"              XXXX        XXXXXX                XXXXXX        XXXX              ",
// // 	"              XXXX        XXXXXX                XXXXXX        XXXX              ",
// // 	"              XXXX        XXXXXX                XXXXXX        XXXX              ",
// // 	"              XXXX        XXXXXX                XXXXXX        XXXX              ",
// // 	"              XXXX        XXXXXX                XXXXXX        XXXX              ",
// // 	"              XXXX        XXXXXX                XXXXXX        XXXX              ",
// // 	"              XXXX        XXXXXX                XXXXXX        XXXX              ",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXX    XXXXXX                XXXXXX    XXXXXXXXXXXXXXXXXXXXXX",
// // 	"                  XXXX    XXXXXX                XXXXXX    XXXX                  ",
// // 	"                  XXXX    XXXXXX                XXXXXX    XXXX                  ",
// // 	"                  XXXX    XXXXXX                XXXXXX    XXXX                  ",
// // 	"                  XXXX    XXXXXX                XXXXXX    XXXX                  ",
// // 	"                  XXXX    XXXXXX                XXXXXX    XXXX                  ",
// // 	"                  XXXX    XXXXXX                XXXXXX    XXXX                  ",
// // 	"                  XXXX    XXXXXX                XXXXXX    XXXX                  ",
// // 	"                  XXXX    XXXXXX                XXXXXX    XXXX                  ",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"XXXXXX    XXXX    XXXX    XXXXXX                XXXXXX    XXXX    XXXX    XXXXXX",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"XXXXXXXXXXXXXX    XXXX    XX                        XX    XXXX    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX    XXXX    XX                        XX    XXXX    XXXXXXXXXXXXXX",
// // }

// // blueBottom := []string{
// // 	"XXXXXXXXXXXXXX    XXXX    XX                        XX    XXXX    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX    XXXX    XX                        XX    XXXX    XXXXXXXXXXXXXX",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"          XXXX    XXXX    XX                        XX    XXXX    XXXX          ",
// // 	"          XXXX            XX                        XX            XXXX          ",
// // 	"          XXXX            XX                        XX            XXXX          ",
// // 	"          XXXX            XX                        XX            XXXX          ",
// // 	"          XXXX            XX                        XX            XXXX          ",
// // 	"          XXXX            XX                        XX            XXXX          ",
// // 	"          XXXX            XX                        XX            XXXX          ",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX      XXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXX      XXXX",
// // 	"XXXX                                                                        XXXX",
// // 	"XXXX                                                                        XXXX",
// // 	"XXXX                                                                        XXXX",
// // 	"XXXX                                                                        XXXX",
// // 	"XXXX                                                                        XXXX",
// // 	"XXXX                                                                        XXXX",
// // 	"XXXX                                                                        XXXX",
// // 	"XXXX                                                                        XXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXX                                                    XXXXXXXXXXXXXX",
// // 	"          XXXX                                                    XXXX          ",
// // 	"          XXXX                                                    XXXX          ",
// // 	"          XXXX                                                    XXXX          ",
// // 	"          XXXX                                                    XXXX          ",
// // 	"          XXXX                                                    XXXX          ",
// // 	"          XXXX                                                    XXXX          ",
// // 	"          XXXX                                                    XXXX          ",
// // 	"          XXXX                                                    XXXX          ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // blueBottomLeft := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX                            XXXXXXXX                            XXXXXXXX",
// // 	"XXXXXXXX                            XXXXXXXX                            XXXXXXXX",
// // 	"XXXXXXXX                            XXXXXXXX                            XXXXXXXX",
// // 	"XXXXXXXX                            XXXXXXXX                            XXXXXXXX",
// // 	"XXXXXXXX                            XXXXXXXX                            XXXXXXXX",
// // 	"XXXXXXXX                            XXXXXXXX                            XXXXXXXX",
// // 	"XXXXXXXX                            XXXXXXXX                            XXXXXXXX",
// // 	"XXXXXXXX                            XXXXXXXX                            XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX    XXXXXXXXXXXXXXXXXXXX    XXXXXXXX",
// // 	"            XXXX            XXXX    XXXXXXXX    XXXX            XXXX            ",
// // 	"            XXXX            XXXX    XXXXXXXX    XXXX            XXXX            ",
// // 	"            XXXX            XXXX    XXXXXXXX    XXXX            XXXX            ",
// // 	"            XXXX            XXXX    XXXXXXXX    XXXX            XXXX            ",
// // 	"            XXXX            XXXX    XXXXXXXX    XXXX            XXXX            ",
// // 	"            XXXX            XXXX    XXXXXXXX    XXXX            XXXX            ",
// // 	"            XXXX            XXXX    XXXXXXXX    XXXX            XXXX            ",
// // 	"            XXXX            XXXX    XXXXXXXX    XXXX            XXXX            ",
// // 	"XXXXXXXXXXXXXXXX    XXXX    XXXX    XXXXXXXX    XXXX    XXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXX    XXXX    XXXXXXXX    XXXX    XXXX    XXXXXXXXXXXXXXXX",
// // }

// // blueMiddleRight := []string{
// // 	"XXXXXXXXXXXXXXXX    XXXX    XXXX    XXXXXXXX    XXXX    XXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXX    XXXX    XXXXXXXX    XXXX    XXXX    XXXXXXXXXXXXXXXX",
// // 	"            XXXX    XXXX    XXXX                XXXX    XXXX    XXXX            ",
// // 	"            XXXX    XXXX    XXXX                XXXX    XXXX    XXXX            ",
// // 	"            XXXX    XXXX    XXXX                XXXX    XXXX    XXXX            ",
// // 	"            XXXX    XXXX    XXXX                XXXX    XXXX    XXXX            ",
// // 	"            XXXX    XXXX    XXXX                XXXX    XXXX    XXXX            ",
// // 	"            XXXX    XXXX    XXXX                XXXX    XXXX    XXXX            ",
// // 	"            XXXX    XXXX    XXXX                XXXX    XXXX    XXXX            ",
// // 	"            XXXX    XXXX    XXXX                XXXX    XXXX    XXXX            ",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"XXXXXX      XXXX    XXXX    XXXXXXXXXXXXXXXXXXXXXXXX    XXXX    XXXX      XXXXXX",
// // 	"            XXXX    XXXX                                XXXX    XXXX            ",
// // 	"            XXXX    XXXX                                XXXX    XXXX            ",
// // 	"            XXXX    XXXX                                XXXX    XXXX            ",
// // 	"            XXXX    XXXX                                XXXX    XXXX            ",
// // 	"            XXXX    XXXX                                XXXX    XXXX            ",
// // 	"            XXXX    XXXX                                XXXX    XXXX            ",
// // 	"            XXXX    XXXX                                XXXX    XXXX            ",
// // 	"            XXXX    XXXX                                XXXX    XXXX            ",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // orangeRightBottom := []string{
// // 	"XXXX    XXXX    XXXX    XXXX        XXXXXXXX        XXXX    XXXX    XXXX    XXXX",
// // 	"XXXX    XXXX    XXXX    XXXX        XXXXXXXX        XXXX    XXXX    XXXX    XXXX",
// // 	"        XXXX            XXXX        XXXXXXXX        XXXX            XXXX        ",
// // 	"        XXXX            XXXX        XXXXXXXX        XXXX            XXXX        ",
// // 	"        XXXX            XXXX        XXXXXXXX        XXXX            XXXX        ",
// // 	"        XXXX            XXXX        XXXXXXXX        XXXX            XXXX        ",
// // 	"        XXXX            XXXX        XXXXXXXX        XXXX            XXXX        ",
// // 	"        XXXX            XXXX        XXXXXXXX        XXXX            XXXX        ",
// // 	"        XXXX            XXXX        XXXXXXXX        XXXX            XXXX        ",
// // 	"        XXXX            XXXX        XXXXXXXX        XXXX            XXXX        ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // }

// // orangeRightTop := []string{
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                    XXXX                                    XXXX",
// // 	"                                    XXXX                                    XXXX",
// // 	"                                    XXXX                                    XXXX",
// // 	"                                    XXXX                                    XXXX",
// // 	"                                    XXXX                                    XXXX",
// // 	"                                    XXXX                                    XXXX",
// // 	"                                    XXXX                                    XXXX",
// // 	"                                    XXXX                                    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXX",
// // 	"                            XXXX                                    XXXX        ",
// // 	"                            XXXX                                    XXXX        ",
// // 	"                            XXXX                                    XXXX        ",
// // 	"                            XXXX                                    XXXX        ",
// // 	"                            XXXX                                    XXXX        ",
// // 	"                            XXXX                                    XXXX        ",
// // 	"                            XXXX                                    XXXX        ",
// // 	"                            XXXX                                    XXXX        ",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          XXXXXXXXXXXXXXXXXXXX",
// // 	"                XXXX                XXXX                    XXXX            XXXX",
// // 	"                XXXX                XXXX                    XXXX            XXXX",
// // 	"                XXXX                XXXX                    XXXX            XXXX",
// // 	"                XXXX                XXXX                    XXXX            XXXX",
// // 	"                XXXX                XXXX                    XXXX            XXXX",
// // 	"                XXXX                XXXX                    XXXX            XXXX",
// // 	"                XXXX                XXXX                    XXXX            XXXX",
// // 	"                XXXX                XXXX                    XXXX            XXXX",
// // 	"XXXX    XXXX    XXXX    XXXX        XXXXXXXX        XXXX    XXXX    XXXX    XXXX",
// // 	"XXXX    XXXX    XXXX    XXXX        XXXXXXXX        XXXX    XXXX    XXXX    XXXX",
// // }

// // orangeLeftTop := []string{
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"                    XXXX                                XXXX                    ",
// // 	"                    XXXX                                XXXX                    ",
// // 	"                    XXXX                                XXXX                    ",
// // 	"                    XXXX                                XXXX                    ",
// // 	"                    XXXX                                XXXX                    ",
// // 	"                    XXXX                                XXXX                    ",
// // 	"                    XXXX                                XXXX                    ",
// // 	"                    XXXX                                XXXX                    ",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXXXXXXXXXX                                XXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"                                                                                ",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXX",
// // 	"            XXXX                                                XXXX            ",
// // 	"            XXXX                                                XXXX            ",
// // 	"            XXXX                                                XXXX            ",
// // 	"            XXXX                                                XXXX            ",
// // 	"            XXXX                                                XXXX            ",
// // 	"            XXXX                                                XXXX            ",
// // 	"            XXXX                                                XXXX            ",
// // 	"            XXXX                                                XXXX            ",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // }

// // orangeLeftBottom := []string{
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXX                                    XXXX                                    ",
// // 	"XXXX                                    XXXX                                    ",
// // 	"XXXX                                    XXXX                                    ",
// // 	"XXXX                                    XXXX                                    ",
// // 	"XXXX                                    XXXX                                    ",
// // 	"XXXX                                    XXXX                                    ",
// // 	"XXXX                                    XXXX                                    ",
// // 	"XXXX                                    XXXX                                    ",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"XXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
// // 	"            XXXX                                    XXXX                        ",
// // 	"            XXXX                                    XXXX                        ",
// // 	"            XXXX                                    XXXX                        ",
// // 	"            XXXX                                    XXXX                        ",
// // 	"            XXXX                                    XXXX                        ",
// // 	"            XXXX                                    XXXX                        ",
// // 	"            XXXX                                    XXXX                        ",
// // 	"            XXXX                                    XXXX                        ",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
// // 	"XXXX                    XXXX            XXXX                    XXXX            ",
// // 	"XXXX                    XXXX            XXXX                    XXXX            ",
// // 	"XXXX                    XXXX            XXXX                    XXXX            ",
// // 	"XXXX                    XXXX            XXXX                    XXXX            ",
// // 	"XXXX                    XXXX            XXXX                    XXXX            ",
// // 	"XXXX                    XXXX            XXXX                    XXXX            ",
// // 	"XXXX                    XXXX            XXXX                    XXXX            ",
// // 	"XXXX                    XXXX            XXXX                    XXXX            ",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // 	"XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXX",
// // }
