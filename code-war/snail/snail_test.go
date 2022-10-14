package codewar

import (
	"fmt"
	"reflect"
	"testing"
)

type point struct {
	x, y int
}

func (p point) at(m [][]int) bool {
	if p.x < 0 || p.x >= len(m[0]) {
		return false
	}
	if p.y < 0 || p.y >= len(m) {
		return false
	}

	return true
}

func (p point) add(p2 point) point {
	return point{p.x + p2.x, p.y + p2.y}
}

type dirsChain struct {
	val  point
	next *dirsChain
}

func newDirsChain() *dirsChain {
	right := &dirsChain{
		val: point{
			x: 1,
			y: 0,
		},
		next: nil,
	}
	down := &dirsChain{
		val: point{
			x: 0,
			y: 1,
		},
		next: nil,
	}
	left := &dirsChain{
		val: point{
			x: -1,
			y: 0,
		},
		next: nil,
	}
	up := &dirsChain{
		val: point{
			x: 0,
			y: -1,
		},
		next: nil,
	}

	right.next = down
	down.next = left
	left.next = up
	up.next = right

	return right
}

func Snail(snaipMap [][]int) []int {
	if len(snaipMap) == 0 {
		return []int{}
	}
	len_x := len(snaipMap[0])
	len_y := len(snaipMap)
	res := make([]int, 0, len_x*len_y)
	dirs := newDirsChain()
	explorMap := make([][]int, len_y)
	for i := range explorMap {
		explorMap[i] = make([]int, len_x)
	}

	p := point{0, 0}
	res = append(res, snaipMap[0][0])
	explorMap[0][0] = 1
	for len(res) < len_x*len_y {
		tmp := p
		p = p.add(dirs.val)

		if p.at(snaipMap) && explorMap[p.y][p.x] == 0 {
			res = append(res, snaipMap[p.y][p.x])
			explorMap[p.y][p.x] = 1
		} else {
			dirs = dirs.next
			p = tmp
		}
	}
	return res
}

func TestSnail(t *testing.T) {
	snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	res := Snail(snailMap)
	want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	fmt.Println(res)
	fmt.Println(reflect.DeepEqual(res, want))

	snailMap = [][]int{}

	res = Snail(snailMap)
	want = []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	fmt.Println(res)
	fmt.Println(reflect.DeepEqual(res, want))
}
