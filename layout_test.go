package hexgrid

import (
	"testing"
	"fmt"
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
func TestHexToPixel(t *testing.T) {

	hex := NewHex(1,0)

	origin := point {0,0}

	size := point {100, 100}

	pixel := HexToPixel(layout{size:size, origin:origin,orientation:orientationFlat},hex)

	fmt.Println(pixel)

}


