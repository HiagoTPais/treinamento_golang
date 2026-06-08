package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	WIDTH    = 31
	HEIGHT   = 31
	TILESIZE = 20
)

var maze [HEIGHT][WIDTH]int

var directions = [][2]int{
	{0, -2},
	{0, 2},
	{-2, 0},
	{2, 0},
}

type Game struct{}

func generateMaze(x, y int) {

	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for _, dir := range directions {

		nx := x + dir[0]
		ny := y + dir[1]

		if nx > 0 && nx < WIDTH-1 &&
			ny > 0 && ny < HEIGHT-1 {

			if maze[ny][nx] == 1 {

				// remove parede
				maze[y+dir[1]/2][x+dir[0]/2] = 0

				// marca caminho
				maze[ny][nx] = 0

				generateMaze(nx, ny)
			}
		}
	}
}

func initMaze() {

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			maze[y][x] = 1
		}
	}

	maze[1][1] = 0

	generateMaze(1, 1)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {

			if maze[y][x] == 1 {

				// parede
				ebitenutil.DrawRect(
					screen,
					float64(x*TILESIZE),
					float64(y*TILESIZE),
					TILESIZE,
					TILESIZE,
					color.RGBA{40, 40, 40, 255},
				)

			} else {

				// chão
				ebitenutil.DrawRect(
					screen,
					float64(x*TILESIZE),
					float64(y*TILESIZE),
					TILESIZE,
					TILESIZE,
					color.RGBA{220, 220, 220, 255},
				)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH * TILESIZE, HEIGHT * TILESIZE
}

func main() {

	rand.Seed(time.Now().UnixNano())

	initMaze()

	ebiten.SetWindowTitle("Maze Generator")

	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
