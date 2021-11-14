/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		index := ch - 'a'
		if node.next[index] == nil {
			node.next[index] = &Trie{}
		}
		node = node.next[index]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	n := this.SearchPrefix(word)
	return n != nil && n.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func (this *Trie) SearchPrefix(s string) *Trie {
	node := this
	for _, ch := range s {
		index := ch - 'a'
		if node.next[index] == nil {
			return nil
		}
		node = node.next[index]
	}

	return node
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

