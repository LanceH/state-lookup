package models

import "container/ring"

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
