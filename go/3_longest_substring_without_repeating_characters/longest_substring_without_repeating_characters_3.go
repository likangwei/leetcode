package main
import "fmt"

func lengthOfLongestSubstring(s string) int {
	max_len := 0
	cur_len := 0
	from_idx := 0
	end_idx := 0
	m := make(map[rune]int)
	for i, c := range s{
		end_idx = end_idx + 1
		fidx, found := m[c]
		if found{
			if fidx >= from_idx{
				from_idx = m[c] + 1
			}
		}
		m[c] = i

		cur_len = end_idx - from_idx
		if cur_len > max_len{
			max_len = cur_len
		}

		// fmt.Printf("%s %c %d %d\n", s, c, from_idx, end_idx)
	}
	return max_len
}

func main() {
	strs := []string{"abcabcbb", "bbbbb", "pwwkew", "aab", "dvdf"}
	for _, s := range strs{
		fmt.Println(lengthOfLongestSubstring(s))
	}
}