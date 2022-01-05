// https://www.hackerrank.com/challenges/day-of-the-programmer/problem

package main

import "fmt"

func dayOfProgrammer(year int32) string {
	// Write your code here
	days_in_months := []int32{31, 28, 31, 30, 31, 30, 31, 31}
	is_leap_year := false
	if year < int32(1919) {
		is_leap_year = year%int32(4) == int32(0)
	} else {
		is_leap_year = year%int32(400) == int32(0) || (year%int32(4) == int32(0) && year%int32(100) != int32(0))
	}

	day_of_the_programmer := int32(256)
	total_of_days := int32(0)

	if is_leap_year {
		days_in_months[1] = 29
	}
	if year == int32(1918) {
		days_in_months[1] = 15
	}

	for _, days_in_month := range days_in_months {
		if total_of_days < day_of_the_programmer {
			total_of_days += days_in_month
		}
	}

	return fmt.Sprintf("%d.09.%d", day_of_the_programmer-total_of_days, year)
}

func main() {
	// 26.09.1918
	fmt.Println(dayOfProgrammer((1918)))
	// 13.09.2017
	fmt.Println(dayOfProgrammer((2017)))
	// 12.09.2016
	fmt.Println(dayOfProgrammer((2016)))
}
