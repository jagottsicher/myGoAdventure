package world

import "github.com/gdamore/tcell/v2"

type Cell struct {
	X, Y   int
	Symbol rune
}

type Room struct {
	RoomData   *[]string
	Background tcell.Color
	Foreground tcell.Color
	BarLeft    bool
	BarRight   bool
	Up         *Room
	Down       *Room
	Left       *Room
	Right      *Room
}
