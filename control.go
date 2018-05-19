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

func MovePlayer(dx, dy int, player *characters.Character, pcMap *charmap) *characters.Character {
	destX, destY := player.Loc.X+dx, player.Loc.Y+dy
	dest, err := AllMaps[player.Loc.MapNum].TileAt(destX, destY)
	if err == nil && maps.IsPassable(dest) {
		target := pcMap.moveNoCollide(player.Loc.X, player.Loc.Y, destX, destY)
		if target == nil {
			player.Loc.X = destX
			player.Loc.Y = destY
		}
		return target
	}
	return nil
}