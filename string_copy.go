package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {

	url, _  := url.Parse("10.234.123")
	fmt.Println(url)

	Metric := "ck.query"
	//Metric := "report.query"
	if strings.Index(Metric, "report.") < 0 && strings.Index(Metric, "statistics.") < 0 && {
		fmt.Println("what?")
		fmt.Println(strings.Index(Metric, "report."))
		fmt.Println(strings.Index(Metric, "statistics."))
	}
	nums := []int {1, 2}
	permute(nums)
}

var result [][]int

func permute(nums []int) [][]int {
	n := len(nums)
	result = [][]int{}
	if n == 0 {
		return result
	}

	path := make([]int, n)
	visited := make([]bool, n)
	backtrack(nums, n, 0, path, visited)
	return result
}

func backtrack(nums []int, n, depth int, path []int, visited []bool) {
	if depth == n {
		temp := make([]int, n)
		copy(temp, path)
		fmt.Println(" path[last]=>", path[len(path) - 1])
		fmt.Println(" temp[last]=>", temp[len(temp) - 1])
		result = append(result, temp)
		return
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		path = append(path, nums[i])
		visited[i] = true
		backtrack(nums, n, depth + 1, path, visited)
		visited[i] = false
		path = path[:len(path) - 1]
	}
}