package simpleui

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image"
	"log"
)

type WidgetMenu struct {
}

func (w *WidgetMenu) Draw(screen *ebiten.Image) {
	//mouseX, mouseY := ebiten.CursorPosition()
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("mouseXY(%d, %d)", mouseX, mouseY))
	mouseClickedUpdateMessage := "Click on the button to update the coordinates"
	if lastMouseClicked {
		mouseClickedUpdateMessage = fmt.Sprintf("Button was last clicked on coordinates: (%d, %d)", lastMouseX, lastMouseY)
	}
	ebitenutil.DebugPrint(screen, mouseClickedUpdateMessage)
	button.Draw(screen)
}

func (w *WidgetMenu) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 640, 480
}

func (w *WidgetMenu) Update(screen *ebiten.Image) error {
	mouseX, mouseY := ebiten.CursorPosition()
	mouseJustPressed := inpututil.IsMouseButtonJustPressed(0)
	mouseJustReleased := inpututil.IsMouseButtonJustReleased(0)
	button.Update(mouseX, mouseY, mouseJustPressed, mouseJustReleased)
	return nil
}

var greenTile *ebiten.Image
var roadTile *ebiten.Image
var skyTile *ebiten.Image
var wallTile *ebiten.Image

var button *Button
var lastMouseX, lastMouseY int
var lastMouseClicked bool

func updateLastMouse(mouseX, mouseY int) {
	lastMouseX = mouseX
	lastMouseY = mouseY
	lastMouseClicked = true
}

func init() {
	var err error
	greenTile, _, err = ebitenutil.NewImageFromFile("resources/images/mapclicker/Green Tile.png", ebiten.FilterDefault)
	roadTile, _, err = ebitenutil.NewImageFromFile("resources/images/mapclicker/Road Tile.png", ebiten.FilterDefault)
	skyTile, _, err = ebitenutil.NewImageFromFile("resources/images/mapclicker/Sky Tile.png", ebiten.FilterDefault)
	wallTile, _, err = ebitenutil.NewImageFromFile("resources/images/mapclicker/Wall Tile.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	lastMouseClicked = false

	button = NewButton(
		&image.Rectangle{
			Min: image.Point{200, 200},
			Max: image.Point{296, 296},
		},
		updateLastMouse,
		map[ButtonStatus]*ebiten.Image{
			ButtonStatusActive:   greenTile,
			ButtonStatusPressed:  roadTile,
			ButtonStatusHover:    skyTile,
			ButtonStatusDisabled: wallTile,
		},
		ButtonStatusActive,
	)
}

type ButtonStatus string

const (
	ButtonStatusActive   = "Active"
	ButtonStatusHover    = "Hover"
	ButtonStatusPressed  = "Clicked"
	ButtonStatusDisabled = "Disabled"
)

type Button struct {
	callbackFn    func(int, int)
	location      *image.Rectangle
	imageByStatus map[ButtonStatus]*ebiten.Image
	status        ButtonStatus
}

func NewButton(location *image.Rectangle, callbackFn func(int, int), imageByStatus map[ButtonStatus]*ebiten.Image, status ButtonStatus) *Button {
	return &Button{
		callbackFn:    callbackFn,
		location:      location,
		imageByStatus: imageByStatus,
		status:        status,
	}
}

func (b *Button) getImage() *ebiten.Image {
	buttonImage, ok := b.imageByStatus[b.status]
	if !ok {
		return nil
	}
	return buttonImage
}

func (b *Button) Draw(screen *ebiten.Image) {
	if b.location == nil {
		return
	}

	buttonDrawOperation := &ebiten.DrawImageOptions{}
	buttonDrawOperation.GeoM.Translate(float64(b.location.Min.X), float64(b.location.Min.Y))
	buttonImage := b.getImage()
	if buttonImage != nil {
		screen.DrawImage(buttonImage, buttonDrawOperation)
	}
}

func (b *Button) Update(mouseX, mouseY int, mouseJustPressed, mouseJustReleased bool) {

	if b.status == ButtonStatusDisabled {
		return
	}

	mouseRect := image.Rectangle{
		Min: image.Point{mouseX, mouseY},
		Max: image.Point{mouseX + 1, mouseY + 1},
	}
	mouseIsHoveringOverButton := mouseRect.In(*b.location)

	switch {
	case mouseJustReleased:
		b.callbackFn(mouseX, mouseY)
		b.status = ButtonStatusActive
	case b.status == ButtonStatusPressed:
		b.status = ButtonStatusPressed
	case mouseJustPressed && mouseIsHoveringOverButton:
		b.status = ButtonStatusPressed
	case mouseIsHoveringOverButton:
		b.status = ButtonStatusHover
	case !mouseIsHoveringOverButton:
		b.status = ButtonStatusActive
	}
}
