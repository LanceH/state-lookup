package geom

import "container/ring"

// Ring is reimplemented to add various math functions to it
type Ring struct {
	*ring.Ring
}

// Tri is what parts get broken into
type Tri struct {
	Abbr   string
	State  *State `json:"-"`
	Part   *Part  `json:"-"`
	MinX   float64
	MinY   float64
	MaxX   float64
	MaxY   float64
	Points [3]Point
}

// State is the top level object
type State struct {
	Abbr      string
	NumPoints int32
	NumParts  int32
	MinX      float64
	MinY      float64
	MaxX      float64
	MaxY      float64
	Parts     []Part
}

// Part is a part of a State
type Part struct {
	Abbr    string
	State   *State `json:"-"`
	NumTris int32
	Points  []Point `json:"-"`
	Tris    []Tri
	R       Ring `json:"-"`
}

// Point is a point
type Point struct {
	X float64
	Y float64
}
