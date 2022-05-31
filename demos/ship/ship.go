package ship

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

const (
	screenWidth, screenHeight = 640, 480
)

var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
	playerOne  player
)

type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerOne.xPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerOne.xPos += playerOne.speed
	}
}

type HelloWorldGame struct{}

func (h *HelloWorldGame) Update(screen *ebiten.Image) error {
	movePlayer()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)
	return nil
}

func (h *HelloWorldGame) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 640, 480
}

// Run this code once at startup
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/spaceship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	playerOne = player{spaceShip, screenWidth / 2, screenHeight / 2, 4}
}
