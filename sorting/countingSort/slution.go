package main

import "fmt"

func main()  {

	var array = []int{4,8,1,2,99,0,22,4,88,5,1}

	fmt.Println(countingSort(array))

}

func countingSort(arr []int) ([]int) {

	max, min := arr[0], arr[0]

	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	 countArr := make([]int, max - min + 1)

	for _, value := range arr {

		countArr[value-min]++

	}

	var finalArr []int

	for index, value := range countArr {

		if value > 0 {
			for i := 0; i < value-1; i++ {
				finalArr = append(finalArr, index+min)
			}
		}
	}

	return finalArr
}
