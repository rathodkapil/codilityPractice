package main

import (
	"fmt"
	"strconv"
)

func triangleAstrik(n int) {

	for i := 0; i < n; i++ {

		for j := 0; j < i+1; j++ {
			fmt.Print("*")
			fmt.Print(" ")
		}
		fmt.Println()

	}
}

func arrayReverse(n [5]int) {

	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	fmt.Println(n)
}

func factor(primes []int, nums int) []int {

	var res []int

	for _, prime := range primes {

		for nums%prime == 0 {
			res = append(res, prime)
			nums = nums / prime
		}
	}

	if nums > 1 {
		res = append(res, nums)
	}

	return res
}

// recurssive way takes longer time
func fibonacci(num int) int {

	if num <= 1 {
		return num
	}

	return fibonacci(num-2) + fibonacci(num-1)
}

func fibonacciLoop(n int) int {
	if n <= 1 {
		return n
	}

	//[0, 1], ?, 2, 3, 5 ,8 ,13
	//0  1  2  3  4  5  6  7

	var num int
	nMinus1 := 1
	nMinus2 := 0
	for i := 2; i <= n; i++ {
		num = nMinus2 + nMinus1
		nMinus2 = nMinus1
		nMinus1 = num

	}
	return num

}

func fibnacchiMap(n int, m map[int]int) int {

	if n <= 2 {
		return 1
	}

	if v, ok := m[n]; ok {
		return v
	}

	m[n] = fibnacchiMap(n-2, m) + fibnacchiMap(n-1, m)

	return m[n]

}

func numInList(nums []int, num int) bool {

	for _, v := range nums {
		if num == v {
			return true
		}
	}
	return false
}

func reverseAString(str string) string {

	// var revStr string
	// for i := len(str) - 1; i >= 0; i-- {
	// 	revStr = revStr + string(str[i])
	// }
	// return revStr

	var revStr string
	for i := 0; i < len(str); i++ {
		//cat --- tac
		revStr = string(str[i]) + revStr
	}

	return revStr
}

func findMax(nums []int) {

	for i := 0; i <= len(nums)-1; i++ {

		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}

	}

	fmt.Println(nums[0])

}

// func arrayRotation(items []int, nos int) {

// 	fmt.Println("before rotation", items)
// 	for i := 0; i < 1; i++ {
// 		for j :=i+1 ; j< nos; j++{

// 		}
// 		var tem int
// 		fmt.Println(items[i])
// 		fmt.Println(items[len(items)-1])
// 		tem = items[i]
// 		items[i] = items[len(items)-1]
// 		items[i+1] = tem

// 		fmt.Println("after rotation i-----", i, items)
// 	}
// 	//fmt.Println("after rotation", items)
// }

func gridTraveller(m, n int, map1 map[string]int) int {

	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}

	key := fmt.Sprintf("%d , %d", m, n)
	fmt.Println(key)

	if v, ok := map1[key]; ok {
		return v
	}

	map1[key] = gridTraveller(m-1, n, map1) + gridTraveller(m, n-1, map1)

	return map1[key]

}

func canSum(targetSum int, nums []int) bool {

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, v := range nums {
		temSum := targetSum - v
		if canSum(temSum, nums) == true {
			return true
		}
	}

	return false
}

func parkingCost(E string, L string) int {

	/*
		The entrance fee of the car parking lot is 2;
		The first full or partial hour costs 3;
		Each successive full or partial hour (after the first) costs 4.
	*/
	var cost int

	entryTimeHH, _ := strconv.Atoi(E[:2])
	exitTimeHH, _ := strconv.Atoi(L[:2])
	entryTimeMM, _ := strconv.Atoi(E[3:])
	exitTimeMM, _ := strconv.Atoi(L[3:])
	totalHHTime := exitTimeHH - entryTimeHH
	totalMMTime := exitTimeMM - entryTimeMM
	totalHHTimeInMins := totalHHTime * 60
	totalMM := totalHHTimeInMins + totalMMTime

	fmt.Println("totalMM---", totalMM)

	timeAfterFistHrPartialHr := totalMM - 60
	totalFullHr := timeAfterFistHrPartialHr / 60
	totalPartialMM := timeAfterFistHrPartialHr % 60

	initialCost := 5
	totalFullHrCost := totalFullHr * 4
	totalPartialMMCost := 0
	if totalPartialMM > 0 {
		totalPartialMMCost = 4
	}

	fmt.Println("totalFullHrCost", totalFullHrCost)
	fmt.Println("totalPartialMMCost", totalPartialMMCost)

	cost = initialCost + totalFullHrCost + totalPartialMMCost
	fmt.Println("cost--", cost)
	return cost

}

func main() {

	//triangleAstrik(5)
	//arrayReverse([...]int{10, 20, 30, 40, 50})

	// res := factor([]int{2, 3, 5}, 720)
	// fmt.Println(res)

	//	fmt.Println(fibonacci(14))
	// fmt.Println(fibonacciLoop(20))
	// m := make(map[int]int)
	// fmt.Println(fibnacchiMap(20, m))
	// //fmt.Println(numInList([]int{1, 2, 3, 4, 5, 6, 6, 7, 8}, 3))

	//fmt.Println(reverseAString("kapil"))
	//findMax([]int{20, 45, 50, 15, 8, 60, 21})

	// arrayRotation([]int{20, 45, 50, 15, 8, 60, 21}, 2)
	// // fmt.Println("hello")
	// m1 := make(map[string]int)
	// fmt.Println(gridTraveller(18, 18, m1))

	s := "09:42"
	l := "11:42"
	parkingCost(s, l)

	//fmt.Println(canSum(300, []int{5, 3, 4, 7}))
}
