package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strings"

	"golang.org/x/exp/slices"
)

const sample = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func priority(r rune) int {
	if r >= 'a' {
		return int(r-'a') + 1
	}
	return int(r-'A') + 27
}

//go:embed input.txt
var input string

// Given sorted rune lists, returns the set of runes that appeared in each list at least once.
func intersect(a, b []rune) []rune {
	var is []rune
	bi := 0
	for ai := 0; ai < len(a); ai++ {
		for bi < len(b) && b[bi] < a[ai] {
			bi++
		}
		if bi >= len(b) {
			break
		}
		if b[bi] == a[ai] {
			is = append(is, a[ai])
		} else {
			continue
		}
		for bi < len(b) && b[bi] == a[ai] {
			bi++
		}
	}
	return is
}

func main() {
	s := bufio.NewScanner(strings.NewReader(input))
	score := 0
	var group [][]rune
	for s.Scan() {
		group = append(group, []rune(s.Text()))
		if len(group) < 3 {
			continue
		}
		slices.Sort(group[0])
		slices.Sort(group[1])
		slices.Sort(group[2])
		badges := intersect(intersect(group[0], group[1]), group[2])
		if len(badges) != 1 {
			log.Fatalf("intersect(intersect(%v, %v), %v)=%v", string(group[0]), string(group[1]), string(group[2]), string(badges))
		}
		score += priority(badges[0])
		group = group[:0]
	}
	fmt.Println(score)
}
