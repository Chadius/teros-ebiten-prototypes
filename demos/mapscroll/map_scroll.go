package mapscroll

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

type mapTile struct {
	r       int
	q       int
	tilePic *ebiten.Image
}

type MapScrollGame struct {
}

func (m *MapScrollGame) Draw(screen *ebiten.Image) {
	m.drawMap(screen)
}

func (m *MapScrollGame) drawMap(screen *ebiten.Image) {
	for _, tile := range mapTiles {
		tileOp := &ebiten.DrawImageOptions{}
		tileWorldX, tileWorldY := m.convertMapToWorldCoordinate(tile.q, tile.r)
		screenX, screenY := m.convertWorldToScreenCoordinates(tileWorldX, tileWorldY, screen)
		tileOp.GeoM.Translate(screenX-48, screenY-48)
		screen.DrawImage(tile.tilePic, tileOp)
	}
}

func (m *MapScrollGame) convertMapToWorldCoordinate(q, r int) (float64, float64) {
	// See Axial Coordinates in:
	// https://www.redblobgames.com/grids/hexagons/
	// q applies the vector (1, 0)
	// r applies the vector (1/2, sqrt(3)/2)
	xPos := float64(q) + float64(r)*0.5
	yPos := float64(r) * 0.866
	return xPos * 96, yPos * 96
}

func (m *MapScrollGame) convertWorldToScreenCoordinates(x, y float64, screen *ebiten.Image) (float64, float64) {
	cameraX := float64(screen.Bounds().Max.X / 2)
	cameraY := float64(screen.Bounds().Max.Y / 2)

	return cameraX + x, cameraY + y
}

func (m *MapScrollGame) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 640, 480
}

func (m *MapScrollGame) Update(screen *ebiten.Image) error {
	return nil
}

var greenTile *ebiten.Image
var roadTile *ebiten.Image
var mapTiles []mapTile

func init() {
	var err error
	greenTile, _, err = ebitenutil.NewImageFromFile("resources/images/mapclicker/Green Tile.png", ebiten.FilterDefault)
	roadTile, _, err = ebitenutil.NewImageFromFile("resources/images/mapclicker/Road Tile.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	mapTiles = []mapTile{
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
}
