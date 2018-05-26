package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/japanoise/gocademy/maps"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	fn := os.Args[1]
	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	m, err := maps.Deserialize(file)
	if err != nil {
		panic(err)
	}
	file.Close()
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			tile, _ := m.TileAt(x, y)
			if tile == maps.GRASS {
				m.SetTileAt(x, y, maps.NewTile(randomRune([]rune{'`', ',', '.'}, rand), termbox.ColorGreen, true, true, 0x00))
			}
		}
	}

	writeMapToBinFile(fn, m)
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

func randomRune(strings []rune, rand *rand.Rand) rune {
	return strings[rand.Intn(len(strings))]
}
