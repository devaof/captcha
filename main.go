package main

import (
	"fmt"
	"os"
	"strconv"
)

func getOperator(number int) string {
	switch number {
	case 0:
		return "+"
	case 1:
		return "-"
	case 2:
		return "*"
	default:
		return "Unknown Operator"
	}
}

func convertNumberToText(number int) string {
	switch number {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return "Number Out Of Scope"
	}
}

func isNumberOutOfScope(number int) bool {
	return number < 1 || number > 9
}

func generateCaptcha(format int, lOper int, operator int, rOper int) string {

	o := getOperator(operator)

	if o == "Unknown Operator" {
		return fmt.Sprintf("%s", o)
	}

	if isNumberOutOfScope(lOper) || isNumberOutOfScope(rOper) {
		return "Number Out Of Scope"
	}

	if format == 0 {
		return fmt.Sprintf("%d %s %s", lOper, o, convertNumberToText(rOper))
	}

	return fmt.Sprintf("%s %s %d", convertNumberToText(lOper), o, rOper)
}

func main() {
	args := os.Args

	format, _ := strconv.Atoi(args[1])
	loper, _ := strconv.Atoi(args[2])
	oper, _ := strconv.Atoi(args[3])
	roper, _ := strconv.Atoi(args[4])

	println(generateCaptcha(format, loper, oper, roper))
}
