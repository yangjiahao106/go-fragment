package main

import "fmt"

type LFUCache struct {
	Cap         int
	Len         int
	key2Node    map[int]*LfuNode
	freq2Bucket map[int]*Bucket

	headBucket *Bucket
	tailBucket *Bucket
}

type Bucket struct {
	Head       *LfuNode
	Tail       *LfuNode
	Freq       int
	PrevBucket *Bucket
	NextBucket *Bucket
}

type LfuNode struct {
	Key  int
	Val  int
	Freq int
	Prev *LfuNode
	Next *LfuNode
}

func testLFU() {

	/**
	 * Your LFUCache object will be instantiated and called as such:
	 * obj := Constructor(capacity);
	 * param_1 := obj.Get(key);
	 * obj.Put(key,value);
	 */

	lfu := Constructor_LFU(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1))
	lfu.Put(3, 3)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))
	lfu.Put(4, 4)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(4))

}

func Constructor_LFU(capacity int) LFUCache {
	c := LFUCache{
		Cap:         capacity,
		key2Node:    map[int]*LfuNode{},
		freq2Bucket: map[int]*Bucket{},
		headBucket:  &Bucket{},
		tailBucket:  &Bucket{},
	}

	c.headBucket.NextBucket = c.tailBucket
	c.tailBucket.PrevBucket = c.headBucket

	return c
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.key2Node[key]
	if !ok {
		return -1
	}

	// 放到node.Freq+1 对应的Bucket 中
	node.Freq += 1
	toBucket, ok := this.freq2Bucket[node.Freq]
	if !ok {
		toBucket = this.creatBucketToPrev(this.freq2Bucket[node.Freq-1], node.Freq)
		this.freq2Bucket[toBucket.Freq] = toBucket
	}

	// remove node from old Bucket,  Bucket 空了需要删除
	this.removeNode(node)
	originBucket := this.freq2Bucket[node.Freq-1]
	if originBucket.Head.Next == originBucket.Tail { // 不能先删除bucket， creatBucketToPrev 要用
		this.removeBucket(originBucket)
		delete(this.freq2Bucket, originBucket.Freq)
	}

	this.addNodeToBucketHead(toBucket, node) // node要先删除再添加，否则没法删除

	return node.Val
}

func (this *LFUCache) Put(key int, value int) {
	node, ok := this.key2Node[key]
	if ok {
		node.Val = value // 更新value
		this.Get(key)    // 更新Freq
		return
	}

	// 先移除 最少最久未使用的key
	if this.Len >= this.Cap {
		lastBucket := this.tailBucket.PrevBucket
		deleteNode := lastBucket.Tail.Prev

		this.removeNode(deleteNode)
		delete(this.key2Node, deleteNode.Key)

		// Bucket 空了需要删除
		if lastBucket.Freq != 1 && lastBucket.Head.Next == lastBucket.Tail {
			this.removeBucket(lastBucket)
			delete(this.freq2Bucket, lastBucket.Freq)
		}
	} else {
		this.Len += 1
	}

	node = &LfuNode{Key: key, Val: value, Freq: 1}

	toBucket, ok := this.freq2Bucket[node.Freq]
	if !ok {
		toBucket = this.creatBucketToPrev(this.tailBucket, node.Freq)
		this.freq2Bucket[toBucket.Freq] = toBucket
	}

	this.addNodeToBucketHead(toBucket, node)
	this.key2Node[key] = node
}

func (this *LFUCache) removeNode(node *LfuNode) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
	node.Prev = nil
	node.Next = nil
}

func (this *LFUCache) addNodeToBucketHead(Bucket *Bucket, node *LfuNode) {
	next := Bucket.Head.Next
	node.Next = next
	next.Prev = node
	Bucket.Head.Next = node
	node.Prev = Bucket.Head
}

func (this *LFUCache) removeBucket(node *Bucket) {
	prev := node.PrevBucket
	next := node.NextBucket
	prev.NextBucket = next
	next.PrevBucket = prev
	node.PrevBucket = nil
	node.NextBucket = nil
}

func (this *LFUCache) creatBucketToPrev(curBucket *Bucket, freq int) *Bucket {
	newBucket := &Bucket{
		Head: &LfuNode{},
		Tail: &LfuNode{},
		Freq: freq,
	}
	newBucket.Head.Next = newBucket.Tail
	newBucket.Tail.Prev = newBucket.Head

	prev := curBucket.PrevBucket
	prev.NextBucket = newBucket
	newBucket.PrevBucket = prev
	newBucket.NextBucket = curBucket
	curBucket.PrevBucket = newBucket

	return newBucket
}
