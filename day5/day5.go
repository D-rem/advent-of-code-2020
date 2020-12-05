package day5

import (
	"errors"
	"sort"

	"github.com/D-rem/advent-of-code-2020/lib"
)

func computeSeatID(seatPlacement []rune) int {
	rowUp := 127
	rowDown := 0
	columnUp := 7
	columnDown := 0

	for _, seatInstruction := range seatPlacement {
		switch seatInstruction {
		case 'F':
			rowUp = rowUp - ((rowUp - rowDown) / 2)

		case 'B':
			rowDown = rowDown + ((rowUp - rowDown + 1) / 2)

		case 'L':
			columnUp = columnUp - ((columnUp - columnDown) / 2)

		case 'R':
			columnDown = columnDown + ((columnUp - columnDown + 1) / 2)
		}
	}
	return (rowDown*8 + columnDown)
}

func Part1() (int, error) {
	t, err := lib.ReadAllLineToRunesSlice("day5/input.txt")

	if err != nil {
		return 0, err
	}

	highestSeatId := -1

	for _, seatPlacement := range t {

		seatId := computeSeatID(seatPlacement)
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	return highestSeatId, nil

}

func Part2() (int, error) {
	t, err := lib.ReadAllLineToRunesSlice("day5/input.txt")

	if err != nil {
		return 0, err
	}

	var seatIds []int
	for _, seatPlacement := range t {
		seatIds = append(seatIds, computeSeatID(seatPlacement))
	}

	sort.Ints(seatIds)

	s1 := seatIds[0]
	for _, s2 := range seatIds[1:] {
		if s2-s1 > 1 {
			return s1 + 1, nil
		}
		s1 = s2
	}
	return 0, errors.New("no seat Id found")
}
