package perturbations

import (
	"regexp"
	"strings"
)

func Perturbations(record string, counts []int) []string {
	if !strings.Contains(record, "?") {
		return []string{record}
	}
	out := []string{}

	pattern := `^\.*`
	for idx, v := range counts {
		for i := 0; i < v; i++ {
			pattern += "#"
		}
		if idx < len(counts)-1 {
			pattern += `\.+`
		}
	}
	pattern += `\.*$`
	validRecord := regexp.MustCompile(pattern)
	toReplace := regexp.MustCompile(`\?`)
	matches := []int{}
	for _, match := range toReplace.FindAllStringIndex(record, -1) {
		matches = append(matches, match[0])
	}

	for _, test := range generateSubsetsWithRemainder(matches) {
		x := replace(record, test[0], test[1])
		if validRecord.MatchString(x) {
			out = append(out, x)
		}
	}
	return out
}

func replace(r string, x, y []int) string {
	out := []rune(r)
	for _, i := range x {
		out[i] = '.'
	}
	for _, i := range y {
		out[i] = '#'
	}
	return string(out)
}

func generateSubsetsWithRemainder(nums []int) [][][]int {
	var result [][][]int
	n := len(nums)

	// Total number of subsets: 2^n
	totalSubsets := 1 << n

	for i := 0; i < totalSubsets; i++ {
		var subset []int
		var remainder []int

		// Check each bit of i to determine the elements in the subset and remainder
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				subset = append(subset, nums[j])
			} else {
				remainder = append(remainder, nums[j])
			}
		}

		result = append(result, [][]int{subset, remainder})
	}

	return result
}
