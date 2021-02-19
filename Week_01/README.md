## 学习笔记

#### 对优先队列的代码的学习

Go 中使用 container/heap 来实现优先队列，要使用 heap，首先要实现一个相关的接口：

```go
type Interface interface {
   sort.Interface
   Push(x interface{}) // add x as element Len()
   Pop() interface{}   // remove and return element Len() - 1.
}

// 其中sort.Interface也是一个接口，相关的定义如下
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

基于上面的接口提供的方法，heap的实现如下：

```go
package heap

import "sort"

type Interface interface {
   sort.Interface
   Push(x interface{}) // add x as element Len()
   Pop() interface{}   // remove and return element Len() - 1.
}

//对堆进行初始化
func Init(h Interface) {
   n := h.Len()
   //堆是使用完全二叉树实现的，所以最后一个非叶子节点是n/2-1，
   //下面的操作就是对所有非叶子节点进行堆化
   for i := n/2 - 1; i >= 0; i-- {
      down(h, i, n)
   }
}

//放入一个元素，然后从下向上堆化
func Push(h Interface, x interface{}) {
   h.Push(x)
   up(h, h.Len()-1)
}

//吐出最小的元素，并维护吐出后的堆
func Pop(h Interface) interface{} {
   n := h.Len() - 1  // n这里-1，让最后一个元素无法被堆化，这个方法很巧妙，实际上就是用h[n]的位置做了临时存储
   h.Swap(0, n)  // 将堆顶放在最后一个叶子，这样从上往下堆化的时候不会造成空洞
   down(h, 0, n)
   return h.Pop()
}

//移除一个元素，并维护堆
func Remove(h Interface, i int) interface{} {
   n := h.Len() - 1
   if n != i {
      h.Swap(i, n)
      if !down(h, i, n) {  // 如果没有向下堆化，就从下往上堆化
         up(h, i)
      }
   }
   return h.Pop()
}

//如果你修改了i这个位置的元素，调用这个方法维护这个堆
func Fix(h Interface, i int) {
   if !down(h, i, h.Len()) {  // 向下堆化 or 向上堆化
      up(h, i)
   }
}

//从下向上堆化，j为堆的最后一个元素
func up(h Interface, j int) {
   for {
      i := (j - 1) / 2 // parent，父节点
      if i == j || !h.Less(j, i) {
         break
      }
      h.Swap(i, j)
      j = i
   }
}

//从上向下堆化
func down(h Interface, i0, n int) bool {
   i := i0
   for {
      j1 := 2*i + 1
      if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
         break
      }
      j := j1 // left child
      if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
         j = j2 // = 2*i + 2  // right child
      }
      if !h.Less(j, i) {
         break
      }
      h.Swap(i, j)
      i = j  // 这个赋值会向下堆化该节点
   }
   return i > i0
}
```

在 learnHeap_test.go 中有一个简单的使用：

```go
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
```



#### 学习碎碎念

- 如何精通一个领域

  - 将知识嚼碎，了解其背后的逻辑
  - 刻意练习
  - 反馈

- 做题的过程

  - 明确题意
  - 考虑各种解题方式，并做比较
  - 编码，这里要注意对边界条件的判断
  - 测试，修改

- 五毒神掌

  1. 知道题目该怎么做
  2. 自己解一遍
  3. 第二天重复一遍
  4. 第五天重复一遍
  5. 在面试前重复一遍

  - tip：五毒神掌和艾宾浩斯遗忘曲线很像，重复会大大提升我们记忆

- 自顶向下的编程方式

  - 先写大逻辑，然后写小逻辑，最后写底层逻辑
  - 这样的写法方便阅读，可以提升代码的可读性
  - 实际上，上面 heap 的实现就遵循了这样的编程方式

- 做题的思路

  - 程序最终是 if、loop、recursion 三种逻辑的组合，所以一个题目既然可以用程序求解，或多或少都会有规律在其中的
  - 双指针：解决特定数组或链表的题目很有用，主要有：
    - 快慢指针：找到或过滤特定的元素
    - 左右夹逼：找最值或找累加数值（接雨水）的时候可以考虑
  - 爬楼梯：经典的斐波那契问题，自底向上使用递推，自顶向下使用递归
  - 三数之和：先排序的操作真的骚气
  - 栈：
    - 如果问题有对称性的特点的时候，可以考虑使用
    - 典型的括号问题
    - 84-柱状图中最大的矩形：一个柱子可以做的最大的面积的宽，是由左边比它小的柱子和右边比它小的柱子确定的。一左一右，有对称的特点，所以，使用递增栈可以解题。
  - 队列：
    - 滑动窗口：遇到类似的题，可以用快慢指针+双端队列试试

