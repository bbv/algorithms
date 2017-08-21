package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Maze struct {
	Width  int
	Height int
	Map    [][]int
}

type Point struct {
	X int
	Y int
}

var nextMoves [][]int = [][]int{[]int{-1, 0}, []int{0, 1}, []int{1, 0}, []int{0, -1}}

func (p Point) Equals(p2 Point) bool {
	if p.X == p2.X && p.Y == p2.Y {
		return true
	}
	return false
}

func (m *Maze) Traverse(start Point, finish Point, path []Point) ([]Point, error) {
	if start.Equals(finish) {
		return path, nil
	}
	next, err := m.findNextMoves(start, path)
	if err != nil {
		return nil, err
	}
	log.Println(next)
	npath := make([]Point, len(path))
	copy(npath, path)
	for _, n := range next {
		p, err := m.Traverse(n, finish, npath)
		if err == nil {
			return p, nil
		}
	}

	return path, nil
}

func (m *Maze) findNextMoves(start Point, path []Point) ([]Point, error) {
	res := make([]Point, 0)
	for _, n := range nextMoves {
		next := Point{X: start.X + n[0], Y: start.Y + n[1]}
		if m.isPointValid(next) && notVisited(next, path) {
			res = append(res, next)
		}
	}

	return res, nil
}

func (m *Maze) isPointValid(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < m.Width && p.Y < m.Height
}

func notVisited(p Point, path []Point) bool {
	for _, pp := range path {
		if p.Equals(pp) {
			return false
		}
	}
	return true
}

func readMazeFromFile(fname string) (Maze, error) {
	str, err := ioutil.ReadFile(fname)
	if err != nil {
		return Maze{}, err
	}
	log.Println("\n" + string(str))

	mmap := make([][]int, 0)
	var mrow []int
	for _, c := range str {
		if mrow == nil {
			mrow = make([]int, 0)
		}
		if c == '\n' {
			if len(mrow) > 0 {
				mmap = append(mmap, mrow)
			}
			mrow = make([]int, 0)
			continue
		}
		mrow = append(mrow, int(c)-'0')
	}
	if len(mrow) > 0 {
		mmap = append(mmap, mrow)
	}
	return Maze{Map: mmap, Width: len(mmap), Height: len(mmap[0])}, nil
}

func main() {
	if len(os.Args) == 1 {

	}
	filename := os.Args[1]
	maze, err := readMazeFromFile(filename)
	if err != nil {
		log.Println(err)
	}
	log.Println(maze)
	path, err := maze.Traverse(Point{0, 0}, Point{19, 9}, []Point{})
	if err != nil {
		log.Println(err)
	}
	log.Println(path)
}
