package main

import tl "github.com/JoelOtter/termloop"

import (
	"fmt"
)

////////////////////PLAYER
type Player struct {
	*tl.Entity
	relativeX int
	relativeY int
	box       *Box
	level     *tl.BaseLevel
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		newx := player.relativeX
		newy := player.relativeY
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			newx = player.relativeX + 1
		case tl.KeyArrowLeft:
			newx = player.relativeX - 1
		case tl.KeyArrowUp:
			newy = player.relativeY - 1
		case tl.KeyArrowDown:
			newy = player.relativeY + 1
		}

		if newx >= player.box.width {
			player.relativeX = newx - player.box.width
			player.Fill(&tl.Cell{Fg: tl.ColorGreen, Ch: '@'})
		} else if newx < 0 {
			player.relativeX = newx + player.box.width
			player.Fill(&tl.Cell{Fg: tl.ColorCyan, Ch: '#'})
		} else {
			player.relativeX = newx
		}

		if newy >= player.box.height {
			player.relativeY = newy - player.box.height
			player.Fill(&tl.Cell{Fg: tl.ColorWhite, Ch: '*'})
		} else if newy < 0 {
			player.relativeY = newy + player.box.height
			player.Fill(&tl.Cell{Fg: tl.ColorYellow, Ch: '$'})
		} else {
			player.relativeY = newy
		}

		msg := fmt.Sprint(player.relativeX, ", ", player.relativeY)
		statusBox.setStatus(msg)
		//player.level.SetOffset(player.relativeX, player.relativeY)
	}
}

func (player *Player) Draw(screen *tl.Screen) {
	//screenWidth, screenHeight := screen.Size()
	x, y := player.relativeX, player.relativeY
	originX, originY := player.box.innerRect.Position()
	player.SetPosition(x+originX, y+originY)

	// We need to make sure and call Draw on the underlying Entity.
	player.Entity.Draw(screen)
}

//////////////// BOUNDING BOX

type Box struct {
	x         int
	y         int
	width     int
	height    int
	outerRect *tl.Rectangle
	innerRect *tl.Rectangle
	level     *tl.BaseLevel
}

func (box *Box) Draw(screen *tl.Screen) {

	////////////////////
	screenWidth, screenHeight := screen.Size()
	centerX := screenWidth / 2
	centerY := screenHeight / 2
	innerWidth, innerHeight := box.innerRect.Size()
	outerWidth, outerHeight := box.outerRect.Size()
	innerX := centerX - innerWidth/2
	outerX := centerX - outerWidth/2
	innerY := centerY - innerHeight/2
	outerY := centerY - outerHeight/2
	box.innerRect.SetPosition(innerX, innerY)
	box.outerRect.SetPosition(outerX, outerY)

	box.outerRect.Draw(screen)
	box.innerRect.Draw(screen)
}

func (box *Box) Tick(event tl.Event) {
	box.innerRect.Tick(event)
	box.outerRect.Tick(event)
}

func makeBox(level *tl.BaseLevel, top, left, width, height, thickness int) Box {
	box := Box{
		left,
		top,
		width,
		height,
		tl.NewRectangle(left-thickness, top-thickness, width+2*thickness, height+2*thickness, tl.ColorWhite),
		tl.NewRectangle(left, top, width, height, tl.ColorBlack),
		level,
	}

	return box
}

//////////////////// STATUS BOX
type StatusBox struct {
	msg   *tl.Text
	level *tl.BaseLevel
}

func (statusBox *StatusBox) Draw(screen *tl.Screen) {
	statusBox.msg.Draw(screen)
}

func (statusBox *StatusBox) Tick(event tl.Event) {
}

func (statusBox *StatusBox) setStatus(msg string) {
	statusBox.msg.SetText(msg)
}

var statusBox StatusBox

///////////////////////////MAIN

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorDefault,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	box := makeBox(level, 5, 5, 160, 30, 1)
	level.AddEntity(&box)

	statusBox = StatusBox{tl.NewText(15, 43, "", tl.ColorWhite, tl.ColorBlack), level}
	level.AddEntity(&statusBox)

	// Set the character at position (0, 0) on the entity.
	player := Player{tl.NewEntity(1, 1, 1, 1), 0, 0, &box, level}
	//player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '#'})
	player.Fill(&tl.Cell{Fg: tl.ColorWhite, Ch: '#'})
	level.AddEntity(&player)

	game.Screen().SetLevel(level)
	game.Start()
}
