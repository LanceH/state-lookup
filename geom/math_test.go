package geom

import (
	"container/ring"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

var tri = Tri{}
var inside = Point{1.0, 1.0}
var outside = Point{5.0, 5.0}
var points Ring

func TestMain(m *testing.M) {
	fmt.Println("Test Main")
	m.Run()
}

func TestTexas(t *testing.T) {
	v := Cross(Point{-97.214885, 27.89744}, Point{-97.13247, 27.89777}, Point{-97.13014, 27.89752})
	if v <= 0 {
		t.Error("expected > 0, got ", v)
	}
}

func TestConvex(t *testing.T) {
	v := points.Convex()
	points.Ring = points.Next()
	if v != true {
		t.Error("expected true, got ", v)
	}

	v = points.Convex()
	points.Ring = points.Next()
	if v == false {
		t.Error("expected true, got ", v)
	}

	v = points.Convex()
	points.Ring = points.Next()
	if v == false {
		t.Error("expected true, got ", v)
	}

	v = points.Convex()
	points.Ring = points.Next()
	if v == true {
		t.Error("expected false, got ", v)
	}
}

func TestInside(t *testing.T) {
	v := tri.InsideTri(inside)
	if v != true {
		t.Error("Expected true, got ", v)
	}
}

func TestOuside(t *testing.T) {
	v := tri.InsideTri(outside)
	if v != false {
		t.Error("Expected false, got ", v)
	}
}

func TestCross(t *testing.T) {
	v := Cross(tri.Points[0], tri.Points[1], tri.Points[2])
	// vectors should be (3,0) and (0,4)
	if v != 12 {
		t.Error("Expected 12, got ", v)
	}
	// vectors should be (0,4) and (3,0)
	v = Cross(tri.Points[2], tri.Points[1], tri.Points[0])
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
		tri.InsideTri(inside)
	}
}

func BenchmarkOutside(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tri.InsideTri(outside)
	}
}

func init() {
	tri.Points[0] = Point{0.0, 0.0}
	tri.Points[1] = Point{0.0, 4.0}
	tri.Points[2] = Point{3.0, 0.0}

	p0 := Point{0.0, 0.0}
	p1 := Point{1.0, 4.0}
	p2 := Point{2.0, 0.0}
	p3 := Point{1.0, 1.0}

	r := ring.New(1)
	r.Value = p0
	points.Ring = r
	points.Value = p0

	r = ring.New(1)
	r.Value = p1
	points.Link(r)
	points.Ring = points.Next()

	r = ring.New(1)
	r.Value = p2
	points.Link(r)
	points.Ring = points.Next()

	r = ring.New(1)
	r.Value = p3
	points.Link(r)
	points.Ring = points.Move(2)

	// used in lookup_test.go

	b, err := ioutil.ReadFile("../data/states.json")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(b, &states)
	if err != nil {
		log.Panic(err)
	}
}
