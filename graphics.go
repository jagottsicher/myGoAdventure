package main

//
// Room graphics
//

// Left of Name Room
var roomGfxLeftOfName = compressedRoom{
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
var roomGfxBelowYellowCastle = compressedRoom{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"1234567890123456789012345678901234567890",
	"1234567890123456789012345678901234567890",
	"1234567890123456789012345678901234567890",
	"1234567890123456789012345678901234567890",
	"1234567890123456789012345678901234567890",
	"1234567890123456789012345678901234567890",
	"1234567890123456789012345678901234567890",
	"1234567890123456789012345678901234567890",
	"1234567890123456789012345678901234567890",
	"1234567890123456789012345678901234567890",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Side CoXXidor
var roomGfxSideCorridor = compressedRoom{
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
var roomGfxNumberRoom = compressedRoom{
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
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// `
var roomGfxTwoExitRoom = compressedRoom{
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
var roomGfxBlueMazeTop = compressedRoom{
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
var roomGfxBlueMaze1 = compressedRoom{
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
var roomGfxBlueMazeBottom = compressedRoom{
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
var roomGfxBlueMazeCenter = compressedRoom{
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
var roomGfxBlueMazeEntry = compressedRoom{
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
var roomGfxMazeMiddle = compressedRoom{
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
var roomGfxMazeSide = compressedRoom{
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
var roomGfxMazeEntry = compressedRoom{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXX  XX    XXXXXXXXXXXXXXXXX   XX  XXXX",
	"XXXX  XX    XXXXXXXXXXXXXXXXX   XX  XXXX",
	"      XX          XXXX          XX      ",
	"      XX          XXXX          XX      ",
	"XXXXXXXX  XX      XXXX      XX  XXXXXXXX",
	"XXXXXXXX  XX      XXXX      XX  XXXXXXXX",
	"          XX      XXXX      XX          ",
	"          XX      XXXX      XX          ",
	"XXXXXXXXXXXX  XX  XXXX  XX  XXXXXXXXXXXX",
}

// Castle
var roomGfxCastle = compressedRoom{
	"XXXXXXXXXXX X X X      X X X XXXXXXXXXXX",
	"X         X X X X      X X X X         X",
	"X         XXXXXXX      XXXXXXX         X",
	"X         XXXXXXXXxxxxXXXXXXXX         X",
	"X           XXXXXXxxxxXXXXXX           X",
	"X           XXXXXX    XXXXXX           X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Red Maze #1
var roomGfxRedMaze1 = compressedRoom{
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
var roomGfxRedMazeBottom = compressedRoom{
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
var roomGfxRedMazeTop = compressedRoom{
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
var roomGfxWhiteCastleEntry = compressedRoom{
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
var roomGfxTopEntryRoom = compressedRoom{
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
var roomGfxBlackMaze1 = compressedRoom{
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
var roomGfxBlackMaze3 = compressedRoom{
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
var roomGfxBlackMaze2 = compressedRoom{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                  XX                  XX",
	"                  XX                  XX",
	"XXXXXXXXXXXXXXXX  XXXXXXXXXXXXXXXXXX  XX",
	"XXXXXXXXXXXXXXXX  XXXXXXXXXXXXXXXXXX  XX",
	"              XX                  XX    ",
	"              XX                  XX    ",
	"XXXX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXX",
	"XXXX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXX",
	"        XXXX      XX        XXXX      XX",
	"        XXXX      XX        XXXX      XX",
	"XX  XX  XXXX  XX  XXXX  XX  XXXX  XX  XX",
}

// Black Maze Entry
var roomGfxBlackMazeEntry = compressedRoom{
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

// Dragons graphics
type size struct {
	width  int
	height int
}

type spriteGfx []string

// Dragon gfx definition
var roomGfxDragonCalm = spriteGfx{
	"    DDD ",
	"DDDDD  D",
	"DDDDDDDD",
	"    DD  ",
	" DDDDDD ",
	"DD    DD",
	"DD    DD",
	" DDDDDD ",
	"    DD  ",
	"DD  DD   ",
	" DD DD   ",
	"  DDD   ",
}

var roomGfxDragonAggressive = spriteGfx{
	"D   DDD ",
	" D DD  D",
	"  DDDDDD",
	" D  DD  ",
	"D DDDDD ",
	"DD    DD",
	"DD    DD",
	" DDDDDD ",
	"  DDD   ",
	" DD    ",
	"  DD    ",
	"   DDDD ",
}

var roomGfxDragonDead = spriteGfx{
	"    DD  ",
	"    DD  ",
	"    DDD ",
	"   D  D ",
	"  DD    ",
	" DDDDDD ",
	"DD    DD",
	"DD    DD",
	" DDDDDD ",
	"  DDD   ",
	" DD   D ",
	"  DDDD  ",
}
