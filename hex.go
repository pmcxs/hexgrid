package hexgrid

import (
	"math"
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

type corner int

const (
	cornerE = iota
	cornerSE
	cornerSW
	cornerW
	cornerNW
	cornerNE
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


func HexAdd(a,b hex) hex {
	return NewHex(a.q + b.q, a.r + b.r)
}

func HexSubtract(a,b hex) hex {
	return NewHex(a.q - b.q, a.r - b.r)
}

func HexMultiply(a hex,k int) hex {
	return NewHex(a.q * k, a.r * k)
}

// The distance between two hexes is the length of the line between them
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

func HexLineDraw(a,b hex) []hex {

	hexLerp := func(a fractionalHex,b fractionalHex,t float64) fractionalHex {
		return NewFractionalHex(a.q * (1 - t) + b.q * t, a.r * (1 - t) + b.r * t);
	}

	N := HexDistance(a, b)

	a_nudge := NewFractionalHex(float64(a.q) + 0.000001, float64(a.r) + 0.000001);
	b_nudge := NewFractionalHex(float64(b.q) + 0.000001, float64(b.r) + 0.000001);


	results := make([]hex, 0);
	step := 1. / math.Max(float64(N), 1);

	for i := 0; i <= N; i++ {
		results = append(results, hexLerp(a_nudge, b_nudge, step * float64(i)).Round());
	}
	return results;

}
