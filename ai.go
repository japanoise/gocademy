package main

import (
	"fmt"

	astar "github.com/beefsack/go-astar"
	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
)

func GenPathToTarget(x, y int, c *characters.Character) string {
	loc := c.Loc
	m := AllMaps[loc.MapNum]
	from := m.GetPather(loc.X, loc.Y)
	to := m.GetPather(x, y)
	mypath, _, found := astar.Path(from, to)
	if found {
		lmp := len(mypath)
		c.Path = make([]*maps.Pather, lmp)
		for i, p := range mypath {
			c.Path[lmp-1-i] = p.(*maps.Pather)
		}
		c.Path = c.Path[1:]
	}
	return fmt.Sprint(c.Path)
}

func Act(c *characters.Character, cmap *charmap) {
	if c.Path != nil {
		tile, _ := AllMaps[c.Loc.MapNum].TileAt(c.Path[0].X, c.Path[0].Y)
		if maps.IsDoor(tile) && !maps.IsPassable(tile) {
			AllMaps[c.Loc.MapNum].SetTileAt(c.Path[0].X, c.Path[0].Y, maps.OpenDoor(tile))
		} else {
			target := cmap.moveNoCollide(c.Loc.X, c.Loc.Y, c.Path[0].X, c.Path[0].Y)
			if target == nil {
				c.Loc.X = c.Path[0].X
				c.Loc.Y = c.Path[0].Y
				if len(c.Path) == 1 {
					c.Path = nil
				} else {
					c.Path = c.Path[1:]
				}
			}
		}
	}
}
