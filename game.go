package main

import "time"

type Game struct {
	gameType uint8
	FPS      time.Duration
}

var game = Game{
	gameType: 2,
	FPS:      60,
}

var err error
