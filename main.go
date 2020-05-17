package main

import (
	"fmt"
	"github.com/rohitjb/datastructure/binaryTree"
)

var t [5][8]int

func main() {

	root := binaryTree.Insert(nil, 30)
	binaryTree.Insert(root,40)
	binaryTree.Insert(root,20)
	binaryTree.Insert(root,10)
	binaryTree.Insert(root,25)
	binaryTree.Insert(root,35)
	binaryTree.Insert(root,45)
	binaryTree.Insert(root,42)
	binaryTree.Insert(root,43)

	root = binaryTree.DeleteNode(root, 40)
	fmt.Println(root)

/*
	binaryTree.LevelOrder(root)
	fmt.Println(binaryTree.SumOfAllNode(root))
	fmt.Println(binaryTree.CountNumberOfNode(root))
	fmt.Println(binaryTree.HeightOfTree(root))
*/
/*
	for i:= 0; i<5; i++ {
		for j:= 0; j<8;j++ {
			t[i][j] = -1
		}
	}

	wt := []int{1, 3, 4, 5}
	val := []int{1, 4, 5, 7}
	w := 7

	fmt.Println(knapsack(wt, val, w, 4))
*/
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func knapsack(wt []int, val []int, w int, n int) int {
	if n == 0 || w == 0 {
		return 0
	}

	if t[n][w] != -1 {
		return t[n][w]
	}

	if wt[n-1] <= w {
		t[n][w] = Max(val[n-1] + knapsack(wt, val, w-wt[n-1], n-1), knapsack(wt, val, w, n-1))
		return t[n][w]
	} else {
		t[n][w] = knapsack(wt, val, w, n-1)
		return t[n][w]
	}
}