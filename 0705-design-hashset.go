const tableSize = 10

type MyHashSet struct {
	hashTable [tableSize]*Bucket
}

type Bucket []int

func hash(key int) int {
	return key % tableSize
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	h := MyHashSet{}
	
	for i := 0; i < tableSize; i++ {
		b := make(Bucket, 0)
		h.hashTable[i] = &b
	}
	
	return h
}

func (this *MyHashSet) Add(key int) {
	bucket := this.hashTable[hash(key)]
	if !bucket.contains(key) {
		*bucket = append(*bucket, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	bucket := this.hashTable[hash(key)]
	this.hashTable[hash(key)] = removeKeyFromBucket(key, bucket)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	bucket := this.hashTable[hash(key)]
	return bucket.contains(key)
}

func (b *Bucket) contains(key int) bool {
	for _, k := range *b {
		if key == k {
			return true
		}
	}
	return false
}

func removeKeyFromBucket(key int, b *Bucket) *Bucket {
	bucket := *b
	for i, k := range *b {
		if key == k {
			bucket = append(bucket[:i], bucket[i+1:]...)
			return &bucket
		}
	}
	return b
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
