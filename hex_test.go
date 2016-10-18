package hexgrid

import (
	"testing"
	"fmt"
)

var hexAddValues = []struct {
	hexA     hex
	hexB     hex
	expected hex
}{
	{NewHex(1, -3), NewHex(3, -7), NewHex(4, -10)},
}

func TestHexAdd(t *testing.T) {

	for _,tt := range hexAddValues {

		actual := HexAdd(tt.hexA, tt.hexB)

		if(actual != tt.expected) {
			t.Error("Expected:",tt.expected,"got:", actual)
		}
	}
}

var hexSubtractValues = []struct {
	hexA     hex
	hexB     hex
	expected hex
}{
	{NewHex(1, -3), NewHex(3, -7), NewHex(-2, 4)},
}

func TestHexSubtract(t *testing.T) {

	for _,tt := range hexSubtractValues {

		actual := HexSubtract(tt.hexA, tt.hexB)

		if(actual != tt.expected) {
			t.Error("Expected:",tt.expected,"got:", actual)
		}
	}
}


//           _ _
//         /     \
//    _ _ /(0,-2) \ _ _
//  /     \       /     \
// /(-1,-1)\ _ _ /(1,-2) \
// \       /     \       /
//  \ _ _ /(0,-1) \ _ _ /
//  /     \       /     \
// /(-1,0) \ _ _ /(1,-1) \
// \       /     \       /
//  \ _ _ / (0,0) \ _ _ /
//        \       /
//         \ _ _ /
var hexNeighborValues = []struct {
	origin    hex
	direction direction
	expected  hex
} {

	{ NewHex(0,-1), directionSE, NewHex(1,-1)},
	{ NewHex(0,-1), directionNE, NewHex(1,-2)},
	{ NewHex(0,-1), directionN,  NewHex(0,-2)},
	{ NewHex(0,-1), directionNW, NewHex(-1,-1)},
	{ NewHex(0,-1), directionSW, NewHex(-1,0)},
	{ NewHex(0,-1), directionS,  NewHex(0,0)},
}

// Tests that the neighbors of a certain hexagon are properly computed for all directions
func TestHexNeighbor(t *testing.T) {

	for _,tt := range hexNeighborValues {

		actual := HexNeighbor(tt.origin, tt.direction)

		if(actual != tt.expected) {
			t.Error("Expected:",tt.expected,"got:", actual)
		}
	}
}


//           _ _
//         /     \
//    _ _ /(0,-2) \ _ _
//  /     \       /     \
// /(-1,-1)\ _ _ /(1,-2) \
// \       /     \       /
//  \ _ _ /(0,-1) \ _ _ /
//  /     \       /     \
// /(-1,0) \ _ _ /(1,-1) \
// \       /     \       /
//  \ _ _ / (0,0) \ _ _ /
//  /     \       /     \
// /(-1,1) \ _ _ / (1,0) \
// \       /     \       /
//  \ _ _ / (0,1) \ _ _ /
//        \       /
//         \ _ _ /
var hexDistanceValues = []struct {
	origin      hex
	destination hex
	expected    int
} {
	{ NewHex(-1,-1), NewHex(1,-1), 2 },
	{ NewHex(-1,-1), NewHex(0,0), 2 },
	{ NewHex(0,-1),  NewHex(0,-2), 1 },
	{ NewHex(-1,-1), NewHex(0,1), 3 },
	{ NewHex(1,0),   NewHex(-1,-1), 3 },
}

func TestHexDistance(t *testing.T) {

	for _,tt := range hexDistanceValues {

		actual := HexDistance(tt.origin,tt.destination)

		if(actual != tt.expected) {
			t.Error("Expected:",tt.expected,"got:", actual)
		}
	}
}


func ExampleHexDistance() {
	source := NewHex(-3,7)
	destination := NewHex(0,0)

	distance := HexDistance(source, destination)
	fmt.Println(distance)
	// Output: 7
}


