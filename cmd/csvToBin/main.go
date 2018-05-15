package main

import (
	"encoding/csv"
	"github.com/japanoise/gocademy/maps"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"strconv"
)

func csvToMap(fn string) *maps.Map {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvr := csv.NewReader(file)
	records, err := csvr.ReadAll()
	height := len(records)
	width := len(records[0])
	log.Println("height:", height, "width:", width)
	ret := maps.NewMap(width, height)
	for y, rec := range records {
		for x, t := range rec {
			log.Println(x, y, t)
			tile, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal(err)
			}
			switch tile {
			case 1:
				// grass
				ret.SetTileAt(x, y, maps.Tile(0x0dac))
			case 2:
				// wall
				ret.SetTileAt(x, y, maps.Tile(0x0023))
			case 3:
				// concrete floor
				ret.SetTileAt(x, y, maps.NewTile('.', termbox.ColorCyan, true, true, 0x00))
			case 4:
				// interior floor
				ret.SetTileAt(x, y, maps.Tile(0x0c2e))
			case 5:
				// door
				ret.SetTileAt(x, y, maps.Tile(0xc02b))
			case 6:
				// Desk
				ret.SetTileAt(x, y, maps.NewTile('[', termbox.ColorRed, true, false, 0x00))
			case 7:
				// chair
				ret.SetTileAt(x, y, maps.NewTile('\\', termbox.ColorRed, true, true, 0x00))
			case 8:
				// lockers
				ret.SetTileAt(x, y, maps.NewTile('%', termbox.ColorCyan, false, false, 0x00))
			case 9:
				// shoe cupboard
				ret.SetTileAt(x, y, maps.NewTile('%', termbox.ColorRed, false, false, 0x00))
			case 10:
				// stairs leading up
				ret.SetTileAt(x, y, maps.Tile(0x0c3c))
			case 11:
				// stairs leading down
				ret.SetTileAt(x, y, maps.Tile(0x0c3e))
			case 12:
				// toilet stall
				ret.SetTileAt(x, y, maps.NewTile(']', termbox.ColorCyan, true, false, 0x00))
			case 13:
				// sink
				ret.SetTileAt(x, y, maps.NewTile('#', termbox.ColorCyan, true, false, 0x00))
			case 14:
				// window
				ret.SetTileAt(x, y, maps.NewTile('#', termbox.ColorBlue, true, false, 0x00))
			case 15:
				// bookcase
				ret.SetTileAt(x, y, maps.NewTile('#', termbox.ColorGreen, false, false, 0x00))
			case 16:
				// fence
				ret.SetTileAt(x, y, maps.NewTile('#', termbox.ColorRed, true, false, 0x00))
			case 17:
				// gate
				ret.SetTileAt(x, y, maps.NewTile('/', termbox.ColorRed, true, true, 0x00))
			case 18:
				// track
				ret.SetTileAt(x, y, maps.NewTile('.', termbox.ColorRed, true, true, 0x00))
			case 19:
				// shower
				ret.SetTileAt(x, y, maps.NewTile('\\', termbox.ColorCyan, true, true, 0x00))
			case 20:
				// swimming pool
				ret.SetTileAt(x, y, maps.NewTile('~', termbox.ColorBlue, true, true, 0x00))
			case 21:
				// pool coping
				ret.SetTileAt(x, y, maps.NewTile('.', termbox.ColorYellow, true, true, 0x00))
			case 22:
				// Male sign
				ret.SetTileAt(x, y, maps.Tile('M'))
			case 23:
				// Female sign
				ret.SetTileAt(x, y, maps.Tile('F'))
			case 24:
				// bed
				ret.SetTileAt(x, y, maps.NewTile('#', termbox.ColorYellow, true, true, 0x00))
			case 25:
				// curtain
				ret.SetTileAt(x, y, maps.NewTile('{', termbox.ColorYellow, false, false, 0x0b))
			case 26:
				// computer
				ret.SetTileAt(x, y, maps.NewTile('=', termbox.ColorDefault, true, false, 0x00))
			case 27:
				// curtain rail
				ret.SetTileAt(x, y, maps.NewTile('|', termbox.ColorCyan, true, false, 0x00))
			case 28:
				// railing
				ret.SetTileAt(x, y, maps.NewTile('%', termbox.ColorDefault, false, false, 0x00))
			default:
				// void
				ret.SetTileAt(x, y, maps.Tile(0x0020))
			}
		}
	}
	log.Println(*ret)
	return ret
}

func writeMapToBinFile(fn string, m *maps.Map) {
	file, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = m.Serialize(file)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ", os.Args[0], " file1 file2")
	}
	fn := os.Args[1]
	m := csvToMap(fn)
	fn = os.Args[2]
	writeMapToBinFile(fn, m)
}
