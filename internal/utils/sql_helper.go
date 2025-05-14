package utils

import ("strings")

func Placeholders(n int) string {
	if n <= 0 {
		return "" // No placeholders
	}
	return strings.TrimRight(strings.Repeat("?,", n), ",")
}
