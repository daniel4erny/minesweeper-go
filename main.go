package main

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell/v2"
)

type playerPos struct {
	X_pos int
	Y_pos int 
}

func resetMines(n int) [][]string{
	res := make([][]string, n)
	for Y := range res{
		res[Y] = make([]string, n)
		for X := range res[Y]{
			res[Y][X] = "?"
		}
	}

	return res
}

func renderMines(mines [][]string, s tcell.Screen, currPos playerPos){
	for Y := range mines{
		for X := range mines[Y]{
			if Y == currPos.Y_pos && X == currPos.X_pos{
				style := tcell.StyleDefault.Underline(true)
				r := []rune(mines[Y][X])[0]
				s.SetContent(X*2, Y, r, nil, style)
			} else {
				s.PutStr(X*2, Y, mines[Y][X])
			}
		}
	}
}

func movePlayer(ev string, currPos *playerPos){
	if ev == "UP"{
		if currPos.Y_pos > 0{
			currPos.Y_pos -= 1
		} 
	} else if ev == "DOWN"{
		if currPos.Y_pos < 10{
			currPos.Y_pos += 1
		}
	} else if ev == "LEFT"{
		if currPos.X_pos > 0{
			currPos.X_pos -= 1
		}
	} else if ev == "RIGHT"{
		if currPos.X_pos < 10{
			currPos.X_pos += 1
		}
	}
}

func main(){
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("idk, je to v pici")
		os.Exit(1)
	}

	err = screen.Init()
	if err != nil {
		fmt.Println("idk, je to v pici")
		os.Exit(1)
	}

	defer screen.Fini()

	screen.SetStyle(tcell.StyleDefault)

	mines := resetMines(10)
	currPos := playerPos{X_pos: 0, Y_pos: 0}

	screen.Clear()

	for {
		renderMines(mines, screen, currPos)
		screen.Show()
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} 
			if ev.Key() == tcell.KeyUp {
				movePlayer("UP", &currPos)
			} 
			if ev.Key() == tcell.KeyDown {
				movePlayer("DOWN", &currPos)
			}
			if ev.Key() == tcell.KeyRight{
				movePlayer("RIGHT", &currPos)
			}
			if ev.Key() == tcell.KeyLeft{
				movePlayer("LEFT", &currPos)
			}
		}
	}
}