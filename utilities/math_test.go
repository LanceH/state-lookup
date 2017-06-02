package utilities

import (
	"fmt"
	"testing"
)

var tri Tri = Tri{}
var inside Point = Point{1.0, 1.0}
var outside Point = Point{5.0, 5.0}

func TestMain(m *testing.M) {
	fmt.Println("Test Main")
	tri.Points[0] = &Point{0.0, 0.0}
	tri.Points[1] = &Point{3.0, 0.0}
	tri.Points[2] = &Point{0.0, 4.0}
	m.Run()
}

func TestInside(t *testing.T) {
	v := tri.InsideTri(&inside)
	if v != true {
		t.Error("Expected true, got ", v)
	}
}

func TestOuside(t *testing.T) {
	v := tri.InsideTri(&outside)
	if v != false {
		t.Error("Expected false, got ", v)
	}
}

func TestCross(t *testing.T) {
	v := Cross(tri.Points[1], tri.Points[0], tri.Points[2])
	// vectors should be (3,0) and (0,4)
	if v != 12 {
		t.Error("Expected 12, got ", v)
	}
	// vectors should be (0,4) and (3,0)
	v = Cross(tri.Points[2], tri.Points[0], tri.Points[1])
	if v != -12 {
		t.Error("Expected -12, got ", v)
	}
}

func TestTri(t *testing.T) {
	v := true
	if false {
		t.Error("Expected true, got ", v)
	}
}

func BenchmarkInside(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tri.InsideTri(&inside)
	}
}

func BenchmarkOutside(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tri.InsideTri(&outside)
	}
}
