package main

import (
	"encoding/json"

	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
	asciiart "github.com/japanoise/termbox-asciiart"
	termutil "github.com/japanoise/termbox-util"
)

func TitleScreen() (bool, *Gamedata) {
	building := &asciiart.Ascii{}
	err := json.Unmarshal(MustAsset("bindata/building.json"), building)
	if err != nil {
		panic(err)
	}
	message := ""
	for {
		choices := []string{"New game", "Load game", "Quit to terminal"}
		choice := termutil.ChoiceIndexCallback("gocademy", choices, 0, func(_, _, sy int) {
			building.DrawAscii(0, 5)
			if message != "" {
				termutil.Printstring(message, 0, sy-1)
			}
		})
		switch choice {
		case 0:
			return true, NewGame()
		case 1:
			g, err := LoadGameChoice()
			if err != nil {
				message = err.Error()
			} else {
				return true, g
			}
		case 2:
			return false, nil
		}
	}
}

func PauseMenu(g *Gamedata) (bool, string) {
	choices := []string{"Continue", "Save game", "Quit to terminal"}
	choice := termutil.ChoiceIndex("Paused", choices, 0)
	switch choice {
	case 0:
		return true, ""
	case 1:
		return true, "Saved: " + SaveGame(g)
	}
	return false, ""
}

func MovePlayer(dx, dy int, player *characters.Character, charmaps []*charmap) (*characters.Character, string) {
	destX, destY := player.Loc.X+dx, player.Loc.Y+dy
	dest, err := AllMaps[player.Loc.MapNum].TileAt(destX, destY)
	if err == nil && maps.IsPassable(dest) {
		target := charmaps[player.Loc.MapNum].moveNoCollide(player.Loc.X, player.Loc.Y, destX, destY)
		if target == nil {
			player.Loc.X = destX
			player.Loc.Y = destY
			if player.Loc.MapNum == maps.GROUNDFLOOR {
				if destY == 0 && destX >= 118 && destY <= 120 {
					// North gate
					return pJumpMap(player, maps.ATHLETICS, destX, destY, charmaps[player.Loc.MapNum], destX-13, 149, charmaps[maps.ATHLETICS])
				} else if destY >= 37 && destY <= 39 && destX >= 68 && destX <= 70 {
					// Central staircase
					return pJumpMap(player, maps.FIRSTFLOOR, destX, destY, charmaps[player.Loc.MapNum], destX-3, destY-32, charmaps[maps.FIRSTFLOOR])
				}
			} else if player.Loc.MapNum == maps.FIRSTFLOOR {
				if destX == 56 && destY == 27 {
					// Roof staircase
					return pJumpMap(player, maps.ROOF, destX, destY, charmaps[player.Loc.MapNum], player.Loc.X, player.Loc.Y, charmaps[maps.ROOF])
				} else if destY >= 5 && destY <= 7 && destX >= 65 && destX <= 67 {
					// Central staircase
					return pJumpMap(player, maps.GROUNDFLOOR, destX, destY, charmaps[player.Loc.MapNum], destX+3, destY+32, charmaps[maps.GROUNDFLOOR])
				}
			} else if player.Loc.MapNum == maps.ATHLETICS {
				// South gate
				if destY == 149 && destX >= 105 && destX <= 107 {
					return pJumpMap(player, maps.GROUNDFLOOR, destX, destY, charmaps[player.Loc.MapNum], destX+13, 0, charmaps[maps.GROUNDFLOOR])
				}
			} else if player.Loc.MapNum == maps.ROOF {
				if destX == 56 && destY == 27 {
					// Roof staircase
					return pJumpMap(player, maps.FIRSTFLOOR, destX, destY, charmaps[player.Loc.MapNum], player.Loc.X, player.Loc.Y, charmaps[maps.FIRSTFLOOR])
				}
			}
		}
		return target, ""
	} else if err == nil && maps.IsDoor(dest) {
		AllMaps[player.Loc.MapNum].SetTileAt(destX, destY, maps.OpenDoor(dest))
		return nil, "You open the door."
	}
	return nil, ""
}
