package hexgrid

import (
	"testing"
	"fmt"
)


func TestHexAdd(t *testing.T) {

	var testCases = []struct {
		hexA     hex
		hexB     hex
		expected hex
	}{
		{NewHex(1, -3), NewHex(3, -7), NewHex(4, -10)},
	}

	for _,tt := range testCases {

		actual := HexAdd(tt.hexA, tt.hexB)

		if(actual != tt.expected) {
			t.Error("Expected:",tt.expected,"got:", actual)
		}
	}
}

func TestHexSubtract(t *testing.T) {

	var testCases = []struct {
		hexA     hex
		hexB     hex
		expected hex
	}{
		{NewHex(1, -3), NewHex(3, -7), NewHex(-2, 4)},
	}


	for _,tt := range testCases {

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


// Tests that the neighbors of a certain hexagon are properly computed for all directions
func TestHexNeighbor(t *testing.T) {

	var testCases = []struct {
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

	for _,tt := range testCases {

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

func TestHexDistance(t *testing.T) {

	var testCases = []struct {
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

	for _,tt := range testCases {

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

func BenchmarkHexDistance(b *testing.B) {

	var testCases = []struct {
		destination hex
	} {
		{ NewHex(0,0)},
		{ NewHex(100,100)},
		{ NewHex(10000,10000)},
	}

	for _,bm := range testCases {

		origin := NewHex(0,0)

		b.Run(fmt.Sprint(origin,":",bm.destination), func(b *testing.B) {
			for i:=0; i < b.N; i++ {
				HexDistance(origin, bm.destination)
			}
		})


	}


}

