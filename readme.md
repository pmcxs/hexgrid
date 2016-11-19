# Hexgrid

This is a GO (Golang) library used to handle regular hexagons.
It's based on the algorithms described at http://www.redblobgames.com/grids/hexagons/implementation.html

## Installation

    go get github.com/pmcxs/hexgrid

## Usage
#### Importing

```go
import "github.com/pmcxs/hexgrid"
```

### Examples

#### Creating hexagons

```go
hexagonA := NewHex(1,2) //at axial coordinates Q=1 R=2
hexagonB := NewHex(2,3) //at axial coordinates Q=2 R=3
```

#### Measuring the distance (in hexagons) between two hexagons

```go
distance := HexDistance(hexagonA, hexagonB)
```

#### Getting the array of hexagons on the path between two hexagons

```go
origin := NewHex(10,20)
destination := NewHex(30,40)
path := HexLineDraw(origin, destination) 
```


#### Creating a layout

```go
origin := point {0,0}     // The coordinate that corresponds to the center of hexagon 0,0
size := point {100, 100}  // The length of an hexagon side => 100
layout: = layout{size, origin, orientationFlat}
```

#### Obtaining the pixel that corresponds to a given hexagon

```go
hex := NewHex(1,0)             
pixel := HexToPixel(layout,hex)  // Pixel that corresponds to the center of hex 1,0 (in the given layout)
```


#### Obtaining the hexagon that contains the given pixel (and rounding it)

```go
point := point {10,20}
hex := PixelToHex(layout, point).Round()
```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request

## History

0.1. First version

## Credits

* Pedro Sousa
* Red Blob Games (http://www.redblobgames.com/grids/hexagons/implementation.html)

## License

MIT