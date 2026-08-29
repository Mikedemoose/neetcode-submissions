type WordDictionary struct {
    letters map[byte] *WordDictionary
	isEnd bool
}

func Constructor() WordDictionary {
    return WordDictionary {
		letters: make(map[byte]*WordDictionary),
	}
}

func (this *WordDictionary) AddWord(word string)  {
    currPtr := this

	for i := range len(word) {
		l := word[i]
		if _, ok := currPtr.letters[l]; !ok {
			newItem := Constructor()
			currPtr.letters[l] = &newItem
		}
		currPtr = currPtr.letters[l]
	}

	currPtr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    pathList := make([]*WordDictionary, 0)
	pathList = append(pathList, this)

	for i := range len(word) {
		if len(pathList) == 0 {
			return false
		}
		l := word[i]

		newPathList := make([]*WordDictionary, 0)
		if l == '.' {
			for _, path := range pathList {
				for _, v := range path.letters {
					newPathList = append(newPathList, v)
				}
			}
		} else {
			for _, path := range pathList {
				if _, ok := path.letters[l]; ok {
					newPathList = append(newPathList, path.letters[l])
				}
			}
		}

		pathList = newPathList
	}

	for _, path := range pathList {
		if path.isEnd {
			return true
		}
	}

	return false
}
