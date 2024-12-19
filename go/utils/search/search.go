package search

import (
	"errors"
	"fmt"
	"math"
)

type SearchIteration struct {
	CurrentPath   Points
	VisitedPoints Points
}

type Searcher interface {
	Search(maze Maze) func(yield func(SearchIteration) bool)
}

func NewSearcher(searchType string) (Searcher, error) {
	switch searchType {
	case "bfs":
		return NewBFS(), nil
	case "dfs":
		return NewDFS(), nil
	case "dijkstra":
		return NewDijkstra(), nil
	case "astar":
		return NewAStar(), nil
	}
	return nil, errors.New(fmt.Sprintf("Unkown Search Type: %s", searchType))
}

type Maze struct {
	Rows  int
	Cols  int
	Start Point
	End   Point
	Walls Points
}

var directions = Points{
	{Row: 1, Col: 0},
	{Row: 0, Col: 1},
	{Row: -1, Col: 0},
	{Row: 0, Col: -1},
}

func NewMaze(input []string) *Maze {
	walls := make(Points, 0)
	rows := len(input)
	var cols int
	var start, end Point
	for r, line := range input {
		if r == 0 {
			cols = len(line)
		}
		for c, char := range line {
			switch char {
			case '#':
				walls = append(walls, Point{Row: r, Col: c})
				break
			case 'S':
				start = Point{Row: r, Col: c}
				break
			case 'E':
				end = Point{Row: r, Col: c}
				break
			default:
				break
			}
		}
	}

	return &Maze{
		Rows:  rows,
		Cols:  cols,
		Start: start,
		End:   end,
		Walls: walls,
	}
}

func (m Maze) GetNeighbors(pt Point) Points {
	neighbors := make(Points, 0, len(directions))
	for _, dir := range directions {
		neighbor := Point{Row: pt.Row + dir.Row, Col: pt.Col + dir.Col}
		if m.Walls.Contains(neighbor) {
			continue
		}
		if neighbor.Row < 0 || neighbor.Row >= m.Rows {
			continue
		}
		if neighbor.Col < 0 || neighbor.Col >= m.Cols {
			continue
		}
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

type Path struct {
	Pt   Point
	From *Path
}

func (p Path) Reconstruct() Points {
	pts := make(Points, 0)
	for current := &p; current != nil; current = current.From {
		pts = append(pts, current.Pt)
	}
	return pts
}

type Point struct {
	Row int
	Col int
}

func (pt Point) String() string {
	return fmt.Sprintf("%d,%d", pt.Row, pt.Col)
}

func (pt Point) Equals(other Point) bool {
	return pt.Row == other.Row && pt.Col == other.Col
}

func (pt Point) DistanceTo(other Point) int {
	return int(math.Sqrt(math.Pow(float64(pt.Row-other.Row), 2) + math.Pow(float64(pt.Col-other.Col), 2)))
}

type Points []Point

func (pts Points) Contains(pt Point) bool {
	for _, p := range pts {
		if p.Equals(pt) {
			return true
		}
	}
	return false
}
