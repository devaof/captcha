package main

import (
	"fmt"
	"os"
	"strconv"
)

// Operator :
var Operator = map[int]string{
	0: "+",
	1: "-",
	2: "*"}

// NumText :
var NumText = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine"}

func getOperator(number int) string {
	return Operator[number]
}

func generateCaptcha(format int, lOper int, operator int, rOper int) string {
	captcha := map[int]string{
		0: fmt.Sprintf("%d %s %s", lOper, Operator[operator], NumText[rOper]),
		1: fmt.Sprintf("%s %s %d", NumText[lOper], Operator[operator], rOper),
	}
	return captcha[format]
}

func main() {
	args := os.Args

	format, _ := strconv.Atoi(args[1])
	loper, _ := strconv.Atoi(args[2])
	oper, _ := strconv.Atoi(args[3])
	roper, _ := strconv.Atoi(args[4])

	println(generateCaptcha(format, loper, oper, roper))
}
