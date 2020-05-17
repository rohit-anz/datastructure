package main

import (
	"fmt"
	"math"
)

func graph() [8][8]int {
	return [8][8]int{{0,0,0,0,0,0,0,0},
		{0,0,25,0,0,0,5,0},
		{0,25,0,12,0,0,0,0},
		{0,0,12,0,8,0,0,0},
		{0,0,0,8,0,16,0,14},
		{0,0,0,0,16,0,20,18},
		{0,5,0,0,0,20,0,0},
		{0,0,10,0,14,18,0,0}}
}

func main() {
	near := [8]int{math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8}
	t := [2][8]int{}
	list := graph()
	min := math.MaxInt8
	var u int
	var v int
	for i:=1; i<8;i++ {
		for j:= i; j<8; j++ {
			if list[i][j] < min && list[i][j] != 0 {
				min = list[i][j]
				u = i
				v = j
			}
		}
	}
	near[u] = 0
	near[v] = 0

	t[0][0] = u
	t[1][0] = v

	for i:=1; i<8; i++ {
		if list[i][u] < list[i][v] && near[i] != 0 {
			near[i] = u
		} else {
			near[i] = v
		}
	}
	var k int
	for i:=1; i<7; i++ {
		min := math.MaxInt8
		for j:=1; j<8; j++ {
			if near[j] != 0 && list[j][near[j]] < min {
				min = list[j][near[j]]
				k = j
			}
		}
		t[0][i] = k
		t[1][i] = near[k]
		near[k] = 0
		for j:=1; j<8; j++ {
			if near[j] != 0 && list[j][k] < list[j][near[j]] {
				near[j] = k
			}
		}
	}

	for i:=0; i<6; i++ {
		fmt.Printf("\n %v %v", t[0][i], t[1][i])
	}
}
