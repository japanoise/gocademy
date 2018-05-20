package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/gocademy/maps"
	"github.com/nsf/termbox-go"
)

const (
	NUMOFMAPS int = 4
)

var (
	AllMaps       []*maps.Map
	Colors        []string            = []string{"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}
	ColorsTermbox []termbox.Attribute = []termbox.Attribute{termbox.ColorBlack, termbox.ColorRed, termbox.ColorGreen, termbox.ColorYellow, termbox.ColorBlue, termbox.ColorMagenta, termbox.ColorCyan, termbox.ColorWhite}
	ConfigDir     string
	DataDir       string
	FrontHair     map[characters.Id]*characters.Element
	BackHair      map[characters.Id]*characters.Element
	HairAccessory map[characters.Id]*characters.Element
	TopicalDetail map[characters.Id]*characters.Element
)

func LoadMaps() {
	AllMaps = make([]*maps.Map, NUMOFMAPS)
	r := bytes.NewReader(MustAsset("bindata/groundfloor.bin"))
	AllMaps[maps.GROUNDFLOOR], _ = maps.Deserialize(r)
	r = bytes.NewReader(MustAsset("bindata/firstfloor.bin"))
	AllMaps[maps.FIRSTFLOOR], _ = maps.Deserialize(r)
	r = bytes.NewReader(MustAsset("bindata/roof.bin"))
	AllMaps[maps.ROOF], _ = maps.Deserialize(r)
	r = bytes.NewReader(MustAsset("bindata/athletics.bin"))
	AllMaps[maps.ATHLETICS], _ = maps.Deserialize(r)
}

func LoadNames() ([]string, []string, []string, []string) {
	enbynames := []string{}
	err := json.Unmarshal(MustAsset("bindata/enbynames.json"), &enbynames)
	if err != nil {
		panic(err)
	}
	boynames := []string{}
	json.Unmarshal(MustAsset("bindata/boynames.json"), &boynames)
	girlnames := []string{}
	json.Unmarshal(MustAsset("bindata/girlnames.json"), &girlnames)
	surnames := []string{}
	json.Unmarshal(MustAsset("bindata/surnames.json"), &surnames)
	return enbynames, boynames, girlnames, surnames
}

func init() {
	LoadMaps()
	var err error
	ConfigDir, DataDir, err = initDirs("gocademy")
	if err != nil {
		panic(err)
	}
	LoadElements()
}

func main() {
	termbox.Init()
	defer termbox.Close()

	playing, gamedata := TitleScreen()

	if !playing {
		return
	}

	var target *characters.Character = nil
	var message = ""
	player := gamedata.Chars[gamedata.PlayerId]
	charmaps := constructCharMaps(gamedata)

	for playing {
		DrawScreen(charmaps[player.Loc.MapNum], player, message)
		message = ""
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				playing, message = PauseMenu(gamedata)
			case termbox.KeyArrowRight:
				target, message = MovePlayer(1, 0, player, charmaps)
			case termbox.KeyArrowLeft:
				target, message = MovePlayer(-1, 0, player, charmaps)
			case termbox.KeyArrowDown:
				target, message = MovePlayer(0, 1, player, charmaps)
			case termbox.KeyArrowUp:
				target, message = MovePlayer(0, -1, player, charmaps)
			case termbox.KeySpace:
				message = fmt.Sprint(player.Loc)
			}
		}
		if target != nil {
			message = target.GetNameString()
		}
	}
}
