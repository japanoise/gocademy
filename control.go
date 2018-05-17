package main

import (
	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
)

func MovePlayer(dx, dy int, player *characters.Character) {
	dest, err := AllMaps[player.Loc.MapNum].TileAt(player.Loc.X+dx, player.Loc.Y+dy)
	if err == nil && maps.IsPassable(dest) {
		player.Loc.X += dx
		player.Loc.Y += dy
	}
}
