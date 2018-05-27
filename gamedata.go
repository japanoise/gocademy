package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
	termutil "github.com/japanoise/termbox-util"
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

	AckMessage("New Game", "Let's start by generating the player character!")
	player := CharGen(ret)
	player.Loc = &characters.Location{5, 5, maps.GROUNDFLOOR}
	player.ID = ret.PlayerId
	ret.Chars[ret.PlayerId] = player

	done := false
	chars := []*characters.Character{player}
	lc := len(chars)
	charsW := termutil.RunewidthStr(player.GetNameString()) + 4
	rfunc := func(sx, sy int) {
		for i := 0; i < charsW; i++ {
			termutil.Printstring("-", (sx-charsW)+i, 1)
			termutil.Printstring("-", (sx-charsW)+i, 2+lc)
		}
		for i, char := range chars {
			termutil.Printstring("| |", sx-charsW, 2+i)
			DrawChar((sx-charsW)+1, 2+i, char)
			termutil.Printstring(char.GetNameString(), (sx-charsW)+3, 2+i)
			termutil.Printstring("|", sx-1, 2+i)
		}
	}
	for !done {
		choice := termutil.ChoiceIndexCallback("New game", []string{"New character", "Edit character", "Done"}, 0, func(_, sx, sy int) { rfunc(sx, sy) })
		switch choice {
		case 0:
			newCh := CharGen(ret)
			newCh.ID = ret.GetNextId()
			newCh.Loc = ret.NextSpawnPoint()
			ret.Chars[newCh.ID] = newCh
			chars = append(chars, newCh)
			lc = len(chars)
			cw := termutil.RunewidthStr(newCh.GetNameString()) + 4
			if cw > charsW {
				charsW = cw
			}
		case 1:
			charCh := make([]string, lc)
			for i, char := range chars {
				charCh[i] = char.GetNameString()
			}
			which := termutil.ChoiceIndexCallback("Edit which character", charCh, 0, func(_, sx, sy int) { rfunc(sx, sy) })
			CharEdit(ret, chars[which])
		default:
			done = true
		}
	}

	ret.genStudents(NUMSTUDENTS)
	return ret
}

func (g *Gamedata) genStudents(numStud int) {
	enames, bnames, gnames, surnames := LoadNames()
	rand := getNewRand()
	backhairids := make([]characters.Id, 0, len(BackHair))
	for key := range BackHair {
		backhairids = append(backhairids, key)
	}
	fronthairids := make([]characters.Id, 0, len(FrontHair))
	for key := range FrontHair {
		fronthairids = append(fronthairids, key)
	}
	for i := 0; i < numStud; i++ {
		student := RandChar(g, rand, enames, bnames, gnames, surnames, backhairids, fronthairids)
		g.Chars[student.ID] = student
		student.Loc = g.NextSpawnPoint()
	}
}

func (g *Gamedata) GetNextId() characters.Id {
	ret := characters.Id(strconv.Itoa(g.IdNum))
	g.IdNum++
	return ret
}

func (g *Gamedata) NextSpawnPoint() *characters.Location {
	return &characters.Location{1, g.IdNum, maps.GROUNDFLOOR}
}

func (g *Gamedata) GetCharacterIds() []string {
	choices := make([]string, 0, len(g.Chars))
	for key := range g.Chars {
		if key != g.PlayerId {
			choices = append(choices, string(key))
		}
	}
	return choices
}

func getNewRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
