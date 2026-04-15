
func convertTemperature(celsius float64) [] float64 {
	cel := celsius
	Kelvin := cel + 273.15
	Fahrenheit := cel * 1.80 + 32.00
	return []float64 {Kelvin, Fahrenheit}
}

