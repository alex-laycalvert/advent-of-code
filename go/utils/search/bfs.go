package search

import "github.com/alex-laycalvert/advent-of-code/utils"

type BFS struct{}

func NewBFS() BFS {
	return BFS{}
}

func (bfs BFS) Search(maze Maze) func(yield func(SearchIteration) bool) {
	return func(yield func(SearchIteration) bool) {
		q := utils.NewQueue[Path]()
		q.Enqueue(Path{
			Pt: maze.Start,
		})
		visited := make(map[string]bool)
		visitedPoints := make(Points, 0)
		for !q.IsEmpty() {
			current := q.Dequeue()
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
				q.Enqueue(Path{
					Pt:   neighbor,
					From: &current,
				})
			}
		}
	}
}
