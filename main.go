package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

var (
	options  []string
	maxLen   int
	selected int

	initPosX int
	initPosY int

	defaultStyle  tcell.Style
	selectedStyle tcell.Style

	isCentered bool
)

func main() {
	isCentered = hasArg("-c")

	defaultStyle = tcell.StyleDefault.
		Background(tcell.ColorReset).
		Foreground(tcell.ColorReset)
	selectedStyle = tcell.StyleDefault.
		Background(tcell.ColorWhite).
		Foreground(tcell.ColorBlack)
	selected = 0

	s := initScreen()
	defer finalizeScreen(s)
	defer os.Exit(1)

	options, maxLen = getOptionsFromStdin()

	if isCentered {
		initPosX, initPosY = s.Size()
		initPosX /= 2
		initPosY /= 2
		initPosX -= maxLen/2 + 1
		initPosY -= len(options)/2 + 1
	} else {
		initPosX, initPosY = 0, 0
	}

	s.SetStyle(defaultStyle)
	for {
		selected = capIntBetweenValues(0, selected, len(options)-1)
		draw(s)
		handleEvents(s)
	}
}

func initScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}

func handleEvents(s tcell.Screen) {
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

func draw(s tcell.Screen) {
	s.Clear()
	drawBox(s, initPosX, initPosY, initPosX+maxLen+1, initPosY+len(options)+1, defaultStyle)
	for i, o := range options {
		style := defaultStyle
		if selected == i {
			style = selectedStyle
		}
		drawText(s, initPosX+1, initPosY+1+i, initPosX+maxLen+1, initPosY+1+i, style, o)
	}
	s.Show()
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
