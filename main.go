package main

import (
	"bytes"

	"github.com/japanoise/gocademy/maps"
	"github.com/nsf/termbox-go"
)

var (
	GroundFloor *maps.Map
	FirstFloor  *maps.Map
	Roof        *maps.Map
	Athletics   *maps.Map
	AllMaps     []*maps.Map
	CurrentMap  int
)

func LoadMaps() {
	r := bytes.NewReader(MustAsset("bindata/groundfloor.bin"))
	GroundFloor, _ = maps.Deserialize(r)
	r = bytes.NewReader(MustAsset("bindata/firstfloor.bin"))
	FirstFloor, _ = maps.Deserialize(r)
	r = bytes.NewReader(MustAsset("bindata/roof.bin"))
	Roof, _ = maps.Deserialize(r)
	r = bytes.NewReader(MustAsset("bindata/athletics.bin"))
	Athletics, _ = maps.Deserialize(r)
	AllMaps = []*maps.Map{GroundFloor, FirstFloor, Roof, Athletics}
}

func init() {
	LoadMaps()
	CurrentMap = 0
}

func main() {
	termbox.Init()
	defer termbox.Close()

	playing := true
	x, y := 0, 0
	for playing {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		sx, sy := termbox.Size()
		AllMaps[CurrentMap].DrawMap(x, y, sx, sy)
		termbox.Flush()
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				playing = false
			case termbox.KeyArrowRight:
				x++
			case termbox.KeyArrowLeft:
				x--
			case termbox.KeyArrowDown:
				y++
			case termbox.KeyArrowUp:
				y--
			case termbox.KeyPgdn:
				CurrentMap = (CurrentMap + 1) % len(AllMaps)
			}
		}
	}
}
