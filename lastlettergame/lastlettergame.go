package lastlettergame

func travel(adjList [][]int, currWord int, seen []bool) (int, []int) {
	maxLen := -1
	maxSet := []int{}

	for _, v := range adjList[currWord] {
		if !seen[v] {
			seen[v] = true
			nextLen, nextSet := travel(adjList, v, seen)
			seen[v] = false

			if nextLen > maxLen {
				maxLen = nextLen
				maxSet = append([]int{v}, nextSet...)
			}
		}
	}

	return maxLen + 1, maxSet
}

func findLongestPath(adjList [][]int, n int) []int {
	// init mark table
	seen := make([]bool, n)
	maxLen := 0
	maxSet := []int{}

	for i := 0; i < n; i++ {
		seen[i] = true
		nextLen, nextSet := travel(adjList, i, seen)
		seen[i] = false

		if nextLen > maxLen {
			maxLen = nextLen
			maxSet = append([]int{i}, nextSet...)
		}
	}

	return maxSet
}

func Sequence(words []string) []string {
	// init memory space
	n := len(words)
	adjList := make([][]int, n)

	// generate graph
	for i := 0; i < n; i++ {
		fw := words[i][0]
		lw := words[i][len(words[i])-1]
		for j := 0; j < i; j++ {
			if lw == words[j][0] {
				adjList[i] = append(adjList[i], j)
			}
			if words[j][len(words[j])-1] == fw {
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	// find the longest path
	longest := findLongestPath(adjList, n)

	// generate answer
	res := make([]string, len(longest))
	for i, v := range longest {
		res[i] = words[v]
	}

	return res
}
