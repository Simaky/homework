package firstHomeWork

import (
	"math"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var mass []int
	n := 100
	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		mass = append(mass, rand.Intn(n)+1)
	}


	fmt.Println("Numbers are multiples of 3 and not multiples of 5: ", findMultipleMembers(mass))
	fmt.Println("Squares of even numbers: ", findSquareEvenMembers(mass))

}

func findMultipleMembers(mass []int) (multipleNum int) {

	multipleNum = 0

	for i := 0; i < len(mass); i++ {

		if isEven(mass[i], 3) && !isEven(mass[i], 5) {
			multipleNum++
		}

	}
	return multipleNum
}

func findSquareEvenMembers(mass []int) (squareEvenNum int) {

	squareEvenNum = 0

	for i := 0; i < len(mass); i++ {
		if condition(mass[i]) {
			squareEvenNum++
		}
	}
	return squareEvenNum
}

func isEven(number int, multipleNum int) (bool) {

	return number%multipleNum == 0

}

func condition(num int) (bool) {

	sqr := math.Sqrt(float64(num))

	n1 := int(sqr)
	sqr = sqr - float64(n1)

	if sqr != 0 {
		return false
	}
	if sqr == 0 && n1%2 == 0 {
		//fmt.Println(num)
		return true
	}
	return false

}
