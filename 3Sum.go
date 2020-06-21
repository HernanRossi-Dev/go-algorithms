package main
import (
	"sort"
	"fmt"
)

func threeSum(nums []int) [][]int {
	length := len(nums)
	solutionSet := make([][]int, 0)
  sort.Ints(nums)
	negMap := make(map[int]int)
	for index, num := range nums {
		negMap[num] = index
	}
	boolCheck := make([]bool, length)
	// a + b + c = 0  or a + b = -c
	for i := 0; i < length; i++ {
		aNum := nums[i]
		if aNum > 0 {
			break
		}
		for j := i + 1; j < length; j++ {
			bNum := nums[j]
			cNum := (aNum + bNum)*(-1)
			if cIndex, isPresent := negMap[cNum]; isPresent {
				if cIndex <=j || cIndex <=i {
					continue
				}
				if !boolCheck[j] && !boolCheck[cIndex] {
					boolCheck[j] = true
					boolCheck[cIndex] = true
					solutionSet = append(solutionSet, []int{aNum, bNum, cNum} )
				}
			}
			for j < (length - 2) && nums[j+1] == bNum {
        j = j + 1
      }
		}
		for index, _ := range boolCheck {
			boolCheck[index] = false
		}
		for i < (length - 2) && nums[i + 1] == aNum {
			i = i + 1
		}
  }
  return solutionSet
}
type test struct {
	input []int
	expected [][3]int
}

func main() {
	tests := []test{}
	testOne := test{input: []int{-1, 0, 1, 2, -1, -4}, expected: [][3]int{{-1, -1, 2}, {-1, 0, 1}}  }
	tests = append(tests, testOne)
	testTwo := test{input: []int{-2, 0, 0, 2, 2}, expected: [][3]int{{-2, 0, 2}} }
	tests = append(tests, testTwo)

	for _, test := range tests {
		result := threeSum(test.input)
		fmt.Print("Test Result: ", result, " Expected: ", test.expected)
	}
}