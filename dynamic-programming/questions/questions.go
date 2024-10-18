package questions

import (
	"fmt"
	"slices"
)

func MinimumDifficultyOfAJobSchedule() {
	jobDifficulty := []int{7, 1, 7, 1, 7, 1}
	d := 3
	if d > len(jobDifficulty) {
		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaa", -1)
		return
	}
	slices.Sort(jobDifficulty)
	sum := jobDifficulty[len(jobDifficulty)-1]
	d--
	for i := 0; i < d; i++ {
		sum += jobDifficulty[i]
	}
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaa", sum)
}
