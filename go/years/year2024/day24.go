package year2024

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 59336987801432
//
// Part 2: Not Implemented
type Day24 struct {
	Input []string
}

func (d Day24) Part1() string {
	_, _, zWires, _ := parseWires(d.Input)
	output := evaluateWireNodes(zWires)
	return strconv.Itoa(output)
}

func (d Day24) Part2() string {
	xWires, yWires, zWires, nodes := parseWires(d.Input)
	keys := make([]string, 0, len(nodes))
	for key := range nodes {
		if key[0] == 'x' || key[0] == 'y' {
			continue
		}
		keys = append(keys, key)
	}

	// triedSwaps := make(map[string][]string)
	swaps := make([]string, 0)

	x := evaluateWireNodes(xWires)
	y := evaluateWireNodes(yWires)
	expected := x + y
	z := evaluateWireNodes(zWires)
	difference := expected ^ z
	differenceStr := strconv.FormatInt(int64(difference), 2)
	numOnes := strings.Count(differenceStr, "1")
	fmt.Println("initial:", differenceStr)

	swapStart := 0
	swapEnd := 1
	for {
		fmt.Println("swapping:", keys[swapStart], keys[swapEnd], "with", swaps)
		swapWireNodes(nodes[keys[swapStart]], nodes[keys[swapEnd]])

		x := evaluateWireNodes(xWires)
		y := evaluateWireNodes(yWires)
		expected := x + y
		actual := evaluateWireNodes(zWires)
		difference := expected ^ actual
		differenceStr := strconv.FormatInt(int64(difference), 2)
		currentNumOnes := strings.Count(differenceStr, "1")
		fmt.Println("  difference:", differenceStr)

		if currentNumOnes < numOnes {
			// add current swap, continue
			numOnes = currentNumOnes
			swaps = append(swaps, keys[swapStart], keys[swapEnd])
		} else {
			swapWireNodes(zWires[swapStart], zWires[swapEnd])
		}

		if numOnes == 0 {
			break
		}

		// find next available swaps
		for {
			if swapStart >= len(keys) {
				break
			}

			swapEnd++
			if swapEnd >= len(keys) {
				swapStart++
				swapEnd = swapStart + 1
			}

			for slices.Contains(swaps, keys[swapStart]) {
				swapStart++
				swapEnd = swapStart + 1
			}

			if !slices.Contains(swaps, keys[swapEnd]) {
				break
			}
		}
	}

	fmt.Println("swaps:", swaps)

	// x := evaluateWireNodes(xWires)
	// fmt.Println("x:", x)
	// y := evaluateWireNodes(yWires)
	// fmt.Println("y:", y)
	// expected := x | y
	// fmt.Println("expected:", expected)
	// actual := evaluateWireNodes(zWires)
	// fmt.Println("actual:", actual)
	// difference := expected ^ actual
	// differenceStr := strconv.FormatInt(int64(difference), 2)
	//
	// fmt.Println(differenceStr)

	return "0"
}

func solveWithSwaps(nodes map[string]*WireNode, zWires []*WireNode, swaps []string, expected int) (correctSwaps []string, found bool) {
	if len(swaps) == 4 {
		swapMultipleWireNodes(nodes, swaps)
		actual := evaluateWireNodes(zWires)
		if actual == expected {
			return swaps, true
		}
		swapMultipleWireNodes(nodes, swaps)
		return nil, false
	}

	for key, value := range nodes {
		if value.hasValue || slices.Contains(swaps, key) {
			continue
		}

		for other, otherValue := range nodes {
			if other == key || otherValue.hasValue || slices.Contains(swaps, other) {
				continue
			}

			if correctSwaps, found := solveWithSwaps(nodes, zWires, append(swaps, key, other), expected); found {
				return correctSwaps, true
			}
		}
	}

	return nil, false
}

type WireNode struct {
	name     string
	hasValue bool
	value    int
	left     *WireNode
	right    *WireNode
	op       rune
}

func (n WireNode) evaluate() int {
	if n.hasValue {
		return n.value
	}

	left := 0
	if n.left != nil {
		left = n.left.evaluate()
	}

	right := 0
	if n.right != nil {
		right = n.right.evaluate()
	}

	switch n.op {
	case '&':
		return left & right
	case '|':
		return left | right
	case '^':
		return left ^ right
	default:
		return 0
	}
}

func (n WireNode) copy() *WireNode {
	copied := WireNode{
		name:     n.name,
		hasValue: n.hasValue,
		value:    n.value,
		op:       n.op,
	}

	if n.left != nil {
		copied.left = n.left.copy()
	}

	if n.right != nil {
		copied.right = n.right.copy()
	}

	return &copied
}

func copyWireNodes(nodes []*WireNode) []*WireNode {
	copied := make([]*WireNode, len(nodes))
	for i, node := range nodes {
		copied[i] = node.copy()
	}
	return copied
}

func swapMultipleWireNodes(nodes map[string]*WireNode, swaps []string) {
	for i := 0; i < len(swaps); i += 2 {
		swapWireNodes(nodes[swaps[i]], nodes[swaps[i+1]])
	}
}

func swapWireNodes(a, b *WireNode) {
	a.left, b.left = b.left, a.left
	a.right, b.right = b.right, a.right
	a.op, b.op = b.op, a.op
}

func evaluateWireNodes(wires []*WireNode) int {
	sort.Slice(wires, func(i, j int) bool {
		a, err := strconv.Atoi(wires[i].name[1:])
		utils.PanicErr(err)
		b, err := strconv.Atoi(wires[j].name[1:])
		utils.PanicErr(err)
		return a < b
	})

	output := make([]string, len(wires))
	for i, node := range wires {
		output[len(wires)-i-1] = strconv.Itoa(node.evaluate())
	}
	outputString := strings.Join(output, "")
	outputNumber, err := strconv.ParseInt(outputString, 2, 64)
	utils.PanicErr(err)
	return int(outputNumber)
}

func parseWires(input []string) (xWires []*WireNode, yWires []*WireNode, zWires []*WireNode, nodes map[string]*WireNode) {
	parsingInitialWires := true
	nodes = make(map[string]*WireNode)
	xWires = make([]*WireNode, 0)
	yWires = make([]*WireNode, 0)
	zWires = make([]*WireNode, 0)

	for _, line := range input {
		if line == "" {
			parsingInitialWires = false
			continue
		}

		if parsingInitialWires {
			wire := line[0:3]
			value, err := strconv.Atoi(line[5:])
			utils.PanicErr(err)
			nodes[wire] = &WireNode{name: wire, hasValue: true, value: value}
			if wire[0] == 'x' {
				xWires = append(xWires, nodes[wire])
			} else {
				yWires = append(yWires, nodes[wire])
			}
			continue
		}

		left := line[0:3]
		out := line[len(line)-3:]
		var right string
		var op rune
		switch line[4] {
		case 'A':
			op = '&'
			right = line[8:11]
			break
		case 'O':
			op = '|'
			right = line[7:10]
			break
		case 'X':
			op = '^'
			right = line[8:11]
			break
		}

		if _, ok := nodes[out]; !ok {
			nodes[out] = &WireNode{name: out}
		}

		if _, ok := nodes[left]; !ok {
			nodes[left] = &WireNode{name: left}
		}

		if _, ok := nodes[right]; !ok {
			nodes[right] = &WireNode{name: right}
		}

		nodes[out].op = op
		nodes[out].left = nodes[left]
		nodes[out].right = nodes[right]

		if out[0] == 'z' {
			zWires = append(zWires, nodes[out])
		}
	}

	return xWires, yWires, zWires, nodes
}
