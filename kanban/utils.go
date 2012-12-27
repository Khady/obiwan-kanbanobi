package main

import (
	"strconv"
)

func SString_of_SUInt32(s []uint32) []string {
	sstring := make([]string, len(s))
	for index, elem := range s {
		sstring[index] = strconv.FormatUint(uint64(elem), 10)
	}
	return sstring
}

func SUInt32_of_SString(s []string) []uint32 {
	suint := make([]uint32, len(s))
	var res uint64
	for index, elem := range s {
		res, _ = strconv.ParseUint(elem, 10, 0)
		suint[index] = uint32(res)
	}
	return suint
}
