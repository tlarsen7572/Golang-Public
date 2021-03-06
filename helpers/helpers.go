package helpers

import (
	"strconv"
	"strings"
)

type StringArray []string

func (array StringArray) Contains(element string) bool {
	for _, value := range array {
		if element == value {
			return true
		}
	}
	return false
}

type IntArray []int

func (array IntArray) Contains(element int) bool {
	for _, value := range array {
		if element == value {
			return true
		}
	}
	return false
}

func WinJoin(elements ...string) string {
	return strings.Join(elements, `\`)
}

func DblToStr(value float64, decimals int) string {
	return strconv.FormatFloat(value, 'f', decimals, 64)
}
