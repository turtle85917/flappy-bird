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
	background     = color.RGBA{66, 154, 208, 255}
	deepbackground = color.RGBA{65, 150, 200, 255}
	ground         = color.RGBA{73, 53, 12, 255}
	deepground     = color.RGBA{64, 46, 11, 255}

	backgrounds = []Background{}

	grounds      = []Ground{}
	groundHeight = float64(70)

	endlessX = float64(0)
)

type Sprite struct {
	x float64
	y float64
	w float64
	h float64
}

type Background struct {
	Sprite
}

type Ground struct {
	Sprite
}

type Game struct{}

func (g *Game) Update() error {
	endlessX -= 10
	if endlessX <= screenSizeX*-1 {
		endlessX = screenSizeX
	}

	refresh()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for idx := 0; idx < len(backgrounds); idx++ {
		bclr := deepbackground
		if idx%2 == 0 {
			bclr = background
		}
		ebitenutil.DrawRect(screen, backgrounds[idx].x, backgrounds[idx].y, backgrounds[idx].w, backgrounds[idx].h, bclr)
	}

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

func refresh() {
	grounds = []Ground{
		{Sprite{x: screenSizeX*-1 + endlessX, y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight}},
		{Sprite{x: endlessX, y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight}},
		{Sprite{x: screenSizeX + endlessX, y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight}}}
	backgrounds = []Background{
		{Sprite{x: initX(-1), y: 0, w: screenSizeX, h: screenSizeY}},
		{Sprite{x: initX(0), y: 0, w: screenSizeX, h: screenSizeY}},
		{Sprite{x: initX(1), y: 0, w: screenSizeX, h: screenSizeY}}}
}

func initX(level float64) float64 {
	return screenSizeX*level + endlessX
}

func init() {
	grounds = append(grounds,
		Ground{Sprite{x: initX(-1), y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight}},
		Ground{Sprite{x: initX(0), y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight}},
		Ground{Sprite{x: initX(1), y: screenSizeY - groundHeight, w: screenSizeX, h: groundHeight}})
	backgrounds = append(backgrounds,
		Background{Sprite{x: initX(-1), y: 0, w: screenSizeX, h: screenSizeY}},
		Background{Sprite{x: initX(0), y: 0, w: screenSizeX, h: screenSizeY}},
		Background{Sprite{x: initX(1), y: 0, w: screenSizeX, h: screenSizeY}})
}

func main() {
	ebiten.SetWindowSize(screenSizeX, screenSizeY)
	ebiten.SetWindowTitle("Flappy bird")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
