package day1

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/D-rem/advent-of-code-2020/lib"
)

func Part1(target int) (string, error) {
	t, err := lib.ReadAllLineToInt("day1/input.txt")

	if err != nil {
		return "", err
	}
	sort.Ints(t)
	for i := len(t) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if t[i]+t[j] == target {
				return strconv.Itoa(t[i] * t[j]), nil
			}
		}
	}

	return "", errors.New(fmt.Sprint("no couple match target : ", target))
}

func Part2(target int) (string, error) {
	t, err := lib.ReadAllLineToInt("day1/input.txt")

	if err != nil {
		return "", err
	}
	for i, x := range t[:len(t)] {
		for j, y := range t[i+1 : len(t)] {
			for _, z := range t[j+2 : len(t)] {
				if x+y+z == target {
					return strconv.Itoa(x * y * z), nil
				}
			}
		}
	}

	return "", errors.New(fmt.Sprint("no couple match target : ", target))
}
