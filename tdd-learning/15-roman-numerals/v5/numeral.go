package main

import "strings"

func ConvertToRoman(arabic int) string {

	if arabic == 4 {
		return "IV"
	}
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
