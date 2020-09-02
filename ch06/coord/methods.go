package coord

import "math"

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i, point := range path {
		if i > 0 {
			sum = sum + path[i-1].Distance(point)
		}
	}
	return sum
}
