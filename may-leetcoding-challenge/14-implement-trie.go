type Trie struct {
	children map[rune]*Trie
	isWord   bool
}
​
/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie)}
}
​
/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, r := range word {
		if _, ok := this.children[r]; !ok {
			this.children[r] = &Trie{children: make(map[rune]*Trie)}
		}
		this = this.children[r]
	}
	this.isWord = true
}
​
/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, r := range word {
		if child, ok := this.children[r]; !ok {
			return false
		} else {
			this = child
		}
	}
​
	return this.isWord
}
​
/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, r := range prefix {
		if child, ok := this.children[r]; !ok {
			return false
		} else {
			this = child
		}
	}
​
	return true
}
​
/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
