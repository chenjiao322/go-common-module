package data_structure

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExpireMap(t *testing.T) {
	e := NewExpireMap()
	if v, ok := e.Get("a"); !(v == nil && ok == false) {
		t.Errorf("error")
	}
	e.Add("a", "b", time.Microsecond)
	if v, ok := e.Get("a"); !(v == "b" && ok == true) {
		t.Errorf("error")
	}
	time.Sleep(time.Microsecond * 2)
	if v, ok := e.Get("a"); !(v == nil && ok == false) {
		t.Errorf("error")
	}
}

func TestExpireMap2(t *testing.T) {
	e := NewExpireMap()

	go func() {
		for i := 0; i < 10; i++ {
			e.Add(strconv.Itoa(i), i, time.Second)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			e.Get(strconv.Itoa(i))
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			e.Remove(strconv.Itoa(i))
		}
	}()
	time.Sleep(time.Millisecond * 999)
	for i := 0; i < 10; i++ {
		fmt.Println(e.Get(strconv.Itoa(i)))
	}
}

func TestExpireMap3(t *testing.T) {
	e := NewExpireMap()
	e = nil
	e.Get("a")
	e.Remove("a")
	e.Add("A", nil, time.Second)
}
