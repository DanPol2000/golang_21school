package main

import (
	"fmt"
	"log"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func fromRightToLeft(queue []TreeNode) []TreeNode {
	var nextQueue []TreeNode

	for len(queue) != 0 {
		x := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if x.Right != nil {
			nextQueue = append(nextQueue, *x.Right)
		}

		if x.Left != nil {
			nextQueue = append(nextQueue, *x.Left)
		}
	}

	return nextQueue
}

func fromLeftToRight(queue []TreeNode) []TreeNode {
	var nextQueue []TreeNode

	for len(queue) != 0 {
		x := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if x.Left != nil {
			nextQueue = append(nextQueue, *x.Left)
		}

		if x.Right != nil {
			nextQueue = append(nextQueue, *x.Right)
		}
	}

	return nextQueue
}

func getGarland(queue []TreeNode) []bool {
	var garland []bool

	for _, q := range queue {
		garland = append(garland, q.HasToy)
	}

	return garland
}

func unrollGarland(root *TreeNode) ([]bool, error) {
	var queue []TreeNode
	var garland []bool

	queue = append(queue, *root)

	for i := 0; ; i++ {
		garland = append(garland, getGarland(queue)...)

		if i%2 == 0 {
			queue = fromLeftToRight(queue)
		} else {
			queue = fromRightToLeft(queue)
		}

		if queue == nil {
			break
		}
	}
	return garland, nil
}

func main() {
	var root TreeNode

	root.HasToy = true
	root.Left = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right = &TreeNode{HasToy: false}
	root.Right.Left = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: true}

	garland, err := unrollGarland(&root)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(garland)
}