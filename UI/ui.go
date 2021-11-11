package ui

import (
	"TurkishDraughts/Board"

	"image/color"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	Width = 1600
	Height = 900
)




func Init() {
	b := board.CreateStartingBoard()

	cfg := pixelgl.WindowConfig{
		Title:  "Turkish Draughts Engine",
		Bounds: pixel.R(0, 0, Width, Height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	/*
	imd.Color = colornames.Blueviolet
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	imd.EndShape = imdraw.SharpEndShape
	imd.Push(pixel.V(100, 500), pixel.V(700, 500))
	imd.Line(30)

	imd.Color = colornames.Limegreen
	imd.Push(pixel.V(500, 500))
	imd.Circle(300, 50)

	imd.Color = colornames.Navy
	imd.Push(pixel.V(200, 500), pixel.V(800, 500))
	imd.Ellipse(pixel.V(120, 80), 0)

	imd.Color = colornames.Red
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(500, 350))
	imd.CircleArc(150, -math.Pi, 0, 30)
	*/

	for !win.Closed() {
		imd := imdraw.New(nil)
		win.Clear(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
		
		drawBoard(imd)
		drawPieces(&b, imd)

		imd.Draw(win)
		win.Update()

		winner, _ := b.PlayerHasWon()
		if !winner {  
			b = *(Analyze(b, 9))
		}
	}
}

