package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ALIVE = iota
	DEAD
)

const CELL_WIDTH = 5
const NUMBER_OF_CELLS_WIDTH = 300
const NUMBER_OF_CELLS_HEIGHT = 200

type Cell struct {
	state int
}

func (c *Cell) IsAlive() bool {
	return c.state == ALIVE
}

type Board struct {
	height int
	width  int
	cells  [][]Cell
}

func (b *Board) deepcopy() Board {
	board := BuildBoard(b.width, b.height)

	for i := range board.cells {
		for j := range board.cells[i] {
			board.cells[i][j] = b.cells[i][j]
		}
	}
	return board
}

func (b *Board) countAliveNeighbours(i, j int) int {
	count := 0
	for di := range 3 {
		for dj := range 3 {

			if (di-1+i < 0) || (di-1+i >= b.width) {
				continue
			}

			if (dj-1+j < 0) || (dj-1+j >= b.height) {
				continue
			}
			if di == 1 && dj == 1 {
				continue
			}

			if b.cells[i+di-1][j+dj-1].IsAlive() {
				count += 1
			}
		}
	}
	return count
}

func BuildBoard(width int, height int) Board {
	board := Board{width: width, height: height}
	board.cells = make([][]Cell, width)
	for i := range board.cells {
		board.cells[i] = make([]Cell, board.height)
		for j := range board.cells[i] {
			if rand.Float32() < 0.5 {
				board.cells[i][j] = Cell{state: ALIVE}
			} else {
				board.cells[i][j] = Cell{state: DEAD}
			}

		}
	}
	return board
}

func (b *Board) DoGameStep() {
	currentBoard := b.deepcopy()
	for i := range b.width {
		for j := range b.height {
			if currentBoard.cells[i][j].IsAlive() {

				aliveNeighbours := currentBoard.countAliveNeighbours(i, j)
				if aliveNeighbours < 2 || aliveNeighbours > 3 {
					b.cells[i][j].state = DEAD
				}
			} else {
				aliveNeighbours := currentBoard.countAliveNeighbours(i, j)
				if aliveNeighbours == 3 {
					b.cells[i][j].state = ALIVE
				}
			}
		}
	}
}

func main() {
	rl.InitWindow(CELL_WIDTH*NUMBER_OF_CELLS_WIDTH, CELL_WIDTH*NUMBER_OF_CELLS_HEIGHT, "Conways Game of Life / Go & raylib")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	gameBoard := BuildBoard(NUMBER_OF_CELLS_WIDTH, NUMBER_OF_CELLS_HEIGHT)
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for i := range gameBoard.width {
			for j := range gameBoard.height {
				if gameBoard.cells[i][j].IsAlive() {
					rl.DrawRectangle(int32(i*CELL_WIDTH), int32(j*CELL_WIDTH), CELL_WIDTH, CELL_WIDTH, rl.Green)
				}
			}
		}
		//rl.DrawFPS(20, 20)

		rl.EndDrawing()

		gameBoard.DoGameStep()
	}
}
