package main

type LFUCache struct {
}

func Constructor_LFU(capacity int) LFUCache {
	return LFUCache{}
}

func (this *LFUCache) Get(key int) int {
	return 0
}

func (this *LFUCache) Put(key int, value int) {

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
