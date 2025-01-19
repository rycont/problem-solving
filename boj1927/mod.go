package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var heap = []int{-1}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var instructions_count int
	fmt.Fscan(reader, &instructions_count)

	for i := 0; i < instructions_count; i++ {
		var instruction int
		fmt.Fscan(reader, &instruction)

		if instruction == 0 {
			remove_min()
		} else {
			insert(instruction)
		}
	}
}

func insert(value int) {
	heap = append(heap, value)

	checking_index := len(heap) - 1
	parent_index := int(math.Floor(float64(checking_index) / 2))

	for heap[checking_index] < heap[parent_index] {
		heap[parent_index], heap[checking_index] = heap[checking_index], heap[parent_index]

		if parent_index == 0 {
			break
		}

		checking_index = parent_index
		parent_index = int(math.Floor(float64(parent_index) / 2))
	}
}

func remove_min() {
	heap_len := len(heap)

	if heap_len == 1 {
		fmt.Fprintln(writer, 0)
		return
	}

	last_node_index := heap_len - 1
	fmt.Fprintln(writer, heap[1])

	heap[1] = heap[last_node_index]
	heap = heap[:last_node_index]

	if heap_len == 2 {
		return
	}

	heap_len--

	anchor_index := 1

	for {
		anchor_left := anchor_index * 2
		anchor_right := anchor_left + 1

		anchor_value := heap[anchor_index]

		if heap_len <= anchor_left {
			break
		}

		anchor_left_value := heap[anchor_left]

		if heap_len <= anchor_right {
			if anchor_left_value < anchor_value {
				heap[anchor_index], heap[anchor_left] = heap[anchor_left], heap[anchor_index]
			}

			break
		}

		anchor_right_value := heap[anchor_right]

		left_valid := anchor_value < anchor_left_value
		right_valid := anchor_value < anchor_right_value

		if left_valid && right_valid {
			break
		}

		if !left_valid && !right_valid {
			left_valid = anchor_right_value < anchor_left_value
			right_valid = !left_valid
		}

		if !left_valid {
			heap[anchor_index], heap[anchor_left] = heap[anchor_left], heap[anchor_index]

			anchor_index = anchor_left
			continue
		}

		if !right_valid {
			heap[anchor_index], heap[anchor_right] = heap[anchor_right], heap[anchor_index]

			anchor_index = anchor_right
			continue
		}
	}
}
