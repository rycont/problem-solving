package main

import (
	"bufio"
	"fmt"
	"os"
)

var tree_quantity int
var tree_heights []int
var goal int

var min_height = 1000000000
var max_height = -1

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &tree_quantity)
	fmt.Fscan(reader, &goal)

	tree_heights = make([]int, tree_quantity)

	var tmp int

	for i := 0; i < tree_quantity; i++ {
		fmt.Fscan(reader, &tmp)

		if tmp < min_height {
			min_height = tmp
		}
		
		if max_height < tmp {
			max_height = tmp
		}

		tree_heights[i] = tmp
	}

	middle := (min_height + max_height) / 2
	unit := middle

	for {
		trimmed_height := get_trimmed_heights(middle)

		if goal == trimmed_height {
			fmt.Fprintln(writer, middle)
			return
		}

		satisfied := goal < trimmed_height

		if unit == 1 {
			if satisfied {
				for goal <= get_trimmed_heights(middle + 1) {
					middle++
				}
			} else {
				for get_trimmed_heights(middle) < goal {
					middle--
				}
			}
			
			fmt.Fprintln(writer, middle)
			return 
		}

		if satisfied {
			middle += unit / 2
		} else {
			middle -= unit / 2
		}

		unit = unit / 2
	}
}

func get_trimmed_heights(cutter_height int) int {
	summary := 0

	for _, height := range tree_heights {
		if height > cutter_height {
			summary += height - cutter_height
		}
	}

	return summary
}
