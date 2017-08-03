package geom

import "fmt"

// Lookup will return which state a point is in, or "XXX" if none found
func Lookup(states map[string]State, lat float64, lng float64) string {
	point := Point{X: lng, Y: lat}
	fmt.Println(point)
	var possible []string
	for s, state := range states {
		if lng < state.MaxX && lng > state.MinX && lat < state.MaxY && lat > state.MinY {
			possible = append(possible, s)
		}
	}
	if len(possible) == 1 {
		return possible[0]
	} else if len(possible) == 0 {
		return "XXX"
	} else {
		for _, s := range possible {
			for _, p := range states[s].Parts {
				for _, t := range p.Tris {
					if t.InsideTri(point) {
						return t.Abbr
					}
				}
			}
		}
	}
	return "XXX"
}
