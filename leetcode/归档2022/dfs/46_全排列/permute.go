package main_old

import "fmt"

func premute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}
	dfs(nums, []int{}, &res, visited)
	return res
}

func dfs(nums []int, path []int, ans *[][]int, visited map[int]bool) {
	if len(path) == len(nums) {
		/*
			为什么加入解集时，要将数组内容拷贝到一个新的数组里，再加入解集？

			因为该 path 变量存的是地址引用，结束当前递归时，
			将它加入 res 后，该算法还要进入别的递归分支继续搜索，
			还要继续将这个 path 传给别的递归调用，它所指向的内存空间还要继续被操作，
			所以 res 中的 path 的内容会被改变，这就不对。 所以要弄一份当前的拷贝，
			放入 res，这样后续对 path 的操作，就不会影响已经放入 res 的内容。
		*/

		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}
	for _, v := range nums {
		if visited[v] {
			continue
		}
		path = append(path, v)
		visited[v] = true
		dfs(nums, path, ans, visited)
		path = path[:len(path)-1]
		visited[v] = false
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(premute(nums))
}
