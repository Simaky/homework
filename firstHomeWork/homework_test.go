package firstHomeWork

import "testing"

func TestFindMultipleMembers(t *testing.T){

	var tests = []struct{
		input []int
		expected int
	}{
		{[]int{6,9}, 2},
		{[]int{25,9}, 1},
		{[]int{6,0}, 1},

	}

	for _,test := range tests{
		if result := findMultipleMembers(test.input); result != test.expected{
			t.Errorf("Find mutiple members{%q} = %v", test.expected, result)
		}
	}

}

func TestFindSquareEvenMembers(t *testing.T){

	var tests = []struct{
		input []int
		expected int
	}{
		{[]int{4,16}, 2},
		{[]int{36,64,4}, 3},
		{[]int{100,3}, 1},
		{[]int{1,-2}, 0},

	}

	for _,test := range tests{
		if result := findSquareEvenMembers(test.input); result != test.expected{
			t.Errorf("Find square even members{%q} = %v", test.expected, result)
		}
	}

}