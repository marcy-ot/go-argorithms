package selection

import "fmt"

// Sort execute selection sort
func Sort(n []int) []int {
	t := 0
	for i := 0; i < len(n); i++ {
		m := i
		for j := i + 1; j < len(n); j++ {
			t++
			if n[j] < n[m] {
				m = j
			}
		}
		if i != m {
			n[i], n[m] = n[m], n[i]
		}
	}
	fmt.Println("number of trial of selection sort:", t)
	return n
}
