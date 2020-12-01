package lib

import (
	"bufio"
	"os"
	"strconv"
)

func ReadAllLineToInt(filepath string) ([]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Stack qui s'execute après la destruction du contexte de la fonction courante
	s := bufio.NewScanner(file)
	var result []int
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}
