package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"path"
	"strings"
)

const sample = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

//go:embed input.txt
var input string

func main() {
	s := bufio.NewScanner(strings.NewReader(input))
	dir := "/"
	fs := make(map[string]int)
	for s.Scan() {
		if strings.HasPrefix(s.Text(), "$ cd") {
			to := s.Text()[len("$ cd "):]
			dir = path.Join(dir, to)
			continue
		}
		if s.Text() == "$ ls" {
			continue
		}
		if strings.HasPrefix(s.Text(), "dir") {
			continue
		}
		var size int
		var f string
		n, err := fmt.Sscanf(s.Text(), "%v %v", &size, &f)
		if err != nil {
			log.Fatalf("scan '%v': %v %v", s.Text(), n, err)
		}
		fs[path.Join(dir, f)] = size
	}
	dirs := make(map[string]int)
	for p, s := range fs {
		for dir := path.Dir(p); dir != ""; dir = path.Dir(dir) {
			dirs[dir] += s
			if dir == "/" {
				// still calculate the size before breaking
				break
			}
		}
	}
	need := 30000000
	total := 70000000
	free := total - dirs["/"]
	target := need - free
	var minp string
	mins := total
	for p, s := range dirs {
		if s > target && s < mins {
			mins = s
			minp = p
		}
	}
	fmt.Println(mins, minp)
}
