package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var sb strings.Builder
	var symbol string
	isNumber := false

	for i, char := range str {
		dig, err := strconv.Atoi(string(char))

		if err != nil { // буква
			symbol = string(char)
			sb.WriteString(string(char))
			isNumber = false
		} else { // число
			if i == 0 || isNumber {
				return str, ErrInvalidString
			}
			isNumber = true
			if dig == 0 {
				// удалить последний добавленный символ, т к его надо повторять ноль раз
				s := sb.String()
				s = s[:len(s)-1]
				sb.Reset()
				sb.WriteString(s)
			} else {
				rep := strings.Repeat(symbol, dig-1)
				sb.WriteString(rep)
			}
		}
	}
	return sb.String(), nil
}
