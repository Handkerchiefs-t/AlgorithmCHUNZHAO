package Week_01

import (
	"container/heap"
	"testing"
)

//定义一个数值为int的heap
type MyHeap []int

//要使用heap，首先要满足一个指定的接口,你需要定义接口中的方法
func (h MyHeap) Len() int {
	return len(h)
}

func (h MyHeap) Less(i, j int) bool {
	return h[i] < h[j]  // 使用<可以构建小顶堆，使用>可以构建大顶堆
}

func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MyHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func Test001(t *testing.T) {
	h := &MyHeap{1,56,98,87,5,98,54,54,51,100}
	heap.Init(h)
	heap.Push(h, 129)
	t.Log(heap.Pop(h))  // 1
	t.Log(heap.Pop(h))  // 5
	t.Log(*h)  // [51 54 54 87 56 98 98 129 100]
}


