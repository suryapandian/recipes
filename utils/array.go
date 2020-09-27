package utils

func RemoveDuplicates(arr []int) []int {
	m := make(map[int]bool)
	var res []int
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = true
		}
	}

	return res
}

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range RemoveDuplicates(a) {
		m[item] = true
	}

	for _, item := range RemoveDuplicates(b) {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
