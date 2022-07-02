package bubble

import "fmt"

// Sort execute bubble sort
func Sort(n []int) []int {
	t := 0
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			t++
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	fmt.Println("number of trial of bubble sort :", t)
	return n
}
