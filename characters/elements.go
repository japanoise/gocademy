package characters

import (
	asciiart "github.com/japanoise/termbox-asciiart"
)

type Element struct {
	Name string
	Pic  *asciiart.Ascii
	ID   Id
}

func GetEmptyElement() *Element {
	return &Element{
		"None",
		&asciiart.Ascii{[][]asciiart.AsciiTile{}},
		"none",
	}
}
