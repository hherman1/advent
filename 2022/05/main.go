package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

const sample = `....[D]....
[N].[C]....
[Z].[M].[P]
.1...2...3.

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

type board [][]rune

var bb board

func (b board) String() string {
	var s strings.Builder
	for _, row := range b {
		for _, r := range row {
			if r == '.' {
				s.WriteString("....")
			} else {
				s.WriteString(fmt.Sprintf("[%v].", string(r)))
			}
		}
		s.WriteRune('\n')
	}
	for i := range b[0] {
		s.WriteString(fmt.Sprintf(".%v..", i+1))
	}
	return s.String()
}

func (b board) top(col int) int {
	top := 0
	for ; top < len(b); top++ {
		if b[top][col-1] != '.' {
			return top
		}
	}
	return top
}

// 1 indexed
func (bp *board) move(reps, from, to int) {
	b := *bp
	ftop := bp.top(from)
	ttop := bp.top(to)
	for ttop-reps < 0 {
		// we must expand the board
		var row []rune
		for range b[0] {
			row = append(row, '.')
		}
		*bp = append(b, []rune{})
		b = *bp
		for i := len(b) - 1; i > 0; i-- {
			b[i] = b[i-1]
		}
		b[0] = row
		ttop++
		ftop++
	}
	ftop += reps - 1
	for i := 0; i < reps; i++ {
		b[ttop-1-i][to-1] = b[ftop-i][from-1]
		b[ftop-i][from-1] = '.'
	}
}

//go:embed input.txt
var input string

func main() {
	s := bufio.NewScanner(strings.NewReader(input))
outer:
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		var row []rune
		chars := []rune(s.Text())
		for i := 0; i < len(chars); i += 4 {
			if chars[i+1] == '1' {
				continue outer
			}
			row = append(row, chars[i+1])
		}
		bb = append(bb, row)
	}

	// now process moves

	for s.Scan() {
		var reps, from, to int
		_, err := fmt.Sscanf(s.Text(), "move %v from %v to %v", &reps, &from, &to)
		if err != nil {
			log.Fatalf("parse line '%v': %v", s.Text(), err)
		}
		bb.move(reps, from, to)
	}
	fmt.Println(bb)
	for i := range bb[0] {
		fmt.Print(string(bb[bb.top(i+1)][i]))
		//fmt.Println(i+1, bb.top(i+1))
	}
	fmt.Println()
}
