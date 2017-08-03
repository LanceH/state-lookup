package geom

import "testing"

var states = make(map[string]State)

func TestLookupSingle(t *testing.T) {
	s := Lookup(states, 32.7767, -96.7970)
	if s != "TX" {
		t.Error("Expected TX got", s)
	}
}

func TestLookupMultiple(t *testing.T) {
	// This point should be Texas/Oklahoma -- but in Texas
	s := Lookup(states, 33.937976, -95.211259)
	if s != "TX" {
		t.Error("Expected TX got", s)
	}
}
