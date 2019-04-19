package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Joy(choco []int, day int) int {
	n := len(choco)
	if n == 1 {
		return day * choco[0]
	}
	left := day*choco[0] + Joy(choco[1:], day+1)
	right := day*choco[n-1] + Joy(choco[:n-1], day+1)
	return max(left, right)
}

type result struct {
	joy      int
	pickLeft bool
}

func JoyMemo(choco []int, i, j int, memo [][]result) int {
	if res := memo[i][j].joy; res != 0 {
		return res
	}

	if j-i == 1 {
		return len(choco) * choco[i]
	}

	day := 1 + len(choco) - (j - i)
	left := day*choco[i] + JoyMemo(choco, i+1, j, memo)
	right := day*choco[j-1] + JoyMemo(choco, i, j-1, memo)
	res := max(left, right)

	memo[i][j].joy = res
	memo[i][j].pickLeft = left > right

	return res
}

func main() {
	c := []int{3, 2, 4, 1, 5, 7}
	// e := []int{2, 3, 5, 1, 4}

	res := Joy(c, 1)
	fmt.Println(res) // 90

}
