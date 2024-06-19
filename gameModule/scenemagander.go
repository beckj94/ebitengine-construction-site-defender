package gameModule

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
  transitionFrom = ebiten.NewImage(ScreenWidth, ScreenHeight)
  transitionTo = ebiten.NewImage(ScreenWidth, ScreenHeight)
)

type Scene interface {
  Update(state *GameState) error
  Draw(screen *ebiten.Image)
}

const transitionMaxCount = 20 // ??? what is this? 20 frames?

type SceneManager struct {
  current Scene
  next Scene
  transitionCount int
}

type GameState struct {
  SceneManager *SceneManager
}

func (s *SceneManager) Update() error {
  if s.transitionCount == 0 {
    return s.current.Update(&GameState{
      SceneManager: s,
    })
  }

  s.transitionCount--
  if s.transitionCount > 0 {
    return nil
  }

  s.current = s.next
  s.next = nil
  return nil
}

func (s *SceneManager) Draw(r *ebiten.Image) {
  if s.transitionCount == 0 {
    s.current.Draw(r)
    return
  }

  transitionFrom.Clear()
  s.current.Draw(transitionFrom)

  transitionTo.Clear()
  s.next.Draw(transitionTo)

  r.DrawImage(transitionFrom, nil)

  alpha := 1 - float32(s.transitionCount)/float32(transitionMaxCount)
  op := &ebiten.DrawImageOptions{}
  op.ColorScale.ScaleAlpha(alpha)
  r.DrawImage(transitionTo, op)
}

func (s *SceneManager) GoTo(scene Scene) {
  log.Printf("scene manager goes to %T", scene)
  if s.current == nil {
    s.current = scene
  } else {
    s.next = scene
    s.transitionCount = transitionMaxCount
  }
}
