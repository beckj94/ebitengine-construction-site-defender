// Copyright 2024 Beck Jonathan and Czufor Bence. All rights reserved.
// Use of this source code is subject to an MIT-style
// licence which can be found in the LICENSE file.

package main

import (
  "errors"
  "image"
  // "image/color"
  _ "image/png"
  _ "image/jpeg"
  "os"
  "log"
  "time"

  "bytes"
  "io/ioutil"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "github.com/hajimehoshi/ebiten/v2/audio"
  "github.com/hajimehoshi/ebiten/v2/audio/wav"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
)

func main() {
  gameWidth, gameHeight := 1280, 720

  ebiten.SetWindowSize(gameWidth, gameHeight)
  ebiten.SetWindowTitle("Construction Site Defender")

  if err := ebiten.RunGame(&gameModule.Game{}); err != nil {
    log.Fatal(err)
  }
}
