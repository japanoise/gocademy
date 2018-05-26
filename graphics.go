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

func AckMessage(title, msg string) {
	looping := true
	mh := 7
	sw := termutil.RunewidthStr(msg)
	wiTitle := termutil.RunewidthStr(title)
	if wiTitle > sw {
		sw = wiTitle
	}
	mw := sw + 8
	if mw < 16 {
		mw = 16
	} else if mw%2 != 0 {
		mw++
	}
	for looping {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		sx, sy := termbox.Size()
		anc := (sx - (mw)) / 2
		yanc := (sy - mh) / 2
		for i := yanc; i < yanc+mh; i++ {
			termutil.PrintRune(anc, i, '|', termbox.ColorDefault)
			termutil.PrintRune(anc+mw, i, '|', termbox.ColorDefault)
		}
		for i := anc; i <= anc+mw; i++ {
			termutil.PrintRune(i, yanc, ' ', termbox.AttrReverse)
			termutil.PrintRune(i, yanc+mh, '-', termbox.ColorDefault)
		}
		termutil.PrintstringColored(termbox.AttrReverse|termbox.AttrBold, "[x]", anc, yanc)
		termutil.PrintstringColored(termbox.AttrReverse|termbox.AttrBold, title, anc+((mw-12)/2), yanc)
		termutil.Printstring(msg, anc+((mw-sw)/2), yanc+2)
		termutil.Printstring("------", anc+((mw-6)/2), yanc+4)
		termutil.Printstring("| OK |", anc+((mw-6)/2), yanc+5)
		termutil.Printstring("------", anc+((mw-6)/2), yanc+6)
		termbox.Flush()
		ev := termbox.PollEvent()
		looping = ev.Type != termbox.EventKey
	}
}
