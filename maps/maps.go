package maps

import (
	"errors"
	"log"
	"os"
)

type Map struct {
	tiles []Tile
	size int
	Width int
	Height int
}

func NewMap(w, h int) *Map {
	size := w*h
	t := make([]Tile, size)
	return &Map{t, size, w, h}
}

func DemoMap(w, h int) *Map {
	ret := NewMap(w,h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ret.tiles[ret.indexOf(x, y)] = Tile(0x20+x+y)
		}
	}
	return ret
}

func (m *Map) TileAt(x, y int) (Tile, error) {
	index := m.indexOf(x, y)
	if m.indexOutOfBounds(index) {
		return 0, errors.New("Out of bounds")
	}
	return m.tiles[index], nil
}

func (m *Map) OutOfBounds(x, y int) bool {
	return x < 0 || x >= m.Width || y < 0 || y >= m.Height
}

func (m *Map) indexOf(x, y int) int {
	return (y*m.Width) + x
}

func (m *Map) indexOutOfBounds(index int) bool {
	return index < 0 || index >= m.size
}

func (m *Map) DrawMap(startx, starty, width, height int) {
	for iy := 0; iy < height; iy++ {
		y := starty + iy
		for ix := 0; ix < width; ix++ {
			x := startx+ix
			if !m.OutOfBounds(x, y) {
				index := m.indexOf(x, y)
				if !m.indexOutOfBounds(index) {
					DrawTile(m.tiles[index], ix, iy)
				}
			}
		}
	}
}
