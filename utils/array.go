package utils

import (
	"strings"
)

func JoinAndRemoveDuplicateStr(strSlice []string) string {
	return strings.Join(RemoveDuplicateStr(strSlice), ", ")
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
