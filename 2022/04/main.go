package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const sample = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func secs(r string) (int, int) {
	s, e, ok := strings.Cut(r, "-")
	if !ok {
		log.Fatalf("failed to cut -: %v", r)
	}
	si, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("convert start: %v", s)
	}
	ei, err := strconv.Atoi(e)
	if err != nil {
		log.Fatalf("convert end: %v", s)
	}
	return si, ei
}

//go:embed input.txt
var input string

func main() {
	s := bufio.NewScanner(strings.NewReader(input))
	c := 0
	for s.Scan() {
		l, r, ok := strings.Cut(s.Text(), ",")
		if !ok {
			log.Fatalf("failed to cut: %v", s.Text())
		}
		ls, le := secs(l)
		rs, re := secs(r)
		if ((ls < rs) && (le < rs)) || ((ls > re) && (le > re)) {
			continue
		}
		c++
	}
	fmt.Println(c)
}
