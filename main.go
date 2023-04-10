package main

import (
	"flag"
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

	isCentered    *bool
	menuTitle     *string
	menuTitlePosX int
)

func main() {
	options, maxLen = GetOptionsFromStdin()

	isCentered = flag.Bool("c", false, "center the menu")
	menuTitle = flag.String("t", "", "title of the menu")
	flag.Parse()

	defaultStyle = tcell.StyleDefault.
		Background(tcell.ColorReset).
		Foreground(tcell.ColorReset)
	selectedStyle = tcell.StyleDefault.
		Background(tcell.ColorWhite).
		Foreground(tcell.ColorBlack)
	selected = 0

	s := InitScreen()
	defer FinalizeScreen(s)
	defer os.Exit(1)

	if *isCentered {
		initPosX, initPosY = s.Size()
		initPosX /= 2
		initPosY /= 2
		initPosX -= maxLen/2 + 1
		initPosY -= len(options)/2 + 1
		menuTitlePosX = initPosX - len(*menuTitle)/2 + 1 + maxLen/2
	} else {
		initPosX, initPosY = 0, 1
		menuTitlePosX = initPosX
	}

	s.SetStyle(defaultStyle)
	for {
		selected = CapIntBetweenValues(0, selected, len(options)-1)
		DrawMenu(s)
		HandleEvents(s)
	}
}
