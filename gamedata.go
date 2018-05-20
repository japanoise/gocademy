package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
)

const NUMSTUDENTS = 10

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
	enames, bnames, gnames, surnames := LoadNames()
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	backhairids := make([]characters.Id, 0, len(BackHair))
	for key := range BackHair {
		backhairids = append(backhairids, key)
	}
	fronthairids := make([]characters.Id, 0, len(FrontHair))
	for key := range FrontHair {
		fronthairids = append(fronthairids, key)
	}
	for i := 0; i < NUMSTUDENTS; i++ {
		student := RandChar(ret, rand, enames, bnames, gnames, surnames, backhairids, fronthairids)
		ret.Chars[student.ID] = student
		student.Loc = ret.NextSpawnPoint()
	}
	return ret
}

func (g *Gamedata) GetNextId() characters.Id {
	ret := characters.Id(strconv.Itoa(g.IdNum))
	g.IdNum++
	return ret
}

func (g *Gamedata) NextSpawnPoint() characters.Location {
	return characters.Location{1, g.IdNum, maps.GROUNDFLOOR}
}
