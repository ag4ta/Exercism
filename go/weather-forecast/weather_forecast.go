// Package weather provides tools to forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents current weather condition as a string.
var CurrentCondition string

// CurrentLocation represents current weather location as a string.
var CurrentLocation string

// Forecast returns a string value describing the current forecast based on given city and condition parameters.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
