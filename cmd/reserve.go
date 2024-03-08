package cmd

import "log"

func Reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
		log.Println(c)
	}
	return result
}
