package main

import (
	"fmt"
	"strings"
)

type BalloonName string

const purple BalloonName = "purple"
const eagle BalloonName = "eagle"
const checker BalloonName = "checker"
const zigzag BalloonName = "zigzag"

type Balloon struct {
	Name   BalloonName
	Basket bool
}

func (b *Balloon) String() string {
	if b.Basket {
		return strings.ToUpper(string(b.Name[0]))
	}
	return string(b.Name[0])
}

func (b *Balloon) matches(other *Balloon) bool {
	return b.Name == other.Name && b.Basket != other.Basket
}

type Piece struct {
	// sides go clockwise
	Sides []*Balloon
}

type Place struct {
	Piece   *Piece
	TopSide int
}

func (pl *Place) Top() *Balloon {
	return pl.Piece.Sides[pl.TopSide]
}

func (pl *Place) Right() *Balloon {
	return pl.Piece.Sides[(pl.TopSide+1)%4]
}

func (pl *Place) Bottom() *Balloon {
	return pl.Piece.Sides[(pl.TopSide+2)%4]
}

func (pl *Place) Left() *Balloon {
	return pl.Piece.Sides[(pl.TopSide+3)%4]
}

func (pl *Place) String() string {
	return fmt.Sprintf("_____\n| %s |\n|%s %s|\n| %s |", pl.Top(), pl.Left(), pl.Right(), pl.Bottom())
}

func (pl *Place) StringRow(r int) string {
	if pl != nil {
		if r == 0 {
			return "_____"
		} else if r == 1 {
			return fmt.Sprintf("| %s |", pl.Top())
		} else if r == 2 {
			return fmt.Sprintf("|%s %s|", pl.Left(), pl.Right())
		} else if r == 3 {
			return fmt.Sprintf("| %s |", pl.Bottom())
		}
	} else {
		if r == 0 {
			return "_____"
		} else {
			return "|   |"
		}
	}
	return ""
}

var order = [][]int{
	[]int{1, 1},
	[]int{1, 0},
	[]int{0, 0},
	[]int{0, 1},
	[]int{0, 2},
	[]int{1, 2},
	[]int{2, 2},
	[]int{2, 1},
	[]int{2, 0},
}

var used = make(map[*Piece]struct{}, 9)

type Board [][]*Place

func (b Board) String() string {
	stringBuilder := ""
	for row := 2; row >= 0; row-- {
		for line := 0; line < 4; line++ {
			for col := 0; col < 3; col++ {
				pl := b[col][row]
				stringBuilder += pl.StringRow(line)
			}
			stringBuilder += "\n"
		}
	}
	return stringBuilder
}

var total = 0
var partials = 0

func main() {
	board := Board{
		make([]*Place, 3),
		make([]*Place, 3),
		make([]*Place, 3),
	}
	fmt.Println(board)
	fillPlace(board, 0)
}

func fillPlace(board Board, nextPos int) bool {
	if nextPos >= len(order) {
		fmt.Printf("Solved! T:%d P:%d\n%s", total, partials, board)
		return true
	}
	x := order[nextPos][0]
	y := order[nextPos][1]
	for _, piece := range pieces {
		if _, isUsed := used[piece]; !isUsed {
			place := fits(board, piece, x, y)
			if place != nil {
				used[piece] = struct{}{}
				board[x][y] = place
				partials++
				fmt.Println(board)
				if !fillPlace(board, nextPos+1) {
					board[x][y] = nil
					delete(used, piece)
				}
			}
		}
	}
	return false
}

func fits(board Board, piece *Piece, x, y int) *Place {
	// rotation
	for i := 0; i < 4; i++ {
		place := &Place{piece, i}
		total++
		if checkAdjacent(board, place, x, y) {
			return place
		}
	}
	return nil
}

// will place fit at x,y
func checkAdjacent(board Board, place *Place, x, y int) bool {
	// x+1, y
	// x-1, y
	// x, y+1
	// x, y-1
	for _, i := range []int{-1, 1} {
		adjx := x + i
		if adjx >= 0 && adjx <= 2 && board[adjx][y] != nil {
			adjacent := board[adjx][y]
			if i == -1 && !place.Left().matches(adjacent.Right()) {
				return false
			}
			if i == 1 && !place.Right().matches(adjacent.Left()) {
				return false
			}
		}
	}
	for _, i := range []int{-1, 1} {
		adjy := y + i
		if adjy >= 0 && adjy <= 2 && board[x][adjy] != nil {
			adjacent := board[x][adjy]
			if i == -1 && !place.Bottom().matches(adjacent.Top()) {
				return false
			}
			if i == 1 && !place.Top().matches(adjacent.Bottom()) {
				return false
			}
		}
	}
	return true
}

var pieces = make([]*Piece, 9)

func init() {
	pieces[0] = &Piece{
		Sides: []*Balloon{
			&Balloon{
				Name:   eagle,
				Basket: false,
			},
			&Balloon{
				Name:   purple,
				Basket: true,
			},
			&Balloon{
				Name:   zigzag,
				Basket: false,
			},
			&Balloon{
				Name:   purple,
				Basket: false,
			},
		},
	}
	pieces[1] = &Piece{
		Sides: []*Balloon{
			&Balloon{
				Name:   eagle,
				Basket: true,
			},
			&Balloon{
				Name:   checker,
				Basket: false,
			},
			&Balloon{
				Name:   purple,
				Basket: true,
			},
			&Balloon{
				Name:   zigzag,
				Basket: false,
			},
		},
	}
	pieces[2] = &Piece{
		Sides: []*Balloon{
			&Balloon{
				Name:   eagle,
				Basket: true,
			},
			&Balloon{
				Name:   purple,
				Basket: true,
			},
			&Balloon{
				Name:   checker,
				Basket: false,
			},
			&Balloon{
				Name:   zigzag,
				Basket: true,
			},
		},
	}
	pieces[3] = &Piece{
		Sides: []*Balloon{
			&Balloon{
				Name:   purple,
				Basket: false,
			},
			&Balloon{
				Name:   zigzag,
				Basket: false,
			},
			&Balloon{
				Name:   checker,
				Basket: false,
			},
			&Balloon{
				Name:   eagle,
				Basket: true,
			},
		},
	}
	pieces[4] = &Piece{
		Sides: []*Balloon{
			&Balloon{
				Name:   checker,
				Basket: true,
			},
			&Balloon{
				Name:   eagle,
				Basket: true,
			},
			&Balloon{
				Name:   eagle,
				Basket: false,
			},
			&Balloon{
				Name:   zigzag,
				Basket: true,
			},
		},
	}
	pieces[5] = &Piece{
		Sides: []*Balloon{
			&Balloon{
				Name:   purple,
				Basket: false,
			},
			&Balloon{
				Name:   checker,
				Basket: false,
			},
			&Balloon{
				Name:   zigzag,
				Basket: false,
			},
			&Balloon{
				Name:   checker,
				Basket: true,
			},
		},
	}
	pieces[6] = &Piece{
		Sides: []*Balloon{
			&Balloon{
				Name:   checker,
				Basket: true,
			},
			&Balloon{
				Name:   purple,
				Basket: false,
			},
			&Balloon{
				Name:   zigzag,
				Basket: false,
			},
			&Balloon{
				Name:   eagle,
				Basket: true,
			},
		},
	}
	pieces[7] = &Piece{
		Sides: []*Balloon{
			&Balloon{
				Name:   checker,
				Basket: true,
			},
			&Balloon{
				Name:   eagle,
				Basket: false,
			},
			&Balloon{
				Name:   zigzag,
				Basket: true,
			},
			&Balloon{
				Name:   purple,
				Basket: true,
			},
		},
	}
	pieces[8] = &Piece{
		Sides: []*Balloon{
			&Balloon{
				Name:   checker,
				Basket: true,
			},
			&Balloon{
				Name:   eagle,
				Basket: false,
			},
			&Balloon{
				Name:   zigzag,
				Basket: false,
			},
			&Balloon{
				Name:   purple,
				Basket: false,
			},
		},
	}
}
