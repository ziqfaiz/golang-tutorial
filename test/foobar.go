package main

import "strconv"

// If the number is divisible by 3 write Foo, otherwise return the number
func Foobar(input int) string {
	isfoo := (input % 3) == 0

	if isfoo {
		return "Foo"
	}
	return strconv.Itoa(input)
}
