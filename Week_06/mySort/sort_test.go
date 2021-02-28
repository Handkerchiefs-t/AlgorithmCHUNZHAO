package mySort

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

//快排
func quickSort(arr []int, begin, end int) {
	if begin >= end {
		return
	}

	pivotIndex := moveAndGetPivot(arr, begin, end)
	quickSort(arr, begin, pivotIndex)
	quickSort(arr, pivotIndex+1, end)
}

func moveAndGetPivot(arr []int, begin, end int) int {
	fast := begin+1
	slow := begin
	pivot := arr[begin]

	for ; fast < end; fast++ {
		if arr[fast] <= pivot {
			slow++
			arr[slow], arr[fast] = arr[fast], arr[slow]
		}
	}

	arr[begin], arr[slow] = arr[slow], arr[begin]

	return slow
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
	i := begin
	j := mid+1
	tmp := make([]int, 0, end-begin+1)

	for i <= mid && j <= end {
		if arr[i] < arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}

	for i <= mid {
		tmp = append(tmp, arr[i])
		i++
	}

	for j <= end {
		tmp = append(tmp, arr[j])
		j++
	}

	fmt.Println(tmp, len(tmp), cap(tmp))
	i = begin
	for _, v := range tmp {
		fmt.Println(arr, i, begin, mid, end)
		arr[i] = v
		i++
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
	x := []int{1,3,2,3,1}
	//heapSort(x)
	//mergeSort(x, 0, len(x)-1)
	//c := true
	quickSort(x, 0, len(x))
	t.Log(x)

	reflect.Copy()
}
