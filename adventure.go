package main

import (
	"time"

	"development/myGoAdventure/internal/game"
	"development/myGoAdventure/internal/input"
	"development/myGoAdventure/internal/render"
)

func main() {
	render.InitScreen()
	render.InitGamestate()
	input.InitUserInput()

	frameDuration := time.Second / game.G.FPS

	render.DrawStage()

	for !isGameOver() {
		startTime := time.Now()

		updateStates()

		render.DrawStage()
		if game.HelpMode {
			render.DrawHelp()
		} else {
			render.DrawAllVisibleObjects()
			render.DrawSpecialRooms()
			render.DrawSelOverlay()
			if game.GodMode {
				render.DrawDebugBat()
			}
		}
		render.DrawEaten()
		render.DrawWinOverlay()
		if game.ConfirmMode {
			render.DrawConfirm()
		}
		render.Screen.Show()

		remainingTime := frameDuration - time.Since(startTime)
		if remainingTime > 0 {
			time.Sleep(remainingTime)
		}
	}
}

func updateStates() {
	if game.NeedFullReset {
		game.ClearForFullReset()
		render.InitGamestate()
	}
	termW, termH := render.Screen.Size()
	input.HandleUserInput()
	game.UpdatePlayerStyle()
	game.UpdateSelOverlay()
	game.GreenDragon.Animate()
	game.YellowDragon.Animate()
	game.RedDragon.Animate()
	game.Bat.Animate()
	game.UpdatePortcullis(game.PortcullisYellow, game.YellowKey, termW, termH)
	game.UpdatePortcullis(game.PortcullisWhite, game.WhiteKey, termW, termH)
	game.UpdatePortcullis(game.PortcullisBlack, game.BlackKey, termW, termH)
	if game.UpdateCastlePortals(termW, termH) {
		render.FillTheScreen()
	}
	game.Sword.Animate()
	game.Magnet.Animate()
	game.TryPickup(termW, termH)
	game.UpdateCarriedObject(termW, termH)
	game.UpdateBat(termW, termH)
	game.UpdateMagnet(termW, termH)
	game.UpdateDragons(termW, termH)
	if game.GreenDS.State != 1 {
		facePlayer(game.GreenDragon, termW)
	}
	if game.YellowDS.State != 1 {
		facePlayer(game.YellowDragon, termW)
	}
	if game.RedDS.State != 1 {
		facePlayer(game.RedDragon, termW)
	}
	game.AdvanceFlashColor()
	game.UpdateChaliceColor()
	game.CheckWinCondition()
	game.UpdateWinState()
	game.UpdateEasterEggBarrier()
}

func facePlayer(dragon *game.Object, termW int) {
	dragonCenterX := dragon.RelX + float64(dragon.Width)/(2*float64(termW))
	playerCenterX := game.Player.RelX + float64(game.Player.Width)/(2*float64(termW))
	dragon.Flipped = dragonCenterX < playerCenterX
}

func isGameOver() bool {
	return false
}
