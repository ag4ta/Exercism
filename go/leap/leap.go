// Package leap provides function to check if given year is a leap year
package leap

// IsLeapYear returns boolean value indicating whether given year is a leap year
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	} else {
		return year%4 == 0
	}
}
