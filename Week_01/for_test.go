package Week_01

import (
	"testing"
)

func TestLS(t *testing.T) {
	s := Stack{[]int{1,2,3}}
	t.Log(s.GetTop())
	t.Log(s.Pop())
	t.Log(s.GetTop())
	s.Push(666)
	t.Log(s.GetTop())
}

type Stack struct {
	s []int
}

func (self *Stack) GetTop() int {
	return self.s[len(self.s)-1]
}

func (self *Stack) Pop() int {
	r := self.s[len(self.s)-1]
	self.s = self.s[:len(self.s)-1]
	return r
}

func (self *Stack) Push(v int) {
	self.s = append(self.s, v)
}
