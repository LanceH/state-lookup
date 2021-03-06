package geom

import (
	"container/ring"
	"fmt"
)

//InsideTri should work this time
func (t Tri) InsideTri(p Point) bool {
	a := Cross(t.Points[0], t.Points[1], p)
	b := Cross(t.Points[1], t.Points[2], p)
	c := Cross(t.Points[2], t.Points[0], p)
	if a > 0 && b > 0 && c > 0 {
		return true
	}
	return false
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

//Straight returns true if the polygon is straight at that point in the Ring
func (r *Ring) Straight() bool {
	if Cross(r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)) == 0 {
		return true
	}
	return false
}

//MakeRing does stuff
func (p *Part) MakeRing() {
	fmt.Println("how many points: ", len(p.Points))
	for _, v := range p.Points {
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
	fmt.Println("Ring: ", r.Len())
	var t Tri
	for r.Len() > 3 {
		if r.Convex() {
			if r.checkEar() {
				t = Tri{Abbr: p.Abbr, State: p.State, Part: p, Points: [3]Point{r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)}}
				t.SetBounds()
				p.Tris = append(p.Tris, t)
				r.Ring = r.Prev()
				r.Unlink(1)
				//fmt.Println("Length Remaining: ", r.Len())
			}
		}
		r.Ring = r.Next()
	}
	t = Tri{Abbr: p.Abbr, State: p.State, Part: p, Points: [3]Point{r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)}}
	t.SetBounds()
	p.Tris = append(p.Tris, t)
}

//SetBounds sets bounds
func (t *Tri) SetBounds() {
	t.MinX = t.Points[0].X
	t.MaxX = t.Points[0].X
	t.MinY = t.Points[0].Y
	t.MaxY = t.Points[0].Y

	for i := 1; i < 2; i++ {
		if t.Points[i].X < t.MinX {
			t.MinX = t.Points[i].X
		}
		if t.Points[i].X > t.MinX {
			t.MinX = t.Points[i].X
		}
		if t.Points[i].Y < t.MinY {
			t.MinY = t.Points[i].Y
		}
		if t.Points[i].Y > t.MinY {
			t.MinY = t.Points[i].Y
		}
	}
}

func (r Ring) checkEar() bool {
	t := Tri{Points: [3]Point{r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)}}
	var q Ring
	q.Ring = r.Move(2)
	for q.Ring != r.Prev() {
		if t.InsideTri(q.Ring.Value.(Point)) {
			return false
		}
		q.Ring = q.Next()
	}
	return true
}
