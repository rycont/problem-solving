package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var cases int

	fmt.Fscanln(reader, &cases)

	for ; cases > 0; cases-- {
		var width, height, cabbages_count int
		fmt.Fscanln(reader, &width, &height, &cabbages_count)

		var cabbages = make([][]int, cabbages_count)

		for i := 0; i < cabbages_count; i++ {
			var x, y int
			fmt.Fscanln(reader, &x, &y)

			cabbages[i] = []int{x, y}
		}

		solve(width, height, cabbages)
	}
}

func solve(width int, height int, cabbages [][]int) {
	field := make([][]bool, width)
	for i := range field {
		field[i] = make([]bool, height)
	}

	for _, cabbage_position := range cabbages {
		field[cabbage_position[0]][cabbage_position[1]] = true
	}

	bugs := 0

	for ; ; bugs++ {
		bug_point_x := -1
		bug_point_y := -1

		for seek_x, line := range field {
			for seek_y, cell := range line {
				if cell {
					bug_point_x = seek_x
					bug_point_y = seek_y

					break
				}
			}
		}

		if bug_point_x == -1 {
			break
		}

		loop(width, height, bug_point_x, bug_point_y, field)
	}

	fmt.Println(bugs)
}

func loop(width int, height int, curepoint_x int, curepoint_y int, field [][]bool) {
	if !field[curepoint_x][curepoint_y] {
		return
	}

	field[curepoint_x][curepoint_y] = false

	if curepoint_x != width-1 {
		loop(width, height, curepoint_x+1, curepoint_y, field)
	}

	if curepoint_y != height-1 {
		loop(width, height, curepoint_x, curepoint_y+1, field)
	}

	if curepoint_x != 0 {
		loop(width, height, curepoint_x-1, curepoint_y, field)
	}

	if curepoint_y != 0 {
		loop(width, height, curepoint_x, curepoint_y-1, field)
	}

}
