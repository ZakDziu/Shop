package main

import "bytes"

func concat(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}

func concatOptimized(str []string) string {
	var buffer bytes.Buffer
	for _, v := range str {
		buffer.WriteString(v)

	}
	return buffer.String()
}
