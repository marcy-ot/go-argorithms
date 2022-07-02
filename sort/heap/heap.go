package heap

func Sort(n []int) []int {
	arr := buildMinHeap(n)
	return arr
}

func buildMinHeap(n []int) []int {
	var arr []int
	for i := ((len(n) / 2) - 1); i >= 0; i-- {
		arr = downHeapify(i, len(n), n)
	}
	return arr
}

func downHeapify(current, size int, arr []int) []int {
	if isLeaf(current, size) {
		return nil
	}
	smallest := current
	leftChildIndex := current*2 + 1
	rightChildIndex := current*2 + 2
	if leftChildIndex < size && arr[leftChildIndex] < arr[smallest] {
		smallest = leftChildIndex
	}
	if rightChildIndex < size && arr[rightChildIndex] < arr[smallest] {
		smallest = rightChildIndex
	}

	if current != smallest {
		// arr = swap(current, smallest, arr)
		arr[current], arr[smallest] = arr[smallest], arr[current]
		downHeapify(smallest, size, arr)
	}
	return arr
}

func swap(first, second int, arr []int) []int {
	arr[first], arr[second] = arr[second], arr[first]
	return arr
}

func isLeaf(index, size int) bool {
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}
