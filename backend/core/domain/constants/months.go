package constants

// Months contains the 12 months in Portuguese abbreviation format
var Months = []string{
	"JAN", "FEV", "MAR", "ABR", "MAI", "JUN",
	"JUL", "AGO", "SET", "OUT", "NOV", "DEZ",
}

// IsValidMonth checks if the provided month string is valid
func IsValidMonth(month string) bool {
	for _, m := range Months {
		if m == month {
			return true
		}
	}
	return false
}

// GetMonthIndex returns the index of the month in the Months array
// Returns -1 if the month is not found
func GetMonthIndex(month string) int {
	for i, m := range Months {
		if m == month {
			return i
		}
	}
	return -1
}
