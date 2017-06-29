package geom

import "fmt"

//LogPlot s
func (p Point) LogPlot() {
	fmt.Printf("points(%f,%f)\n", p.X, p.Y)
}

//LogPlot s
func (t Tri) LogPlot() {
	fmt.Printf("lines(c(%f,%f,%f,%f), c(%f,%f,%f,%f))\n", t.Points[0].X, t.Points[1].X, t.Points[2].X, t.Points[0].X, t.Points[0].Y, t.Points[1].Y, t.Points[2].Y, t.Points[0].Y)
}

//LogPlot s
func (r Ring) LogPlot() {
	x := ""
	y := ""
	for i := 0; i < r.Len(); i++ {
		r.Ring = r.Next()
		x = fmt.Sprintf("%s,%f", x, r.Value.(Point).X)
		y = fmt.Sprintf("%s,%f", y, r.Value.(Point).Y)
	}
	fmt.Printf("lines(c(%s),c(%s))\n", x, y)
}
