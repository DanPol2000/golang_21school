package main

import (
	"errors"
	"fmt"
	"log"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func countToys(root *TreeNode, count *int) {
	if root == nil {
		return
	}

	if root.HasToy == true {
		*count += 1
	}

	countToys(root.Left, count)
	countToys(root.Right, count)
}

func areToysBalanced(root *TreeNode) (bool, error) {
	var leftBranchToys int
	var rightBranchToys int

	if root == nil {
		return false, errors.New("root node is nil")
	}

	countToys(root.Left, &leftBranchToys)
	countToys(root.Right, &rightBranchToys)

	if leftBranchToys == rightBranchToys {
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	var root TreeNode

	root.HasToy = true
	root.Left = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right = &TreeNode{HasToy: true}
	root.Right.Left = &TreeNode{HasToy: true}

	garland, err := areToysBalanced(&root)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(garland)
}