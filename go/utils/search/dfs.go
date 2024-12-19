package search

import "github.com/alex-laycalvert/advent-of-code/utils"

type DFS struct{}

func NewDFS() DFS {
	return DFS{}
}

func (dfs DFS) Search(maze Maze) func(yield func(SearchIteration) bool) {
	return func(yield func(SearchIteration) bool) {
		s := utils.NewStack[Path]()
		s.Push(Path{
			Pt: maze.Start,
		})
		visited := make(map[string]bool)
		visitedPoints := make(Points, 0)
		for !s.IsEmpty() {
			current := s.Pop()
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
				s.Push(Path{
					Pt:   neighbor,
					From: &current,
				})
			}
		}
	}
}
