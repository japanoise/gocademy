package main

import (
	"fmt"
	"math/rand"

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
		to = findWarp(loc.MapNum, mapNum)
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

func findWarp(from, to int) *maps.Pather {
	ret := Warps[warpId(from, to)]
	if ret == nil {
		switch from {
		case maps.ROOF:
			// Roof's only exit.
			return Warps[warpId(maps.ROOF, maps.FIRSTFLOOR)]
		case maps.FIRSTFLOOR:
			// The first floor lacks a warp to the athletics field, so return the G warp.
			return Warps[warpId(maps.FIRSTFLOOR, maps.GROUNDFLOOR)]
		case maps.GROUNDFLOOR:
			// G floor lacks a warp to the roof, so return the staircase warp.
			return Warps[warpId(maps.GROUNDFLOOR, maps.FIRSTFLOOR)]
		case maps.ATHLETICS:
			// Athletics field's only exit.
			return Warps[warpId(maps.ATHLETICS, maps.GROUNDFLOOR)]
		}
	}
	return ret
}

// Main AI entry point. Called every turn.
func Act(g *Gamedata, rand *rand.Rand, c *characters.Character, cmaps []*charmap) {
	if c.Path == nil {
		actOnWarpPoints(c, cmaps)
		if c.Target == "" {
			switch c.CurrentMood {
			case characters.MoodAngry:
				// Nothing yet. Just stew for now.
			case characters.MoodSad:
				if c.PartnerId == "" {
					// Mope by yourself
					goToQuietArea(g, rand, c)
				} else {
					// Find partner for comfort
					c.Target = c.PartnerId
				}
			case characters.MoodBlush:
				if c.PartnerId == "" {
					// Nothing yet. Find crush to talk to, etc.
				} else {
					// Find partner to talk to.
					c.Target = c.PartnerId
				}
			case characters.MoodLewd:
				if c.PartnerId == "" {
					// find somewhere quiet to... y'know.
					goToQuietArea(g, rand, c)
				} else {
					// find partner for relief :lenny:
					c.Target = c.PartnerId
				}
			default:
				// For now, just loiter.
			}
		} else {
			tc := g.Chars[c.Target]
			GenPathToTarget(tc.Loc.X, tc.Loc.Y, tc.Loc.MapNum, c)
		}
	} else if len(c.Path) == 0 {
		c.Path = nil
	} else {
		followPath(c, cmaps)
	}
}

// Warp the npc if they're standing on a warp point.
func actOnWarpPoints(c *characters.Character, cmaps []*charmap) {
	for warpID, warpPoint := range Warps {
		if warpPoint.X == c.Loc.X && warpPoint.Y == c.Loc.Y && c.Loc.MapNum == getWarpSource(warpID) {
			newMapId := getWarpDest(warpID)
			endPather := Warps[warpId(newMapId, c.Loc.MapNum)]
			x, y := endPather.X, endPather.Y
			jumpMap(c.Loc.X, c.Loc.Y, cmaps[c.Loc.MapNum], x, y, cmaps[newMapId])
			c.Loc.X = x
			c.Loc.Y = y
			c.Loc.MapNum = newMapId
			c.Path = nil
		}
	}
}

// Find the NPC somewhere quiet to go
func goToQuietArea(g *Gamedata, rand *rand.Rand, c *characters.Character) {
	switch c.Loc.MapNum {
	case maps.ATHLETICS:
		// Supply closet
		GenPathToTarget(162, 116, maps.ATHLETICS, c)
	case maps.GROUNDFLOOR:
		switch rand.Intn(4) {
		case 0:
			// West janitor closet
			GenPathToTarget(24, 33, maps.GROUNDFLOOR, c)
		case 1:
			// East janitor closet
			GenPathToTarget(114, 33, maps.GROUNDFLOOR, c)
		case 2:
			// East edge of the building
			GenPathToTarget(136, 54, maps.GROUNDFLOOR, c)
		case 3:
			// West edge of the building
			GenPathToTarget(1, 54, maps.GROUNDFLOOR, c)
		}
	case maps.FIRSTFLOOR:
		switch rand.Intn(2) {
		case 0:
			// West janitor closet
			GenPathToTarget(21, 1, maps.FIRSTFLOOR, c)
		case 1:
			// East janitor closet
			GenPathToTarget(111, 1, maps.FIRSTFLOOR, c)
		}
	case maps.ROOF:
		// There's not really anywhere quiet on the roof, so just go back down.
		warp := Warps[warpId(maps.ROOF, maps.FIRSTFLOOR)]
		GenPathToTarget(warp.X, warp.Y, maps.ROOF, c)
	}
}

func followPath(c *characters.Character, cmaps []*charmap) {
	tile, _ := AllMaps[c.Loc.MapNum].TileAt(c.Path[0].X, c.Path[0].Y)
	if maps.IsDoor(tile) && !maps.IsPassable(tile) {
		AllMaps[c.Loc.MapNum].SetTileAt(c.Path[0].X, c.Path[0].Y, maps.OpenDoor(tile))
	} else {
		target := cmaps[c.Loc.MapNum].moveNoCollide(c.Loc.X, c.Loc.Y, c.Path[0].X, c.Path[0].Y)
		targetWantsMySpace := false
		if target != nil && target.Path != nil && len(target.Path) != 0 {
			targetWantsMySpace = target.Path[0].X == c.Loc.X && target.Path[0].Y == c.Loc.Y
		}
		if target == nil || target == c || targetWantsMySpace {
			if targetWantsMySpace {
				cmaps[c.Loc.MapNum].data[c.Loc.X][c.Loc.Y] = target
				cmaps[c.Loc.MapNum].data[c.Path[0].X][c.Path[0].Y] = c
				target.Loc.X = c.Loc.X
				target.Loc.Y = c.Loc.Y
			}
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
