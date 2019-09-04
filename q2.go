package main

import (
	"fmt"
	"math/rand"
)
// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Tree definition
// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(20) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

// wirtten by Zhili Mai
// ID: 25234842
func max(a int, b int) (int){
    if a > b {
        return a
    }
	return b
}

func distanceHelper(a *Tree, b *Tree, root *Tree) (int, int) {
	left, right := -1, -1
	if root == a {left = 0}
	if root == b {right = 0}
	if left == 0 && right == 0 {return left, right} // same root

	// var ll, lr, rl, rr int
	if root.Left != nil {
		l, r := distanceHelper(a, b, root.Left)
		left = max(left, l)
		right = max(right, r)
	}
	if left >= 0 && right >= 0 {return left, right} // found two

	if root.Right != nil {
		l, r := distanceHelper(a, b, root.Right)
		left = max(left, l)
		right = max(right, r)
	}
	if left >= 0 && right >= 0 {return left, right} // found two

	// increase distance if other is not found
	if left >= 0 {left++}
	if right >= 0 {right++}
	return left, right
}

func distance(a *Tree, b *Tree, root *Tree) int {
	left, right := distanceHelper(a, b, root)
	if left == -1 || right == -1 {return -1}
	return left + right
}

func maxDistHelper(root *Tree) (int, int, int){
	l1, r1, l2, r2, d1, d2 := 0, 0, 0, 0, 0, 0
	if root.Left != nil {
		l1, r1, d1 = maxDistHelper(root.Left)
        d1 += 1
	}
	if root.Right != nil {
		l2, r2, d2 = maxDistHelper(root.Right)
        d2 += 1
	}
	maxdist := max(l1, r1) + d1 + max(l2, r2) + d2
	if maxdist > max(l1 + r1, l2 + r2) {
		return max(l1, r1) + d1, max(l2, r2) + d2, 0
	}
	if l1 + r1 > l2 + r2 {
		return l1, r1, d1
	} else if l1 + r1 < l2 + r2 {
		return l2, r2, d2
	}
	return max(l1, r1) + d1, max(l2, r2) + d2, 0
}

func maxDist(root *Tree) int {
	l, r, _ := maxDistHelper(root)
	return l + r
}

func build() (*Tree){
	root := &Tree{nil, 10, nil}
	t := []int{4,8,6,2,30,18,14,12,16,26,24,28,36,34,38}
	for _, v := range t {
		insert(root, v)
	}
	return root
}

func testDistance(root *Tree) {
	// var root *Tree
	var a, b *Tree
	a = root.Right.Left.Left.Right
	b = root.Left.Left
	// b = root.Right.Right.Right
	fmt.Println(root.String())
	fmt.Println(a.Value)
	fmt.Println(b.Value)
	fmt.Println(distance(a, b, root))
	fmt.Println(distance(a, nil, root))
}

func main() {
	root := build()
	fmt.Println(root.String())
	fmt.Printf("length: %d", maxDist(root))
}