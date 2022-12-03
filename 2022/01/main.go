package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var list string

func main() {
	s := bufio.NewScanner(strings.NewReader(list))
	cur := 0
	var vals []int
	for s.Scan() {
		if s.Text() == "" {
			vals = append(vals, cur)
			cur = 0
			continue
		}
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("parse %v: %v", s.Text(), err)
		}
		cur += val
	}
	sort.Ints(vals)

	out := 0
	for _, v := range vals[len(vals)-3:] {
		out += v
	}
	fmt.Println(out)
}
