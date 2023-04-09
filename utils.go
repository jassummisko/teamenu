package main

import (
	"io"
	"math"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}
}

func readStdin() string {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(stdin), "\n")
}

func capIntBetweenValues(min int, val int, max int) int {
	return int(math.Max(math.Min(float64(max), float64(val)), float64(min)))
}

func getOptionsFromStdin() ([]string, int) {
	opt := strings.Split(readStdin(), "\n")
	ml := 0
	for _, o := range opt {
		if len(o) > ml {
			ml = len(o)
		}
	}
	return opt, ml
}

func hasArg(arg string) bool {
	for _, s := range os.Args {
		if s == arg {
			return true
		}
	}
	return false
}
