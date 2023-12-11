package main

import (
	"fmt"
	"slices"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/fatih/color"
)

type Direction int
type Directions [2]Direction

const (
	DirectionUnspecified Direction = iota
	DirectionNorth
	DirectionEast
	DirectionSouth
	DirectionWest
)

var (
	DirectionsSouthEast  = Directions{DirectionSouth, DirectionEast}
	DirectionsSouthWest  = Directions{DirectionSouth, DirectionWest}
	DirectionsNorthEast  = Directions{DirectionNorth, DirectionEast}
	DirectionsNorthWest  = Directions{DirectionNorth, DirectionWest}
	DirectionsNorthSouth = Directions{DirectionNorth, DirectionSouth}
	DirectionsWestEast   = Directions{DirectionWest, DirectionEast}
)

var DirectionsMap = map[string]Directions{
	"F": DirectionsSouthEast,
	"7": DirectionsSouthWest,
	"L": DirectionsNorthEast,
	"J": DirectionsNorthWest,
	"|": DirectionsNorthSouth,
	"-": DirectionsWestEast,
}

type TileCoords struct {
	X, Y int
}

type Tile struct {
	TileCoords
	Directions Directions
	IsStart    bool
}

type Row []*Tile
type Map []Row

func ParseMap(lines []string) Map {
	result := make(Map, len(lines))
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		result[y] = make(Row, len(line))
		for x := 0; x < len(line); x++ {
			v := string(line[x])
			if v == "." {
				continue
			}
			isStart := v == "S"
			d := DirectionsMap[v]
			result[y][x] = &Tile{
				TileCoords: TileCoords{X: x, Y: y},
				Directions: d,
				IsStart:    isStart,
			}
		}
	}
	return result
}

func DebugPrintMap(m Map, loop []*Tile, insideTiles []TileCoords) {
	colLoop := color.New(color.FgCyan)
	colInside := color.New(color.BgHiRed, color.FgWhite)
	print := func(str string, isInLoop, isInside bool) {
		if isInside {
			colInside.Print(str)
		} else if isInLoop {
			colLoop.Print(str)
		} else {
			fmt.Print(str)
		}
	}
	mapPretty := map[Directions]string{
		DirectionsNorthEast:  "┗",
		DirectionsNorthWest:  "┛",
		DirectionsNorthSouth: "┃",
		DirectionsSouthEast:  "┏",
		DirectionsSouthWest:  "┓",
		DirectionsWestEast:   "━",
	}
	for y := 0; y < len(m); y++ {
		row := m[y]
		for x := 0; x < len(row); x++ {
			tile := row[x]
			if tile == nil {
				print("·", false, slices.Contains(insideTiles, TileCoords{x, y}))
				continue
			}
			isInLoop := slices.Contains(loop, tile)
			isInside := slices.Contains(insideTiles, tile.TileCoords)
			if tile.IsStart {
				print("⊙", isInLoop, isInside)
			} else {
				print(mapPretty[tile.Directions], isInLoop, isInside)
			}
		}
		fmt.Print("\n")
	}
}

func FindStartTile(m Map) *Tile {
	for y := 0; y < len(m); y++ {
		row := m[y]
		for x := 0; x < len(row); x++ {
			if row[x] != nil && row[x].IsStart {
				return row[x]
			}
		}
	}
	panic("no start tile")
}

func InvertDirection(d Direction) Direction {
	switch d {
	case DirectionNorth:
		return DirectionSouth
	case DirectionSouth:
		return DirectionNorth
	case DirectionEast:
		return DirectionWest
	case DirectionWest:
		return DirectionEast
	}
	panic("invalid direction")
}

func FindLoop(m Map) []*Tile {
	startTile := FindStartTile(m)
	loop := []*Tile{startTile}
	lastDirection := DirectionUnspecified
	didStart := false
	for {
		currTile := loop[len(loop)-1]
		if currTile.IsStart {
			if didStart {
				// found the loop
				break
			}
			var nextTile *Tile
			if currTile.Y > 0 && m[currTile.Y-1][currTile.X] != nil {
				nextTile = m[currTile.Y-1][currTile.X]
				lastDirection = DirectionNorth
			} else if currTile.Y < len(m)-1 && m[currTile.Y+1][currTile.X] != nil {
				nextTile = m[currTile.Y+1][currTile.X]
				lastDirection = DirectionSouth
			} else if currTile.X > 0 && m[currTile.Y][currTile.X-1] != nil {
				nextTile = m[currTile.Y][currTile.X-1]
				lastDirection = DirectionWest
			} else if currTile.X < len(m[0])-1 && m[currTile.Y][currTile.X+1] != nil {
				nextTile = m[currTile.Y][currTile.X+1]
				lastDirection = DirectionEast
			} else {
				panic("nowhere to go from start")
			}
			loop = append(loop, nextTile)
			didStart = true
			continue
		}

		inputDirection := InvertDirection(lastDirection)
		var nextDirection Direction
		for _, d := range currTile.Directions {
			if d != inputDirection {
				nextDirection = d
				break
			}
		}

		var nextTile *Tile
		switch nextDirection {
		case DirectionNorth:
			nextTile = m[currTile.Y-1][currTile.X]
		case DirectionSouth:
			nextTile = m[currTile.Y+1][currTile.X]
		case DirectionWest:
			nextTile = m[currTile.Y][currTile.X-1]
		case DirectionEast:
			nextTile = m[currTile.Y][currTile.X+1]
		default:
			panic("invalid next direction")
		}
		loop = append(loop, nextTile)
		lastDirection = nextDirection
	}
	return loop
}

type day10 struct{}

func (day10) Part1(lines []string) any {
	m := ParseMap(lines)
	loop := FindLoop(m)
	return len(loop) / 2
}

func (day10) Part2(lines []string) any {
	m := ParseMap(lines)
	loop := FindLoop(m)

	insideTileCoords := []TileCoords{}

	for y := 0; y < len(m); y++ {
		isInside := false
		for x := 0; x < len(m[y]); x++ {
			tile := m[y][x]
			isPartOfLoop := tile != nil && slices.Contains(loop, tile)
			if isPartOfLoop && (tile.IsStart ||
				tile.Directions == DirectionsSouthEast ||
				tile.Directions == DirectionsSouthWest ||
				tile.Directions == DirectionsNorthSouth) {
				isInside = !isInside

			}
			if !isPartOfLoop && isInside {
				insideTileCoords = append(insideTileCoords, TileCoords{x, y})
			}
		}
	}

	DebugPrintMap(m, loop, insideTileCoords)

	return len(insideTileCoords)
}

func main() {
	aoc.Program(&day10{})
}
