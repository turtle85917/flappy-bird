package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenSizeX = 1024
	screenSizeY = 860
)

var (
	background = color.RGBA{66, 154, 208, 255}
	ground     = color.RGBA{73, 53, 12, 255}
	deepground = color.RGBA{64, 46, 11, 255}

	grounds      = []Ground{}
	groundHeight = float64(70)
	groundX      = float64(0)
)

type Ground struct {
	x float64
	y float64
	w float64
	h float64
}

type Game struct{}

func (g *Game) Update() error {
	groundX -= 10
	if groundX <= screenSizeX*-1 {
		groundX = screenSizeX
	}

	groundRefresh()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(background)
	for idx := 0; idx < len(grounds); idx++ {
		gclr := ground
		if idx%2 == 0 {
			gclr = deepground
		}

		ebitenutil.DrawRect(screen, grounds[idx].x, grounds[idx].y, grounds[idx].w, grounds[idx].h, gclr)
	}
}

func (g *Game) Layout(outsidewidth, outsideheight int) (screenwidth, screenheight int) {
	return screenSizeX, screenSizeY
}

func groundRefresh() {
	grounds = []Ground{
		{x: screenSizeX*-1 + groundX, y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight},
		{x: groundX, y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight},
		{x: screenSizeX + groundX, y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight}}
}

func init() {
	grounds = append(grounds,
		Ground{x: screenSizeX*-1 + groundX, y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight},
		Ground{x: groundX, y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight},
		Ground{x: screenSizeX + groundX, y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight})
}

func main() {
	ebiten.SetWindowSize(screenSizeX, screenSizeY)
	ebiten.SetWindowTitle("Flappy bird")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
