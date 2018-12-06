package practice

const (
	visited = 1
	sorted  = 2
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := makeGraph(numCourses, prerequisites)
	// inDegree := calInDegree(graph)
	marked := make(map[int]int)
	for i := 0; i < numCourses; i++ {

		if !dfs(i, marked, graph) {
			return false
		}
	}
	return true
}

func dfs(c int, marked map[int]int, graph [][]int) bool {

	if status, ok := marked[c]; ok {
		if status == sorted {
			return true
		}
		if status == visited {
			return false
		}
	}
	marked[c] = visited
	for _, v := range graph[c] {
		if !dfs(v, marked, graph) {
			return false
		}
	}
	marked[c] = sorted
	return true
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
