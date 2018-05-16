package main

import (
	"os"

	"github.com/japanoise/gocademy/maps"
	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()

	file, _ := os.Open("map.bin")
	m, _ := maps.Deserialize(file)

	playing := true
	x, y := 0, 0
	for playing {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		sx, sy := termbox.Size()
		m.DrawMap(x, y, sx, sy)
		termbox.Flush()
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				playing = false
			case termbox.KeyArrowRight:
				x++
			case termbox.KeyArrowLeft:
				x--
			case termbox.KeyArrowDown:
				y++
			case termbox.KeyArrowUp:
				y--
			}
		}
	}
}
