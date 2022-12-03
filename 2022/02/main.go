package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
	Win      = 6
	Loss     = 0
	Draw     = 3
)

func beats(win, lose int) func(int, int) (int, int) {
	return func(left, right int) (int, int) {
		if left == win && right == lose {
			return Win, Loss
		}
		if right == win && left == lose {
			return Loss, Win
		}
		return Draw, Draw
	}
}

func game(rules ...func(int, int) (int, int)) func(int, int) (int, int) {
	return func(left, right int) (int, int) {
		for _, r := range rules {
			l, r := r(left, right)
			if l != Draw {
				return left + l, right + r
			}
		}
		return left + Draw, right + Draw
	}
}

func mover(rules ...func(int, int) (int, int)) func(int, int) int {
	moves := []int{Rock, Paper, Scissors}
	return func(opponent, outcome int) int {
		for _, move := range moves {
			mres := Draw
			for _, rule := range rules {
				_, res := rule(opponent, move)
				if res != Draw {
					mres = res
					break
				}
			}
			if mres == outcome {
				return move
			}
		}
		log.Fatalf("outcome %v cannot be achieved against move %v", outcome, opponent)
		return 0
	}
}

var round = game(beats(Rock, Scissors), beats(Scissors, Paper), beats(Paper, Rock))
var move = mover(beats(Rock, Scissors), beats(Scissors, Paper), beats(Paper, Rock))

//go:embed input.txt
var list string

func main() {
	s := bufio.NewScanner(strings.NewReader(list))
	score := 0
	for s.Scan() {
		o, me, ok := strings.Cut(s.Text(), " ")
		if !ok {
			log.Fatalf("no space: '%v'", s.Text())
		}
		oplay := int([]rune(o)[0]-'A') + 1
		outcome := int([]rune(me)[0]-'X') * 3
		mplay := move(oplay, outcome)
		_, mscore := round(oplay, mplay)
		score += mscore
	}
	fmt.Println(score)
}
