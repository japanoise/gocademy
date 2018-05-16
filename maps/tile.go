package maps

import "github.com/nsf/termbox-go"

/*
Ladies and gentlemen, a roguelike tile stored within 16 bits.

Supports seven colours, ascii for tile sprite, passable & transparency
as separate values, and very rudimentary doors.

DOOO PVCC CTTT TTTT
T = sprite data
C = color data
V = seethrough?
P = passable/door closed?
O = offset (+ ascii in T) if door open
D = door?

eg.
0x0c2e = roguelike floor (white passable seethrough . that isn't a door)
0x0dac = roguelike grass (green passable seethrough , that isn't a door)
0x0023 = roguelike wall (white non-passable solid # that isn't a door)
0xc02b = roguelike closed door (white non-passable solid + that's a door with offset 4 - a /)
0xcc2b = roguelike open door (white passable seethrough + that's a door with offset 4 - drawn as a /)
*/
type Tile uint16

// Masks for extracting data from tiles
const (
	SPRITEMASK   Tile = 0x7f
	COLORMASK         = 0x07
	PASSMASK          = 0x800
	VISMASK           = 0x400
	DOORMASK          = 0x8000
	DOORSPRMASK       = 0x7000
	OPENDOORMASK      = 0x8800
)

func NewTile(sprite rune, color termbox.Attribute, seethrough bool, passable bool, DoorData byte) Tile {
	var ret Tile = 0
	ret |= Tile(sprite & rune(SPRITEMASK))
	ret |= (Tile(color) & COLORMASK) << 7
	if seethrough {
		ret |= VISMASK
	}
	if passable {
		ret |= PASSMASK
	}
	ret |= Tile(DoorData&0x0F) << 12
	return ret
}

// Extract the ascii from the tile; or the open door offset, if appropriate.
func GetSprite(t Tile) rune {
	if OPENDOORMASK&t == OPENDOORMASK {
		//door open
		return rune((t & SPRITEMASK) + ((t & DOORSPRMASK) >> 12))
	}
	return rune(t & SPRITEMASK)
}

// Get the color data stored in the tile (as a termbox color)
func GetColor(t Tile) termbox.Attribute {
	return termbox.Attribute((t >> 7) & COLORMASK)
}

// Is the tile passable?
func IsPassable(t Tile) bool {
	return t&PASSMASK != 0
}

// Can you see through the tile?
func IsSeeThrough(t Tile) bool {
	return t&VISMASK != 0
}

// Is the tile a door?
func IsDoor(t Tile) bool {
	return t&DOORMASK != 0
}

// Open the door
func OpenDoor(t Tile) Tile {
	if IsDoor(t) {
		t = t &^ (VISMASK)
		t = t &^ (PASSMASK)
	}
	return t
}

// Close the door
func CloseDoor(t Tile) Tile {
	if IsDoor(t) {
		t |= (VISMASK)
		t |= (PASSMASK)
	}
	return t
}

// Draw given tile at x, y on the screen
func DrawTile(t Tile, x, y int) {
	termbox.SetCell(x, y, GetSprite(t), GetColor(t), termbox.ColorDefault)
}
