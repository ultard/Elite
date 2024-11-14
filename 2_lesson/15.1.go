package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Grid [][]int

func initializeGrid(rows, cols int) Grid {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	return grid
}

func printGrid(grid Grid) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("■ ")
			} else {
				fmt.Print("□ ")
			}
		}
		fmt.Println()
	}
}

func countAliveNeighbors(grid Grid, x, y int) int {
	rows, cols := len(grid), len(grid[0])
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := x+i, y+j
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
				count++
			}
		}
	}
	return count
}

func step(grid Grid) Grid {
	rows, cols := len(grid), len(grid[0])
	newGrid := initializeGrid(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			liveNeighbors := countAliveNeighbors(grid, i, j)
			if grid[i][j] == 1 && (liveNeighbors == 2 || liveNeighbors == 3) {
				newGrid[i][j] = 1
			} else if grid[i][j] == 0 && liveNeighbors == 3 {
				newGrid[i][j] = 1
			}
		}
	}
	return newGrid
}

func loadInitialState(grid Grid, scanner *bufio.Scanner) {
	fmt.Println("Введите начальное состояние (формат: x y для каждой живой клетки, пустая строка для завершения):")
	for {
		fmt.Print("Клетка: ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		coords := strings.Fields(line)
		if len(coords) != 2 {
			fmt.Println("Ошибка: введите две координаты.")
			continue
		}
		x, err1 := strconv.Atoi(coords[0])
		y, err2 := strconv.Atoi(coords[1])
		if err1 != nil || err2 != nil || x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			fmt.Println("Ошибка: некорректные координаты.")
			continue
		}
		grid[x][y] = 1
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var rows, cols int
	var delay int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите количество строк: ")
	fmt.Scan(&rows)
	fmt.Print("Введите количество столбцов: ")
	fmt.Scan(&cols)
	//fmt.Print("Введите количество шагов симуляции: ")
	//fmt.Scan(&steps)
	fmt.Print("Введите задержку между шагами (в миллисекундах): ")
	fmt.Scan(&delay)

	grid := initializeGrid(cols, rows)
	loadInitialState(grid, scanner)

	for {
		clearScreen()
		//fmt.Printf("Шаг %d:\n", i+1)
		printGrid(grid)
		grid = step(grid)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
