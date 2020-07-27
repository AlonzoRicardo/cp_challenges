package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, t, a, b, c, fpt int
	fmt.Scan(&n, &t, &a, &b, &c, &fpt)
	fmt.Println(n, t, a, b, c, fpt)

	// Array to hold the times required to solve each problem
	times := make([]int, n)
	times[0] = fpt

	// Insert problem times
	for i := 1; i < n; i++ {
		times[i] = (a*times[i-1]+b)%c + 1
	}

	// Sort array ascending order
	sort.Ints(times)

	var problemsSolved int
	var penalty int
	var fromContestStart int
	fmt.Println("times", times)
	for _, time := range times {
		if fromContestStart+time > t {
			break
		}
		// When submiting a solution to the judge, the time in minutes from contest start was added to a penalty counter.
		penalty = (fromContestStart + penalty + time)
		fromContestStart = (fromContestStart + time)
		problemsSolved++
	}

	fmt.Println(problemsSolved, penalty%1000000007)
}
