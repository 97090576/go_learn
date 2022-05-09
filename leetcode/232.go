package main

import "fmt"

// 题目要求是用两个栈实现一个队列，这个在剑指 offer 里已经实现过了，所以这里我直接使用切片实现一个队列
type MyQueue struct {
	s []int
}

func Constructor() MyQueue {
	var q MyQueue
	q.s = make([]int, 0)
	return q
}

func (this *MyQueue) Push(x int) {
	this.s = append(this.s, x)
	fmt.Println(this.s)
}

func (this *MyQueue) Pop() int {
	res := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]
	return res
}

func (this *MyQueue) Peek() int {
	res := this.s[len(this.s)-1]
	return res
}

func (this *MyQueue) Empty() bool {
	return len(this.s) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
