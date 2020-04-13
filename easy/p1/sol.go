package p1

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		a := target - v
		if j, ok := m[a]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return []int{0, 0}
}
