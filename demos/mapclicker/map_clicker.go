package mapclicker

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

type mapTile struct {
	r       int
	q       int
	tilePic *ebiten.Image
}

type MapClickerGame struct {
	userClickedOnScreen bool
}

func (m *MapClickerGame) Update(screen *ebiten.Image) error {
	return nil
}

func (m *MapClickerGame) convertMapToWorldCoordinate(r, q int) (float64, float64) {
	// See Axial Coordinates in:
	// https://www.redblobgames.com/grids/hexagons/
	// r applies the vector (1, 0)
	// q applies the vector (1/2, sqrt(3)/2)
	xPos := float64(r) + float64(q)*0.5
	yPos := float64(q) * 0.866
	return xPos * 96, yPos * 96
}

func (m *MapClickerGame) drawMap(screen *ebiten.Image) {
	mapTiles := []mapTile{
		{
			r:       0,
			q:       0,
			tilePic: greenTile,
		},
		{
			r:       1,
			q:       0,
			tilePic: roadTile,
		},
		{
			r:       0,
			q:       1,
			tilePic: greenTile,
		},
	}

	for _, tile := range mapTiles {
		tileOp := &ebiten.DrawImageOptions{}
		tileOp.GeoM.Translate(m.convertMapToWorldCoordinate(tile.r, tile.q))
		screen.DrawImage(tile.tilePic, tileOp)
	}
}

func (m *MapClickerGame) Draw(screen *ebiten.Image) {
	m.drawMap(screen)

	m.userClickedOnScreen = ebiten.IsMouseButtonPressed(0)

	if m.userClickedOnScreen {
		mouseX, mouseY := ebiten.CursorPosition()
		tileX := mouseX / 96
		tileY := mouseY / 96
		ebitenutil.DebugPrint(screen, fmt.Sprintf("tile(%d, %d), screen(%d, %d)", tileX, tileY, mouseX, mouseY))
	} else {
		ebitenutil.DebugPrint(screen, "Click on the screen")
	}
}

func (m *MapClickerGame) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 640, 480
}

var greenTile *ebiten.Image
var roadTile *ebiten.Image

func init() {
	var err error
	greenTile, _, err = ebitenutil.NewImageFromFile("resources/images/mapclicker/Green Tile.png", ebiten.FilterDefault)
	roadTile, _, err = ebitenutil.NewImageFromFile("resources/images/mapclicker/Road Tile.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}
