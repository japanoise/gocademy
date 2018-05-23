package main

import (
	"fmt"

	astar "github.com/beefsack/go-astar"
	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
)

func GenPathToTarget(x, y, mapNum int, c *characters.Character) string {
	loc := c.Loc
	m := AllMaps[loc.MapNum]
	from := m.GetPather(loc.X, loc.Y)
	var to *maps.Pather
	if mapNum == loc.MapNum {
		to = m.GetPather(x, y)
	} else {
		to = Warps[warpId(loc.MapNum, mapNum)]
	}
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

func Act(c *characters.Character, cmaps []*charmap) {
	if c.Path != nil {
		tile, _ := AllMaps[c.Loc.MapNum].TileAt(c.Path[0].X, c.Path[0].Y)
		if maps.IsDoor(tile) && !maps.IsPassable(tile) {
			AllMaps[c.Loc.MapNum].SetTileAt(c.Path[0].X, c.Path[0].Y, maps.OpenDoor(tile))
		} else {
			target := cmaps[c.Loc.MapNum].moveNoCollide(c.Loc.X, c.Loc.Y, c.Path[0].X, c.Path[0].Y)
			if target == nil || target == c {
				c.Loc.X = c.Path[0].X
				c.Loc.Y = c.Path[0].Y
				if len(c.Path) == 1 {
					// Check warp points
					for warpID, warpPoint := range Warps {
						if warpPoint.X == c.Path[0].X && warpPoint.Y == c.Path[0].Y {
							newMapId := getWarpDest(warpID)
							endPather := Warps[warpId(newMapId, c.Loc.MapNum)]
							x, y := endPather.X, endPather.Y
							if cmaps[newMapId].data[x][y] == nil {
								jumpMap(c.Loc.X, c.Loc.Y, cmaps[c.Loc.MapNum], x, y, cmaps[newMapId])
								c.Loc.X = x
								c.Loc.Y = y
								c.Loc.MapNum = newMapId
								c.Path = nil
							}
							return // Quit early; we'll wait at the warp point
						}
					}
					// Evidently we're not trying to warp, so we're at our destination.
					c.Path = nil
				} else {
					c.Path = c.Path[1:]
				}
			}
		}
	}
}
