package utils

import (
	"strings"
)

//MakeExcited transform a setence to all caps with an exclametion point
func MakeExcited(sentence string) string {
	return strings.ToUpper(sentence) + "!"
}
