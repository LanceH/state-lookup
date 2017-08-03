package geom

import (
	"fmt"
	"testing"
)

var states = make(map[string]State)

func TestLookupSingle(t *testing.T) {
	s := Lookup(states, 32.7767, -96.7970)
	if s != "TX" {
		t.Error("Expected TX got", s)
	}
}

func TestLookupMultiple(t *testing.T) {
	// This point should be Texas/Oklahoma -- but in Texas
	fmt.Println("test multiple")
	s := Lookup(states, 33.937976, -95.211259)
	if s != "TX" {
		t.Error("Expected TX got", s)
	}
}

func BenchmarkMultipleLookups(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Lookup(states, 33.937976, -95.211259)
	}
}
