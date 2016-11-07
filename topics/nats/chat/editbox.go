// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"sync"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const editBoxWidth = 80
const editBoxHeight = 15

var (
	eb  editBox
	out []string
	mu  sync.Mutex
)

// =============================================================================

// WriteMessage adds a message to the current view of messages.
func WriteMessage(who, s string) {
	mu.Lock()
	{
		out = append(out, who+": "+s)
		if len(out) > editBoxHeight-2 {
			out = out[1:]
		}
	}
	mu.Unlock()
	redrawAll()
}

// Draw will draw the client UI and hold the application
// from terminating.
func Draw(event func(string)) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	redrawAll()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowLeft, termbox.KeyCtrlB:
				eb.MoveCursorOneRuneBackward()
			case termbox.KeyArrowRight, termbox.KeyCtrlF:
				eb.MoveCursorOneRuneForward()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				eb.DeleteRuneBackward()
			case termbox.KeyDelete, termbox.KeyCtrlD:
				eb.DeleteRuneForward()
			case termbox.KeyTab:
				eb.InsertRune('\t')
			case termbox.KeySpace:
				eb.InsertRune(' ')
			case termbox.KeyCtrlK:
				eb.DeleteTheRestOfTheLine()
			case termbox.KeyHome, termbox.KeyCtrlA:
				eb.MoveCursorToBeginningOfTheLine()
			case termbox.KeyEnd, termbox.KeyCtrlE:
				eb.MoveCursorToEndOfTheLine()
			case termbox.KeyEnter:
				event(string(eb.text))

				eb.MoveCursorToBeginningOfTheLine()
				eb.DeleteTheRestOfTheLine()
			default:
				if ev.Ch != 0 {
					eb.InsertRune(ev.Ch)
				}
			}

		case termbox.EventError:
			panic(ev.Err)
		}

		redrawAll()
	}
}

// =============================================================================

func redrawAll() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()

	midy := (h / 2) + (h / 3)
	midx := (w - editBoxWidth) / 2

	termbox.SetCell(midx-1, midy, '│', coldef, coldef)
	termbox.SetCell(midx+editBoxWidth, midy, '│', coldef, coldef)
	termbox.SetCell(midx-1, midy-1, '┌', coldef, coldef)
	termbox.SetCell(midx-1, midy+1, '└', coldef, coldef)
	termbox.SetCell(midx+editBoxWidth, midy-1, '┐', coldef, coldef)
	termbox.SetCell(midx+editBoxWidth, midy+1, '┘', coldef, coldef)
	fill(midx, midy-1, editBoxWidth, 1, termbox.Cell{Ch: '─'})
	fill(midx, midy+1, editBoxWidth, 1, termbox.Cell{Ch: '─'})

	eb.Draw(midx, midy, editBoxWidth, 1)
	termbox.SetCursor(midx+eb.CursorX(), midy)

	var i int
	var o int
	for i = 2; i < editBoxHeight; i++ {
		o++
		termbox.SetCell(midx-1, midy-i, '│', coldef, coldef)
		if len(out) >= o {
			tbprint(midx+1, midy-i, coldef, coldef, out[len(out)-o])
		}
		termbox.SetCell(midx+editBoxWidth, midy-i, '│', coldef, coldef)
	}
	fill(midx, midy-i, editBoxWidth, 1, termbox.Cell{Ch: '─'})

	tbprint(midx+32, midy+3, coldef, coldef, "Press ESC to quit")
	termbox.Flush()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func runeAdvanceLen(r rune, pos int) int {
	if r == '\t' {
		return tabstopLength - pos%tabstopLength
	}
	return runewidth.RuneWidth(r)
}

func voffsetCoffset(text []byte, boffset int) (voffset, coffset int) {
	text = text[:boffset]
	for len(text) > 0 {
		r, size := utf8.DecodeRune(text)
		text = text[size:]
		coffset++
		voffset += runeAdvanceLen(r, voffset)
	}
	return
}

func byteSliceGrow(s []byte, desiredCap int) []byte {
	if cap(s) < desiredCap {
		ns := make([]byte, len(s), desiredCap)
		copy(ns, s)
		return ns
	}
	return s
}

func byteSliceRemove(text []byte, from, to int) []byte {
	size := to - from
	copy(text[from:], text[to:])
	text = text[:len(text)-size]
	return text
}

func byteSliceInsert(text []byte, offset int, what []byte) []byte {
	n := len(text) + len(what)
	text = byteSliceGrow(text, n)
	text = text[:n]
	copy(text[offset+len(what):], text[offset:])
	copy(text[offset:], what)
	return text
}

// =============================================================================

const preferredHorizontalThreshold = 5
const tabstopLength = 8

type editBox struct {
	text          []byte
	lineVoffset   int
	cursorBoffset int // cursor offset in bytes
	cursorVoffset int // visual cursor offset in termbox cells
	cursorCoffset int // cursor offset in unicode code points
}

func (eb *editBox) Draw(x, y, w, h int) {
	eb.AdjustVOffset(w)

	const coldef = termbox.ColorDefault
	fill(x, y, w, h, termbox.Cell{Ch: ' '})

	t := eb.text
	lx := 0
	tabstop := 0
	for {
		rx := lx - eb.lineVoffset
		if len(t) == 0 {
			break
		}

		if lx == tabstop {
			tabstop += tabstopLength
		}

		if rx >= w {
			termbox.SetCell(x+w-1, y, '→',
				coldef, coldef)
			break
		}

		r, size := utf8.DecodeRune(t)
		if r == '\t' {
			for ; lx < tabstop; lx++ {
				rx = lx - eb.lineVoffset
				if rx >= w {
					goto next
				}

				if rx >= 0 {
					termbox.SetCell(x+rx, y, ' ', coldef, coldef)
				}
			}
		} else {
			if rx >= 0 {
				termbox.SetCell(x+rx, y, r, coldef, coldef)
			}
			lx += runewidth.RuneWidth(r)
		}
	next:
		t = t[size:]
	}

	if eb.lineVoffset != 0 {
		termbox.SetCell(x, y, '←', coldef, coldef)
	}
}

func (eb *editBox) AdjustVOffset(width int) {
	ht := preferredHorizontalThreshold
	maxHThreshold := (width - 1) / 2
	if ht > maxHThreshold {
		ht = maxHThreshold
	}

	threshold := width - 1
	if eb.lineVoffset != 0 {
		threshold = width - ht
	}
	if eb.cursorVoffset-eb.lineVoffset >= threshold {
		eb.lineVoffset = eb.cursorVoffset + (ht - width + 1)
	}

	if eb.lineVoffset != 0 && eb.cursorVoffset-eb.lineVoffset < ht {
		eb.lineVoffset = eb.cursorVoffset - ht
		if eb.lineVoffset < 0 {
			eb.lineVoffset = 0
		}
	}
}

func (eb *editBox) MoveCursorTo(boffset int) {
	eb.cursorBoffset = boffset
	eb.cursorVoffset, eb.cursorCoffset = voffsetCoffset(eb.text, boffset)
}

func (eb *editBox) RuneUnderCursor() (rune, int) {
	return utf8.DecodeRune(eb.text[eb.cursorBoffset:])
}

func (eb *editBox) RuneBeforeCursor() (rune, int) {
	return utf8.DecodeLastRune(eb.text[:eb.cursorBoffset])
}

func (eb *editBox) MoveCursorOneRuneBackward() {
	if eb.cursorBoffset == 0 {
		return
	}
	_, size := eb.RuneBeforeCursor()
	eb.MoveCursorTo(eb.cursorBoffset - size)
}

func (eb *editBox) MoveCursorOneRuneForward() {
	if eb.cursorBoffset == len(eb.text) {
		return
	}
	_, size := eb.RuneUnderCursor()
	eb.MoveCursorTo(eb.cursorBoffset + size)
}

func (eb *editBox) MoveCursorToBeginningOfTheLine() {
	eb.MoveCursorTo(0)
}

func (eb *editBox) MoveCursorToEndOfTheLine() {
	eb.MoveCursorTo(len(eb.text))
}

func (eb *editBox) DeleteRuneBackward() {
	if eb.cursorBoffset == 0 {
		return
	}

	eb.MoveCursorOneRuneBackward()
	_, size := eb.RuneUnderCursor()
	eb.text = byteSliceRemove(eb.text, eb.cursorBoffset, eb.cursorBoffset+size)
}

func (eb *editBox) DeleteRuneForward() {
	if eb.cursorBoffset == len(eb.text) {
		return
	}
	_, size := eb.RuneUnderCursor()
	eb.text = byteSliceRemove(eb.text, eb.cursorBoffset, eb.cursorBoffset+size)
}

func (eb *editBox) DeleteTheRestOfTheLine() {
	eb.text = eb.text[:eb.cursorBoffset]
}

func (eb *editBox) InsertRune(r rune) {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	eb.text = byteSliceInsert(eb.text, eb.cursorBoffset, buf[:n])
	eb.MoveCursorOneRuneForward()
}

func (eb *editBox) CursorX() int {
	return eb.cursorVoffset - eb.lineVoffset
}
