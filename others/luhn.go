package others

import (
	"errors"
	"strconv"
	"strings"
)
/*
	Given a number determine whether it is valid per the Luhn formula.
	The Luhn algorithm is a simple checksum formula used to validate a variety of
	identification numbers, such as credit card numbers, for example.
*/
func luhn(number string) (bool, error) {
	number = strings.ReplaceAll(number, " ", "")
	var sum int
	for i := 0; i < len(number); i++ {
		newNum, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false, errors.New("all characters must be numbers")
		}
		if i == 0 || i % 2 == 0 {
			double := newNum * 2
			if double > 9 {
				double = double - 9
			}
			newNum = double
		}
		sum += newNum
	}
	if sum % 10 == 0 {
		return true, nil
	}
	return false, nil
}
