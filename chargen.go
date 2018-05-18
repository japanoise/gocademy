package main

import (
	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/termbox-util"
)

var pronounStrings []string = []string{
	characters.GenderFmt(characters.ENBY, characters.ENBY, "%s/%o/%r"),
	characters.GenderFmt(characters.MALE, characters.ENBY, "%s/%o/%r"),
	characters.GenderFmt(characters.FEMALE, characters.ENBY, "%s/%o/%r"),
}

func GetChargenRefreshFunc(c *characters.Character) func(int, int) {
	return func(sx, sy int) {
		name := c.Surname + " " + c.GivenName
		pronouns := pronounStrings[c.Gender]
		termutil.Printstring("Name:", sx-5, 1)
		termutil.Printstring(name, sx-(len(name)), 2)
		termutil.Printstring("Pronouns:", sx-9, 3)
		termutil.Printstring(pronouns, sx-(len(pronouns)), 4)
	}
}

// Generate a character
func CharGen(g *Gamedata) *characters.Character {
	ret := &characters.Character{}
	rfunc := GetChargenRefreshFunc(ret)
	for ret.GivenName == "" {
		ret.GivenName = termutil.Prompt("Character's given name?", rfunc)
	}
	for ret.Surname == "" {
		ret.Surname = termutil.Prompt(ret.GivenName+"'s family name?", rfunc)
	}
	ret.Gender = characters.CGender(termutil.ChoiceIndexCallback("Character's pronouns?", pronounStrings, 0, func(choice int, sx int, sy int) { ret.Gender = characters.CGender(choice); rfunc(sx, sy) }))
	ret.ID = g.GetNextId()
	return ret
}
