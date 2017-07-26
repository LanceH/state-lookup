package main

import (
	"container/ring"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gitlab.com/LanceH/state-lookup/geom"

	shp "github.com/jonas-p/go-shp"
)

var shapeFile = "../../data/cb_2013_us_state_500k.shp"
var states = make(map[string]geom.State)
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
	build(shape, states)

	fmt.Println("States:")
	for _, v := range states {
		fmt.Println(v.Abbr)
	}
}

func build(shape *shp.Reader, states map[string]geom.State) {
	fmt.Println(states)
	for shape.Next() {
		n, p := shape.Shape()
		box := p.BBox()
		abbr := shape.ReadAttribute(n, 4)

		ps := p.(*shp.Polygon)

		state := geom.State{Abbr: abbr, NumPoints: ps.NumPoints, NumParts: ps.NumParts, MinX: box.MinX, MinY: box.MinY, MaxX: box.MaxX, MaxY: box.MaxY, Parts: make([]geom.Part, 0)}
		if abbr != "XX" {
			fmt.Println("\n\nState:", abbr)
			var end int32
			for k, v := range ps.Parts {
				fmt.Println(k, v)
			}
			for k, start := range ps.Parts {
				if int32(k) < ps.NumParts-1 {
					end = ps.Parts[k+1] - 1
				} else {
					end = ps.NumPoints - 1
				}
				points := makePoints(ps, start, end)
				part := geom.Part{Abbr: abbr, State: &state, NumTris: ps.NumPoints - 2, Points: points, Tris: nil, R: geom.Ring{Ring: ring.New(0)}}
				fmt.Println("Making Ring...")
				part.MakeRing()
				fmt.Println(part.R.Len())
				fmt.Println("Making Triangles...")
				part.MakeTri()
				fmt.Println("Triangles: ", len(part.Tris))
				state.Parts = append(state.Parts, part)
				fmt.Println("Added Part")
			}
			//*states = append(*states, state)
			states[abbr] = state
		}
	}
	s, _ := json.MarshalIndent(states, "", "  ")
	err := ioutil.WriteFile("../../data/states.json", s, 0644)
	if err != nil {
		log.Panic(err)
	}
}

func makePoints(ps *shp.Polygon, start, end int32) (points []geom.Point) {
	fmt.Printf("Index: %d - Num %d\n", start, end)
	for i := start; i < end; i++ {
		points = append(points, geom.Point{X: ps.Points[i].X, Y: ps.Points[i].Y})
	}
	return points
}
