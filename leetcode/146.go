package main

//
//type DoubleLinkedNode struct {
//	val   int
//	left  *DoubleLinkedNode
//	right *DoubleLinkedNode
//}
//
//type LRUCache struct {
//	head *DoubleLinkedNode
//	tail *DoubleLinkedNode
//	m    map[int]int
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		values:   make(map[int]int, capacity),
//		times:    make(map[int]int, capacity),
//		capacity: capacity,
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	v, ok := this.values[key]
//	if ok {
//		return v
//	}
//	return -1
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if len(this.values) == this.capacity {
//		for k, v := range this.times {
//			if v == this.capacity {
//				delete(this.times, k)
//				delete(this.values, k)
//			}
//		}
//	}
//	this.values[key] = value
//	for k := range this.times {
//		this.times[k]++
//	}
//	this.times[key] = 0
//}
//
///**
// * Your LRUCache object will be instantiated and called as such:
// * obj := Constructor(capacity);
// * param_1 := obj.Get(key);
// * obj.Put(key,value);
// */
