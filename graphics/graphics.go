// TO-DO
// own struct for sizes
// add all screens
// add and initialize all directions

package graphics

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/gdamore/tcell/v2"

	"github.com/mattn/go-runewidth"
)

type CompressedRoom []string

type RoomID int

type Rooms struct {
	Width               int
	Height              int
	defaultXFactor      int
	defaultYFactor      int
	defaultYInnerFactor int
	compressedRoomData  CompressedRoom
	Up                  *Rooms
	Down                *Rooms
	Left                *Rooms
	Right               *Rooms
}

var YellowCastle = Rooms{
	Width:               160,
	Height:              44,
	defaultXFactor:      4,
	defaultYFactor:      2,
	defaultYInnerFactor: 4,
	compressedRoomData:  roomGfxCastle,
	Up:                  nil,
	Down:                &TopEntryRoom,
	Left:                nil,
	Right:               nil,
}

var TopEntryRoom = Rooms{
	Width:               160,
	Height:              44,
	defaultXFactor:      4,
	defaultYFactor:      2,
	defaultYInnerFactor: 4,
	compressedRoomData:  roomGfxTopEntryRoom,
	Up:                  nil,
	Down:                nil,
	Left:                nil,
	Right:               nil,
}

func InitDirections(r *Rooms) {
	TopEntryRoom.Up = &YellowCastle
}

func Display(s tcell.Screen, r *Rooms) {

	// line by line
	// top and last lines repeat twice, the inner lines 8 times
	for lines, content := range r.compressedRoomData[0:1] {
		for yScale := 0; yScale < r.defaultYFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.defaultXFactor; xScale++ {
					emitStr(s, ((r.defaultXFactor * columns) + xScale), (lines + yScale), Player.Style, string(char))
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
					emitStr(s, ((r.defaultXFactor * columns) + xScale), ((r.defaultYInnerFactor * (lines)) + r.defaultYFactor + yScale), Player.Style, string(char))
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
					emitStr(s, ((r.defaultXFactor * columns) + xScale), ((r.Height) - r.defaultYFactor + yScale), Player.Style, string(char))
					// r.defaultYFactor + ((r.defaultYInnerFactor*r.defaultYFactor)*(r.Height-(2*r.defaultYFactor)) + yScale)), Player.Style, string(char))
				}
			}
		}
	}
}

type PlayBall struct {
	width  int
	height int
	pos_x  int
	pos_y  int
	Style  tcell.Style
}

func (player *PlayBall) Init() {
	Player.width = 4
	Player.height = 2
	Player.pos_x = 80
	Player.pos_y = 40
	Player.Style = tcell.StyleDefault.Foreground(tcell.ColorDarkBlue.TrueColor()).Background(tcell.ColorYellow)
}

func (Player PlayBall) Display(s tcell.Screen) {
	// s.Clear()
	emitStr(s, Player.pos_x, Player.pos_y, Player.Style, "PPPP")
	emitStr(s, Player.pos_x, Player.pos_y+1, Player.Style, "PPPP")
	emitStr(s, 80, 24, Player.Style, fmt.Sprintf("%d/%d", Player.pos_x, Player.pos_y))
	s.Show()
}

func (Player *PlayBall) Movement(s tcell.Screen, deltaX, deltaY int) {
	w, h := s.Size()
	if Player.pos_x+deltaX >= w {
		Player.pos_x -= w
	}
	if Player.pos_y+deltaY >= h {
		Player.pos_y -= h
	}
	if Player.pos_x+deltaX < 0 {
		Player.pos_x += w
	}
	if Player.pos_y+deltaY < 0 {
		Player.pos_y += h
	}
	Player.pos_x += deltaX
	Player.pos_y += deltaY
}

var Player PlayBall

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
func ClearScreen() {
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

//
// Room graphics
//

// Left of Name Room
var roomGfxLeftOfName = CompressedRoom{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Below Yellow Castle
var roomGfxBelowYellowCastle = CompressedRoom{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Side CoXXidor
var roomGfxSideCoXXidor = CompressedRoom{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Number Room Definition
var roomGfxNumberRoom = CompressedRoom{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// `
var roomGfxTwoExitRoom = CompressedRoom{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Top of Blue Maze
var roomGfxBlueMazeTop = CompressedRoom{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"        XX    XX        XX    XX        ",
	"        XX    XX        XX    XX        ",
	"XXXX    XX    XXXX    XXXX    XX    XXXX",
	"XXXX    XX    XXXX    XXXX    XX    XXXX",
	"XXXX    XX                    XX    XXXX",
	"XXXX    XX                    XX    XXXX",
	"XXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXX",
	"      XX        XX    XX        XX      ",
	"      XX        XX    XX        XX      ",
	"XXXX  XX  XXXXXXXX    XXXXXXXX  XX  XXXX",
}

// Blue Maze #1
var roomGfxBlueMaze1 = CompressedRoom{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXX  XXXXXXXXXXXXXXXX  XXXXXXXXXX",
	"XXXXXXXXXX  XXXXXXXXXXXXXXXX  XXXXXXXXXX",
	"XXXX              XXXX              XXXX",
	"XXXX              XXXX              XXXX",
	"XXXX  XXXXXXXXXX  XXXX  XXXXXXXXXX  XXXX",
	"XXXX  XXXXXXXXXX  XXXX  XXXXXXXXXX  XXXX",
	"      XX      XX  XXXX  XX      XX      ",
	"      XX      XX  XXXX  XX      XX      ",
	"XXXXXXXX  XX  XX  XXXX  XX  XX  XXXXXXXX",
}

// Bottom of Blue Maze
var roomGfxBlueMazeBottom = CompressedRoom{
	"XXXXXXXX  XX  XX        XX  XX  XXXXXXXX",
	"      XX      XX        XX      XX      ",
	"      XX      XX        XX      XX      ",
	"XXXX  XXXXXXXXXX        XXXXXXXXXX  XXXX",
	"XXXX  XXXXXXXXXX        XXXXXXXXXX  XXXX",
	"XXXX                                XXXX",
	"XXXX                                XXXX",
	"XXXXXXXX                        XXXXXXXX",
	"XXXXXXXX                        XXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Center of Blue Maze
var roomGfxBlueMazeCenter = CompressedRoom{
	"XXXX  XX  XXXXXXXX    XXXXXXXX  XX  XXXX",
	"      XX      XXXX    XXXX      XX      ",
	"      XX      XXXX    XXXX      XX      ",
	"XXXXXXXXXXXX  XXXX    XXXX  XXXXXXXXXXXX",
	"XXXXXXXXXXXX  XXXX    XXXX  XXXXXXXXXXXX",
	"          XX  XXXX    XXXX  XX          ",
	"          XX  XXXX    XXXX  XX          ",
	"XXXX  XX  XX  XXXX    XXXX  XX  XX  XXXX",
	"XXXX  XX  XX  XXXX    XXXX  XX  XX  XXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXXXXXX  XX  XX        XX  XX  XXXXXXXX",
}

// Blue Maze Entry
var roomGfxBlueMazeEntry = CompressedRoom{
	"XXXXXXXX  XX  XX  XXXX  XX  XX  XXXXXXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXX  XX  XX  XXXXXXXXXXXX  XX  XX  XXXX",
	"XXXX  XX  XX  XXXXXXXXXXXX  XX  XX  XXXX",
	"      XX  XX                XX  XX      ",
	"      XX  XX                XX  XX      ",
	"XXXXXXXX  XXXXXXXXXXXXXXXXXXXX  XXXXXXXX",
	"XXXXXXXX  XXXXXXXXXXXXXXXXXXXX  XXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Maze Middle
var roomGfxMazeMiddle = CompressedRoom{
	"XXXXXXXXXXXX  XX  XXXX  XX  XXXXXXXXXXXX",
	"              XX  XXXX  XX              ",
	"              XX  XXXX  XX              ",
	"XXXX      XXXXXX  XXXX  XXXXXX      XXXX",
	"XXXX      XXXXXX  XXXX  XXXXXX      XXXX",
	"          XX                XX          ",
	"          XX                XX          ",
	"XXXXXXXX  XX  XXXXXXXXXXXX  XX  XXXXXXXX",
	"XXXXXXXX  XX  XXXXXXXXXXXX  XX  XXXXXXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXX  XX  XX  XX  XXXX  XX  XX  XX  XXXX",
}

// Maze Side
var roomGfxMazeSide = CompressedRoom{
	"XXXX  XX  XX  XX  XXXX  XX  XX  XX  XXXX",
	"      XX      XX  XXXX  XX      XX      ",
	"      XX      XX  XXXX  XX      XX      ",
	"      XXXXXX  XX  XXXX  XX  XXXXXX      ",
	"      XXXXXX  XX  XXXX  XX  XXXXXX      ",
	"                  XXXX                  ",
	"                  XXXX                  ",
	"      XXXXXXXX    XXXX    XXXXXXXX      ",
	"      XXXXXXXX    XXXX    XXXXXXXX      ",
	"      XX          XXXX          XX      ",
	"      XX          XXXX          XX      ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Maze Entry
var roomGfxMazeEntry = CompressedRoom{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXX  XX    XXXXXXXXXXXXXXXXR   XX  XXXX",
	"XXXX  XX    XXXXXXXXXXXXXXXXR   XX  XXXX",
	"      XX          XXXX          XX      ",
	"      XX          XXXX          XX      ",
	"XXXXXXXX  XX      XXXX      XX  XXXXXXXX",
	"XXXXXXXX  XX      XXXX      XX  XXXXXXXX",
	"          XX      XXXX      XX          ",
	"          XX      XXXX      XX          ",
	"XXXXXXXXXXXX  XX  XXXX  XX  XXXXXXXXXXXX",
}

// Castle
var roomGfxCastle = CompressedRoom{
	"XXXXXXXXXXX X X X      X X X XXXXXXXXXXX",
	"XX        X X X X      X X X X        XX",
	"XX        XXXXXXX      XXXXXXX        XX",
	"XX        XXXXXXXXXXXXXXXXXXXX        XX",
	"XX          XXXXXXXXXXXXXXXX          XX",
	"XX          XXXXXX    XXXXXX          XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XXXXXXXXXXXXXX            XXXXXXXXXXXXXX",
}

// Red Maze #1
var roomGfxRedMaze1 = CompressedRoom{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"              XX        XX              ",
	"              XX        XX              ",
	"XXXXXXXXXXXX  XX        XX  XXXXXXXXXXXX",
	"XXXXXXXXXXXX  XX        XX  XXXXXXXXXXXX",
	"XXXX      XX  XX  XXXX  XX  XX      XXXX",
	"XXXX      XX  XX  XXXX  XX  XX      XXXX",
	"XXXX  XX  XXXXXX  XXXX  XXXXXX  XX  XXXX",
}

// Bottom of Red Maze
var roomGfxRedMazeBottom = CompressedRoom{
	"XXXX  XX  XXXXXX  XXXX  XXXXXX  XX  XXXX",
	"XXXX  XX                        XX  XXXX",
	"XXXX  XX                        XX  XXXX",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"      XX  XX                XX  XX  XXXX",
	"      XX  XX                XX  XX  XXXX",
	"XXXXXXXXXXXX                XXXXXXXXXXXX",
	"XXXXXXXXXXXX                XXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Top of Red Maze
var roomGfxRedMazeTop = CompressedRoom{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                  XXXX                  ",
	"                  XXXX                  ",
	"XXXXXXXXXXXXXXXX  XXXX  XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX  XXXX  XXXXXXXXXXXXXXXX",
	"              XX  XXXX  XX              ",
	"              XX  XXXX  XX              ",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"XXXX  XX  XX                XX  XX  XXXX",
	"XXXX  XX  XX                XX  XX  XXXX",
	"XXXX  XXXXXX  XX        XX  XXXXXX  XXXX",
}

// White Castle Entry
var roomGfxWhiteCastleEntry = CompressedRoom{
	"XXXX  XXXXXX  XX        XX  XXXXXX  XXXX",
	"XXXX          XX        XX          XXXX",
	"XXXX          XX        XX          XXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXX  XX                        XX  XXXX",
	"XXXX  XX                        XX  XXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Top Entry Room
var roomGfxTopEntryRoom = CompressedRoom{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Black Maze #1
var roomGfxBlackMaze1 = CompressedRoom{
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"            XX            XX            ",
	"            XX            XX            ",
	"XXXXXXXXXXXXXX            XXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXX            XXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX",
	"XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
}

// Black Maze #3
var roomGfxBlackMaze3 = CompressedRoom{
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"XX                  XX                  ",
	"XX                  XX                  ",
	"XX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXX",
	"XX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXX",
	"      XX                  XX            ",
	"      XX                  XX            ",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"XX          XX      XX          XX      ",
	"XX          XX      XX          XX      ",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
}

// Black Maze #2
var roomGfxBlackMaze2 = CompressedRoom{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                  XX                  XX",
	"                  XX                  XX",
	"XXXXXXXXXXXXXXXX  XXXXXXXXXXXXXXXXXX  XX",
	"XXXXXXXXXXXXXXXX  XXXXXXXXXXXXXXXXXX  XX",
	"                  XX                  XX",
	"                  XX                  XX",
	"XXXX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXX",
	"XXXX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXX",
	"        XXXX      XX        XXXX      XX",
	"        XXXX      XX        XXXX      XX",
	"XX  XX  XXXX  XX  XXXX  XX  XXXX  XX  XX",
}

// Black Maze Entry
var roomGfxBlackMazeEntry = CompressedRoom{
	"XX  XX  XXXX  XX  XXXX  XX  XXXX  XX  XX",
	"    XX        XX  XXXX  XX        XX    ",
	"    XX        XX  XXXX  XX        XX    ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}
