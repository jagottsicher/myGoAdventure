package main

//
// player graphics
//

var playerGfx = []*cell{
	{x: 0, y: 0, symbol: 'L'},
	{x: 1, y: 0, symbol: 'R'},
}

var playerGfxBefore = []*cell{
	{x: 0, y: 0, symbol: 'B'},
	{x: 1, y: 0, symbol: '4'},
}

//
// Room graphics
//

// Castle
var roomGfxCastle = &[]string{
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

func getWidth() int {

	screenWidth, _ := screen.Size()
	return screenWidth
}

func convertToBinary(data string) int64 {
	binary := int64(0)
	for _, char := range data {
		binary <<= 1
		if char == 'X' {
			binary |= 1
		}
	}
	return binary
}
