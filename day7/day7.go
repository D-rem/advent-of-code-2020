package day7

import (
	"fmt"
	"strings"

	"github.com/D-rem/advent-of-code-2020/lib"
)


type bag struct {
	color     string
	number int
	childrens []bag
}



func getBags(bags map[string]bag, bagToFind string) int{

	count := 0
	for _,v :range bags {
		for color := range v.childrens {
			if color  == bagToFind{
				count++
			} else {

			}
		}
	}
	return count

}

func Part1(bagToFind string) (int, error) {
	t, err := lib.ReadAllLineToString("day7/input.txt")

	if err != nil {
		return 0, err
	}

	var bags map[string][]bag = make(map[string][]bag)
	for _, line := range t {
		if strings.Contains(line, "contain no others bags.") {
			s := strings.Split(line, " contain")
			bags[s[0]]= nil
		} else {
			line = line[:len(line)]
			s := strings.Replace(line, " contain", ",", 1)
			sTab := strings.Split(s, ", ")
			var tone, color string
			fmt.Sscanf(sTab[0], "%s %s bags", &tone, &color)
			key := tone + " " + color
			var c []bags
			for _, child := range sTab[1:] {
				var nbBags int
				fmt.Sscanf(child, "%d %s %s bags", &nbBags, &tone, &color)
				c = append(c, 
					bag{
						color:     tone+" "+color,
						number: nbBags,
						childrens: nil,
					})
		}
		bags[key] = c
	}

	fmt.Printf("%+v", bags)
	return 0, nil

}
