package year2024

import (
	"slices"
	"strconv"
	"strings"
)

// ANSWERS:
//
// Part 1: 1366
//
// Part 2: bs,cf,cn,gb,gk,jf,mp,qk,qo,st,ti,uc,xw
type Day23 struct {
	Input []string
}

func (d Day23) Part1() string {
	connectedComputers := buildConnectionsMap(d.Input)
	unique := make(map[string]string)
	answer := 0
	for key, value := range connectedComputers {
		if key[0] != 't' {
			continue
		}
		for _, edge := range value {
			for _, otherEdge := range connectedComputers[edge] {
				if otherOtherEdges, ok := connectedComputers[otherEdge]; ok && slices.Contains(otherOtherEdges, key) {
					keys := []string{key, edge, otherEdge}
					slices.Sort(keys)
					key := strings.Join(keys, ",")
					if _, ok := unique[key]; !ok {
						unique[key] = key
						answer++
					}
				}
			}
		}
	}

	return strconv.Itoa(answer)
}

func (d Day23) Part2() string {
	connectedComputers := buildConnectionsMap(d.Input)
	keys := make([]string, len(connectedComputers))
	i := 0
	for k := range connectedComputers {
		keys[i] = k
		i++
	}

	maxClique := bronKerbosch([]string{}, keys, []string{}, connectedComputers)
	slices.Sort(maxClique)
	return strings.Join(maxClique, ",")
}

func bronKerbosch(r []string, p []string, x []string, neighbors map[string][]string) []string {
	if len(p) == 0 && len(x) == 0 {
		return r
	}

	var maxClique []string
	for _, v := range p {
		vSet := []string{v}
		neighborsV := neighbors[v]
		rV := SetUnion(r, vSet)
		pV := SetIntersection(p, neighborsV)
		xV := SetIntersection(x, neighborsV)
		clique := bronKerbosch(rV, pV, xV, neighbors)
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
		p = SetDifference(p, vSet)
		x = SetUnion(x, vSet)
	}

	return maxClique
}

func buildConnectionsMap(input []string) map[string][]string {
	connectedComputers := make(map[string][]string)
	for _, line := range input {
		left := line[0:2]
		right := line[3:]
		if connections, ok := connectedComputers[left]; ok {
			if slices.Contains(connections, right) {
				continue
			}
		}
		connectedComputers[left] = append(connectedComputers[left], right)

		if connections, ok := connectedComputers[right]; ok {
			if slices.Contains(connections, left) {
				continue
			}
		}
		connectedComputers[right] = append(connectedComputers[right], left)
	}
	return connectedComputers
}

func SetUnion(a []string, b []string) []string {
	union := make([]string, 0)
	for _, value := range a {
		if !slices.Contains(union, value) {
			union = append(union, value)
		}
	}
	for _, value := range b {
		if !slices.Contains(union, value) {
			union = append(union, value)
		}
	}
	return union
}

func SetDifference(a []string, b []string) []string {
	difference := make([]string, 0)
	for _, value := range a {
		if !slices.Contains(b, value) {
			difference = append(difference, value)
		}
	}
	return difference
}

func SetIntersection(a []string, b []string) []string {
	intersection := make([]string, 0)
	for _, value := range a {
		if slices.Contains(b, value) {
			intersection = append(intersection, value)
		}
	}
	return intersection
}
