package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}

	defaultStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	selectedStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
	s.SetStyle(defaultStyle)
	s.Clear()

	defer os.Exit(0)

	options := strings.Split(readStdin(), "\n")
	maxLen := 0
	for _, o := range options {
		if len(o) > maxLen {
			maxLen = len(o)
		}
	}
	maxLen += 1
	selected := 0
	for {
		selected = capIntBetweenValues(0, selected, len(options)-1)
		s.Clear()
		drawBox(s, 0, 0, maxLen, len(options)+1, defaultStyle)
		for i, o := range options {
			style := defaultStyle
			if selected == i {
				style = selectedStyle
			}
			drawText(s, 1, 1+i, maxLen, 1+i, style, o)
		}
		s.Show()

		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				finalizeScreen(s)
				os.Exit(0)
			} else if ev.Key() == tcell.KeyDown {
				selected += 1
			} else if ev.Key() == tcell.KeyUp {
				selected -= 1
			} else if ev.Key() == tcell.KeyEnter {
				outputAndExit(s, options[selected])
			}
		}
	}
}

func finalizeScreen(s tcell.Screen) {
	err := recover()
	s.Fini()
	if err != nil {
		panic(err)
	}
}

func outputAndExit(s tcell.Screen, out string) {
	finalizeScreen(s)
	fmt.Println(out)
	os.Exit(0)
}
