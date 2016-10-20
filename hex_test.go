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

func TestHexScale(t *testing.T) {

	var testCases = []struct {
		hexA     hex
		factor	 int
		expected hex
	}{
		{NewHex(1, -3), 2, NewHex(2, -6)},
		{NewHex(-2, 3), 2, NewHex(-4, 6)},
	}

	for _,tt := range testCases {

		actual := HexScale(tt.hexA, tt.factor)

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

// DISTANCE TESTS


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



//          _____         _____         _____
//         /     \       /     \       /     \
//   _____/ -2,-2 \_____/  0,-3 \_____/  2,-4 \_____
//  /     \       /     \       /     \       /     \
// / -3,-1 \_____/ -1,-2 \_____/  1,-3 \_____/  3,-4 \
// \       /     \       /     \       /     \       /
//  \_____/ -2,-1 \_____/  0,-2 \_____/  2,-3 \_____/
//  /     \       /     \       /     \       /     \
// / -3,0  \_____/ -1,-1 \_____/  1,-2 \_____/  3,-3 \
// \       /     \       /     \       /     \       /
//  \_____/ -2,0  \_____/  0,-1 \_____/  2,-2 \_____/
//  /     \       /     \       /     \       /     \
// / -3,1  \_____/ -1,0  \_____/  1,-1 \_____/  3,-2 \
// \       /     \       /     \       /     \       /
//  \_____/       \_____/       \_____/       \_____/
func TestHexLineDraw(t *testing.T) {


	var testCases = []struct {
		origin      hex
		destination hex
		expected    string	// the expected path serialized to string
	} {
		{ NewHex(-3,-1), NewHex(3,-3), "[(-3,-1) (-2,-1) (-1,-2) (0,-2) (1,-2) (2,-3) (3,-3)]" },
		{ NewHex(-2, 0), NewHex(2,-2), "[(-2,0) (-1,0) (0,-1) (1,-1) (2,-2)]" },
		{ NewHex(1, -1), NewHex(1,-3), "[(1,-1) (1,-2) (1,-3)]" },
	}


	for _,tt := range testCases {

		actual := fmt.Sprint(HexLineDraw(tt.origin,tt.destination))

		if(actual != tt.expected) {
			t.Error("Expected:",tt.expected,"got:", actual)
		}
	}
}


//                 _____
//                /     \
//          _____/ -1,-2 \_____
//         /     \       /     \
//   _____/ -2,-1 \_____/  0,-2 \_____
//  /     \       /     \       /     \
// / -3,-1 \_____/ -1,-2 \_____/  1,-3 \
// \       /     \       /     \       /
//  \_____/ -2,-2 \_____/  0,-3 \_____/
//  /     \       /     \       /     \
// / -3,-1 \_____/ -1,-2 \_____/  1,-3 \
// \       /     \ CENTR /     \       /
//  \_____/ -2,-1 \_____/  0,-2 \_____/
//  /     \       /     \       /     \
// / -3,0  \_____/ -1,-1 \_____/  1,-2 \
// \       /     \       /     \       /
//  \_____/ -2,0  \_____/  0,-1 \_____/
//        \       /     \       /
//         \_____/ -1,0  \_____/
//               \       /
//                \_____/
func TestHexRange(t *testing.T) {


	var testCases = []struct {
		radius                   int
		expectedNumberOfHexagons int
	} {
		{ 0, 1},
		{ 1, 7},
		{ 2, 19},
	}

	for _,tt := range testCases {

		actual := HexRange(NewHex(1,-2),tt.radius)

		if(len(actual) != tt.expectedNumberOfHexagons) {
			t.Error("Expected:",tt.expectedNumberOfHexagons,"got:", len(actual))
		}
	}
}

//    _ _           _ _           _ _
//  /     \       /     \       /     \
// /  0 0  \ _ _ /  2-1  \ _ _ /  4-2  \ _ _
// \       /     \   X   /     \   X   /     \
//  \ _ _ /  1 0  \ _ _ /  3-1  \ _ _ /  5-2  \
//  /     \       /# # #\   X   /     \   X   /
// /  0 1  \ _ _ /# 2 0 #\ _ _ /  4-1  \ _ _ /
// \       /     \#     #/# # #\   X   /     \
//  \ _ _ /  1 1  \#_#_#/# 3 0 #\ _ _ /  5-1  \
//  /     \  |P|  /     \#  X  #/     \   X   /
// /  0 2  \ _ _ /  2 1  \#_#_#/  4 0  \ _ _ /
// \       /     \       /     \   X   /     \
//  \ _ _ /  1 2  \ _ _ /  3 1  \ _ _ /  5 0  \
//  /     \       /     \       /     \       /
// /  0 3  \ _ _ /  2 2  \ _ _ /  4 1  \ _ _ /
// \       /     \       /     \       /     \
//  \ _ _ /  1 3  \ _ _ /  3 2  \ _ _ /  5 1  \
//        \       /     \       /     \       /
//         \ _ _ /       \ _ _ /       \ _ _ /
//
// The FOV measured from the central hex at 1,1, assuming blocking hexagons at 2,0 and 3,0.
// The hexagons marked with an X are non-visible. The remaining 16 are visible.
func TestHexFieldOfView(t *testing.T) {

	universe := []hex {
		NewHex(0,0),
		NewHex(0,1),
		NewHex(0,2),
		NewHex(0,3),
		NewHex(1,0),
		NewHex(1,1),
		NewHex(1,2),
		NewHex(1,3),
		NewHex(2,-1),
		NewHex(2,0),
		NewHex(2,1),
		NewHex(2,2),
		NewHex(3,-1),
		NewHex(3,0),
		NewHex(3,1),
		NewHex(3,2),
		NewHex(4,-2),
		NewHex(4,-1),
		NewHex(4,0),
		NewHex(4,1),
		NewHex(5,-2),
		NewHex(5,-1),
		NewHex(5,0),
		NewHex(5,1),
	}

	losBlockers := []hex { NewHex(2,0),NewHex(3,0)}

	actual := HexFieldOfView(NewHex(1,1), universe, losBlockers)

	if len(actual) != 16 {
		t.Error("Expected: 16 got:", len(actual))
	}
}

////////////////
// Benchmarks //
////////////////

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

func BenchmarkHexLineDraw(b *testing.B) {

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
				HexLineDraw(origin, bm.destination)
			}
		})
	}
}

func BenchmarkHexRange(b *testing.B) {

	var testCases = []struct {
		radius int
	} {
		{ 0 },
		{ 10 },
		{ 100 },
	}

	for _,bm := range testCases {

		origin := NewHex(0,0)

		b.Run(fmt.Sprint(origin,":",bm.radius), func(b *testing.B) {
			for i:=0; i < b.N; i++ {
				HexRange(NewHex(1,-2),bm.radius)
			}
		})
	}
}

func BenchmarkHexHasLineOfSight(b *testing.B) {

	for i:=0; i < b.N; i++ {
		HexHasLineOfSight(NewHex(1,1), NewHex(4,-1), []hex { NewHex(2,0),NewHex(3,0)})
	}
}