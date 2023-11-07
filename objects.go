package main

import "github.com/gdamore/tcell"

// Gates
type gate struct {
	buildRune   rune
	pos_x       int
	pos_y       int
	width       int
	height      int
	unlocked    bool
	existsIn    *rooms
	up          *rooms
	objectStyle tcell.Style
}

var yellowCastleGate = gate{
	buildRune: '┼',
	pos_x:     73,
	pos_y:     18,
	width:     14,
	height:    4,
	unlocked:  false,
	existsIn:  &roomYellowCastle,
	up:        &roomOnTopOfYellowCastle,
	objectStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlack),
}

var whiteCastleGate = gate{
	buildRune: '┼',
	pos_x:     73,
	pos_y:     18,
	width:     14,
	height:    4,
	unlocked:  false,
	existsIn:  &roomWhiteCastle,
	up:        &roomRedMazeEntryBottomRight,
	objectStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlack),
}

var blackCastleGate = gate{
	buildRune: '┼',
	pos_x:     73,
	pos_y:     18,
	width:     14,
	height:    4,
	unlocked:  false,
	existsIn:  &roomBlackCastle,
	up:        &roomOnTopOfBlackCastle,
	objectStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlack),
}

// Bars
type bar struct {
	buildRune   rune
	pos_x       int
	width       int
	height      int
	unlocked    bool
	objectStyle tcell.Style
}

var leftBar = bar{
	buildRune: 'X',
	pos_x:     6,
	width:     2,
	height:    stageHeight,
	unlocked:  false,
	objectStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlack),
}

var rightBar = bar{
	buildRune: 'X',
	pos_x:     stageWidth - 8,
	width:     2,
	height:    stageHeight,
	unlocked:  false,
	objectStyle: tcell.StyleDefault.
		Background(tcell.ColorDarkGray).
		Foreground(tcell.ColorBlack),
}
