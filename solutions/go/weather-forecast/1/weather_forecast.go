// Package weather tools to show current weather condition.
// shows weather condition.
package weather

// CurrentCondition holds current condition.
var CurrentCondition string
// CurrentLocation holds current location.
var CurrentLocation string

// Forecast outputs current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
