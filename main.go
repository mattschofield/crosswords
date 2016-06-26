package main

import (
	"fmt"
	"log"
)

var (
	BLACK               Colour = 0
	WHITE               Colour = 1
	rows                [][]int
	grid                Grid = Grid{}
	visitedWhiteSquares int  = 0
)

type Colour int
type Grid struct {
	squares [][]Square
}
type Square struct {
	x      int
	y      int
	seen   bool
	colour Colour
}

func (s *Square) toString() string {
	var colour string
	if s.colour == 1 {
		colour = "white"
	} else {
		colour = "black"
	}
	return fmt.Sprintf("[%d,%d] COLOUR=%s SEEN=%v", s.x, s.y, colour, s.seen)
}

func (s *Square) valid() bool {
	if s.colour == WHITE {
		return true
	} else {
		return false
	}
}

func (g *Grid) getSquareAtCoords(x, y int) Square {
	return g.squares[x][y]
}

func (g *Grid) findFirstValidSquare() *Square {
	for i := 0; i < len(g.squares); i++ {
		for j := 0; j < len(g.squares[i]); j++ {
			sq := g.squares[i][j]
			if sq.valid() {
				return &sq
			}
		}
	}
	return nil
}

func (g *Grid) getValidNeighbours(s Square) []Square {
	sqs := []Square{}

	if s.x >= 1 {
		sq := g.getSquareAtCoords(s.x-1, s.y)
		if sq.valid() {
			sqs = append(sqs, sq)
		}
	}

	if s.x < len(g.squares)-1 {
		sq := g.getSquareAtCoords(s.x+1, s.y)
		if sq.valid() {
			sqs = append(sqs, sq)
		}
	}

	if s.y >= 1 {
		sq := g.getSquareAtCoords(s.x, s.y-1)
		if sq.valid() {
			sqs = append(sqs, sq)
		}
	}

	if s.y < len(g.squares)-1 {
		sq := g.getSquareAtCoords(s.x, s.y+1)
		if sq.valid() {
			sqs = append(sqs, sq)
		}
	}

	return sqs
}

func (g *Grid) traverse(s Square) {
	g.squares[s.x][s.y].seen = true
	neighbours := g.getValidNeighbours(s)

	for i := 0; i < len(neighbours); i++ {
		if !neighbours[i].seen {
			g.traverse(neighbours[i])
		}
	}
}

func init() {
	rows = make([][]int, 13, 13)

	rows[0] = []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0}
	rows[1] = []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}
	rows[2] = []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	rows[3] = []int{0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0}
	rows[4] = []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	rows[5] = []int{0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0}
	rows[6] = []int{1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	rows[7] = []int{0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0}
	rows[8] = []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1}
	rows[9] = []int{0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0}
	rows[10] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	rows[11] = []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}
	rows[12] = []int{0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}

	grid.squares = make([][]Square, len(rows))

	for i := 0; i < len(rows); i++ {
		grid.squares[i] = make([]Square, len(rows[i]))

		for j := 0; j < len(rows[i]); j++ {
			var c Colour
			if rows[i][j] == 1 {
				c = WHITE
			} else {
				c = BLACK
			}

			grid.squares[i][j] = Square{
				x:      i,
				y:      j,
				seen:   false,
				colour: c,
			}
		}
	}
}

func main() {
	run(grid)
}

func run(g Grid) {
	// First count the number of white squares in the crossword
	totalWhiteSquares := 0

	for i := 0; i < len(g.squares); i++ {
		for j := 0; j < len(g.squares[i]); j++ {
			if len(g.squares) != len(g.squares[i]) {
				log.Printf(
					"invalid crossword, not square - row %d has %d squares, expected %d",
					i+1,
					len(g.squares[i]),
					len(g.squares),
				)
				return
			}

			sq := g.squares[i][j]
			if sq.colour == WHITE {
				totalWhiteSquares++
			}

			sq.seen = true
		}
	}

	log.Printf(
		"crossword is %dx%d with %d white spaces",
		len(g.squares),
		len(g.squares),
		totalWhiteSquares,
	)

	g.traverse(*g.findFirstValidSquare())

	var count int = 0
	for i := 0; i < len(g.squares); i++ {
		for j := 0; j < len(g.squares[i]); j++ {
			sq := g.squares[i][j]
			if sq.colour == WHITE && sq.seen {
				count++
			}
		}
	}

	if count == totalWhiteSquares {
		log.Printf("Success - traversed %d white squares out of %d\n", count, totalWhiteSquares)
	} else {
		log.Printf("Failed - traversed %d white squares out of %d\n", count, totalWhiteSquares)
	}
}
