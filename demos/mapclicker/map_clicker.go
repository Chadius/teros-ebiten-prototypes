package mapclicker

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"math"
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

func (m *MapClickerGame) convertScreenToWorldCoordinates(x, y int, screen *ebiten.Image) (int, int) {
	worldX := x - screen.Bounds().Max.X/2
	worldY := y - screen.Bounds().Max.Y/2
	return worldX, worldY
}

func (m *MapClickerGame) convertWorldToScreenCoordinates(x, y float64, screen *ebiten.Image) (float64, float64) {
	cameraX := float64(screen.Bounds().Max.X / 2)
	cameraY := float64(screen.Bounds().Max.Y / 2)

	return cameraX + x, cameraY + y
}

func (m *MapClickerGame) convertMapToWorldCoordinate(q, r int) (float64, float64) {
	// See Axial Coordinates in:
	// https://www.redblobgames.com/grids/hexagons/
	// q applies the vector (1, 0)
	// r applies the vector (1/2, sqrt(3)/2)
	xPos := float64(q) + float64(r)*0.5
	yPos := float64(r) * 0.866
	return xPos * 96, yPos * 96
}

func (m *MapClickerGame) convertWorldCoordinatesToMapCoordinates(x, y int) (int, int) {
	xScaled := float64(x) / 96.0
	yScaled := float64(y) / 96.0

	// r = 2 * yScaled / sqrt(3)
	r := yScaled * 1.154

	// q = x - (y / sqrt(3))
	q := xScaled - (yScaled / 1.732)

	return int(math.Round(q)), int(math.Round(r))
}

func (m *MapClickerGame) drawMap(screen *ebiten.Image) {
	mapTiles := []mapTile{
		{
			q:       0,
			r:       0,
			tilePic: greenTile,
		},
		{
			q:       1,
			r:       0,
			tilePic: roadTile,
		},
		{
			q:       0,
			r:       1,
			tilePic: greenTile,
		},
	}

	for _, tile := range mapTiles {
		tileOp := &ebiten.DrawImageOptions{}
		tileWorldX, tileWorldY := m.convertMapToWorldCoordinate(tile.q, tile.r)
		screenX, screenY := m.convertWorldToScreenCoordinates(tileWorldX, tileWorldY, screen)
		tileOp.GeoM.Translate(screenX-48, screenY-48)
		screen.DrawImage(tile.tilePic, tileOp)
	}
}

func (m *MapClickerGame) Draw(screen *ebiten.Image) {
	m.drawMap(screen)

	m.userClickedOnScreen = ebiten.IsMouseButtonPressed(0)

	if m.userClickedOnScreen {
		screenX, screenY := ebiten.CursorPosition()
		worldX, worldY := m.convertScreenToWorldCoordinates(screenX, screenY, screen)
		tileQ, tileR := m.convertWorldCoordinatesToMapCoordinates(worldX, worldY)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("screenXY(%d, %d)\nworldXY(%d, %d)\ntileQR(%d, %d)", screenX, screenY, worldX, worldY, tileQ, tileR))
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
