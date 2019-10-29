package main

import (
	"fmt"
	"testing"
)

func TestGetOperator(t *testing.T) {
	t.Run("0", testGetOperator(0, "+"))
	t.Run("1", testGetOperator(1, "-"))
	t.Run("2", testGetOperator(2, "*"))
	t.Run("3", testGetOperator(3, "Unknown Operator"))
}

func testGetOperator(number int, expect string) func(t *testing.T) {
	return func(t *testing.T) {
		actual := getOperator(number)
		if actual != expect {
			t.Error(fmt.Sprintf("Expected the operator %d to be %s but instead got %s!", number, expect, actual))
		}
	}
}

func TestConvertNumberToText(t *testing.T) {
	t.Run("1", testConvertNumberToText(1, "one"))
	t.Run("2", testConvertNumberToText(2, "two"))
	t.Run("3", testConvertNumberToText(3, "three"))
	t.Run("4", testConvertNumberToText(4, "four"))
	t.Run("5", testConvertNumberToText(5, "five"))
	t.Run("6", testConvertNumberToText(6, "six"))
	t.Run("7", testConvertNumberToText(7, "seven"))
	t.Run("8", testConvertNumberToText(8, "eight"))
	t.Run("9", testConvertNumberToText(9, "nine"))
	t.Run("10", testConvertNumberToText(10, "Number Out Of Scope"))
}

func testConvertNumberToText(number int, expect string) func(t *testing.T) {
	return func(t *testing.T) {
		actual := convertNumberToText(number)
		if actual != expect {
			t.Error(fmt.Sprintf("Expected the number %d to be %s but instead got %s!", number, expect, actual))
		}
	}
}

func TestGenerateCaptcha(t *testing.T) {
	t.Run("1 1 1 1", testGenerateCaptcha(1, 1, 1, 1, "one - 1"))
	t.Run("1 1 0 1", testGenerateCaptcha(1, 1, 0, 1, "one + 1"))
	t.Run("0 1 1 1", testGenerateCaptcha(0, 1, 1, 1, "1 - one"))
	t.Run("0 1 0 1", testGenerateCaptcha(0, 1, 0, 1, "1 + one"))
	t.Run("1 2 2 1", testGenerateCaptcha(1, 2, 2, 1, "two * 1"))
	t.Run("0 2 2 1", testGenerateCaptcha(0, 2, 2, 1, "2 * one"))
	t.Run("0 2 3 1", testGenerateCaptcha(0, 2, 3, 1, "Unknown Operator"))
	t.Run("0 10 2 1", testGenerateCaptcha(0, 10, 2, 1, "Number Out Of Scope"))
}

func testGenerateCaptcha(format int, lOper int, operator int, rOper int, expect string) func(t *testing.T) {
	return func(t *testing.T) {
		actual := generateCaptcha(format, lOper, operator, rOper)
		if actual != expect {
			t.Error(fmt.Sprintf("Expected the captcha to be %s but instead got %s!", expect, actual))
		}
	}
}

func TestIsNumberOutOfScope(t *testing.T) {
	var expect = true
	actual := isNumberOutOfScope(10)
	if actual != expect {
		t.Error(fmt.Sprintf("Expected to be %t but instead got %t!", expect, actual))
	}
}
