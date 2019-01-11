package hexgrid

import "math"

type point struct {
	x float64
	y float64
}

type layout struct {
	orientation orientation
	size        point // multiplication factor relative to the canonical hexagon, where the points are on a unit circle
	origin      point // center point for hexagon 0,0
}

type orientation struct {
	f0, f1, f2, f3, b0, b1, b2, b3, startAngle float64
}

var orientationPointy orientation = orientation{math.Sqrt(3.), math.Sqrt(3.) / 2., 0., 3. / 2., math.Sqrt(3.) / 3., -1. / 3., 0., 2. / 3., 0.5}

var orientationFlat orientation = orientation{3. / 2., 0., math.Sqrt(3.) / 2., math.Sqrt(3.), 2. / 3., 0., -1. / 3., math.Sqrt(3.) / 3., 0.}

// HexToPixel returns the center pixel for a given hexagon an a certain layout
func HexToPixel(l layout, h hex) point {

	M := l.orientation
	size := l.size
	origin := l.origin
	x := (M.f0*float64(h.q) + M.f1*float64(h.r)) * size.x
	y := (M.f2*float64(h.q) + M.f3*float64(h.r)) * size.y
	return point{x + origin.x, y + origin.y}
}

// PixelToHex returns the corresponding hexagon axial coordinates for a given pixel on a certain layout
func PixelToHex(l layout, p point) fractionalHex {

	M := l.orientation
	size := l.size
	origin := l.origin

	pt := point{(p.x - origin.x) / size.x, (p.y - origin.y) / size.y}
	q := M.b0*pt.x + M.b1*pt.y
	r := M.b2*pt.x + M.b3*pt.y
	return fractionalHex{q, r, -q - r}
}

func HexCornerOffset(l layout, c int) point {

	M := l.orientation
	size := l.size
	angle := 2. * math.Pi * (M.startAngle - float64(c)) / 6.
	return point{size.x * math.Cos(angle), size.y * math.Sin(angle)}
}

// Gets the corners of the hexagon for the given layout, starting at the E vertex and proceeding in a CCW order
func HexagonCorners(l layout, h hex) []point {

	corners := make([]point, 0)
	center := HexToPixel(l, h)

	for i := 0; i < 6; i++ {
		offset := HexCornerOffset(l, i)
		corners = append(corners, point{center.x + offset.x, center.y + offset.y})
	}
	return corners
}
