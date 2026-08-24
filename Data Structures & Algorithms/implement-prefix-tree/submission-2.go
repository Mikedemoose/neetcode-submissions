type PrefixTree struct {
	letters []*PrefixTree
	isEnd bool
}

func Constructor() PrefixTree {
    return PrefixTree {
		letters: make([]*PrefixTree, 26),
	}
}

func (this *PrefixTree) Insert(word string) {
	temp := this
	for _, r := range word {
		if temp.letters[r-'a'] != nil {
			temp = temp.letters[r-'a']
		} else {
			newTree := Constructor()
			temp.letters[r-'a'] = &newTree
			temp = temp.letters[r-'a']
		}
	}
	temp.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	temp := this
	for _, r := range word {
		if temp.letters[r-'a'] == nil {
			return false
		}
		temp = temp.letters[r-'a']
	}
	return temp.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	temp := this
	for _, r := range prefix {
		if temp.letters[r-'a'] == nil {
			return false
		}
		temp = temp.letters[r-'a']
	}
	return true
}
