package day2

import (
	"fmt"
	"strings"

	"github.com/D-rem/advent-of-code-2020/lib"
)

func Part1() (int, error) {
	var result int
	var m, err = initMap()

	if err != nil {
		return 0, err
	}

	for key, element := range m {
		var min, max int
		var letter string

		fmt.Sscanf(key, "%d-%d %s", &min, &max, &letter)

		for _, y := range element {

			if (strings.Count(y, letter) >= min) && (strings.Count(y, letter) <= max) {
				result++
			}
		}
	}
	return result, nil
}

func Part2() (int, error) {
	var result int
	var m, err = initMap()

	if err != nil {
		return 0, err
	}

	for key, element := range m {
		var min, max int
		var letter string

		fmt.Sscanf(key, "%d-%d %s", &min, &max, &letter)

		for _, y := range element {
			r := []rune(y)
			if (strings.ContainsRune(letter, r[min-1])) && !(strings.ContainsRune(letter, r[max-1])) {
				result++
			}
			if !(strings.ContainsRune(letter, r[min-1])) && (strings.ContainsRune(letter, r[max-1])) {
				result++
			}
		}
	}
	return result, nil
}

func initMap() (map[string][]string, error) {
	m := make(map[string][]string)

	t, err := lib.ReadAllLineToString("day2/input.txt")

	if err != nil {
		return nil, err
	}
	for _, x := range t {
		key := strings.Split(x, ": ")
		m[key[0]] = append(m[key[0]], key[1])
	}
	return m, nil

}
