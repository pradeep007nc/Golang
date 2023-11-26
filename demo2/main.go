package main

import (
	"demo2/leetcode"
	"fmt"
	"reflect"
	"time"
)

// type car struct {
// 	name  string
// 	color string
// 	year  int
// 	model string
// }

// func addReturn(val1 int, val2 int) (int, int) {
// 	return val1 + val2, val1 - val2
// }

// func demo2(val1 int) (int, int) {
// 	return val1, val1 + 10
// }

// func changeVals(c *car) {
// 	c.year = 2020
// 	c.color = "blue"
// 	c.model = "bmw"
// 	c.name = "test"
// }

func main() {
	// fmt.Println("Hello World")

	// val1, val2 := addReturn(10, 20)
	// val3, _ := demo2(1)

	// data := fmt.Sprintf("the val is %d and %d\n", val1, val2)

	// fmt.Println(data, val3)

	// myCar := car{}

	// cars := [3]car{
	// 	{name: "Toyota", color: "Blue", year: 2020, model: "Camry"},
	// 	{name: "Ford", color: "Red", year: 2021, model: "Mustang"},
	// 	{name: "Honda", color: "Green", year: 2019, model: "Accord"},
	// }

	// changeVals(&myCar)

	// for i, c := range cars {
	// 	fmt.Println(i, c)
	// }

	// fmt.Println(myCar.name, myCar.color, myCar.year, myCar.model)

	fmt.Println(twoSum.TwoSum([]int{2, 7, 11, 15}, 9))

	startTime := time.Now()

	twoSum.TwoSumTime(10000, 10000)

	endTime := time.Now()

	total := endTime.Sub(startTime)

	fmt.Println(reflect.TypeOf(total))
	fmt.Println(total)
}
