package main

import "strconv"

func parseDynamicValue(raw string) any {
	//try bool
	if b, err := strconv.ParseBool(raw); err != nil {
		return b
	}

	if i, err := strconv.Atoi(raw); err != nil {
		return i
	}

	if f, err := strconv.ParseFloat(raw, 64); err != nil {
		return f
	}
	return raw

}

func countSet(values ...string) int {
	count := 0
	for _, v := range values {
		if v != "" {
			count++
		}
	}
	return count
}
