package main

import (
	"fmt"
	"testing"
)

func TestGenerateCaptcha(t *testing.T) {
	t.Run("1 1 1 1", testGenerateCaptcha(1, 1, 1, 1, "one - 1"))
	t.Run("1 1 0 1", testGenerateCaptcha(1, 1, 0, 1, "one + 1"))
	t.Run("0 1 1 1", testGenerateCaptcha(0, 1, 1, 1, "1 - one"))
	t.Run("0 1 0 1", testGenerateCaptcha(0, 1, 0, 1, "1 + one"))
	t.Run("1 2 2 1", testGenerateCaptcha(1, 2, 2, 1, "two * 1"))
	t.Run("0 2 2 1", testGenerateCaptcha(0, 2, 2, 1, "2 * one"))
}

func testGenerateCaptcha(format int, lOper int, operator int, rOper int, expect string) func(t *testing.T) {
	return func(t *testing.T) {
		actual := generateCaptcha(format, lOper, operator, rOper)
		if actual != expect {
			t.Error(fmt.Sprintf("Expected the captcha to be %s but instead got %s!", expect, actual))
		}
	}
}
