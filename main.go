package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

// constant values
const (
	// The screen Width and Screen Height for our simualtion
	screenWidth, screenHeight = 640, 360

	// Number of boids in the sky
	noBoids = 400
	adjRate = 0.020
	boidRaduis = 10
)

// The colour of our boid
var (
	pix   = color.RGBA{249, 105, 14, 255}
	boids = make([]*Boid, 400, 500)
	screenWidthSlice = make([]int, screenWidth)
	screenHeightSlice = make([]int, screenHeight)
	boidMap = make([][]int, (screenWidth * screenHeight))
)

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	// Draw a Diamond shape to represent the boids
	for _, v := range boids {
		screen.Set(int(v.Position.x+1), int(v.Position.y), pix)
		screen.Set(int(v.Position.x-1), int(v.Position.y), pix)
		screen.Set(int(v.Position.x), int(v.Position.y+1), pix)
		screen.Set(int(v.Position.x), int(v.Position.y-1), pix)

	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	// Create a boid for the max number of boids
	for i := 0; i < noBoids; i++ {
		boidConstructor(i)
	}

	// Set up the array for the screen
	for row:=0; row< screenWidth; row++ {
		for column:=0; column< screenHeight; column++ {
			boidMap[row][column] = -1
		}
	}


	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Go Go Boids!!!")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
