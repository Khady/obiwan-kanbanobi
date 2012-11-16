package main

import (
	"strconv"
)

func SString_of_SInt(s []int) []string {
	sstring := make([]string, len(s))
	for index, elem := range s {
		sstring[index] = strconv.Itoa(elem)
	}
	return sstring
}

func SInt_of_SString(s []string) []int {
	sint := make([]int, len(s))
	for index, elem := range s {
		sint[index], _ = strconv.Atoi(elem)
	}
	return sint
}
