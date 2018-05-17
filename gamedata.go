package main

import (
	"strconv"

	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
)

type Gamedata struct {
	Chars    map[characters.Id]*characters.Character
	PlayerId characters.Id
	IdNum    int
}

func NewGame() *Gamedata {
	ret := &Gamedata{}
	ret.Chars = make(map[characters.Id]*characters.Character)
	ret.PlayerId = ret.GetNextId()
	player := characters.NewCharacter("Player", "One", ret.PlayerId)
	player.Loc = characters.Location{5, 5, maps.GROUNDFLOOR}
	ret.Chars[ret.PlayerId] = player
	return ret
}

func (g *Gamedata) GetNextId() characters.Id {
	return characters.Id(strconv.Itoa(g.IdNum))
}