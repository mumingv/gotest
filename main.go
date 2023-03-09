package main

import (
	"fmt"

	"github.com/mumingv/golib/leetcode"
)

type TreeNode = leetcode.TreeNode

func main() {
	a := []any{3, 9, 20, nil, nil, 15, 7}
	fmt.Println(a)
	root := leetcode.NewTree(a)
	root.Print()
}
