package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseDynamicValue(raw string) any {
	//try bool
	if b, err := strconv.ParseBool(raw); err == nil {
		return b
	}

	if i, err := strconv.Atoi(raw); err == nil {
		return i
	}

	if f, err := strconv.ParseFloat(raw, 64); err == nil {
		return f
	}

	if strings.ContainsAny(raw, ":=") {
		fmt.Printf("⚠️  This looks like Go code, not a value: %q\n", raw)
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
