func findOrder(numCourses int, prerequisites [][]int) []int {
	depMap := make(map[int]map[int]struct{})
	for _, item := range prerequisites {
		if _, ok := depMap[item[0]]; !ok {
			depMap[item[0]] = make(map[int]struct{})
		}
		depMap[item[0]][item[1]] = struct{}{}
	}


	courseOrder := make([]int, 0)
	courseDone := make(map[int]struct{})
	for i := range numCourses {
		if _, ok := courseDone[i]; ok {
			continue
		}
		if _, ok := depMap[i]; !ok {
			courseDone[i] = struct{}{}
			courseOrder = append(courseOrder, i)
		} else {
			courseOrderForI := getCourseOrderDfs(depMap, i, courseDone, make(map[int]struct{}))
			if len(courseOrderForI) == 0 {
				return []int{}
			}
			courseOrder = append(courseOrder, courseOrderForI...)
		}
	}

	return courseOrder
}


func getCourseOrderDfs(depMap map[int]map[int]struct{}, courseId int, courseDone map[int]struct{}, visited map[int]struct{}) []int {
	// get order of courses to be taken to complete course i
	// remove the items which are already marked as done from the final output array
	// mark the course as done in the courseDone array

	deps, _ := depMap[courseId]
	res := make([]int, 0)

	for k, _ := range deps {
		if _, ok := courseDone[k]; ok {
			continue
		}
		if _, ok := visited[k]; ok {
			return []int{}
		}
		visited[k] = struct{}{}
		order := getCourseOrderDfs(depMap, k, courseDone, visited)
		if len(order) == 0 {
			return []int{}
		}
		delete(visited, k)
		res = append(res, order...)
	}

	res = append(res, courseId)
	courseDone[courseId] = struct{}{}
	return res
}
