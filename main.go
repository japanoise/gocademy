package main

import (
	"bytes"

	"github.com/japanoise/gocademy/maps"
	"github.com/nsf/termbox-go"
)

var (
	AllMaps []*maps.Map
)

func LoadMaps() {
	AllMaps = make([]*maps.Map, 4)
	r := bytes.NewReader(MustAsset("bindata/groundfloor.bin"))
	AllMaps[maps.GROUNDFLOOR], _ = maps.Deserialize(r)
	r = bytes.NewReader(MustAsset("bindata/firstfloor.bin"))
	AllMaps[maps.FIRSTFLOOR], _ = maps.Deserialize(r)
	r = bytes.NewReader(MustAsset("bindata/roof.bin"))
	AllMaps[maps.ROOF], _ = maps.Deserialize(r)
	r = bytes.NewReader(MustAsset("bindata/athletics.bin"))
	AllMaps[maps.ATHLETICS], _ = maps.Deserialize(r)
}

func init() {
	LoadMaps()
}

func main() {
	termbox.Init()
	defer termbox.Close()

	playing := true

	gamedata := NewGame()
	player := gamedata.Chars[gamedata.PlayerId]

	for playing {
		DrawScreen(player)
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				playing = false
			case termbox.KeyArrowRight:
				MovePlayer(1, 0, player)
			case termbox.KeyArrowLeft:
				MovePlayer(-1, 0, player)
			case termbox.KeyArrowDown:
				MovePlayer(0, 1, player)
			case termbox.KeyArrowUp:
				MovePlayer(0, -1, player)
			case termbox.KeyPgdn:
				player.Loc.MapNum = (player.Loc.MapNum + 1) % len(AllMaps)
			}
		}
	}
}
