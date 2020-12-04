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

func ReadAllLineToString(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Stack qui s'execute après la destruction du contexte de la fonction courante
	s := bufio.NewScanner(file)
	var result []string
	for s.Scan() {
		if err != nil {
			return nil, err
		}
		result = append(result, s.Text())
	}
	return result, nil
}

func ReadAllLineToRunesSlice(filepath string) ([][]rune, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Stack qui s'execute après la destruction du contexte de la fonction courante
	s := bufio.NewScanner(file)
	var result [][]rune
	for s.Scan() {
		if err != nil {
			return nil, err
		}
		result = append(result, []rune(s.Text()))
	}
	return result, nil
}

//func ReadAllLineToMap(filepath string) ([]map[string][]string , error) {
//	file, err := os.Open(filepath)
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close() // Stack qui s'execute après la destruction du contexte de la fonction courante
//	s := bufio.NewScanner(file)
//	var result [][]rune
//	for s.Scan() {
//		if err != nil {
//			return nil, err
//		}
//		result = append(result, []rune(s.Text()))
//	}
//	return result, nil
//}
