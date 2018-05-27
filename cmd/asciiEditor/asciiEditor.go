package main

import (
	"encoding/json"
	"os"
	"strconv"

	asciiart "github.com/japanoise/termbox-asciiart"
	"github.com/japanoise/termbox-util"
	"github.com/nsf/termbox-go"
)

const (
	WIDTH        int = 11
	HEIGHT       int = 9
	INSTRUCTIONS int = WIDTH + 9
	PALLETTE     int = HEIGHT + 4
)

func drawScreen(a *asciiart.Ascii, cbg, cfg termbox.Attribute, cx, cy int) {
	termbox.SetCursor(3+cx, 2+cy)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termutil.Printstring("01234567890", 3, 0)
	termutil.Printstring("+-----------+", 2, 1)
	for i := 0; i < HEIGHT; i++ {
		termutil.Printstring(strconv.Itoa(i), 0, 2+i)
		termbox.SetCell(2, 2+i, '|', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(WIDTH+3, 2+i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}
	termutil.Printstring("+-----------+", 2, HEIGHT+2)
	termutil.PrintstringColored(240, "   _____", 3, 3)
	termutil.PrintstringColored(240, "  / _ _ \\", 3, 4)
	termutil.PrintstringColored(240, "  | ^ * |", 3, 5)
	termutil.PrintstringColored(240, "  |     |", 3, 6)
	termutil.PrintstringColored(240, "  \\  v  /", 3, 7)
	termutil.PrintstringColored(240, "   \\___/", 3, 8)
	termutil.PrintstringColored(240, "  __| |__", 3, 9)
	termutil.PrintstringColored(240, " |   v   |", 3, 10)
	a.DrawAsciiNoClobber(3, 2)

	termutil.Printstring("gocademy ascii editor", INSTRUCTIONS, 0)
	termutil.Printstring("^C          - Exit", INSTRUCTIONS, 2)
	termutil.Printstring("^F, →       - Cursor forward", INSTRUCTIONS, 3)
	termutil.Printstring("^B, ←       - Cursor back", INSTRUCTIONS, 4)
	termutil.Printstring("^P, ↑       - Cursor up", INSTRUCTIONS, 5)
	termutil.Printstring("^N, ↓       - Cursor down", INSTRUCTIONS, 6)
	termutil.Printstring("^A, HOME    - Start of line", INSTRUCTIONS, 7)
	termutil.Printstring("^E, END     - End of line", INSTRUCTIONS, 8)
	termutil.Printstring("^S          - Save to file", INSTRUCTIONS, 9)
	termutil.Printstring("^D, DEL, BS - Delete char at point", INSTRUCTIONS, 10)
	termutil.Printstring("^J          - Next BG color", INSTRUCTIONS, 11)
	termutil.Printstring("^K          - Next FG color", INSTRUCTIONS, 12)
	termutil.Printstring("^Q          - Erase your work", INSTRUCTIONS, 13)
	termutil.Printstring("Other characters self insert.", INSTRUCTIONS, 15)

	termutil.Printstring("bg color", 0, PALLETTE)
	termbox.SetCell(int(cbg)*2, PALLETTE+1, '>', termbox.ColorDefault, termbox.ColorDefault)
	termutil.Printstring("fg color", 0, PALLETTE+2)
	termbox.SetCell(int(cfg)*2, PALLETTE+3, '>', termbox.ColorDefault, termbox.ColorDefault)
	for i := 0; i < 9; i++ {
		termbox.SetCell(1+(2*i), HEIGHT+5, ' ', termbox.ColorDefault, termbox.Attribute(i))
		termbox.SetCell(1+(2*i), HEIGHT+7, ' ', termbox.Attribute(i)|termbox.AttrReverse, termbox.ColorDefault)
	}

	termbox.Flush()
}

func resetArt() *asciiart.Ascii {
	art := &asciiart.Ascii{make([][]asciiart.AsciiTile, HEIGHT)}
	for i := range art.Data {
		art.Data[i] = make([]asciiart.AsciiTile, WIDTH)
		for j := range art.Data[i] {
			art.Data[i][j] = asciiart.AsciiTile{C: " ", Fg: termbox.ColorDefault, Bg: termbox.ColorDefault}
		}
	}
	return art
}

func main() {
	art := resetArt()
	termbox.Init()
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputMouse)

	cx, cy := 0, 0
	cbg, cfg := termbox.ColorDefault, termbox.ColorDefault
	for {
		drawScreen(art, cbg, cfg, cx, cy)
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventMouse {
			x, y := ev.MouseX, ev.MouseY
			if x >= 3 && x <= 14 && y >= 2 && y <= 10 {
				cx, cy = x-3, y-2
			} else if y >= PALLETTE && y <= PALLETTE+3 && x < 18 {
				col := (x - (x % 2)) / 2
				if y == PALLETTE || y == PALLETTE+1 {
					cbg = termbox.Attribute(col)
				} else {
					cfg = termbox.Attribute(col)
				}
			}
		} else if ev.Type == termbox.EventKey {
			if ev.Ch == 0 {
				switch ev.Key {
				case termbox.KeyCtrlC:
					return
				case termbox.KeyArrowLeft, termbox.KeyCtrlB:
					if cx > 0 {
						cx--
					}
				case termbox.KeyArrowRight, termbox.KeyCtrlF:
					if cx < WIDTH-1 {
						cx++
					}
				case termbox.KeyArrowUp, termbox.KeyCtrlP:
					if cy > 0 {
						cy--
					}
				case termbox.KeyArrowDown, termbox.KeyCtrlN:
					if cy < HEIGHT-1 {
						cy++
					}
				case termbox.KeyCtrlA, termbox.KeyHome:
					cx = 0
				case termbox.KeyCtrlE, termbox.KeyEnd:
					cx = WIDTH - 1
				case termbox.KeyCtrlJ:
					cbg = (cbg + 1) % 9
				case termbox.KeyCtrlK:
					cfg = (cfg + 1) % 9
				case termbox.KeyCtrlS:
					fn := termutil.Prompt("Filename?", func(int, int) {
						drawScreen(art, cbg, cfg, cx, cy)
					})
					if fn != "" {
						file, err := os.Create(fn)
						if err != nil {
							termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
							termutil.Printstring(err.Error(), 0, 0)
							termbox.Flush()
							termbox.PollEvent()
						} else {
							enc := json.NewEncoder(file)
							enc.Encode(art)
							file.Close()
						}
					}
				case termbox.KeyBackspace, termbox.KeyBackspace2, termbox.KeyDelete, termbox.KeyCtrlD:
					art.Data[cy][cx] = asciiart.AsciiTile{C: " ", Fg: termbox.ColorDefault, Bg: termbox.ColorDefault}
				case termbox.KeySpace:
					art.Data[cy][cx] = asciiart.AsciiTile{C: " ", Fg: cfg, Bg: cbg}
				case termbox.KeyCtrlQ:
					art = resetArt()
				}
			} else {
				art.Data[cy][cx] = asciiart.AsciiTile{C: string(ev.Ch), Fg: cfg, Bg: cbg}
			}
		}
	}
}
