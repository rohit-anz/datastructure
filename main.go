package main

import (
	"fmt"
	"github.com/rohitjb/datastructure/binaryTree"
)

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
}