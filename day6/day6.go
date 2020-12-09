package day6

import (
	"fmt"

	"github.com/D-rem/advent-of-code-2020/lib"
)

type question struct {
	a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z, voters int
}

type questions []question

func newQuestion() question {
	return question{}
}

func Part1() (int, error) {
	t, err := lib.ReadAllLineToString("day6/input.txt")

	if err != nil {
		return 0, err
	}

	var questions []question
	question := question{}

	for _, line := range t {
		if line == "" {
			questions = append(questions, question)
			question = newQuestion()
			continue
		}
		for _, response := range []rune(line) {
			question.voters++
			switch response {
			case 'a':
				question.a++
			case 'b':
				question.b++
			case 'c':
				question.c++
			case 'd':
				question.d++
			case 'e':
				question.e++
			case 'f':
				question.f++
			case 'g':
				question.g++
			case 'h':
				question.h++
			case 'i':
				question.i++
			case 'j':
				question.j++
			case 'k':
				question.k++
			case 'l':
				question.l++
			case 'm':
				question.m++
			case 'n':
				question.n++
			case 'o':
				question.o++
			case 'p':
				question.p++
			case 'q':
				question.q++
			case 'r':
				question.r++
			case 's':
				question.s++
			case 't':
				question.t++
			case 'u':
				question.u++
			case 'v':
				question.v++
			case 'w':
				question.w++
			case 'x':
				question.x++
			case 'y':
				question.y++
			case 'z':
				question.z++
			}
		}

	}
	question.voters++
	questions = append(questions, question)
	count := 0

	for _, q := range questions {
		//fmt.Println(countValidAnswers(&q))
		count += countValidAnswers(&q)
	}

	return count, nil

}

func countValidAnswers(q *question) int {
	count := 0
	if q.a > 0 {
		count++
	}
	if q.b > 0 {
		count++
	}
	if q.c > 0 {
		count++
	}
	if q.d > 0 {
		count++
	}
	if q.e > 0 {
		count++
	}
	if q.f > 0 {
		count++
	}
	if q.g > 0 {
		count++
	}
	if q.h > 0 {
		count++
	}
	if q.i > 0 {
		count++
	}
	if q.j > 0 {
		count++
	}
	if q.k > 0 {
		count++
	}
	if q.l > 0 {
		count++
	}
	if q.m > 0 {
		count++
	}
	if q.n > 0 {
		count++
	}
	if q.o > 0 {
		count++
	}
	if q.p > 0 {
		count++
	}
	if q.q > 0 {
		count++
	}
	if q.r > 0 {
		count++
	}
	if q.s > 0 {
		count++
	}
	if q.t > 0 {
		count++
	}
	if q.u > 0 {
		count++
	}
	if q.v > 0 {
		count++
	}
	if q.w > 0 {
		count++
	}
	if q.x > 0 {
		count++
	}
	if q.y > 0 {
		count++
	}
	if q.z > 0 {
		count++
	}

	return count
}

func countAllValidAnswers(q *question) int {
	count := 0
	if q.a == q.voters {
		count++
	}
	if q.b == q.voters {
		count++
	}
	if q.c == q.voters {
		count++
	}
	if q.d == q.voters {
		count++
	}
	if q.e == q.voters {
		count++
	}
	if q.f == q.voters {
		count++
	}
	if q.g == q.voters {
		count++
	}
	if q.h == q.voters {
		count++
	}
	if q.i == q.voters {
		count++
	}
	if q.j == q.voters {
		count++
	}
	if q.k == q.voters {
		count++
	}
	if q.l == q.voters {
		count++
	}
	if q.m == q.voters {
		count++
	}
	if q.n == q.voters {
		count++
	}
	if q.o == q.voters {
		count++
	}
	if q.p == q.voters {
		count++
	}
	if q.q == q.voters {
		count++
	}
	if q.r == q.voters {
		count++
	}
	if q.s == q.voters {
		count++
	}
	if q.t == q.voters {
		count++
	}
	if q.u == q.voters {
		count++
	}
	if q.v == q.voters {
		count++
	}
	if q.w == q.voters {
		count++
	}
	if q.x == q.voters {
		count++
	}
	if q.y == q.voters {
		count++
	}
	if q.z == q.voters {
		count++
	}

	return count
}

func Part2() (int, error) {
	t, err := lib.ReadAllLineToString("day6/input.txt")

	if err != nil {
		return 0, err
	}

	var questions []question
	question := newQuestion()

	for _, line := range t {
		if line == "" {
			questions = append(questions, question)
			question = newQuestion()
			continue
		}
		for _, response := range []rune(line) {
			switch response {
			case 'a':
				question.a++
			case 'b':
				question.b++
			case 'c':
				question.c++
			case 'd':
				question.d++
			case 'e':
				question.e++
			case 'f':
				question.f++
			case 'g':
				question.g++
			case 'h':
				question.h++
			case 'i':
				question.i++
			case 'j':
				question.j++
			case 'k':
				question.k++
			case 'l':
				question.l++
			case 'm':
				question.m++
			case 'n':
				question.n++
			case 'o':
				question.o++
			case 'p':
				question.p++
			case 'q':
				question.q++
			case 'r':
				question.r++
			case 's':
				question.s++
			case 't':
				question.t++
			case 'u':
				question.u++
			case 'v':
				question.v++
			case 'w':
				question.w++
			case 'x':
				question.x++
			case 'y':
				question.y++
			case 'z':
				question.z++
			}
		}
		question.voters++

	}
	questions = append(questions, question)
	count := 0

	for _, q := range questions {
		fmt.Printf("%+v\n", q)
		count += countAllValidAnswers(&q)
	}

	return count, nil

}
