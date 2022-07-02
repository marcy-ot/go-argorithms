package insert

func Sort(n []int) []int {
	for i := 1; i < len(n); i++ {
		for j := i - 1; j > 0; j-- {
			if n[i] < n[j] {
				n[i], n[j] = n[j], n[i]
			} else {
				break
			}
		}
	}
	return n
}
