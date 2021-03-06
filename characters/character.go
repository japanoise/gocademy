package characters

import (
	"unicode/utf8"

	"github.com/japanoise/gocademy/maps"
	asciiart "github.com/japanoise/termbox-asciiart"
	termbox "github.com/nsf/termbox-go"
)

type Character struct {
	GivenName       string
	Surname         string
	Gender          CGender
	Relationships   map[Id]*Relation
	ID              Id
	PersonalityId   Id
	Loc             *Location
	Face            asciiart.Ascii
	Sprite          rune
	FrontHairId     Id
	BackHairId      Id
	HairAccessoryId Id
	TopicalDetailId Id
	HairColor       termbox.Attribute
	EyeColor        termbox.Attribute
	Path            []*maps.Pather
	Target          Id
	Leye            string `json:"-"`
	Reye            string `json:"-"`
	Mouth           string `json:"-"`
	CurrentMood     Mood
	PartnerId       Id
}

func (c *Character) SayDialogue(id Id) string {
	return c.GetNameString() + ": " + c.GetDialogue(id)
}

func (c *Character) GetDialogue(id Id) string {
	if Personalities[c.PersonalityId] != nil {
		return Personalities[c.PersonalityId].GetDialogue(id)
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
