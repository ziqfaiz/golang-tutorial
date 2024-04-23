// https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/
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
