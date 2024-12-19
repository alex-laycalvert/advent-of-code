package search

import "github.com/alex-laycalvert/advent-of-code/utils"

type AStar struct{}

func NewAStar() AStar {
	return AStar{}
}

type AStarPath struct {
	Path

	Weight int
}

func (a AStarPath) Less(other AStarPath) bool {
	return a.Weight < other.Weight
}

func (a AStar) Search(maze Maze) func(yield func(SearchIteration) bool) {
	return func(yield func(SearchIteration) bool) {
		pq := utils.NewPriorityQueue[AStarPath]()
		pq.Push(AStarPath{
			Path: Path{
				Pt: maze.Start,
			},
			Weight: maze.Start.DistanceTo(maze.End),
		})
		visited := make(map[string]bool)
		visitedPoints := make(Points, 0)
		for !pq.IsEmpty() {
			current, _ := pq.Pop()
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
				pq.Push(AStarPath{
					Path: Path{
						Pt:   neighbor,
						From: &current.Path,
					},
					Weight: neighbor.DistanceTo(maze.End),
				})
			}
		}
	}
}
