// Copyright 2024 Beck Jonathan and Czufor Bence. All rights reserved.
// Use of this source code is subject to an MIT-style
// licence which can be found in the LICENSE file.

package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/beckj94/ebitengine-construction-site-defender/gameModule"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
  gameWidth, gameHeight := 1280, 720

  ebiten.SetWindowSize(gameWidth, gameHeight)
  ebiten.SetWindowTitle("Construction Site Defender")

  if err := ebiten.RunGame(&gameModule.Game{}); err != nil {
    log.Fatal(err)
  }
}
