package main

import "fmt"

func miniMaxSum(arr []int32) (int32,int32) {
	var sum int32 
	sum = 0
	min := arr[0]
	max := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	
	for _, v := range arr {
		sum += v
	}
	out_min := sum-max
	out_max := sum-min
    return out_min, out_max
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	min := arr[0]
	max := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum-max, sum-min)

}
