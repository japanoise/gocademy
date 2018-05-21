package maps

import (
	"fmt"

	"github.com/beefsack/go-astar"
)

type Pather struct {
	X, Y int
	T    Tile
	M    *Map `json:"-"` // It's safe not to export this; it'll only be used at calculation time.
}

func (p *Pather) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := p.M.GetPather(p.X+offset[0], p.Y+offset[1]); n != nil &&
			!IsWall(p.T) {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (t *Pather) PathNeighborCost(to astar.Pather) float64 {
	return 1.0
}

func (t *Pather) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Pather)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func (p *Pather) String() string {
	return fmt.Sprintf("{%d, %d}", p.X, p.Y)
}
