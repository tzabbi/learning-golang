package main

import "fmt"

func main() {
	twoDArray1 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	twoDArray2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	twoDArray3 := [][]int{{1,2,3,4,5},{6,7,8,9,10},{11,12,13,14,15},{16,17,18,19,20},{21,22,23,24,25}}
	fmt.Println(snailSort(twoDArray1))
	fmt.Println(snailSort(twoDArray2))
	fmt.Println(snailSort(twoDArray3))
}

func snailSort(arr [][]int) []int {
	arrayLength := len(arr)
	snailArray := arr[0]
	//snailArray = snailArray[:len(snailArray)-1]
	for i := 0; i < arrayLength-1; i++ {
		if i%2 == 0 {
			for j := 1 + i/2; j <= arrayLength-1-i/2; j++ {
				snailArray = append(snailArray, arr[j][arrayLength-1-i/2])
			}
			for k := arrayLength - 2 - i; k >= 0; k-- {
				snailArray = append(snailArray, arr[arrayLength-1-i/2][k+i/2])
			}
		} else {
			for j := arrayLength - 1 - ((i+1)/2); j >= ((i+1)/2); j-- {
				snailArray = append(snailArray, arr[j][((i+1)/2)-1])
			}
			for k := ((i + 1) / 2); k <= arrayLength-1-((i+1)/2); k++ {
				snailArray = append(snailArray, arr[((i+1)/2)][k])
			}
		}
	}
	return snailArray
}
