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
	player := CharGen(ret)
	player.Loc = characters.Location{5, 5, maps.GROUNDFLOOR}
	ret.Chars[ret.PlayerId] = player
	return ret
}

func (g *Gamedata) GetNextId() characters.Id {
	ret := characters.Id(strconv.Itoa(g.IdNum))
	g.IdNum++
	return ret
}
