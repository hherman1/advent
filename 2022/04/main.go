package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

const sample = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

//go:embed input.txt
var input string

func main() {
	s := bufio.NewScanner(strings.NewReader(input))
	c := 0
	for s.Scan() {
		var ls, le, rs, re int
		_, err := fmt.Sscanf(s.Text(), "%v-%v,%v-%v", &ls, &le, &rs, &re)
		if err != nil {
			log.Fatalf("failed to parse line '%v': %w", s.Text(), err)
		}
		if le < rs || ls > re {
			continue
		}
		c++
	}
	fmt.Println(c)
}
