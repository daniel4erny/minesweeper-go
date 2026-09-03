package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"math/rand/v2"

	"github.com/gdamore/tcell/v2"
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
		helper[i] = "#"
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

	var mine_count int

	counts := [][2]int{{-1, -1}, {-1, 0}, {-1, +1}, {0, -1}, {0, +1}, {+1, -1}, {+1, 0}, {+1, +1}}

	for Y := range res{
		for X := range res{
			if res[Y][X] == "M"{
				continue
			} 

			mine_count = 0
			for _, val := range counts{
				if Y + val[0] < 0 || X + val[1] < 0 || Y + val[0] >= n  || X + val[1] >= n{
					continue
				}
				if res[Y+val[0]][X+val[1]] == "M"{
					mine_count += 1
				}
			}

			if mine_count > 0{
				res[Y][X] = strconv.FormatInt(int64(mine_count), 10)
			}
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
					style = style.Foreground(tcell.ColorYellow)
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

			if val == "F"{
				style = style.Foreground(tcell.ColorGreen)
			}

			if val == "#"{
				style = style.Foreground(tcell.ColorGray)
			}

			if Y == currPos.Y_pos && X == currPos.X_pos {
				style = style.Underline(true).Bold(true)
			}

			r := []rune(val)[0]
			s.SetContent(X*2, Y, r, nil, style)
		}
	}
}

func renderStatus(currPos playerPos, n int, time int, remainingMines int, s tcell.Screen, game_status string){
	pos := fmt.Sprintf("POS: X: %d Y: %d", currPos.X_pos, currPos.Y_pos)
	s.PutStr((2*n)+5, 0, pos)
	timestr := fmt.Sprintf("TIME: %d", time)
	s.PutStr((2*n)+5, 1, timestr)
	remainingMinesStr := fmt.Sprintf("REMAINING MINES: %d", remainingMines)
	s.PutStr((2*n)+5, 2, remainingMinesStr)
	status := fmt.Sprintf("STATUS: %s", game_status)
	s.PutStr((2*n)+5, 3, status)
}

func movePlayer(ev string, currPos *playerPos, n int){
	if ev == "UP"{
		if currPos.Y_pos > 0{
			currPos.Y_pos -= 1
		} 
	} else if ev == "DOWN"{
		if currPos.Y_pos < n-1{
			currPos.Y_pos += 1
		}
	} else if ev == "LEFT"{
		if currPos.X_pos > 0{
			currPos.X_pos -= 1
		}
	} else if ev == "RIGHT"{
		if currPos.X_pos < n-1{
			currPos.X_pos += 1
		}
	}
}

func flagMine(mines [][]string, currPos playerPos, remainNumber *int){
	if *remainNumber == 0 && mines[currPos.Y_pos][currPos.X_pos] != "F"{
		return
	}
	if mines[currPos.Y_pos][currPos.X_pos] == "F"{
		mines[currPos.Y_pos][currPos.X_pos] = "?"
		*remainNumber += 1
	} else {
		mines[currPos.Y_pos][currPos.X_pos] = "F"
		*remainNumber -= 1
	}
}

func evalShowMine(mines_hidden *[][]string, mines_to_show *[][]string, currPos playerPos, game_status *string) {
	if *game_status == "LOST" {
		return
	}

	hidden := *mines_hidden
	shown := *mines_to_show

	height := len(hidden)
	if height == 0 {
		return
	}
	width := len(hidden[0])

	startY := currPos.Y_pos
	startX := currPos.X_pos

	if shown[startY][startX] == "F" {
		return
	}

	if hidden[startY][startX] == "M" {
		*game_status = "LOST"
		shown[startY][startX] = "M"
		return
	}

	if shown[startY][startX] != "?" {
		return
	}

	if hidden[startY][startX] != "" && hidden[startY][startX] != "#" {
		shown[startY][startX] = hidden[startY][startX]
		return
	}

	queue := []playerPos{currPos}

	directions := []playerPos{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		cY, cX := curr.Y_pos, curr.X_pos
		if shown[cY][cX] != "?" {
			continue
		}
		if hidden[cY][cX] == "" || hidden[cY][cX] == "#" {
			shown[cY][cX] = "#"
		} else {
			shown[cY][cX] = hidden[cY][cX]
		}
		if hidden[cY][cX] == "" || hidden[cY][cX] == "#" {
			for _, dir := range directions {
				nY, nX := cY+dir.Y_pos, cX+dir.X_pos
				if nY >= 0 && nY < height && nX >= 0 && nX < width {
					if shown[nY][nX] == "?" {
						queue = append(queue, playerPos{Y_pos: nY, X_pos: nX})
					}
				}
			}
		}
	}
}

func main(){
	var mode string
	var mine_count int
	var map_size int
	fmt.Print("Type type of game ('begginer', 'normal', 'expert') press enter to make custom game: ") 
	fmt.Scan(&mode)
	if mode == "begginer"{
		mine_count = 10
		map_size = 8
	} else if mode == "normal"{
		mine_count = 20
		map_size = 16
	} else if mode == "expert"{
		mine_count = 60
		map_size = 21
	} else {
		fmt.Print("mine count: ")
		_, err := fmt.Scan(&mine_count)
		if err != nil{
			log.Panic("IDK")
		}
		fmt.Print("map size: ")
		_, err = fmt.Scan(&map_size)
		if err != nil{
			log.Panic("JE TO V PICI")
		} else if mine_count >= map_size*map_size || map_size > 35{
			log.Panic("ty jsi blbej lol")
		}
	}
	game_status := "PlAYING"

	remaining_mines := mine_count

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

	mines_hidden := resetMines(map_size,mine_count)
	_ = mines_hidden
	mines_to_show :=  make([][]string, map_size)

	for Y := range mines_to_show{
		mines_to_show[Y] = make([]string, map_size)
		for X := range mines_to_show[Y]{
			mines_to_show[Y][X] = "?"
		}
	}

	currPos := playerPos{X_pos: 0, Y_pos: 0}

	screen.Clear()

	for {
		renderMines(mines_to_show, screen, currPos)
		renderStatus(currPos, map_size, 67, remaining_mines, screen, game_status)
		screen.Show()
		ev := screen.PollEvent()
		switch ev := ev.(type) {
			case *tcell.EventKey:
			switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					return
				case tcell.KeyUp:
					movePlayer("UP", &currPos, map_size)
				case tcell.KeyDown:
					movePlayer("DOWN", &currPos, map_size)
				case tcell.KeyRight:
					movePlayer("RIGHT", &currPos, map_size)
				case tcell.KeyLeft:
					movePlayer("LEFT", &currPos, map_size)
				case tcell.KeyRune:
					switch ev.Rune(){
						case 'R', 'r':
							mines_hidden = resetMines(map_size,mine_count)
							mines_to_show =  make([][]string, map_size)

							for Y := range mines_to_show{
								mines_to_show[Y] = make([]string, map_size)
								for X := range mines_to_show[Y]{
									mines_to_show[Y][X] = "?"
								}
							}
							game_status = "PLAYING"

						case 'f', 'F':
							flagMine(mines_to_show, currPos, &remaining_mines)

						case ' ':
							evalShowMine(&mines_hidden, &mines_to_show, currPos, &game_status)
				}
			}
		}
		screen.Clear()
	}
}
