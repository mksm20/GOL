package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"time"
)

type Game struct {
	grid [][]Cell
}

func GetCellState(x int, y int, grid [][]Cell) Cell {
	if x < 0 || x >= len(grid) {
		return Cell{Alive: false}
	}
	if y < 0 || y >= len(grid[0]) {
		return Cell{Alive: false}
	}
	return grid[x][y]
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

func nextGeneration(grid [][]Cell) [][]Cell {
	newGrid := make([][]Cell, len(grid))
	for y := 0; y < len(grid); y++ {
		newGrid[y] = make([]Cell, len(grid[0]))
	}
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			newGrid[y][x] = nextCellState(x, y, grid)
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
				fmt.Print("█*█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	gridSize := 40
	noChange := bool(false)
	grid := initGrid(gridSize)
	for !noChange {
		render(grid)
		newGrid := nextGeneration(grid)
		noChange = reflect.DeepEqual(grid, newGrid)
		grid = newGrid
		time.Sleep(20 * time.Millisecond)
	}
}
