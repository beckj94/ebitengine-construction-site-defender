package gameModule

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

type Game struct {
	sceneManager	*SceneManager
	// input        Input
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	// TODO: pass the keys to the active scene somehow
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		log.Printf("Space pressed on game")
	}
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoTo(&TitleScene{})
	}

	// g.input.Update()
	// if err := g.sceneManager.Update(&g.input); err != nil {
	// 	return err
	// }
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}
