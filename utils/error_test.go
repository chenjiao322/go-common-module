package utils

import (
	"errors"
	"testing"
)

func TestCrash1(t *testing.T) {
	Crash(nil)
}

func TestCrash2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	Crash(errors.New("123"))
	t.Errorf("no panic")
}
