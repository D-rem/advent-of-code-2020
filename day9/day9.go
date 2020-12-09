package day9

import (
	"errors"
	"fmt"

	"github.com/D-rem/advent-of-code-2020/lib"
)

func execPart(preamble int, t []int) (int, error) {
outLoop:
	for i := preamble; i < len(t); i++ {
		for j := i - preamble; j < i-1; j++ {
			for k := i - preamble + 1; k < i; k++ {
				if t[i] == t[j]+t[k] {
					continue outLoop
				}
			}
		}
		return t[i], nil
	}
	return 0, errors.New("error")
}

func Part1(preamble int) (int, error) {
	t, err := lib.ReadAllLineToInt("day9/input.txt")
	if err != nil {
		return 0, err
	}

	return execPart(preamble, t)
}

func Part2(preamble int) (int, error) {
	t, err := lib.ReadAllLineToInt("day9/input.txt")
	if err != nil {
		return 0, err
	}

	numberToFind, err := execPart(preamble, t)
	if err != nil {
		return 0, err
	}
	var numbers []int
outLoop:
	for i := 0; i < len(t); i++ {
		total := t[i]
		numbers = append(numbers, t[i])
		for j := i + 1; j < len(t); j++ {
			total += t[j]
			numbers = append(numbers, t[j])
			if total == numberToFind && len(numbers) > 1 {
				break outLoop
			} else if total > numberToFind {
				numbers = []int{}
				break
			}
		}
	}
	fmt.Println(numbers)
	var min, max int
	for i, x := range numbers {
		if i == 0 || x > max {
			max = x
		}
		if i == 0 || x < min {
			min = x
		}
	}
	return min + max, nil
}
