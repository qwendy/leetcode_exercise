package practice

type DepthFirstSearch struct {
	marked map[int]bool
	edgeTo []int
	graph  [][]int
}

func NewDepthFirstSearch() *DepthFirstSearch {
	return &DepthFirstSearch{
		marked: make(map[int]bool),
	}
}

func (dfs *DepthFirstSearch) MakeGraph(num int, infos [][]int) {
	dfs.graph = make([][]int, num)
	for _, pair := range infos {
		// undirected graph
		dfs.graph[pair[0]] = append(dfs.graph[pair[0]], pair[1])
		dfs.graph[pair[1]] = append(dfs.graph[pair[1]], pair[0])
	}
}

func (dfs *DepthFirstSearch) dfs(v int) {
	dfs.marked[v] = true
	for _, nextVertex := range dfs.graph[v] {
		if !dfs.marked[nextVertex] {
			dfs.edgeTo[nextVertex] = v
			dfs.dfs(nextVertex)
			return
		}
	}
}

func (dfs *DepthFirstSearch) PathTo(s, v int) (path []int) {
	dfs.edgeTo = make([]int, len(dfs.graph))
	dfs.dfs(s)
	if !dfs.marked[v] {
		return path
	}
	for dfs.edgeTo[v] != s {
		path = append(path, dfs.edgeTo[v])
		v = dfs.edgeTo[v]
	}
	path = append(path, s)
	return
}
