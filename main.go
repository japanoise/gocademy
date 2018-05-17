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
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		sx, sy := termbox.Size()
		AllMaps[player.Loc.MapNum].DrawMap(player.Loc.X-(sx/2), player.Loc.Y-(sy/2), sx, sy)
		termbox.SetCell(sx/2, sy/2, '@', termbox.ColorDefault, termbox.ColorDefault)
		termbox.Flush()
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				playing = false
			case termbox.KeyArrowRight:
				player.Loc.X++
			case termbox.KeyArrowLeft:
				player.Loc.X--
			case termbox.KeyArrowDown:
				player.Loc.Y++
			case termbox.KeyArrowUp:
				player.Loc.Y--
			case termbox.KeyPgdn:
				player.Loc.MapNum = (player.Loc.MapNum + 1) % len(AllMaps)
			}
		}
	}
}
