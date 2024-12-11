package ezaoc

// func DFS[T comparable](g *Graph[T], start T) []T {
// 	visited := make(map[T]bool)
// 	var result []T
// 	var dfs func(T)
// 	dfs = func(v T) {
// 		visited[v] = true
// 		result = append(result, v)
// 		for _, u := range g.AdjacencyList[v] {
// 			if !visited[u] {
// 				dfs(u)
// 			}
// 		}
// 	}
// 	dfs(start)
// 	return result
// }
