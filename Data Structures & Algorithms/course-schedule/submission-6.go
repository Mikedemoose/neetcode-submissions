func canFinish(numCourses int, prerequisites [][]int) bool {
    depMap := make(map[int]map[int]struct{})
	for _, item := range prerequisites {
		if _, ok := depMap[item[0]]; !ok {
			depMap[item[0]] = make(map[int]struct{})
		}
		depMap[item[0]][item[1]] = struct{}{}
	}

	finalCourseMap := make(map[int]bool)

	for i := range numCourses {
		if isOk, ok := finalCourseMap[i]; ok {
			if !isOk {
				return false
			}
			continue
		}

		if !checkCourseFinishableDfs(depMap, make(map[int]struct{}), i, finalCourseMap) {
			return false
		}
	}
	return true
}




func checkCourseFinishableDfs(depMap map[int]map[int]struct{}, visited map[int]struct{}, courseId int, finalCourseMap map[int]bool) bool {
	isFinishable := false
	if depsMap, ok := depMap[courseId]; !ok {
		// fmt.Println("Course", courseId, "has no dependencies")
		isFinishable = true
	} else if _, ok1 := visited[courseId]; ok1 {
		// fmt.Println("Course", courseId, "has already been visited in this cycle. So returning false")
		isFinishable = false
	} else if isOk, ok2 := finalCourseMap[courseId]; ok2 {
		// fmt.Println("Course", courseId, "has already been visited. Value is", isOk)
		isFinishable = isOk
	} else {
		// fmt.Println("Course", courseId, "Needs computation")

		visited[courseId] = struct{}{}
		isFinishable = true

		for k, _ := range depsMap {
			if !checkCourseFinishableDfs(depMap, visited, k, finalCourseMap) {
				isFinishable = false
				break
			}
		}

	}
	if isFinishable {
		delete(visited, courseId)
	}

	finalCourseMap[courseId] = isFinishable
	// fmt.Println("Course", courseId, ":", isFinishable)
	return isFinishable
}
