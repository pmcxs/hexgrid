package hexgrid

import (
	"math"
	"fmt"
)

type direction int

const (
	directionSE = iota
	directionNE
	directionN
	directionNW
	directionSW
	directionS
)

var directions = []hex {
	NewHex(1,0),
	NewHex(1,-1),
	NewHex(0,-1),
	NewHex(-1,0),
	NewHex(-1,+1),
	NewHex(0,+1),
}

// hex describes a regular hexagon with Cube Coordinates (although the S coordinate is computed on the constructor)
// It's also easy to reference them as axial (trapezoidal coordinates):
// - R represents the vertical axis
// - Q the diagonal one
// - S can be ignored
// For additional reference on these coordinate systems: http://www.redblobgames.com/grids/hexagons/#coordinates
//           _ _
//         /     \
//    _ _ /(0,-1) \ _ _
//  /     \  -R   /     \
// /(-1,0) \ _ _ /(1,-1) \
// \  -Q   /     \       /
//  \ _ _ / (0,0) \ _ _ /
//  /     \       /     \
// /(-1,1) \ _ _ / (1,0) \
// \       /     \  +Q   /
//  \ _ _ / (0,1) \ _ _ /
//        \  +R   /
//         \ _ _ /
type hex struct {
	q int	// x axis
	r int 	// y axis
	s int	// z axis
}

func NewHex(q,r int) hex {

	h := hex{q:q, r:r, s:-q-r}
	return h

}

func (h hex) String() string {
	return fmt.Sprintf("(%d,%d)",h.q,h.r)
}

// fractionHex provides a more precise representation for hexagons when precision is required.
// It's also represented in Cube Coordinates
type fractionalHex struct {
	q float64
	r float64
	s float64
}

func NewFractionalHex(q,r float64) fractionalHex {

	h := fractionalHex{q:q, r:r, s:-q-r}
	return h

}

// Rounds a FractionalHex to a Regular Hex
func (h fractionalHex) Round() hex {

	roundToInt := func (a float64) int {
		if a < 0 {
			return int(a - 0.5)
		}
		return int(a + 0.5)
	}

	q := roundToInt(h.q)
	r := roundToInt(h.r)
	s := roundToInt(h.s)


	q_diff := math.Abs(float64(q) - h.q)
	r_diff := math.Abs(float64(r) - h.r)
	s_diff := math.Abs(float64(s) - h.s)

	if  q_diff > r_diff && q_diff > s_diff {
		q = -r - s
	} else if (r_diff > s_diff) {
		r = -q - s;
	} else {
		s = -q - r;
	}
	return hex{q, r, s};

}

// Adds two hexagons
func HexAdd(a,b hex) hex {
	return NewHex(a.q + b.q, a.r + b.r)
}

// Subtracts two hexagons
func HexSubtract(a,b hex) hex {
	return NewHex(a.q - b.q, a.r - b.r)
}

// Scales an hexagon by a k factor. If factor k is 1 there's no change
func HexScale(a hex,k int) hex {
	return NewHex(a.q * k, a.r * k)
}

func HexLength(hex hex) int {
	return int((math.Abs(float64(hex.q)) + math.Abs(float64(hex.r)) + math.Abs(float64(hex.s))) / 2.);
}

func HexDistance(a,b hex) int {
	sub := HexSubtract(a, b)
	return HexLength(sub)
}

// Returns the neighbor hexagon at a certain direction
func HexNeighbor(h hex, direction direction) hex {
	directionOffset := directions[direction]
	return NewHex(h.q + directionOffset.q, h.r + directionOffset.r)
}

// Returns the slice of hexagons that exist on a line that goes from hexagon a to hexagon b
func HexLineDraw(a,b hex) []hex {

	hexLerp := func(a fractionalHex,b fractionalHex,t float64) fractionalHex {
		return NewFractionalHex(a.q * (1 - t) + b.q * t, a.r * (1 - t) + b.r * t);
	}

	N := HexDistance(a, b)

	// Sometimes the hexLerp will output a point that’s on an edge.
	// On some systems, the rounding code will push that to one side or the other,
	// somewhat unpredictably and inconsistently.
	// To make it always push these points in the same direction, add an “epsilon” value to a.
	// This will “nudge” things in the same direction when it’s on an edge, and leave other points unaffected.

	a_nudge := NewFractionalHex(float64(a.q) + 0.000001, float64(a.r) + 0.000001);
	b_nudge := NewFractionalHex(float64(b.q) + 0.000001, float64(b.r) + 0.000001);


	results := make([]hex, 0);
	step := 1. / math.Max(float64(N), 1);

	for i := 0; i <= N; i++ {
		results = append(results, hexLerp(a_nudge, b_nudge, step * float64(i)).Round());
	}
	return results;
}

// Returns the set of hexagons around a certain center for a given radius
func HexRange(center hex, radius int) []hex {

	var results = make([]hex,0)

	if radius >= 0 {
		for dx := -radius; dx <= radius; dx++ {

			for dy := math.Max(float64(-radius), float64(-dx - radius)); dy <= math.Min(float64(radius), float64(-dx + radius)); dy++ {
				results = append(results, HexAdd(center, NewHex(int(dx), int(dy))))
			}
		}
	}

	return results

}


// Returns the set of hexagons that form a rectangle with the specified width and height
func HexRectangleGrid(width,height int) []hex {

	results := make([]hex, 0)

	for q:=0; q < width; q++ {
		qOffset := int(math.Floor(float64(q) / 2.))

		for r := -qOffset; r < height - qOffset; r++ {

			results = append(results, NewHex(q,r))
		}
	}

	return results
}


// Determines if a given hexagon is visible from another hexagon, taking into consideration a set of blocking hexagons
func HexHasLineOfSight(center hex, target hex, blocking []hex) bool {

	contains := func(s []hex, e hex) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}
		return false
	}

	for _,h := range HexLineDraw(center, target) {

		if contains(blocking, h) {
			return false
		}
	}

	return true
}

// Returns the list of hexagons that are visible from a given hexagon
func HexFieldOfView(source hex, candidates []hex, blocking []hex) []hex {

	results := make([]hex, 0)

	for _,h := range candidates {

		distance := HexDistance(source, h);

		if len(blocking) == 0 || distance <= 1 || HexHasLineOfSight(source, h, blocking) {
			results = append(results, h)
		}
	}

	return results
}
