package main

// CelsiusToFahrenheit converts celsius to fahrenheit
// T(°F) = T(°C) × 9/5 + 32
func CelsiusToFahrenheit(celsius float32) float32 {
	return celsius*(9/5) + 32
}

// FahrenheitToCelsius converts fahrenheit to celsius
// T(°C) = (T(°F) - 32) × 5/9
func FahrenheitToCelsius(fahrenheit float32) float32 {
	return (fahrenheit - 32) * (5 / 9)
}
