package main

import (
	"github.com/cserrant/teros-ebiten-prototypes/demos/simpleui"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth, screenHeight = 640, 480
)

func main() {
	//h := &ship.HelloWorldGame{}
	//h := &mapclicker.MapClickerGame{}
	//h := &mapscroll.MapScrollGame{}
	h := &simpleui.WidgetMenu{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Teros demos")
	if err := ebiten.RunGame(h); err != nil {
		log.Fatal(err)
	}
}
