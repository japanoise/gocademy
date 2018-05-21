package maps

import (
	"encoding/gob"
	"errors"
	"io"
)

type Map struct {
	Tiles   []Tile
	pathers []*Pather
	Size    int
	Width   int
	Height  int
}

func NewMap(w, h int) *Map {
	size := w * h
	t := make([]Tile, size)
	return &Map{t, nil, size, w, h}
}

func DemoMap(w, h int) *Map {
	ret := NewMap(w, h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ret.Tiles[ret.indexOf(x, y)] = Tile(0x20 + x + y)
		}
	}
	return ret
}

func (m *Map) TileAt(x, y int) (Tile, error) {
	index := m.indexOf(x, y)
	if m.indexOutOfBounds(index) {
		return 0, errors.New("Out of bounds")
	}
	return m.Tiles[index], nil
}

func (m *Map) SetTileAt(x, y int, t Tile) {
	index := m.indexOf(x, y)
	if !m.indexOutOfBounds(index) {
		m.Tiles[index] = t
	}
}

func (m *Map) OutOfBounds(x, y int) bool {
	return x < 0 || x >= m.Width || y < 0 || y >= m.Height
}

func (m *Map) indexOf(x, y int) int {
	return (y * m.Width) + x
}

func (m *Map) indexOutOfBounds(index int) bool {
	return index < 0 || index >= m.Size
}

func (m *Map) DrawMap(startx, starty, width, height int) {
	for iy := 0; iy < height; iy++ {
		y := starty + iy
		for ix := 0; ix < width; ix++ {
			x := startx + ix
			if !m.OutOfBounds(x, y) {
				index := m.indexOf(x, y)
				if !m.indexOutOfBounds(index) {
					DrawTile(m.Tiles[index], ix, iy)
				}
			}
		}
	}
}

func (m *Map) GetPather(x, y int) *Pather {
	if m.pathers == nil {
		m.pathers = make([]*Pather, len(m.Tiles))
	}
	index := m.indexOf(x, y)
	if m.indexOutOfBounds(index) {
		return nil
	}
	if m.pathers[index] == nil {
		m.pathers[index] = &Pather{x, y, m.Tiles[index], m}
	}
	return m.pathers[index]
}

func (m *Map) Serialize(w io.Writer) error {
	enc := gob.NewEncoder(w)
	return enc.Encode(*m)
}

func Deserialize(r io.Reader) (*Map, error) {
	dec := gob.NewDecoder(r)
	ret := &Map{}
	err := dec.Decode(ret)
	return ret, err
}
