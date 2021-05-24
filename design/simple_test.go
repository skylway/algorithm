package design

import (
	"fmt"
	"testing"
)

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("skylway")
	if s != "Hi, skylway" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("skylway")
	fmt.Println(s)
	if s != "Hello, skylway" {
		t.Fatal("Type2 test fail")
	}
}