package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	fadeon     bool
	color      = rl.White
	fadeamount = float32(0.0)

	noiseLineX1, noiseLineX2, noiseLineX3, noiseLineX4, noiseLineX5, noiseLineX6, noiseLineX7, noiseLineX8, noiseLineX9, noiseLineX10, noiseLineX11, noiseLineX12 int32

	noiseLineX1Change, noiseLineX2Change, noiseLineX3Change, noiseLineX4Change, noiseLineX5Change, noiseLineX6Change, noiseLineX7Change, noiseLineX8Change, noiseLineX9Change, noiseLineX10Change, noiseLineX11Change, noiseLineX12Change int32

	noiseLineDistance1, noiseLineDistance2, noiseLineDistance3, noiseLineDistance4, noiseLineDistance5, noiseLineDistance6, noiseLineDistance7, noiseLineDistance8, noiseLineDistance9, noiseLineDistance10, noiseLineDistance11, noiseLineDistance12 int

	noiseLineLR1, noiseLineLR2, noiseLineLR3, noiseLineLR4, noiseLineLR5, noiseLineLR6, noiseLineLR7, noiseLineLR8, noiseLineLR9, noiseLineLR10, noiseLineLR11, noiseLineLR12 bool

	pixelNoiseOn, switchScanLines, noiseLinesOn, noiseLinesScreenOn, scanlineson bool
	goimg                                                                        = rl.NewRectangle(0, 0, 300, 500)
	imgs                                                                         rl.Texture2D
	frameCountGameStart                                                          int
	noiseLinesMAP                                                                = make([]int32, 12)
	noiseLinesDistanceMAP                                                        = make([]int, 12)
	noiseLinesLRMAP                                                              = make([]bool, 12)
	pixelNoiseMAP                                                                = make([]bool, 65472)
	screenW                                                                      = int32(1366)
	screenH                                                                      = int32(768)
)

func main() { // MARK: main()
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides INFO window
	pixelNoiseOn = true
	scanlineson = true
	noiseLinesOn = true

	cNOISELINES()
	cPIXELNOISE()
	raylib()
}

func cNOISELINES() {

	for a := 0; a < 12; a++ {
		noiseLinesMAP[a] = rInt32(0, int(screenW))
		noiseLinesLRMAP[a] = flipcoin()
		noiseLinesDistanceMAP[a] = rInt(100, 300)
	}
}

func cPIXELNOISE() {

	for a := 0; a < 880; a++ {
		pixelNoiseMAP[a] = false
	}
	for a := 0; a < 880; a++ {
		placePixelNoise := rolldice() + rolldice()
		if placePixelNoise == 12 {
			pixelNoiseMAP[a] = true
		}

	}

}

func raylib() {
	rl.InitWindow(screenW, screenH, "displayfx raylib / golang")
	imgs = rl.LoadTexture("gologo.png")

	rl.SetTargetFPS(60)

	// MARK: WindowShouldClose
	for !rl.WindowShouldClose() {
		frameCountGameStart++

		if frameCountGameStart%4 == 0 {
			if fadeon {
				if fadeamount > 0.0 {
					fadeamount -= 0.1
				} else {
					fadeon = false
				}

			} else {
				if fadeamount < 1.0 {
					fadeamount += 0.1
				} else {
					fadeon = true
				}

			}

		}
		if rl.IsKeyPressed(rl.KeyF1) {
			if pixelNoiseOn {
				pixelNoiseOn = false
			} else {
				pixelNoiseOn = true
			}
		}
		if rl.IsKeyPressed(rl.KeyF2) {
			if scanlineson {
				scanlineson = false
			} else {
				scanlineson = true
			}
		}
		if rl.IsKeyPressed(rl.KeyF3) {
			if noiseLinesOn {
				noiseLinesOn = false
			} else {
				noiseLinesOn = true
			}
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		logov2 := rl.NewVector2(533, 134)
		rl.DrawTextureRec(imgs, goimg, logov2, rl.White)

		rl.DrawRectangle(0, 0, screenW, screenH, rl.Fade(color, fadeamount))

		rl.DrawText("F1 - turn pixel noise on/off", 18, 22, 20, rl.Yellow)
		rl.DrawText("F2 - turn scan lines on/off", 18, 52, 20, rl.Yellow)
		rl.DrawText("F3 - turn noise lines on/off", 18, 82, 20, rl.Yellow)

		rl.DrawText("F1 - turn pixel noise on/off", 20, 20, 20, rl.Blue)
		rl.DrawText("F2 - turn scan lines on/off", 20, 50, 20, rl.Blue)
		rl.DrawText("F3 - turn noise lines on/off", 20, 80, 20, rl.Blue)

		// draw pixel noise
		if pixelNoiseOn {
			if frameCountGameStart%2 == 0 {
				cPIXELNOISE()
			}

			lineCountPixelNoise := 0
			pixelNoiseY := int32(0)
			pixelNoiseX := int32(0)
			for a := 0; a < 880; a++ {

				if pixelNoiseMAP[a] == true {
					rl.DrawRectangle(pixelNoiseX, pixelNoiseY, 2, 2, rl.Black)
				}

				lineCountPixelNoise += 34
				pixelNoiseX += 34
				if lineCountPixelNoise > 1350 {
					lineCountPixelNoise = 0
					pixelNoiseX = 0
					pixelNoiseY += 34

				}
			}
		}
		// draw scan lines
		if scanlineson {
			linesY := int32(0)
			for a := 0; a < int(screenH); a++ {
				rl.DrawLine(0, linesY, screenW, linesY, rl.Fade(rl.Black, 0.5))
				linesY += 2
				a++
			}
		}

		// noise lines
		if noiseLinesOn {
			if frameCountGameStart%60 == 0 {
				if noiseLinesScreenOn {
					noiseLinesScreenOn = false
				} else {
					noiseLinesScreenOn = true
				}
			}

			if noiseLinesScreenOn {
				for a := 0; a < noiseLineDistance1; a++ {
					noiseLineX1Change++
				}
				for a := 0; a < noiseLineDistance2; a++ {
					noiseLineX2Change++
				}
				for a := 0; a < noiseLineDistance3; a++ {
					noiseLineX3Change++
				}
				for a := 0; a < noiseLineDistance4; a++ {
					noiseLineX4Change++
				}
				if noiseLineLR1 {
					rl.DrawLine(noiseLineX1+noiseLineX1Change, 0, noiseLineX1+noiseLineX1Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX1-noiseLineX1Change, 0, noiseLineX1-noiseLineX1Change, screenH, rl.Black)
				}
				if noiseLineLR2 {
					rl.DrawLine(noiseLineX2+noiseLineX2Change, 0, noiseLineX1+noiseLineX2Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX2-noiseLineX2Change, 0, noiseLineX2-noiseLineX2Change, screenH, rl.Black)
				}
				if noiseLineLR3 {
					rl.DrawLine(noiseLineX3+noiseLineX3Change, 0, noiseLineX3+noiseLineX3Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX3-noiseLineX3Change, 0, noiseLineX3-noiseLineX3Change, screenH, rl.Black)
				}
				if noiseLineLR4 {
					rl.DrawLine(noiseLineX4+noiseLineX4Change, 0, noiseLineX4+noiseLineX4Change, screenH, rl.Black)
				} else {
					rl.DrawLine(noiseLineX4-noiseLineX4Change, 0, noiseLineX4-noiseLineX4Change, screenH, rl.Black)
				}

			} else {
				cNOISELINES()
				noiseLineX1Change = 0
				noiseLineX2Change = 0
				noiseLineX3Change = 0
				noiseLineX4Change = 0
				noiseLineX1 = noiseLinesMAP[0]
				noiseLineX2 = noiseLinesMAP[1]
				noiseLineX3 = noiseLinesMAP[2]
				noiseLineX4 = noiseLinesMAP[3]
				noiseLineDistance1 = noiseLinesDistanceMAP[0]
				noiseLineDistance2 = noiseLinesDistanceMAP[1]
				noiseLineDistance3 = noiseLinesDistanceMAP[2]
				noiseLineDistance4 = noiseLinesDistanceMAP[3]
				noiseLineLR1 = noiseLinesLRMAP[0]
				noiseLineLR2 = noiseLinesLRMAP[1]
				noiseLineLR3 = noiseLinesLRMAP[2]
				noiseLineLR4 = noiseLinesLRMAP[3]
			}
		}
		rl.EndDrawing()
	} // end WindowShouldClose
	rl.CloseWindow()
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
