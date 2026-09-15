type Twitter struct {
    users map[int]*User
	currTimestamp int
	maxTweetSize int
}

type User struct {
	userId int
	tweetIds []Tweet
	following map[int]struct{}
}

type Tweet struct {
	tweetId int
	timestamp int
}


func Constructor() Twitter {
    return Twitter {
		users: make(map[int]*User),
		currTimestamp: 0,
		maxTweetSize: 10,
	}
}

func (this *Twitter) CheckAndCreateUser(userId int) *User {
	if user, ok := this.users[userId]; ok {
		return user
	}
	newUser := &User {
		userId: userId,
		following: make(map[int]struct{}),
	}
	this.users[userId] = newUser
	return newUser
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	user := this.CheckAndCreateUser(userId)
	user.tweetIds = append(user.tweetIds, Tweet{
		tweetId: tweetId,
		timestamp: this.currTimestamp,
	})
	if len(user.tweetIds) > this.maxTweetSize {
		start := len(user.tweetIds)-this.maxTweetSize
		user.tweetIds = user.tweetIds[start:]
	}
	this.currTimestamp++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
	var user *User
	var ok bool
	newsFeed := make([]int, 0)
	if user, ok = this.users[userId]; !ok {
		return newsFeed
	}
    
	newsFeedHeap := MinHeap{maxSize: 10}
	for _, tweetId := range user.tweetIds {
		newsFeedHeap.CheckAndInsert(tweetId)
	}
	for followeeId, _ := range user.following {
		if followee, ok := this.users[followeeId]; ok {
			for _, tweetId := range followee.tweetIds {
				newsFeedHeap.CheckAndInsert(tweetId)
			}
		}
	}

	sort.Slice(newsFeedHeap.heap, func(i, j int)bool {
		return newsFeedHeap.heap[i].timestamp > newsFeedHeap.heap[j].timestamp
	})

	for _, item := range newsFeedHeap.heap {
		newsFeed = append(newsFeed, item.tweetId)
	}

	return newsFeed
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	follower, _ := this.CheckAndCreateUser(followerId), this.CheckAndCreateUser(followeeId)
	if followerId == followeeId {
		return
	}
	follower.following[followeeId] = struct{}{}
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	follower, _ := this.CheckAndCreateUser(followerId), this.CheckAndCreateUser(followeeId)
	if followerId == followeeId {
		return
	}
	delete(follower.following, followeeId)
}




type MinHeap struct {
	heap []Tweet
	maxSize int
}
func getParentIndex(childIndex int) int {
	return (childIndex-1)/2
}

func (m *MinHeap) CheckAndInsert(val Tweet) {
	if len(m.heap) == m.maxSize {
		if val.timestamp < m.heap[0].timestamp {
			return
		}
		m.Pop()
	}
	m.Insert(val)
}

func (m *MinHeap) Insert(val Tweet) {
	m.heap = append(m.heap, val)
	curr := len(m.heap)-1
	par := getParentIndex(curr)

	for par >= 0 {
		if m.heap[curr].timestamp < m.heap[par].timestamp {
			m.heap[curr], m.heap[par] = m.heap[par], m.heap[curr]
			curr = par
			par = getParentIndex(curr)
		} else {
			break
		}
	}
}

func (m *MinHeap) Pop() {
	m.heap[0] = m.heap[len(m.heap)-1]
	m.heap = m.heap[:len(m.heap)-1]

	m.heapify(0)
}

func (m *MinHeap) heapify(index int) {
	for {
		minIndex := index
		left, right := 2*index+1, 2*index+2

		if left < len(m.heap) && m.heap[left].timestamp < m.heap[minIndex].timestamp {
			minIndex = left
		}
		if right < len(m.heap) && m.heap[right].timestamp < m.heap[minIndex].timestamp {
			minIndex = right
		}

		if minIndex == index {
			break
		}

		m.heap[minIndex], m.heap[index] = m.heap[index], m.heap[minIndex]
		index = minIndex
	}
}
