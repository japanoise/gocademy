package main

import (
	"github.com/japanoise/gocademy/maps"
	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()

	maps.DrawTile(0x0c2e, 0, 0)
	maps.DrawTile(0x0dac, 0, 1)
	maps.DrawTile(0x0023, 0, 2)
	maps.DrawTile(0xc02b, 0, 3)
	maps.DrawTile(0xcc2b, 0, 4)
	termbox.Flush()

	termbox.PollEvent()
}
