package main

import (
	"bytes"

	"github.com/japanoise/gocademy/maps"
	"github.com/nsf/termbox-go"
)

const (
	NUMOFMAPS int = 4
)

var (
	AllMaps []*maps.Map
)

func LoadMaps() {
	AllMaps = make([]*maps.Map, NUMOFMAPS)
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
	charmaps := constructCharMaps(gamedata)

	for playing {
		DrawScreen(charmaps[player.Loc.MapNum], player)
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				playing = false
			case termbox.KeyArrowRight:
				MovePlayer(1, 0, player, charmaps[player.Loc.MapNum])
			case termbox.KeyArrowLeft:
				MovePlayer(-1, 0, player, charmaps[player.Loc.MapNum])
			case termbox.KeyArrowDown:
				MovePlayer(0, 1, player, charmaps[player.Loc.MapNum])
			case termbox.KeyArrowUp:
				MovePlayer(0, -1, player, charmaps[player.Loc.MapNum])
			case termbox.KeyPgdn:
				player.Loc.MapNum = (player.Loc.MapNum + 1) % len(AllMaps)
				// This is suboptimal (may cause OOB panic), but it's the fast way for now.
				charmaps = constructCharMaps(gamedata)
			}
		}
	}
}
