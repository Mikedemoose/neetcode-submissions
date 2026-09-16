type TaskFreq struct {
	freq int
	latestCycle int
}

func leastInterval(tasks []byte, n int) int {
	currCycle := 0
	remainingTasks := len(tasks)

	// save count of tasks and latestCycleTime
	taskCount := make([]TaskFreq, 26)
	for _, task := range tasks {
		taskCount[task-'A'].freq++
	}

	// helper to get next task
	getTask := func() int {
		currTaskIndex := -1
		maxFreq := 0
		for i, taskFreq := range taskCount {
			if taskFreq.freq > maxFreq {
				if taskFreq.latestCycle == 0 || currCycle-taskFreq.latestCycle > n {
					maxFreq = taskFreq.freq
					currTaskIndex = i
				}
			}
		}
		return currTaskIndex
	}

	for remainingTasks > 0 {
		currCycle++
		if currTask := getTask(); currTask == -1 {
			continue
		} else {
			taskCount[currTask].freq--
			taskCount[currTask].latestCycle = currCycle
			remainingTasks--
		}
	}

	return currCycle

}

