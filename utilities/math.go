package utilities

func (t Tri) InsideTri(p *Point) bool {
	c := Cross(t.Points[0], p, t.Points[1])
	switch {
	case c > 0.0:
		if (Cross(t.Points[1], p, t.Points[2]) > 0.0) &&
			(Cross(t.Points[2], p, t.Points[0]) > 0.0) {
			return true
		} else {
			return false
		}
	case c < 0.0:
		if (Cross(t.Points[1], p, t.Points[2]) < 0.0) &&
			(Cross(t.Points[2], p, t.Points[0]) < 0.0) {
			return true
		} else {
			return false
		}
	case c == 0:
		return true
	}
	return true
}

func Cross(a, b, c *Point) float64 {
	return (a.x-b.x)*(c.y-b.y) - (a.y-b.y)*(c.x-b.x)
}

type Tri struct {
	abbr   string
	State  *State
	Part   *Part
	MinX   float64
	MinY   float64
	MaxX   float64
	MaxY   float64
	Points [3]*Point
}

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

type Part struct {
	abbr    string
	State   *State
	NumTris int32
	Tris    []*Tri
}

type Point struct {
	x float64
	y float64
}
