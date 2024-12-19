package search

import (
	_ "fmt"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

type Dijkstra struct{}

func NewDijkstra() Dijkstra {
	return Dijkstra{}
}

type DijkstraPath struct {
	Path

	Length int
}

func (dp DijkstraPath) Less(other DijkstraPath) bool {
	return dp.Length < other.Length
}

func (d Dijkstra) Search(maze Maze) func(yield func(SearchIteration) bool) {
	return func(yield func(SearchIteration) bool) {
		pq := utils.NewPriorityQueue[DijkstraPath]()
		pq.Push(DijkstraPath{
			Path: Path{
				Pt: maze.Start,
			},
			Length: 0,
		})
		visited := make(map[string]bool)
		visitedPoints := make(Points, 0)
		for !pq.IsEmpty() {
			current, _ := pq.Pop()
			// fmt.Println(len(current.Path.Reconstruct()))
			if visited[current.Pt.String()] {
				continue
			}
			visited[current.Pt.String()] = true
			visitedPoints = append(visitedPoints, current.Pt)

			if !yield(SearchIteration{
				CurrentPath:   current.Reconstruct(),
				VisitedPoints: visitedPoints,
			}) {
				break
			}

			if current.Pt.Equals(maze.End) {
				break
			}
			neighbors := maze.GetNeighbors(current.Pt)
			for _, neighbor := range neighbors {
				pq.Push(DijkstraPath{
					Path: Path{
						Pt:   neighbor,
						From: &current.Path,
					},
					Length: current.Length + 1,
				})
			}
		}
	}
}
