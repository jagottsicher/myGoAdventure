package game

import "time"

// Game difficulty switch values (matching original Atari hardware).
const (
	DifficultyA = 0 // harder
	DifficultyB = 1 // easier (dragons hesitate before biting)
)

// Game state machine values.
const (
	StateGameSelect = 0
	StateActive1    = 1 // Game variation 1
	StateActive2    = 2 // Game variation 2
	StateActive3    = 3 // Game variation 3
	StateWin        = 4
)

// Game holds top-level configuration and the current finite-state.
type Game struct {
	state      int
	level      int           // 0=easy/1=medium/2=hard (zero-indexed)
	fps        time.Duration // target frames per second
	diffLeft   int           // left difficulty switch
	diffRight  int           // right difficulty switch
}

var game = &Game{
	state:     StateActive2,
	level:     1,           // medium (game 2)
	fps:       30,
	diffLeft:  DifficultyB,
	diffRight: DifficultyB,
}

// IsGameOver returns true when the game has reached a terminal state.
func IsGameOver() bool {
	return false // placeholder – win condition not yet implemented
}
