package hexgrid

import (
	"fmt"
	"testing"
)

//              50  50  50
//            |---|---|---|
//            |   |   *********
//            |   |  * U:1,V:-1*
//            |   | *           *
//            |    * (150,-86.6) *
//        *********       +       *********
//       * U:0,V:0 *             * U:2,V:-1*
//      *           *           *           *
//     *    (0,0)    *         *   (300,0)   *
//    *       +       *********       +       *
//     *             * U:1,V:0 *             *
//      *           *           *           *
//       *         *  (150,86.6) *         *
//        *********       +       *********
//                 *             *
//                  *           *
//                   *         *
//                    *********
//  Width: 	200 	2 * size
//  Height:	86.6	sqrt(3)/2 * width


var hexToPixelValues = []struct {
	hexA     hex
	expected string
}{
	{NewHex(1, 0), "150.0;86.6"},
	{NewHex(2, -1), "300.0;0.0"},
}



func TestHexToPixel(t *testing.T) {

	origin := point {0,0}

	size := point {100, 100}

	for _,tt := range hexToPixelValues {

		pixel := HexToPixel(layout{size:size, origin:origin,orientation:orientationFlat},tt.hexA)

		actual := fmt.Sprintf("%.1f;%.1f", pixel.x, pixel.y )

		if(actual != tt.expected) {
			t.Error("Expected:",tt.expected,"got:", actual)
		}
	}


}


