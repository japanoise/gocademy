package main

import (
	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
	"github.com/japanoise/termbox-util"
	termbox "github.com/nsf/termbox-go"
)

func DrawChar(x, y int, char *characters.Character) {
	if char.HairColor == termbox.ColorBlack || char.HairColor == termbox.ColorWhite {
		termbox.SetCell(x, y, char.Sprite, char.HairColor, termbox.ColorRed)
	} else {
		termbox.SetCell(x, y, char.Sprite, char.HairColor, termbox.ColorDefault)
	}
}

func DrawChars(sx, sy, originx, originy int, m *maps.Map, charmap *charmap) {
	for ix := 0; ix < sx; ix++ {
		for iy := 0; iy < sy; iy++ {
			x, y := originx+ix, originy+iy
			if charmap.inBounds(x, y) {
				if charmap.data[x][y] == nil {
					tile, _ := m.TileAt(x, y)
					maps.DrawTile(tile, ix, iy)
				} else {
					DrawChar(ix, iy, charmap.data[x][y])
				}
			}
		}
	}
}

func DrawScreen(charmap *charmap, player *characters.Character, message string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	sx, sy := termbox.Size()
	originx, originy := player.Loc.X-(sx/2), player.Loc.Y-(sy/2)
	DrawChars(sx, sy, originx, originy, AllMaps[player.Loc.MapNum], charmap)
	if message != "" {
		termutil.Printstring(message, 0, sy-1)
	}
	termbox.Flush()
}
