package Week_05

type Trie struct {
	isEnd bool
	nextList [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	root := this

	for _, char := range word {
		if root.nextList[char - 'a'] == nil {
			root.nextList[char - 'a'] = &Trie{}
		}

		root = root.nextList[char - 'a']
	}

	root.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this

	for _, char := range word {
		if root.nextList[char - 'a'] == nil {
			return false
		}

		root = root.nextList[char - 'a']
	}

	return root.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this

	for _, char := range prefix {
		if root.nextList[char - 'a'] == nil {
			return false
		}

		root = root.nextList[char - 'a']
	}

	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
