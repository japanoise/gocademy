package main

import (
	"github.com/japanoise/gocademy/maps"
	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()

	m := maps.DemoMap(80,24)
	m.DrawMap(-5,-5,80,24)
	termbox.Flush()

	termbox.PollEvent()
}
