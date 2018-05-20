package main

import (
	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
)

type charmap struct {
	width  int
	height int
	data   charmapData
}

type charmapData [][]*characters.Character

func constructCharMaps(g *Gamedata) []*charmap {
	ret := make([]*charmap, NUMOFMAPS)
	ret[maps.GROUNDFLOOR] = makeCharMap(AllMaps[maps.GROUNDFLOOR])
	ret[maps.FIRSTFLOOR] = makeCharMap(AllMaps[maps.FIRSTFLOOR])
	ret[maps.ROOF] = makeCharMap(AllMaps[maps.ROOF])
	ret[maps.ATHLETICS] = makeCharMap(AllMaps[maps.ATHLETICS])
	for _, char := range g.Chars {
		ret[char.Loc.MapNum].data[char.Loc.X][char.Loc.Y] = char
	}
	return ret
}

// Construct a charmap for the given maps.Map
func makeCharMap(m *maps.Map) *charmap {
	ret := make(charmapData, m.Width)
	for i := 0; i < m.Width; i++ {
		ret[i] = make([]*characters.Character, m.Height)
	}
	return &charmap{m.Width, m.Height, ret}
}

func (c *charmap) inBounds(x, y int) bool {
	return x < c.width && y < c.height && x >= 0 && y >= 0
}

// Attempts to move entity at sourcex,sourcey to destx,desty; returns non-nil if there's something there
func (c *charmap) moveNoCollide(sourceX, sourceY, destX, destY int) *characters.Character {
	if c.inBounds(destX, destY) {
		target := c.data[destX][destY]
		if target == nil {
			c.data[destX][destY] = c.data[sourceX][sourceY]
			c.data[sourceX][sourceY] = nil
		}
		return target
	} else {
		return nil
	}
}

// Jumps the source character to the destination character. Doesn't test bounds, so be responsible.
func jumpMap(sourcex, sourcey int, sourcec *charmap, destx, desty int, destc *charmap) {
	destc.data[destx][desty] = sourcec.data[sourcex][sourcey]
	sourcec.data[sourcex][sourcey] = nil
}
