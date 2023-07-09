// Package weather is a package that tells the future.
package weather

// CurrentCondition is the current weather, easy.
var CurrentCondition string

// CurrentLocation is where we are now.
var CurrentLocation string

// Forecast tells the future, its basically magic.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
