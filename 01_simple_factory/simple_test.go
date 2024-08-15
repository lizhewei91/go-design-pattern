package simplefactory

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "hi, Tom" {
		t.Fatalf("Type1 test failed, got: '%v', except: 'hi, Tom'", s)
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "hello, Tom" {
		t.Fatalf("Type2 test failed, got: '%v', except: 'hello, Tom'", s)
	}
}
