package maps

import (
	"testing"

	astar "github.com/beefsack/go-astar"
)

func TestPath(testing *testing.T) {
	m := Map{[]Tile{GRASS, GRASS, GRASS, GRASS, GRASS, GRASS, GRASS, GRASS, GRASS}, nil, 9, 3, 3}
	astar.Path(m.GetPather(0, 0), m.GetPather(1, 1))
}
