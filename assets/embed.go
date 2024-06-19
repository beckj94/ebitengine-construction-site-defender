package assets

import (
	_ "embed"
)

var (
	//go:embed img/backgrounds/scene-title.png
	MainMenu_png []byte

	//go:embed img/backgrounds/scene-gameplay.png
	GameplaySceneBackground_png []byte
)

