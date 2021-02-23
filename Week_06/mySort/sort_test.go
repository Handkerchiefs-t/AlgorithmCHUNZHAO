package mySort

import (
	"container/heap"
	"testing"
)

//快排
func quickSort(arr []int, begin, end int, useTimer *bool) {

	if begin >= end {
		return
	}
	p := moveAndGetPivot(arr, begin, end)
	quickSort(arr, begin, p, useTimer)
	quickSort(arr, p+1, end, useTimer)
}

func moveAndGetPivot(arr []int, begin, end int) int {
	quick := begin + 1
	slow := begin + 1
	pivot := arr[begin]

	for ; quick < end; quick++ {
		if arr[quick] < pivot {
			arr[quick], arr[slow] = arr[slow], arr[quick]
			slow++
		}
	}

	arr[slow-1], arr[begin] = arr[begin], arr[slow-1]

	return slow-1
}



//归并排序
func mergeSort(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	mid := (begin + end) >> 1
	mergeSort(arr, begin, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, begin, mid, end)
}

func merge(arr []int, begin, mid, end int) {
	tmp := make([]int, 0, end-begin)
	i := begin
	j := mid

	for i < mid && j < end {
		if arr[i] < arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}

	for i < mid {
		tmp = append(tmp, arr[i])
		i++
	}

	for j < end {
		tmp = append(tmp, arr[j])
		j++
	}

	for i, v := range tmp {
		arr[begin+i] = v
	}
}



//堆排序
func heapSort(arr []int) {
	h := heapType(arr)
	tmp := make([]int, 0, len(arr))
	heap.Init(&h)

	for h.Len() > 0 {
		tmp = append(tmp, heap.Pop(&h).(int))
	}

	copy(arr, tmp)
}

type heapType []int

func (h *heapType) Len() int {
	return len(*h)
}

func (h *heapType) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *heapType) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heapType) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapType) Pop() interface{} {
	res := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]

	return res
}


func Test00001(t *testing.T) {
	x := []int{1,0,6,0,7,0}
	heapSort(x)
	//mergeSort(x, 0, len(x))
	//c := true
	//quickSort(x, 0, len(x), &c)
	t.Log(x)
}
