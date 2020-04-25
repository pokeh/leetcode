type LRUCache struct {
	cache    map[int]int
	keyOrder []int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]int),
		keyOrder: make([]int, 0, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.cache[key]
	if !ok {
		return -1
	}

	// update cache order
	for i, k := range this.keyOrder {
		if k == key {
			this.keyOrder = append(this.keyOrder[:i], this.keyOrder[i+1:]...)
		}
	}
	this.keyOrder = append(this.keyOrder, key)

	return val
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.cache[key]
	if ok {
		// value exists in cache
		// remove from keyOrder
		for i, k := range this.keyOrder {
			if k == key {
				this.keyOrder = append(this.keyOrder[:i], this.keyOrder[i+1:]...)
			}
		}
	} else {
		// value doesn't exist in cache
		// if cache is at capacity, remove oldest key
		if len(this.keyOrder) == this.capacity {
			oldest := this.keyOrder[0]
			this.keyOrder = this.keyOrder[1:]
			delete(this.cache, oldest)
		}
	}

	// add new key as most recent
	this.cache[key] = value
	this.keyOrder = append(this.keyOrder, key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
