package main

import (
	"github.com/japanoise/gocademy/characters"
	termbox "github.com/nsf/termbox-go"
)

func DrawScreen(player *characters.Character) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	sx, sy := termbox.Size()
	AllMaps[player.Loc.MapNum].DrawMap(player.Loc.X-(sx/2), player.Loc.Y-(sy/2), sx, sy)
	termbox.SetCell(sx/2, sy/2, '@', termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}
