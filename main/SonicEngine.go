package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

/**
Engine of the game with the logic of the animation of each sprites.
*/

func stopAnimation(win *pixelgl.Window) {
	sonicStopPic.drawPicture(win, 0, pixelRightRect)
}

func waitingAnimation(win *pixelgl.Window) {
	if frame < 2 {
		minWaitX = minWaitX + pixelWaitRect
		maxWaitX = maxWaitX + pixelWaitRect
		frame = frame + 1
	} else {
		frame = 0
		minWaitX = 0
		maxWaitX = pixelWaitRect
	}
	sonicWaitingPic.drawPicture(win, minWaitX, maxWaitX)
	time.Sleep(200 * time.Millisecond)
}

func upAnimation(win *pixelgl.Window) {
	if frame < 5 {
		minUpX = minUpX + pixelUpRect
		maxUpX = maxUpX + pixelUpRect
		frame = frame + 1
	}
	sonicUpPic.drawPicture(win, minUpX, maxUpX)
}

func downAnimation(win *pixelgl.Window) {
	if frame < 5 {
		minDownX = minDownX + pixelDownRect
		maxDownX = maxDownX + pixelDownRect
		frame = frame + 1
	}
	sonicDownPic.drawPicture(win, minDownX, maxDownX)
}

func ballAnimation(win *pixelgl.Window) {
	if frame < 5 {
		minBallX = minBallX + pixelBallRect
		maxBallX = maxBallX + pixelBallRect
		frame = frame + 1
	} else {
		frame = 0
		minBallX = 0
		maxBallX = pixelBallRect
	}
	sonicBallPic.drawPicture(win, minBallX, maxBallX)
}

func leftAnimation(win *pixelgl.Window) {
	if frame < 9 {
		minLeftX = minLeftX - (pixelLeftRect + 1.8)
		maxLeftX = maxLeftX - pixelLeftRect
	} else {
		frame = 0
		minLeftX = (pixelLeftRect + 1.8) * 8
		maxLeftX = pixelLeftRect * 9
	}
	frame = frame + 1
	xStep = xStep - 5
	sonicLeftPic.drawPicture(win, minLeftX, maxLeftX)
}

func rightAnimation(win *pixelgl.Window) {
	if frame < 9 {
		minRightX = minRightX + pixelRightRect
		maxRightX = maxRightX + pixelRightRect
	} else {
		frame = 0
		minRightX = 0
		maxRightX = pixelRightRect
	}
	frame = frame + 1
	xStep = xStep + 5
	sonicRightPic.drawPicture(win, minRightX, maxRightX)
}

func jumpAnimation(win *pixelgl.Window) {
	if frame < 8 {
		minJumpX = minJumpX + pixelJumpRect
		maxJumpX = maxJumpX + pixelJumpRect
	} else {
		frame = 0
		minJumpX = 0
		maxJumpX = pixelJumpRect
	}
	frame = frame + 1
	sonicJumpPic.drawPicture(win, minJumpX, maxJumpX)
}

/**
Implementation of [SonicSprites] to draw the specific sprite in the window.
Using interface each group of sprites invokes this method, avoiding type mismatch
*/
func (pic SonicSprites) drawPicture(win *pixelgl.Window, minX float64, maxX float64) {
	fmt.Printf("Moving... frame:%d .xMin:%f xMax:%f\n", frame, minRightX, maxRightX)
	sonicSprite := pixel.NewSprite(pic.picture, pixel.R(minX, 0, maxX, 50))
	movement = pixel.Vec{X: xStep, Y: yStep}
	sonicSprite.Draw(win, pixel.IM.Moved(movement))
}
