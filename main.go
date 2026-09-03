package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"math/rand/v2"
)

type playerPos struct {
	X_pos int
	Y_pos int 
}

func popRandom(s *[]string) string{
	idx := rand.IntN(len(*s))
	item := (*s)[idx]
	lastIdx := len(*s) - 1
	(*s)[idx] = (*s)[lastIdx]
	*s = (*s)[:lastIdx]
	return item
}

func resetMines(n int, m int) [][]string{
	helper := make([]string, (n*n) - m)
	helper_mines := make([]string, m)
	for i := range helper{
		helper[i] = "?"
	}

	for i := range helper_mines{
		helper_mines[i] = "M"
	}

	var final_mines []string

	final_mines = append(final_mines, helper...)
	final_mines = append(final_mines, helper_mines...)

	res := make([][]string, n)
	for Y := range res{
		res[Y] = make([]string, n)
		for X := range res[Y]{
			res[Y][X] = popRandom(&final_mines)
		}
	}

	return res
}

func renderMines(mines [][]string, s tcell.Screen, currPos playerPos) {
	for Y := range mines {
		for X := range mines[Y] {
			val := mines[Y][X]
			if len(val) == 0 {
				continue
			}
			style := tcell.StyleDefault

			if val == "M" {
				style = style.Foreground(tcell.ColorRed)
			} else if val >= "1" && val <= "8" {
				switch val {
				case "1":
					style = style.Foreground(tcell.ColorBlue)
				case "2":
					style = style.Foreground(tcell.ColorGreen)
				case "3":
					style = style.Foreground(tcell.ColorDarkCyan)
				case "4":
					style = style.Foreground(tcell.ColorDarkBlue)
				case "5":
					style = style.Foreground(tcell.ColorLightCyan)
				default:
					style = style.Foreground(tcell.ColorViolet)
				}
			}

			if Y == currPos.Y_pos && X == currPos.X_pos {
				style = style.Underline(true).Bold(true)
			}

			r := []rune(val)[0]
			s.SetContent(X*2, Y, r, nil, style)
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

	mines_hidden := resetMines(20,50)
	mines_to_show :=  make([][]string, 20)

	for Y := range mines_to_show{
		mines_to_show[Y] = make([]string, 20)
		for X := range mines_to_show[Y]{
			mines_to_show[Y][X] = "?"
		}
	}

	currPos := playerPos{X_pos: 0, Y_pos: 0}

	screen.Clear()

	for {
		renderMines(mines_hidden, screen, currPos)
		screen.Show()
		ev := screen.PollEvent()
		switch ev := ev.(type) {
			case *tcell.EventKey:
			switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					return
				case tcell.KeyUp:
					movePlayer("UP", &currPos)
				case tcell.KeyDown:
					movePlayer("DOWN", &currPos)
				case tcell.KeyRight:
					movePlayer("RIGHT", &currPos)
				case tcell.KeyLeft:
					movePlayer("LEFT", &currPos)
				case tcell.KeyRune:
					switch ev.Rune(){
						case 'R', 'r':
						mines_hidden = resetMines(20,50)
						mines_to_show =  make([][]string, 20)

						for Y := range mines_to_show{
							mines_to_show[Y] = make([]string, 20)
							for X := range mines_to_show[Y]{
								mines_to_show[Y][X] = "?"
							}
						}
					}
			}
		}
	}
}