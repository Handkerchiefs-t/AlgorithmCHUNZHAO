package Week_02

import "container/heap"

func getLeastNumbers(arr []int, k int) []int {
	res := []int{}

	x := MyHeap(arr)
	h := (*MyHeap)(nil)
	h = &x
	heap.Init(h)

	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(int))
	}

	return res
}

type MyHeap []int

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
