package main

import (
	"container/ring"
	"fmt"

	shp "github.com/jonas-p/go-shp"
)

var shapeFile = "../../data/cb_2013_us_state_500k.shp"
var states []State
var numStates int

func main() {
	fmt.Println("Starting...")
	fmt.Printf("Opening %s", shapeFile)
	shape, err := shp.Open(shapeFile)
	if err != nil {
		panic(err)
	}
	defer shape.Close()

	numStates = shape.AttributeCount()

	fmt.Printf("Processing %d states\n", numStates)
	build(shape, &states)

	fmt.Println("States:")
	for _, v := range states {
		fmt.Println(v.abbr)
	}
}

func build(shape *shp.Reader, states *[]State) {
	fmt.Println(states)
	for shape.Next() {
		n, p := shape.Shape()
		box := p.BBox()
		abbr := shape.ReadAttribute(n, 4)

		ps := p.(*shp.Polygon)

		state := State{abbr, ps.NumPoints, ps.NumParts, box.MinX, box.MinY, box.MaxX, box.MaxY, make([]Part, 0)}
		if abbr == "WY" {
			var end int32
			for k, start := range ps.Parts {
				if int32(k) < ps.NumParts-2 {
					end = ps.Parts[k+1]
				} else {
					end = ps.NumPoints
				}
				points := makePoints(ps, start, end)
				part := Part{abbr, &state, ps.NumPoints - 2, points, nil, ring.New(0)}
				part.makeRing()
				state.Parts = append(state.Parts, part)
			}
			*states = append(*states, state)
		}
	}
}

func (p *Part) makeRing() {
	fmt.Println(p)
	for _, v := range p.Tris {
		fmt.Println(v)
	}
}

func makePoints(ps *shp.Polygon, start, end int32) (points []Point) {
	fmt.Printf("Index: %d - Num %d\n", start, end)
	for i := start; i < end; i++ {
		points = append(points, Point{ps.Points[i].X, ps.Points[i].Y})
		fmt.Println(ps.Points[i])
	}
	return points
}

// Tri is what parts get broken into
type Tri struct {
	abbr   string
	State  *State
	Part   *Part
	MinX   float64
	MinY   float64
	MaxX   float64
	MaxY   float64
	Points [3]Point
}

// State is the top level object
type State struct {
	abbr      string
	NumPoints int32
	NumParts  int32
	x1        float64
	y1        float64
	x2        float64
	y2        float64
	Parts     []Part
}

// Part is a part of a State
type Part struct {
	abbr    string
	State   *State
	NumTris int32
	Points  []Point
	Tris    []Tri
	R       *ring.Ring
}

// Point is a point
type Point struct {
	x float64
	y float64
}
