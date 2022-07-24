package main 

import (
	"fmt"
)

func main() {

	//1.定义一个数组
	var hens = [6]float64{3.0,5.0,1.0,3.4,2.0,50.0}
	//遍历数组
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	avgWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Printf("totalWeight=%v avgWeight=%v\n", totalWeight, avgWeight)
}