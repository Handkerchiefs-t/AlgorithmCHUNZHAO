package Week_02

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	m := map[int]int{}
	h := &MyHeap{}
	heap.Init(h)

	// 使用 map 统计频次
	for _, value := range nums {
		m[value]++
	}

	// 加入堆
	for value, count := range m {
		heap.Push(h, Node{
			value: value,
			count: count,
		})
	}

	// 从堆中取k个值
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(Node).value)
	}

	return res
}

type Node struct{
	value int
	count int
}

//定义一个数值为int的heap
type MyHeap []Node

func (h MyHeap) Len() int {
	return len(h)
}

func (h MyHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MyHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
