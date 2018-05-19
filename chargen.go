package main

import (
	"encoding/json"
	"math/rand"

	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/termbox-util"
	termbox "github.com/nsf/termbox-go"
)

var pronounStrings []string = []string{
	characters.GenderFmt(characters.ENBY, characters.ENBY, "%s/%o/%r"),
	characters.GenderFmt(characters.MALE, characters.ENBY, "%s/%o/%r"),
	characters.GenderFmt(characters.FEMALE, characters.ENBY, "%s/%o/%r"),
}

func GetChargenRefreshFunc(c *characters.Character) func(int, int) {
	return func(sx, sy int) {
		name := c.GetNameString()
		pronouns := pronounStrings[c.Gender]
		termutil.Printstring("Name:", sx-5, 1)
		termutil.Printstring(name, sx-(len(name)), 2)
		termutil.Printstring("Pronouns:", sx-9, 3)
		termutil.Printstring(pronouns, sx-(len(pronouns)), 4)
		c.Face.DrawAscii(sx-12, 5)
	}
}

// Generate a character
func CharGen(g *Gamedata) *characters.Character {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	ret := &characters.Character{}
	err := json.Unmarshal(MustAsset("bindata/face.json"), &ret.Face)
	if err != nil {
		panic(err)
	}
	rfunc := GetChargenRefreshFunc(ret)

	for ret.GivenName == "" {
		ret.GivenName = termutil.Prompt("Character's given name?", rfunc)
	}
	for ret.Surname == "" {
		ret.Surname = termutil.Prompt(ret.GivenName+"'s family name?", rfunc)
	}

	ret.Gender = characters.CGender(termutil.ChoiceIndexCallback(ret.GivenName+"'s pronouns?", pronounStrings, 0, func(choice int, sx int, sy int) { ret.Gender = characters.CGender(choice); rfunc(sx, sy) }))

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	mouth := termutil.Prompt(ret.GivenName+"'s mouth shape? (e.g. _, w, v, u)", rfunc)
	if mouth == "" {
		mouth = "-"
	}
	ret.Face.Data[5][5].C = mouth

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	eye := termutil.Prompt(ret.GivenName+"'s left eye shape? (currently ^)", rfunc)
	if eye == "" {
		eye = "^"
	}
	ret.Face.Data[3][4].C = eye

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	eye = termutil.Prompt(ret.GivenName+"'s right eye shape? (currently *)", rfunc)
	if eye == "" {
		eye = "^"
	}
	ret.Face.Data[3][6].C = eye

	ret.HairColor = ColorsTermbox[termutil.ChoiceIndexCallback(ret.GivenName+"'s hair color?", Colors, 0, func(_, sx, sy int) { rfunc(sx, sy) })]

	eyecolor := ColorsTermbox[termutil.ChoiceIndexCallback(ret.GivenName+"'s eye color?", Colors, 0, func(choice, sx, sy int) {
		col := ColorsTermbox[choice]
		ret.Face.Data[3][4].Fg = col
		ret.Face.Data[3][6].Fg = col
		if col == termbox.ColorBlack || col == termbox.ColorWhite {
			ret.Face.Data[3][4].Bg = termbox.ColorRed
			ret.Face.Data[3][6].Bg = termbox.ColorRed
		} else {
			ret.Face.Data[3][4].Bg = termbox.ColorDefault
			ret.Face.Data[3][6].Bg = termbox.ColorDefault
		}
		rfunc(sx, sy)
	})]
	ret.EyeColor = eyecolor

	ret.TopicalDetailId = elemSelec(TopicalDetail, "topical detail(s)", ret, rfunc)
	ret.BackHairId = elemSelec(BackHair, "back hair", ret, rfunc)
	ret.FrontHairId = elemSelec(FrontHair, "front hair (bangs)", ret, rfunc)
	ret.HairAccessoryId = elemSelec(HairAccessory, "hair accessory/detail", ret, rfunc)

	ret.ID = g.GetNextId()
	ret.GenSprite()
	return ret
}

func elemSelec(smap map[characters.Id]*characters.Element, toSelec string, ret *characters.Character, rfunc func(int, int)) characters.Id {
	fronthairs, fronthairIds := GetSelectionLists(smap)
	retId := fronthairIds[termutil.ChoiceIndexCallback("Select "+ret.GivenName+"'s "+toSelec, fronthairs, 0, func(choice, sx, sy int) {
		rfunc(sx, sy)
		DrawElement(sx-12, 5, smap[fronthairIds[choice]], ret.HairColor)
	})]
	ElementMerge(ret, smap[retId], ret.HairColor)
	return retId
}

func RandChar(g *Gamedata, rand *rand.Rand, enbynames, boynames, girlnames, surnames []string) *characters.Character {
	ret := &characters.Character{}
	ret.Gender = randGender(rand)
	switch ret.Gender {
	case characters.ENBY:
		ret.GivenName = randomString(enbynames, rand)
	case characters.MALE:
		ret.GivenName = randomString(boynames, rand)
	case characters.FEMALE:
		ret.GivenName = randomString(girlnames, rand)
	}
	ret.Surname = randomString(surnames, rand)
	ret.HairColor = randomColor(rand)
	ret.EyeColor = randomColor(rand)
	ret.ID = g.GetNextId()
	ret.GenSprite()
	return ret
}

func randGender(rand *rand.Rand) characters.CGender {
	return characters.CGender(rand.Intn(3))
}

func randomString(strings []string, rand *rand.Rand) string {
	return strings[rand.Intn(len(strings))]
}

func randomColor(rand *rand.Rand) termbox.Attribute {
	return termbox.Attribute(2 + rand.Intn(6))
}
