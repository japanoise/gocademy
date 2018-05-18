package main

import (
	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
	termbox "github.com/nsf/termbox-go"
)

func DrawChar(x, y int, char *characters.Character) {
	termbox.SetCell(x, y, '@', termbox.ColorDefault, termbox.ColorDefault)
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

func DrawScreen(charmap *charmap, player *characters.Character) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	sx, sy := termbox.Size()
	originx, originy := player.Loc.X-(sx/2), player.Loc.Y-(sy/2)
	DrawChars(sx, sy, originx, originy, AllMaps[player.Loc.MapNum], charmap)
	termbox.Flush()
}
