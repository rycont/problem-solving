package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var field [][]byte

var white_pieces = 0
var blue_pieces = 0

func main() {
	defer writer.Flush()

	var field_size int

	fmt.Fscan(reader, &field_size)

	field = make([][]byte, field_size)

	reader.ReadLine()

	for i := 0; i < field_size; i++ {
		bytes, _, _ := reader.ReadLine()
		field[i] = []byte{}

		for _, cell := range bytes {
			if cell == 32 {
				continue
			}

			field[i] = append(field[i], cell-48)
		}
	}

	figure_pieces(0, 0, field_size-1, field_size-1)

	fmt.Fprintln(writer, white_pieces)
	fmt.Fprintln(writer, blue_pieces)
}

func figure_pieces(start_x int, start_y int, end_x int, end_y int) {
	if start_x == end_x {
		cell_item := int(field[start_y][start_x])

		if cell_item == 1 {
			blue_pieces++
		} else if cell_item == 0 {
			white_pieces++
		}

		return
	}

	filled_items := get_filled_item(start_x, start_y, end_x, end_y)

	if filled_items == 1 {
		blue_pieces++
		return
	} else if filled_items == 0 {
		white_pieces++
		return
	}

	x_mid := (start_x + end_x) / 2
	y_mid := (start_y + end_y) / 2

	figure_pieces(
		start_x,
		start_y,
		x_mid,
		y_mid,
	)

	figure_pieces(
		x_mid+1,
		start_y,
		end_x,
		y_mid,
	)

	figure_pieces(
		start_x,
		y_mid+1,
		x_mid,
		end_y,
	)

	figure_pieces(
		x_mid+1,
		y_mid+1,
		end_x,
		end_y,
	)

}

func get_filled_item(start_x int, start_y int, end_x int, end_y int) int {
	first_seq := field[start_y][start_x:(end_x + 1)]
	repetitive_item := get_repetitive_item(first_seq)

	if repetitive_item == -1 {
		return -1
	}

	for i := start_y; i <= end_y; i++ {
		comp_seq := field[i][start_x:(end_x + 1)]
		is_equal := bytes.Compare(first_seq, comp_seq) == 0

		if !is_equal {
			return -1
		}
	}

	return repetitive_item
}

func get_repetitive_item(seq []byte) int {
	init_value := seq[0]

	for _, v := range seq {
		if v != init_value {
			return -1
		}
	}

	return int(init_value)
}
