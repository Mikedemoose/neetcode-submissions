type TimeMap struct {
	timeMapDict map[string][]TimeValue
}

type TimeValue struct {
	ts int
	val string
}

func Constructor() TimeMap {
	return TimeMap {
		timeMapDict: make(map[string][]TimeValue),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timeMapDict[key] = append(this.timeMapDict[key], TimeValue{
		ts: timestamp,
		val: value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if val, ok := this.timeMapDict[key]; ok {
		left, right := 0, len(val)-1

		for left < right {
			mid := (left+right)/2 + 1
			midTimestamp := val[mid].ts
			if midTimestamp == timestamp {
				return val[mid].val
			} else if midTimestamp > timestamp {
				right = mid-1
			} else {
				left = mid
			}
		}
		if val[left].ts <= timestamp {
			return val[left].val
		}
	}
	return ""
}
