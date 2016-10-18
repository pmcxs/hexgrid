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

#### Measuring distance between hexagons

```go
distance := HexDistance(hexagonA, hexagonB)
```

#### Obtaining the pixel that corresponds to a given hexagon

```go
hex := NewHex(1,0)        // We'll be calculating the pixel that corresponds to the center of hex 1,0     
origin := point {0,0}     // The coordinate that corresponds to the center of hexagon 0,0
size := point {100, 100}  // The length of an hexagon side => 100
pixel := HexToPixel(layout{size:size, origin:origin,orientation:orientationFlat},hex)
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