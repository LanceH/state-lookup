package main

import (
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

		for _, v := range ps.Parts {
			fmt.Println("Part Index:", v)
			part := Part{abbr, &state, ps.NumPoints - 2, make([]Tri, 0)}
			state.Parts = append(state.Parts, part)
		}
		*states = append(*states, state)
	}
}

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

type Part struct {
	abbr    string
	State   *State
	NumTris int32
	Tris    []Tri
}

type Point struct {
	x float64
	y float64
}
