package main

import (
	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
)

func MovePlayer(dx, dy int, player *characters.Character, pcMap *charmap) {
	destX, destY := player.Loc.X+dx, player.Loc.Y+dy
	dest, err := AllMaps[player.Loc.MapNum].TileAt(destX, destY)
	if err == nil && maps.IsPassable(dest) && pcMap.moveNoCollide(player.Loc.X, player.Loc.Y, destX, destY) == nil {
		player.Loc.X = destX
		player.Loc.Y = destY
	}
}
