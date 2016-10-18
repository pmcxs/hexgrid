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
