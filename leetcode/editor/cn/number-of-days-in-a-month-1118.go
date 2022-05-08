package main

func numberOfDays(year int, month int) int {
	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	isLeapYear := func(year int) bool {
		if year%400 == 0 {
			return true
		} else if year%4 == 0 && year%100 != 0 {
			return true
		}

		return false
	}
	if isLeapYear(year) && month == 2 {
		return monthDays[month-1] + 1
	}
	return monthDays[month-1]

}
