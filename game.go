package main

import "time"

// Game holds top-level game configuration such as variant and target frame rate.
type Game struct {
	gameType uint8
	FPS      time.Duration
}

var game = Game{
	gameType: 2,
	FPS:      60,
}

var err error
