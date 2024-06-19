package gameModule

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/beckj94/ebitengine-construction-site-defender/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var imageBackground *ebiten.Image

type TitleScene struct {}

func loadSceneBackground() *ebiten.Image {
	file, _, err := image.Decode(bytes.NewReader(assets.MainMenu_png))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(file)
}

func init() {
	img := loadSceneBackground()
	imageBackground = ebiten.NewImageFromImage(img)
	log.Printf("Titlescene init")
}


func (s *TitleScene) Update(state *GameState) error {

	log.Printf("TitleScene update")
	// Start the game on space press
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
	// if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		log.Printf("space pressed on title scene")
		state.SceneManager.GoTo(GameplayScene())
		return nil
	}

	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) {
	s.drawTitleBackground(r)
	message := "PRESS SPACE TO START"
	x := ScreenWidth / 2
	y := ScreenHeight - 120
	drawTextWithShadow(r, message, x, y, 1, color.RGBA{0x80, 0, 0, 0xff}, text.AlignCenter, text.AlignStart)
}

func (s *TitleScene) drawTitleBackground(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	r.DrawImage(imageBackground, op)
}
