package main

import (
	"fmt"
	"strings"
)

type CandyName string

const CandyCane CandyName = "aCandyCane"
const JellyBean CandyName = "bJellyBean"
const CandyCorn CandyName = "cCandyCorn"
const CitrusJelly CandyName = "jCitrusJelly"

type Candy struct {
	Name    CandyName
	TopLeft bool
}

func (b *Candy) String() string {
	if b.TopLeft {
		return strings.ToUpper(string(b.Name[0]))
	}
	return string(b.Name[0])
}

func (b *Candy) matches(other *Candy) bool {
	return b.Name == other.Name && b.TopLeft != other.TopLeft
}

type Piece struct {
	Index int
	// sides go clockwise
	Sides []*Candy
}

type Place struct {
	Piece   *Piece
	TopSide int
}

func (pl *Place) Top() *Candy {
	return pl.Piece.Sides[pl.TopSide]
}

func (pl *Place) Right() *Candy {
	return pl.Piece.Sides[(pl.TopSide+1)%4]
}

func (pl *Place) Bottom() *Candy {
	return pl.Piece.Sides[(pl.TopSide+2)%4]
}

func (pl *Place) Left() *Candy {
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
			return fmt.Sprintf("|%s%d%s|", pl.Left(), pl.Piece.Index, pl.Right())
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
		Index: 0,
		Sides: []*Candy{
			&Candy{
				Name:    JellyBean,
				TopLeft: false,
			},
			&Candy{
				Name:    CandyCane,
				TopLeft: true,
			},
			&Candy{
				Name:    CandyCorn,
				TopLeft: true,
			},
			&Candy{
				Name:    CitrusJelly,
				TopLeft: false,
			},
		},
	}
	pieces[1] = &Piece{
		Index: 1,
		Sides: []*Candy{
			&Candy{
				Name:    JellyBean,
				TopLeft: true,
			},
			&Candy{
				Name:    CitrusJelly,
				TopLeft: false,
			},
			&Candy{
				Name:    CandyCorn,
				TopLeft: true,
			},
			&Candy{
				Name:    JellyBean,
				TopLeft: false,
			},
		},
	}
	pieces[2] = &Piece{
		Index: 2,
		Sides: []*Candy{
			&Candy{
				Name:    CandyCane,
				TopLeft: false,
			},
			&Candy{
				Name:    CitrusJelly,
				TopLeft: false,
			},
			&Candy{
				Name:    CandyCane,
				TopLeft: true,
			},
			&Candy{
				Name:    JellyBean,
				TopLeft: false,
			},
		},
	}
	pieces[3] = &Piece{
		Index: 3,
		Sides: []*Candy{
			&Candy{
				Name:    JellyBean,
				TopLeft: false,
			},
			&Candy{
				Name:    CitrusJelly,
				TopLeft: true,
			},
			&Candy{
				Name:    CandyCorn,
				TopLeft: false,
			},
			&Candy{
				Name:    CandyCane,
				TopLeft: false,
			},
		},
	}
	pieces[4] = &Piece{
		Index: 4,
		Sides: []*Candy{
			&Candy{
				Name:    CandyCorn,
				TopLeft: true,
			},
			&Candy{
				Name:    CandyCane,
				TopLeft: false,
			},
			&Candy{
				Name:    JellyBean,
				TopLeft: true,
			},
			&Candy{
				Name:    CitrusJelly,
				TopLeft: true,
			},
		},
	}
	pieces[5] = &Piece{
		Index: 5,
		Sides: []*Candy{
			&Candy{
				Name:    CandyCorn,
				TopLeft: false,
			},
			&Candy{
				Name:    CitrusJelly,
				TopLeft: false,
			},
			&Candy{
				Name:    JellyBean,
				TopLeft: false,
			},
			&Candy{
				Name:    CandyCane,
				TopLeft: true,
			},
		},
	}
	pieces[6] = &Piece{
		Index: 6,
		Sides: []*Candy{
			&Candy{
				Name:    CandyCane,
				TopLeft: false,
			},
			&Candy{
				Name:    CandyCorn,
				TopLeft: true,
			},
			&Candy{
				Name:    CandyCorn,
				TopLeft: false,
			},
			&Candy{
				Name:    CitrusJelly,
				TopLeft: true,
			},
		},
	}
	pieces[7] = &Piece{
		Index: 7,
		Sides: []*Candy{
			&Candy{
				Name:    CandyCorn,
				TopLeft: true,
			},
			&Candy{
				Name:    CitrusJelly,
				TopLeft: false,
			},
			&Candy{
				Name:    JellyBean,
				TopLeft: false,
			},
			&Candy{
				Name:    CandyCane,
				TopLeft: true,
			},
		},
	}
	pieces[8] = &Piece{
		Index: 8,
		Sides: []*Candy{
			&Candy{
				Name:    CitrusJelly,
				TopLeft: false,
			},
			&Candy{
				Name:    CandyCane,
				TopLeft: false,
			},
			&Candy{
				Name:    JellyBean,
				TopLeft: true,
			},
			&Candy{
				Name:    CandyCorn,
				TopLeft: true,
			},
		},
	}
}
