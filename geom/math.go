package geom

import (
	"container/ring"
	"fmt"
)

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
	return (a.X-b.X)*(c.Y-b.Y) - (a.Y-b.Y)*(c.X-b.X)
}

//Convex returns true if the polygon is convex at that point in the Ring
func (r *Ring) Convex() bool {
	if Cross(r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)) > 0 {
		return true
	}
	return false
}

//MakeRing does stuff
func (p *Part) MakeRing() {
	//fmt.Println(p)
	for _, v := range p.Points {
		// fmt.Println(k)
		if p.R.Len() == 0 {
			p.R = Ring{ring.New(1)}
			p.R.Value = v
		} else {
			r := ring.New(1)
			r.Value = v
			p.R.Link(r)
		}
		p.R = Ring{p.R.Next()}
	}
	p.R = Ring{p.R.Next()}
}

// MakeTri is destructive to the ring
func (p *Part) MakeTri() {
	r := p.R
	for r.Len() > 3 {
		if r.Convex() {
			if r.checkEar() {
				p.Tris = append(p.Tris, Tri{Points: [3]Point{r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)}})
				r.Ring = r.Prev()
				r.Unlink(1)
			}
		}
		r.Ring = r.Next()
	}
	p.Tris = append(p.Tris, Tri{Points: [3]Point{r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)}})
}

func (r Ring) checkEar() bool {
	t := Tri{Points: [3]Point{r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)}}
	fmt.Println("Checking: ", t)
	var q Ring
	q.Ring = r.Move(2)
	for i := 0; i < q.Len()-3; i++ {
		if t.InsideTri(q.Ring.Value.(Point)) {
			return false
		}
	}
	return true
}
