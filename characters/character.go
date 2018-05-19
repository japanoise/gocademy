package characters

import (
	"unicode/utf8"

	asciiart "github.com/japanoise/termbox-asciiart"
	termbox "github.com/nsf/termbox-go"
)

type Character struct {
	GivenName       string
	Surname         string
	Gender          CGender
	Relationships   map[Id]Relation
	ID              Id
	PersonalityId   Id
	Loc             Location
	Face            asciiart.Ascii
	Sprite          rune
	FrontHairId     Id
	BackHairId      Id
	HairAccessoryId Id
	TopicalDetailId Id
	HairColor       termbox.Attribute
	EyeColor        termbox.Attribute
}

func (c *Character) GetDialogue(id Id) string {
	if Personalities[id] != nil {
		return Personalities[id].GetDialogue(id)
	} else {
		return Default.GetDialogue(id)
	}
}

func (c *Character) GetNameString() string {
	return c.Surname + " " + c.GivenName
}

func NewCharacter(forename, surname string, id Id) *Character {
	ret := &Character{}
	ret.GivenName = forename
	ret.Surname = surname
	ret.GenSprite()
	return ret
}

func (c *Character) GenSprite() {
	c.Sprite, _ = utf8.DecodeRuneInString(c.GivenName)
}
