package isbnconv

import (
	"fmt"
	"strconv"
	"strings"
)

var InvalidLengthError = fmt.Errorf("Invalid Length")

func ISBN13to10(isbn string) (string, error) {
	//sanitize
	isbn = strings.Replace(isbn, "-", "", -1)

	if len(isbn) == 10 {
		return isbn, nil
	}

	if len(isbn) != 13 {
		return "", InvalidLengthError
	}

	var sum = 0
	runes := []rune(isbn)[3:12]

	for i, r := range runes {
		item, err := strconv.Atoi(string(r))
		if err != nil {
			return "", err
		}
		sum += item * (10 - i)
	}

	var digit = ""

	var c = 11 - (sum % 11)
	if c == 10 {
		digit = "X"
	} else if c == 11 {
		digit = "0"
	} else {
		digit = strconv.Itoa(c)
	}

	return string(runes) + digit, nil
}
