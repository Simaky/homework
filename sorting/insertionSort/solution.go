package main

func insertionSort(array []int) ([]int) {

	var newValue int

	for i := 0; i < len(array); i++ {
		newValue = array[i]
		j := i
		for j > 0 && array[j-1] > newValue {
			array[j] = array[j-1]
			j--
		}
		array[j] = newValue

	}

	return array
}
