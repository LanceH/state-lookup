package geom

import (
	"container/ring"
	"fmt"
)

//InsideTri returns true if point is inside Tri
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
	return false
}

//Cross returns a cross product of two points
func Cross(a, b, c Point) float64 {
	return (a.X-b.X)*(c.Y-b.Y) - (a.Y-b.Y)*(c.X-b.X)
}

//Convex returns true if the polygon is convex at that point in the Ring
func (r *Ring) Convex() bool {
	//fmt.Println(r.Value.(Point))
	//fmt.Println("Cross: ", Cross(r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)))
	if Cross(r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)) > 0 {
		return true
	}
	return false
}

//MakeRing does stuff
func (p *Part) MakeRing() {
	fmt.Println("how many points: ", len(p.Points))
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
	fmt.Println("Ring: ", r.Len())
	var t Tri
	for r.Len() > 3 {
		// if r.Len() == 28 {
		// 	xmin := p.Tris[0].Points[0].X
		// 	xmax := p.Tris[0].Points[0].X
		// 	ymin := p.Tris[0].Points[0].Y
		// 	ymax := p.Tris[0].Points[0].Y
		// 	fmt.Println("plot.new()")
		// 	for _, v := range p.Tris {
		// 		if v.Points[0].X < xmin {
		// 			xmin = v.Points[0].X
		// 		}
		// 		if v.Points[0].X > xmax {
		// 			xmax = v.Points[0].X
		// 		}
		// 		if v.Points[0].Y < ymin {
		// 			ymin = v.Points[0].Y
		// 		}
		// 		if v.Points[0].Y > ymax {
		// 			ymax = v.Points[0].Y
		// 		}
		// 		fmt.Printf("x2 <- c(%f,%f,%f,%f)\n", v.Points[0].X, v.Points[1].X, v.Points[2].X, v.Points[0].X)
		// 		fmt.Printf("y2 <- c(%f,%f,%f,%f)\n", v.Points[0].Y, v.Points[1].Y, v.Points[2].Y, v.Points[0].Y)
		// 		fmt.Printf("lines(x2,y2)\n")
		// 	}
		// 	fmt.Printf("plot(1,1,xlim=c(%f,%f), ylim=c(%f,%f))\n", xmin, xmax, ymin, ymax)
		// 	// fmt.Println("Tris: ", r.Len())
		// 	// fmt.Println(r.Len())
		// 	// fmt.Println("x")
		// 	// for i := 0; i < r.Len(); i++ {
		// 	// 	fmt.Print(r.Value.(Point).X, ",")
		// 	// 	r = Ring{r.Next()}
		// 	// }
		// 	// fmt.Println("\n\ny")
		// 	// for i := 0; i < r.Len(); i++ {
		// 	// 	fmt.Print(r.Value.(Point).Y, ",")
		// 	// 	r = Ring{r.Next()}
		// 	// }
		// 	// fmt.Println("")
		// 	//os.Exit(0)
		// }
		//fmt.Println(r.Len(), r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point), r.Convex(), Cross(r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)))
		//fmt.Println(r.Len())
		if r.Convex() {
			if r.checkEar() {
				t = Tri{Abbr: p.Abbr, State: p.State, Part: p, Points: [3]Point{r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)}}
				t.SetBounds()
				p.Tris = append(p.Tris, t)
				//fmt.Println("\nTesting  ", r.Value.(Point))
				r.Ring = r.Prev()
				r.Unlink(1)
				//fmt.Println("Removing ", a.Value.(Point))
				r.Ring = r.Next()
				fmt.Println("Length Remaining: ", r.Len())
			}
		}
		r.Ring = r.Next()
	}
	t = Tri{Abbr: p.Abbr, State: p.State, Part: p, Points: [3]Point{r.Prev().Value.(Point), r.Value.(Point), r.Next().Value.(Point)}}
	t.SetBounds()
	p.Tris = append(p.Tris, t)
}

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
	for i := 0; i < q.Len()-3; i++ {
		if t.InsideTri(q.Ring.Value.(Point)) {
			fmt.Println("point inside tri")
			return false
		}
		q.Ring = q.Next()
	}
	return true
}
