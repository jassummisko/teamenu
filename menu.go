package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

func InitScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}

func HandleEvents(s tcell.Screen) {
	ev := s.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventResize:
		s.Sync()
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			FinalizeScreen(s)
			os.Exit(0)
		} else if ev.Key() == tcell.KeyDown {
			selected += 1
		} else if ev.Key() == tcell.KeyUp {
			selected -= 1
		} else if ev.Key() == tcell.KeyEnter {
			OutputAndExit(s, options[selected])
		}
	}
}

func DrawMenu(s tcell.Screen) {
	s.Clear()
	DrawBox(s, initPosX, initPosY, initPosX+maxLen+1, initPosY+len(options)+1, defaultStyle)
	DrawText(s, menuTitlePosX, initPosY-1, menuTitlePosX+len(*menuTitle), initPosY-1, defaultStyle, *menuTitle)
	for i, o := range options {
		style := defaultStyle
		if selected == i {
			style = selectedStyle
		}
		DrawText(s, initPosX+1, initPosY+1+i, initPosX+maxLen+1, initPosY+1+i, style, o)
	}
	s.Show()
}

func FinalizeScreen(s tcell.Screen) {
	err := recover()
	s.Fini()
	if err != nil {
		panic(err)
	}
}

func OutputAndExit(s tcell.Screen, out string) {
	FinalizeScreen(s)
	fmt.Println(out)
	os.Exit(0)
}
