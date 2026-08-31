type TimeMap struct {
	timeMapDict map[string][][]any
}

func Constructor() TimeMap {
	return TimeMap {
		timeMapDict: make(map[string][][]any),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timeMapDict[key] = append(this.timeMapDict[key], []any{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if val, ok := this.timeMapDict[key]; ok {
		left, right := 0, len(val)-1

		for left < right {
			mid := (left+right)/2 + 1
			midTimestamp := val[mid][0].(int)
			if midTimestamp == timestamp {
				return val[mid][1].(string)
			} else if midTimestamp > timestamp {
				right = mid-1
			} else {
				left = mid
			}
		}
		if val[left][0].(int) <= timestamp {
			return val[left][1].(string)
		}
	}
	return ""
}
