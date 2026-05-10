package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	WIDTH  = 15
	HEIGHT = 15
)

var maze [HEIGHT][WIDTH]int

var directions = [][2]int{
	{0, -2},
	{0, 2},
	{-2, 0},
	{2, 0},
}

func generateMaze(x, y int) {
	for a := 0; a < HEIGHT; a++ {
		for b := 0; b < WIDTH; b++ {
			maze[a][b] = 1
		}
	}

	maze[y][x] = 0

	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for _, dir := range directions {
		nx := x + dir[0]
		ny := y + dir[1]

		if nx > 0 && nx < WIDTH-1 && ny > 0 && ny < HEIGHT-1 {
			if maze[ny][nx] == 1 {
				maze[y+dir[1]/2][x+dir[0]/2] = 0
				generateMaze(nx, ny)
			}
		}
	}
}

func printMaze() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			fmt.Print(maze[y][x], " ")
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	generateMaze(1, 1)
	printMaze()
}
