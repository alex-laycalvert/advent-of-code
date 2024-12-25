package year2024_test

import (
	"testing"

	"github.com/alex-laycalvert/advent-of-code/years/year2024"
)

func TestDay1(t *testing.T) {
	day1 := year2024.Day1{Input: []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}}

	expectedPart1 := "11"
	part1 := day1.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "31"
	part2 := day1.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay2(t *testing.T) {
	day2 := year2024.Day2{Input: []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}}

	expectedPart1 := "2"
	part1 := day2.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "4"
	part2 := day2.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay3(t *testing.T) {
	day3 := year2024.Day3{Input: []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}}

	expectedPart1 := "161"
	part1 := day3.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "48"
	part2 := day3.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay4(t *testing.T) {
	day4 := year2024.Day4{Input: []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}}

	expectedPart1 := "18"
	part1 := day4.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "9"
	part2 := day4.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay5(t *testing.T) {
	day5 := year2024.Day5{Input: []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}}

	expectedPart1 := "143"
	part1 := day5.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "123"
	part2 := day5.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay6(t *testing.T) {
	day6 := year2024.Day6{Input: []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}}

	expectedPart1 := "41"
	part1 := day6.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "6"
	part2 := day6.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay7(t *testing.T) {
	day7 := year2024.Day7{Input: []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}}

	expectedPart1 := "3749"
	part1 := day7.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "11387"
	part2 := day7.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay8(t *testing.T) {
	day8 := year2024.Day8{Input: []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}}

	expectedPart1 := "14"
	part1 := day8.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "34"
	part2 := day8.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay9(t *testing.T) {
	day9 := year2024.Day9{Input: []string{
		"2333133121414131402",
	}}

	expectedPart1 := "1928"
	part1 := day9.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "2858"
	part2 := day9.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay10(t *testing.T) {
	day10 := year2024.Day10{Input: []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}}

	expectedPart1 := "36"
	part1 := day10.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "81"
	part2 := day10.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay11(t *testing.T) {
	day11 := year2024.Day11{Input: []string{
		"125 17",
	}}

	expectedPart1 := "55312"
	part1 := day11.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	// expectedPart2 := ""
	// part2 := day11.Part2()
	// if part2 != expectedPart2 {
	// 	t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	// }
}

func TestDay12(t *testing.T) {
	day12 := year2024.Day12{Input: []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}}

	expectedPart1 := "1930"
	part1 := day12.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "1206"
	part2 := day12.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay13(t *testing.T) {
	day13 := year2024.Day13{Input: []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
	}}

	expectedPart1 := "480"
	part1 := day13.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "875318608908"
	part2 := day13.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay14(t *testing.T) {
	day14 := year2024.Day14{Input: []string{
		"11,7",
		// "p=2,4 v=2,-3", // Test Value
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}}

	expectedPart1 := "12"
	part1 := day14.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	// Part 2 doesn't apply to sample
	// expectedPart2 := "Not Implemented"
	// part2 := day14.Part2()
	// if part2 != expectedPart2 {
	// 	t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	// }
}

func TestDay15(t *testing.T) {
	day15 := year2024.Day15{Input: []string{
		"##########",
		"#..O..O.O#",
		"#......O.#",
		"#.OO..O.O#",
		"#..O@..O.#",
		"#O#..O...#",
		"#O..O..O.#",
		"#.OO.O.OO#",
		"#....O...#",
		"##########",
		"",
		"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
		"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
		"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
		"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
		"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
		"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
		">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
		"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
		"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
		"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	}}

	expectedPart1 := "10092"
	part1 := day15.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "9021"
	part2 := day15.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay16(t *testing.T) {
	day16 := year2024.Day16{Input: []string{
		"#################",
		"#...#...#...#..E#",
		"#.#.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#",
		"#.#.#.#.###.#.#.#",
		"#...#.#.#.....#.#",
		"#.#.#.#.#.#####.#",
		"#.#...#.#.#.....#",
		"#.#.#####.#.###.#",
		"#.#.#.......#...#",
		"#.#.###.#####.###",
		"#.#.#...#.....#.#",
		"#.#.#.#####.###.#",
		"#.#.#.........#.#",
		"#.#.#.#########.#",
		"#S#.............#",
		"#################",
	}}

	expectedPart1 := "11048"
	part1 := day16.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "64"
	part2 := day16.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}

	day16 = year2024.Day16{Input: []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}}

	expectedPart1 = "7036"
	part1 = day16.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 = "45"
	part2 = day16.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}

	day16 = year2024.Day16{Input: []string{
		"###########################",
		"#######################..E#",
		"######################..#.#",
		"#####################..##.#",
		"####################..###.#",
		"###################..##...#",
		"##################..###.###",
		"#################..####...#",
		"################..#######.#",
		"###############..##.......#",
		"##############..###.#######",
		"#############..####.......#",
		"############..###########.#",
		"###########..##...........#",
		"##########..###.###########",
		"#########..####...........#",
		"########..###############.#",
		"#######..##...............#",
		"######..###.###############",
		"#####..####...............#",
		"####..###################.#",
		"###..##...................#",
		"##..###.###################",
		"#..####...................#",
		"#.#######################.#",
		"#S........................#",
		"###########################",
	}}

	expectedPart1 = "21148"
	part1 = day16.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}
}

func TestDay17(t *testing.T) {
	day17 := year2024.Day17{Input: []string{
		"Register A: 729",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,1,5,4,3,0",
	}}

	expectedPart1 := "4,6,3,5,6,3,5,2,1,0"
	part1 := day17.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	day17 = year2024.Day17{Input: []string{
		"Register A: 2024",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,3,5,4,3,0",
	}}

	expectedPart2 := "117440"
	part2 := day17.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

// func TestDay18(t *testing.T) {
// 	day18 := year2024.Day18{Input: []string{
// 		"5,4",
// 		"4,2",
// 		"4,5",
// 		"3,0",
// 		"2,1",
// 		"6,3",
// 		"2,4",
// 		"1,5",
// 		"0,6",
// 		"3,3",
// 		"2,6",
// 		"5,1",
// 		"1,2",
// 		"5,5",
// 		"2,5",
// 		"6,5",
// 		"1,4",
// 		"0,4",
// 		"6,4",
// 		"1,1",
// 		"6,1",
// 		"1,0",
// 		"0,5",
// 		"1,6",
// 		"2,0",
// 	}}
//
// 	expectedPart1 := "22"
// 	part1 := day18.Part1()
// 	if part1 != expectedPart1 {
// 		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
// 	}
//
// 	expectedPart2 := "Not Implemented"
// 	part2 := day18.Part2()
// 	if part2 != expectedPart2 {
// 		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
// 	}
// }

func TestDay19(t *testing.T) {
	day19 := year2024.Day19{Input: []string{
		"r, wr, b, g, bwu, rb, gb, br",
		"",
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}}

	expectedPart1 := "6"
	part1 := day19.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "16"
	part2 := day19.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay20(t *testing.T) {
	day20 := year2024.Day20{Input: []string{
		"###############",
		"#...#...#.....#",
		"#.#.#.#.#.###.#",
		"#S#...#.#.#...#",
		"#######.#.#.###",
		"#######.#.#...#",
		"#######.#.###.#",
		"###..E#...#...#",
		"###.#######.###",
		"#...###...#...#",
		"#.#####.#.###.#",
		"#.#...#.#.#...#",
		"#.#.#.#.#.#.###",
		"#...#...#...###",
		"###############",
	}}

	expectedPart1 := "0"
	part1 := day20.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "0"
	part2 := day20.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay21(t *testing.T) {
	day21 := year2024.Day21{Input: []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}}

	expectedPart1 := "126384"
	part1 := day21.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	// expectedPart2 := "Not Implemented"
	// part2 := day21.Part2()
	// if part2 != expectedPart2 {
	// 	t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	// }
}

func TestDay22(t *testing.T) {
	day22 := year2024.Day22{Input: []string{
		"1",
		"10",
		"100",
		"2024",
	}}

	expectedPart1 := "37327623"
	part1 := day22.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	day22 = year2024.Day22{Input: []string{
		"1",
		"2",
		"3",
		"2024",
	}}

	expectedPart2 := "23"
	part2 := day22.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay23(t *testing.T) {
	day23 := year2024.Day23{Input: []string{
		"kh-tc",
		"qp-kh",
		"de-cg",
		"ka-co",
		"yn-aq",
		"qp-ub",
		"cg-tb",
		"vc-aq",
		"tb-ka",
		"wh-tc",
		"yn-cg",
		"kh-ub",
		"ta-co",
		"de-co",
		"tc-td",
		"tb-wq",
		"wh-td",
		"ta-ka",
		"td-qp",
		"aq-cg",
		"wq-ub",
		"ub-vc",
		"de-ta",
		"wq-aq",
		"wq-vc",
		"wh-yn",
		"ka-de",
		"kh-ta",
		"co-tc",
		"wh-qp",
		"tb-vc",
		"td-yn",
	}}

	expectedPart1 := "7"
	part1 := day23.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "co,de,ka,ta"
	part2 := day23.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}

func TestDay24(t *testing.T) {
	day24 := year2024.Day24{Input: []string{
		"x00: 1",
		"x01: 0",
		"x02: 1",
		"x03: 1",
		"x04: 0",
		"y00: 1",
		"y01: 1",
		"y02: 1",
		"y03: 1",
		"y04: 1",
		"",
		"ntg XOR fgs -> mjb",
		"y02 OR x01 -> tnw",
		"kwq OR kpj -> z05",
		"x00 OR x03 -> fst",
		"tgd XOR rvg -> z01",
		"vdt OR tnw -> bfw",
		"bfw AND frj -> z10",
		"ffh OR nrd -> bqk",
		"y00 AND y03 -> djm",
		"y03 OR y00 -> psh",
		"bqk OR frj -> z08",
		"tnw OR fst -> frj",
		"gnj AND tgd -> z11",
		"bfw XOR mjb -> z00",
		"x03 OR x00 -> vdt",
		"gnj AND wpb -> z02",
		"x04 AND y00 -> kjc",
		"djm OR pbm -> qhw",
		"nrd AND vdt -> hwm",
		"kjc AND fst -> rvg",
		"y04 OR y02 -> fgs",
		"y01 AND x02 -> pbm",
		"ntg OR kjc -> kwq",
		"psh XOR fgs -> tgd",
		"qhw XOR tgd -> z09",
		"pbm OR djm -> kpj",
		"x03 XOR y03 -> ffh",
		"x00 XOR y04 -> ntg",
		"bfw OR bqk -> z06",
		"nrd XOR fgs -> wpb",
		"frj XOR qhw -> z04",
		"bqk OR frj -> z07",
		"y03 OR x01 -> nrd",
		"hwm AND bqk -> z03",
		"tgd XOR rvg -> z12",
		"tnw OR pbm -> gnj",
	}}

	expectedPart1 := "2024"
	part1 := day24.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	day24 = year2024.Day24{Input: []string{
		"x00: 0",
		"x01: 1",
		"x02: 0",
		"x03: 1",
		"x04: 0",
		"x05: 1",
		"y00: 0",
		"y01: 0",
		"y02: 1",
		"y03: 1",
		"y04: 0",
		"y05: 1",
		"",
		"x00 AND y00 -> z05",
		"x01 AND y01 -> z02",
		"x02 AND y02 -> z01",
		"x03 AND y03 -> z03",
		"x04 AND y04 -> z04",
		"x05 AND y05 -> z00",
	}}

	//   101010
	// & 101100
	// --------
	//   101000

	// expectedPart2 := "Not Implemented"
	// part2 := day24.Part2()
	// if part2 != expectedPart2 {
	// 	t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	// }
}

func TestDay25(t *testing.T) {
	day25 := year2024.Day25{Input: []string{
		"#####",
		".####",
		".####",
		".####",
		".#.#.",
		".#...",
		".....",
		"",
		"#####",
		"##.##",
		".#.##",
		"...##",
		"...#.",
		"...#.",
		".....",
		"",
		".....",
		"#....",
		"#....",
		"#...#",
		"#.#.#",
		"#.###",
		"#####",
		"",
		".....",
		".....",
		"#.#..",
		"###..",
		"###.#",
		"###.#",
		"#####",
		"",
		".....",
		".....",
		".....",
		"#....",
		"#.#..",
		"#.#.#",
		"#####",
	}}

	expectedPart1 := "3"
	part1 := day25.Part1()
	if part1 != expectedPart1 {
		t.Fatalf("Expected %s, got %s", expectedPart1, part1)
	}

	expectedPart2 := "Not Implemented"
	part2 := day25.Part2()
	if part2 != expectedPart2 {
		t.Fatalf("Expected %s, got %s", expectedPart2, part2)
	}
}
