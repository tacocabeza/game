package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	"image"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 0
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8

	ActionMoveLeft input.Action = iota
	ActionMoveRight
	ActionMoveUp
	ActionMoveDown
)

var (
	runnerImage *ebiten.Image
)

type Game struct {
	count       int
	p           *player
	inputSystem input.System
}

type player struct {
	input *input.Handler
	pos   image.Point
}

func newExampleGame() *Game {
	g := &Game{}
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	keymap := input.Keymap{
		ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
		ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
		ActionMoveUp:    {input.KeyGamepadUp, input.KeyUp, input.KeyW},
		ActionMoveDown:  {input.KeyGamepadDown, input.KeyDown, input.KeyS},
	}
	g.p = &player{
		input: g.inputSystem.NewHandler(0, keymap),
		pos:   image.Point{X: 96, Y: 96},
	}

	return g
}

func (g *Game) Update() error {
	g.inputSystem.Update()
	g.p.Update()
	g.count++
	return nil
}

func (p *player) Update() {
	if p.input.ActionIsPressed(ActionMoveLeft) {
		p.pos.X -= 4
	}
	if p.input.ActionIsPressed(ActionMoveRight) {
		p.pos.X += 4
	}
	if p.input.ActionIsPressed(ActionMoveUp) {
		p.pos.Y -= 4
	}
	if p.input.ActionIsPressed(ActionMoveDown) {
		p.pos.Y += 4
	}
}

func (p *player) Draw(screen *ebiten.Image, sprite *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.pos.X), float64(p.pos.Y))
	screen.DrawImage(sprite, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	/* 	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op) */
	player := runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
	g.p.Draw(screen, player)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	img, _, err := ebitenutil.NewImageFromFile("assets/char_walk_right.png")

	if err != nil {
		log.Fatal(err)
	}

	runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("animation")
	if err := ebiten.RunGame(newExampleGame()); err != nil {
		log.Fatal(err)
	}

}
