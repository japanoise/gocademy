package maps

import (
	"testing"
	"github.com/nsf/termbox-go"
)

func TestRoguelikeGrass(t *testing.T) {
	tile := NewTile('.', termbox.ColorGreen, true, true, 0)
	if GetColor(tile) != termbox.ColorGreen {
		t.Error("isn't green")
	}
	if GetSprite(tile) != '.' {
		t.Error("wrong sprite")
	}
	if !IsPassable(tile) {
		t.Error("not passable")
	}
	if !IsSeeThrough(tile) {
		t.Error("not see through")
	}
	if IsDoor(tile) {
		t.Error("is a door")
	}
}

func TestRoguelikeFloor(t *testing.T) {
	tile := NewTile('.', termbox.ColorDefault, true, true, 0)
	if GetSprite(tile) != '.' {
		t.Error("wrong sprite")
	}
	if !IsPassable(tile) {
		t.Error("not passable")
	}
	if !IsSeeThrough(tile) {
		t.Error("not see through")
	}
	if IsDoor(tile) {
		t.Error("is a door")
	}
}

func TestRoguelikeDoorClosed(t *testing.T) {
	tile := NewTile('+', termbox.ColorDefault, false, false, 0x0c)
	if GetSprite(tile) != '+' {
		t.Error("wrong sprite")
	}
	if IsPassable(tile) {
		t.Error("passable")
	}
	if IsSeeThrough(tile) {
		t.Error("see through")
	}
	if !IsDoor(tile) {
		t.Error("isn't a door")
	}
}

func TestRoguelikeDoorOpen(t *testing.T) {
	tile := NewTile('+', termbox.ColorDefault, true, true, 0x0c)
	if GetSprite(tile) != '/' {
		t.Error("wrong sprite")
	}
	if !IsPassable(tile) {
		t.Error("not passable")
	}
	if !IsSeeThrough(tile) {
		t.Error("not see through")
	}
	if !IsDoor(tile) {
		t.Error("isn't a door")
	}
}
