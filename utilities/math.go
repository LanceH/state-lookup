package utilities

import "container/ring"

//InsideTri returns true if point is isn't Tri
func (t Tri) InsideTri(p Point) bool {
	c := Cross(t.Points[0], p, t.Points[1])
	switch {
	case c > 0.0:
		if (Cross(t.Points[1], p, t.Points[2]) > 0.0) &&
			(Cross(t.Points[2], p, t.Points[0]) > 0.0) {
			return true
		}
		return false
	case c < 0.0:
		if (Cross(t.Points[1], p, t.Points[2]) < 0.0) &&
			(Cross(t.Points[2], p, t.Points[0]) < 0.0) {
			return true
		}
		return false
	case c == 0:
		return true
	}
	return true
}

//Cross returns a cross product of two points
func Cross(a, b, c Point) float64 {
	return (a.x-b.x)*(c.y-b.y) - (a.y-b.y)*(c.x-b.x)
}

//Convex returns true if the polygon is convex at that point in the Ring
func (r *Ring) Convex() bool {
	if Cross(r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)) > 0 {
		return true
	}
	return false
}

type Ring struct {
	*ring.Ring
}

//Tri is a triangle
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

//State is the top level map object
type State struct {
	abbr      string
	NumPoints int32
	NumParts  int32
	x1        float64
	y1        float64
	x2        float64
	y2        float64
	Parts     []*Part
}

//Part is a Part
type Part struct {
	abbr    string
	State   *State
	NumTris int32
	Tris    []*Tri
}

//Point is a Point
type Point struct {
	x float64
	y float64
}
