package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"time"
)

func GetCellState(x int, y int, grid [][]Cell) Cell {
	if x < 0 || x >= len(grid) {
		return Cell{Alive: false}
	}
	if y < 0 || y >= len(grid[0]) {
		return Cell{Alive: false}
	}
	return grid[x][y]
}

func outSideAlive(grid [][]Cell, expandX []int, expandY []int) {
	x1Complete, x2Complete := false, false
	y1Complete, y2Complete := false, false
	for _, val := range grid {
		if val[0].Alive && !x1Complete {
			expandX[0] = -1
			x1Complete = true
		}
		if val[len(grid[0])-1].Alive && !x2Complete {
			expandX[1] = 1
			x2Complete = true
		}
		if x1Complete && x2Complete {
			break
		}
	}
	for _, val := range grid[0] {
		if val.Alive && !y1Complete {
			expandY[0] = -1
			y1Complete = true
		}
		if y1Complete {
			break
		}
	}

	for _, val := range grid[len(grid)-1] {
		if val.Alive && !y2Complete {
			expandY[1] = 1
			y2Complete = true
		}
		if y2Complete {
			break
		}
	}
	fmt.Println(expandX[0], expandY[0])
}

func nextCellState(x int, y int, grid [][]Cell) Cell {
	current := GetCellState(x, y, grid)
	numAlive := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			neighbor := GetCellState(x+dx, y+dy, grid)
			if neighbor.Alive {
				numAlive++
			}
		}
	}
	if (numAlive == 2 || numAlive == 3) && current.Alive {
		current.Alive = true
		return current
	}
	if (numAlive > 3) && current.Alive {
		current.Alive = false
		return current
	}
	if numAlive == 3 && !current.Alive {
		current.Alive = true
		return current
	}
	if (numAlive < 2) && current.Alive {
		current.Alive = false
		return current
	}
	return current
}

func nextGeneration(grid [][]Cell, expandX []int, expandY []int) [][]Cell {
	outSideAlive(grid, expandX, expandY)
	shiftX, shiftY := false, false
	sizeX, sizeY := len(grid), len(grid[0])

	if expandX[0] == -1 {
		shiftX = true
		sizeX = sizeX + 1
	}
	if expandY[0] == -1 {
		shiftY = true
		sizeY = sizeY + 1
	}
	if expandX[1] == 1 {
		sizeX = sizeX + 1
	}
	if expandY[1] == 1 {
		sizeY = sizeY + 1
	}

	newGrid := make([][]Cell, sizeX)

	for i := range newGrid {
		newGrid[i] = make([]Cell, sizeY)
	}

	fmt.Println(len(newGrid), len(newGrid[0]))

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if !shiftX && !shiftY {
				newGrid[x][y] = nextCellState(x, y, grid)
			}
			if shiftX && !shiftY {
				newGrid[x+1][y] = nextCellState(x, y, grid)
			}
			if !shiftX && shiftY {
				newGrid[x][y+1] = nextCellState(x, y, grid)
			}
			if shiftX && shiftY {
				newGrid[x+1][y+1] = nextCellState(x, y, grid)
			}
		}
	}
	return newGrid
}

func initGrid(grid_size int) [][]Cell {
	initGrid := make([][]Cell, grid_size)
	for y := 0; y < grid_size; y++ {
		initGrid[y] = make([]Cell, grid_size)
	}
	for x := 0; x < grid_size; x++ {
		for y := 0; y < grid_size; y++ {
			initGrid[x][y] = Cell{Alive: rand.IntN(5) == 1}
		}
	}
	return initGrid
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func render(grid [][]Cell) {
	clearScreen()
	for _, row := range grid {
		for _, cell := range row {
			if cell.Alive {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	expand_y := make([]int, 2)
	expand_x := make([]int, 2)
	gridSize := 20
	noChange := bool(false)
	grid := initGrid(gridSize)
	for !noChange {
		render(grid)
		newGrid := nextGeneration(grid, expand_x, expand_y)
		expand_x[0], expand_x[1] = 0, 0
		expand_y[0], expand_y[1] = 0, 0
		noChange = reflect.DeepEqual(grid, newGrid)
		grid = newGrid
		time.Sleep(200 * time.Millisecond)
	}
}
