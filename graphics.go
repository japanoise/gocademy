package main

import (
	"github.com/japanoise/gocademy/characters"
	termbox "github.com/nsf/termbox-go"
)

func DrawChar(x, y int, char *characters.Character) {
	if char != nil {
		termbox.SetCell(x, y, '@', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func DrawChars(sx, sy, originx, originy int, charmaps []*charmap, player *characters.Character) {
	for x := 0; x < sx; x++ {
		for y := 0; y < sy; y++ {
			if charmaps[player.Loc.MapNum].inBounds(originx+x, originy+y) {
				DrawChar(x, y, charmaps[player.Loc.MapNum].data[originx+x][originy+y])
			}
		}
	}
}

func DrawScreen(charmaps []*charmap, player *characters.Character) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	sx, sy := termbox.Size()
	originx, originy := player.Loc.X-(sx/2), player.Loc.Y-(sy/2)
	AllMaps[player.Loc.MapNum].DrawMap(originx, originy, sx, sy)
	DrawChars(sx, sy, originx, originy, charmaps, player)
	termbox.Flush()
}
