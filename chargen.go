package main

import (
	"encoding/json"
	"fmt"
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

func ApplyEyeColor(col termbox.Attribute, ret *characters.Character) {
	ret.Face.Data[3][4].Fg = col
	ret.Face.Data[3][6].Fg = col
	if col == termbox.ColorBlack || col == termbox.ColorWhite {
		ret.Face.Data[3][4].Bg = termbox.ColorRed
		ret.Face.Data[3][6].Bg = termbox.ColorRed
	} else {
		ret.Face.Data[3][4].Bg = termbox.ColorDefault
		ret.Face.Data[3][6].Bg = termbox.ColorDefault
	}
}

// Generate a character. This function doesn't set an ID so do that yourself, fuccboi.
func CharGen(g *Gamedata) *characters.Character {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	ret := &characters.Character{}
	err := json.Unmarshal(MustAsset("bindata/face.json"), &ret.Face)
	if err != nil {
		panic(err)
	}
	rfunc := GetChargenRefreshFunc(ret)

	namePrompt(ret, rfunc)

	selPronouns(ret, rfunc)

	mouthSel(ret, rfunc)

	eyePrompt(ret, rfunc)

	selHairColor(ret, rfunc)

	selEyeColor(ret, rfunc)

	ret.TopicalDetailId = elemSelec(TopicalDetail, "topical detail(s)", ret, rfunc)
	ret.BackHairId = elemSelec(BackHair, "back hair", ret, rfunc)
	ret.FrontHairId = elemSelec(FrontHair, "front hair (bangs)", ret, rfunc)
	ret.HairAccessoryId = elemSelec(HairAccessory, "hair accessory/detail", ret, rfunc)

	ret.ID = g.GetNextId()
	ret.GenSprite()
	return ret
}

func namePrompt(ret *characters.Character, rfunc func(int, int)) {
	for ret.GivenName == "" {
		ret.GivenName = termutil.Prompt("Character's given name?", rfunc)
	}
	for ret.Surname == "" {
		ret.Surname = termutil.Prompt(ret.GivenName+"'s family name?", rfunc)
	}
}

func selPronouns(ret *characters.Character, rfunc func(int, int)) {
	ret.Gender = characters.CGender(termutil.ChoiceIndexCallback(ret.GivenName+"'s pronouns?", pronounStrings, 0, func(choice int, sx int, sy int) { ret.Gender = characters.CGender(choice); rfunc(sx, sy) }))
}

func mouthSel(ret *characters.Character, rfunc func(int, int)) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	ret.Mouth = termutil.Prompt(ret.GivenName+"'s mouth shape? (e.g. _, w, v, u)", rfunc)
	if ret.Mouth == "" {
		ret.Mouth = ret.Face.Data[5][5].C
	}
	ret.Face.Data[5][5].C = ret.Mouth
}

func eyePrompt(ret *characters.Character, rfunc func(int, int)) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	ret.Leye = termutil.Prompt(fmt.Sprintf("%s's left eye shape? (currently %s)", ret.GivenName, ret.Face.Data[3][4].C), rfunc)
	if ret.Leye == "" {
		ret.Leye = ret.Face.Data[3][4].C
	}
	ret.Face.Data[3][4].C = ret.Leye

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	ret.Reye = termutil.Prompt(fmt.Sprintf("%s's right eye shape? (currently %s)", ret.GivenName, ret.Face.Data[3][6].C), rfunc)
	if ret.Reye == "" {
		ret.Reye = ret.Face.Data[3][6].C
	}
	ret.Face.Data[3][6].C = ret.Reye
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

func selHairColor(ret *characters.Character, rfunc func(int, int)) {
	ret.HairColor = ColorsTermbox[termutil.ChoiceIndexCallback(ret.GivenName+"'s hair color?", Colors, 0, func(_, sx, sy int) { rfunc(sx, sy) })]
}

func selEyeColor(ret *characters.Character, rfunc func(int, int)) {
	eyecolor := ColorsTermbox[termutil.ChoiceIndexCallback(ret.GivenName+"'s eye color?", Colors, 0, func(choice, sx, sy int) {
		col := ColorsTermbox[choice]
		ApplyEyeColor(col, ret)
		rfunc(sx, sy)
	})]
	ret.EyeColor = eyecolor
}

// Edit a character.
func CharEdit(g *Gamedata, c *characters.Character) {
	rfunc := GetChargenRefreshFunc(c)
	done := false
	for !done {
		choice := termutil.ChoiceIndexCallback("What do you want to do with "+c.GetNameString(), []string{
			"Change name",
			"Change pronouns",
			"Change mouth",
			"Change eyes",
			"Change hair color",
			"Change eye color",
			"Change front hair",
			"Change back hair",
			"Change topical details",
			"Change hair accessory",
			"Done",
		}, 0, func(_, sx, sy int) { rfunc(sx, sy) })
		switch choice {
		case 0:
			namePrompt(c, rfunc)
		case 1:
			selPronouns(c, rfunc)
		case 2:
			mouthSel(c, rfunc)
		case 3:
			eyePrompt(c, func(sx, sy int) {
				rebuildPortrait(c)
				rfunc(sx, sy)
			})
			rebuildPortrait(c)
		case 4:
			selHairColor(c, func(sx, sy int) {
				rebuildPortrait(c)
				rfunc(sx, sy)
			})
			rebuildPortrait(c)
		case 5:
			selEyeColor(c, func(sx, sy int) {
				rebuildPortrait(c)
				rfunc(sx, sy)
			})
			rebuildPortrait(c)
		case 6:
			c.TopicalDetailId = elemSelec(TopicalDetail, "topical detail(s)", c, func(sx, sy int) {
				rebuildPortrait(c)
				rfunc(sx, sy)
			})
			rebuildPortrait(c)
		case 7:
			c.BackHairId = elemSelec(BackHair, "back hair", c, func(sx, sy int) {
				rebuildPortrait(c)
				rfunc(sx, sy)
			})
			rebuildPortrait(c)
		case 8:
			c.FrontHairId = elemSelec(FrontHair, "front hair (bangs)", c, func(sx, sy int) {
				rebuildPortrait(c)
				rfunc(sx, sy)
			})
			rebuildPortrait(c)
		case 9:
			c.HairAccessoryId = elemSelec(HairAccessory, "hair accessory/detail", c, func(sx, sy int) {
				rebuildPortrait(c)
				rfunc(sx, sy)
			})
			rebuildPortrait(c)
		default:
			c.GenSprite()
			done = true
		}
	}
}

func rebuildPortrait(c *characters.Character) {
	err := json.Unmarshal(MustAsset("bindata/face.json"), &c.Face)
	if err != nil {
		panic(err)
	}
	c.Face.Data[3][4].C = c.Leye
	c.Face.Data[3][6].C = c.Reye
	c.Face.Data[5][5].C = c.Mouth
	ApplyEyeColor(c.EyeColor, c)
	ElementMerge(c, TopicalDetail[c.TopicalDetailId], c.HairColor)
	ElementMerge(c, BackHair[c.BackHairId], c.HairColor)
	ElementMerge(c, FrontHair[c.FrontHairId], c.HairColor)
	ElementMerge(c, HairAccessory[c.HairAccessoryId], c.HairColor)
}

func RandChar(g *Gamedata, rand *rand.Rand, enbynames, boynames, girlnames, surnames []string, backhairids, fronthairids []characters.Id) *characters.Character {
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

	err := json.Unmarshal(MustAsset("bindata/face.json"), &ret.Face)
	if err != nil {
		panic(err)
	}

	eyes := []string{"^", "@", "I", "*", "~", "O", "n", "="}
	eye := randomString(eyes, rand)
	ret.Face.Data[3][4].C = eye
	ret.Face.Data[3][6].C = eye
	mouths := []string{"w", "^", "~", "v", "u", "3", "_", "-"}
	ApplyEyeColor(ret.EyeColor, ret)

	ret.Face.Data[5][5].C = randomString(mouths, rand)
	ret.BackHairId = randomId(backhairids, rand)
	ElementMerge(ret, BackHair[ret.BackHairId], ret.HairColor)
	ret.FrontHairId = randomId(fronthairids, rand)
	ElementMerge(ret, FrontHair[ret.FrontHairId], ret.HairColor)
	return ret
}

func randGender(rand *rand.Rand) characters.CGender {
	return characters.CGender(rand.Intn(3))
}

func randomString(strings []string, rand *rand.Rand) string {
	return strings[rand.Intn(len(strings))]
}

func randomId(strings []characters.Id, rand *rand.Rand) characters.Id {
	return strings[rand.Intn(len(strings))]
}

func randomColor(rand *rand.Rand) termbox.Attribute {
	return termbox.Attribute(2 + rand.Intn(6))
}
