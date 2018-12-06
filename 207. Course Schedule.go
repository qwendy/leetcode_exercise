package practice

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := makeGraph(numCourses, prerequisites)
	inDegree := calInDegree(graph)
	marked := make(map[int]int)
	visited := 1
	sorted := 2
	dfs := func(c int) {
		marked[c] = visited
		for _, v := range graph[c] {

		}
	}

	return false
}

func makeGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for _, pair := range prerequisites {
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}
	return graph
}

func calInDegree(graph [][]int) []int {
	inDegree := make([]int, len(graph))
	for _, infos := range graph {
		for _, v := range infos {
			inDegree[v]++
		}
	}
	return inDegree
}
