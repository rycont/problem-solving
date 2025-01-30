package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var population int
var relations int

var distance_map [][]int

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &population)
	fmt.Fscan(reader, &relations)

	distance_map = make([][]int, population+1)

	for i := 1; i <= population; i++ {
		distance_map[i] = make([]int, population+1)

		for j := 1; j <= population; j++ {
			distance_map[i][j] = population + 1
		}

		distance_map[i][i] = 0
	}

	for i := 0; i < relations; i++ {
		var t1, t2 int

		fmt.Fscan(reader, &t1)
		fmt.Fscan(reader, &t2)

		distance_map[t1][t2] = 1
		distance_map[t2][t1] = 1
	}

	for via := 1; via <= population; via++ {
		for i := 1; i <= population; i++ {
			for j := 1; j <= population; j++ {
				via_distance := distance_map[i][via] + distance_map[via][j]

				if distance_map[i][j] <= via_distance {
					continue
				}

				distance_map[i][j] = via_distance
				distance_map[j][i] = via_distance
			}
		}
	}

	inssa_id := 1
	inssa_bacon_sum := 0

	for target := 1; target <= population; target++ {
		inssa_bacon_sum += distance_map[1][target]
	}

	for key := 1; key <= population; key++ {
		bacon_sum := 0

		for target := 1; target <= population; target++ {
			bacon_sum += distance_map[key][target]
		}

		if bacon_sum < inssa_bacon_sum {
			inssa_id = key
			inssa_bacon_sum = bacon_sum
		}
	}
	// print_map()
	fmt.Fprintln(writer, inssa_id)
}

func print_map() {

	fmt.Fprint(writer, "┏")

	for j := 1; j < population; j++ {
		fmt.Fprint(writer, "━━━━")
		fmt.Fprint(writer, "┳")
	}

	fmt.Fprint(writer, "━━━━┓")
	fmt.Fprintf(writer, "\n")

	for i := 1; i <= population; i++ {
		fmt.Fprint(writer, "┃")
		for j := 1; j <= population; j++ {
			fmt.Fprintf(writer, "%3d ", distance_map[i][j])
			if j < population {
				fmt.Fprintf(writer, "│")
			} else {
				fmt.Fprintf(writer, "┃")
			}
		}

		fmt.Fprintf(writer, "\n")

		if i == population {
			break
		}

		fmt.Fprint(writer, "┣")

		for j := 1; j < population; j++ {
			fmt.Fprint(writer, "────")
			fmt.Fprint(writer, "┼")
		}

		fmt.Fprint(writer, "────┫")
		fmt.Fprintf(writer, "\n")
	}

	fmt.Fprint(writer, "┗")

	for j := 1; j < population; j++ {
		fmt.Fprint(writer, "━━━━")
		fmt.Fprint(writer, "┻")
	}

	fmt.Fprint(writer, "━━━━┛\n")
}
