package main

import "github.com/nsf/termbox-go"

func main() {
	termbox.Init()
	defer termbox.Close()

	DrawTile(0x0c2e, 0, 0)
	DrawTile(0x0dac, 0, 1)
	DrawTile(0x0023, 0, 2)
	DrawTile(0xc02b, 0, 3)
	DrawTile(0xcc2b, 0, 4)
	termbox.Flush()

	termbox.PollEvent()
}
