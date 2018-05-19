package main

import (
	"encoding/json"
	"strings"

	"github.com/japanoise/gocademy/characters"
	termbox "github.com/nsf/termbox-go"
)

// Handling elements

func GetSelectionLists(smap map[characters.Id]*characters.Element) ([]string, []characters.Id) {
	retSelList := []string{}
	retKeyList := []characters.Id{}
	for key, value := range smap {
		retSelList = append(retSelList, value.Name)
		retKeyList = append(retKeyList, key)
	}
	return retSelList, retKeyList
}

func DrawElement(x, y int, e *characters.Element, c termbox.Attribute) {
	for iy, tileList := range e.Pic.Data {
		for ix, tile := range tileList {
			if tile.C == " " && tile.Fg == termbox.ColorDefault && tile.Bg == tile.Fg {
				continue
			}
			if tile.Fg == termbox.ColorBlack {
				termbox.SetCell(x+ix, y+iy, rune(tile.C[0]), c, tile.Bg)
			} else if tile.Bg == termbox.ColorBlack {
				termbox.SetCell(x+ix, y+iy, rune(tile.C[0]), tile.Fg, c)
			} else {
				termbox.SetCell(x+ix, y+iy, rune(tile.C[0]), tile.Fg, tile.Bg)
			}
		}
	}
}

func ElementMerge(c *characters.Character, elem *characters.Element, col termbox.Attribute) {
	for iy, tileList := range elem.Pic.Data {
		for ix, tile := range tileList {
			if tile.C == " " && tile.Fg == termbox.ColorDefault && tile.Bg == tile.Fg {
				continue
			}
			if tile.Fg == termbox.ColorBlack {
				c.Face.Data[iy][ix].C = tile.C
				c.Face.Data[iy][ix].Fg = col
				c.Face.Data[iy][ix].Bg = tile.Bg
			} else if tile.Bg == termbox.ColorBlack {
				c.Face.Data[iy][ix].C = tile.C
				c.Face.Data[iy][ix].Fg = tile.Fg
				c.Face.Data[iy][ix].Bg = col
			} else {
				c.Face.Data[iy][ix].C = tile.C
				c.Face.Data[iy][ix].Fg = tile.Fg
				c.Face.Data[iy][ix].Bg = tile.Bg
			}
		}
	}
}

func LoadElements() {
	FrontHair = make(map[characters.Id]*characters.Element)
	BackHair = make(map[characters.Id]*characters.Element)
	HairAccessory = make(map[characters.Id]*characters.Element)
	HairAccessory["none"] = characters.GetEmptyElement()
	TopicalDetail = make(map[characters.Id]*characters.Element)
	TopicalDetail["none"] = characters.GetEmptyElement()
	for _, fname := range AssetNames() {
		if strings.HasPrefix(fname, "bindata/fronthair/") {
			e := characters.Element{}
			err := json.Unmarshal(MustAsset(fname), &e)
			if err != nil {
				panic(err)
			}
			FrontHair[e.ID] = &e
		} else if strings.HasPrefix(fname, "bindata/backhair/") {
			e := characters.Element{}
			err := json.Unmarshal(MustAsset(fname), &e)
			if err != nil {
				panic(err)
			}
			BackHair[e.ID] = &e
		} else if strings.HasPrefix(fname, "bindata/hairacc/") {
			e := characters.Element{}
			err := json.Unmarshal(MustAsset(fname), &e)
			if err != nil {
				panic(err)
			}
			HairAccessory[e.ID] = &e
		} else if strings.HasPrefix(fname, "bindata/topical/") {
			e := characters.Element{}
			err := json.Unmarshal(MustAsset(fname), &e)
			if err != nil {
				panic(err)
			}
			TopicalDetail[e.ID] = &e
		}
	}
}
