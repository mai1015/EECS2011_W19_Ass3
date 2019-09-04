package main

import (
	"fmt"
	"math"
)


func parent(e int) (int) {
	return (e-1)/2
}

func left(e int) (int) {
	return (e*2)+1
}

func right(e int) (int) {
	return (e*2)+2
}

func compare(a int, b int) (int) {
	return b - a;
}

func swap(tree []int, a int, b int) {
	tem := tree[a]
	tree[a] = tree[b]
	tree[b] = tem
}

func downheap(tree []int, e int) {
	for left(e) < len(tree){
		l, r := left(e), right(e)
		small := l
		if r < len(tree) && compare(tree[l], tree[r]) > 0{
			small = r
		}
		if compare(tree[small], tree[e]) >= 0 {
			break
		} 
		swap(tree, small, e)
		e = small
	}
}

func construct(tree []int) {
	for i := len(tree) - 1; i >= 0; i-- {
		downheap(tree, i)
	}
}

func testTree() {
	heap := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(heap)
	construct(heap)
	fmt.Println(heap)
	n := int(math.Log2(float64(len(heap))))
	for i := 0; i < n; i++ {
		fmt.Println(heap[0])
		heap[0] = heap[len(heap) - 1]
		heap = heap[0:len(heap) - 1]
		downheap(heap, 0)
		fmt.Println(heap)
	}
}

func main() {
	testTree()
}