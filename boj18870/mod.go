package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var list_size int
var numbers []int

var unique_numbers_map = make(map[int]int)

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &list_size)

	numbers = make([]int, list_size)

	for i := 0; i < list_size; i++ {
		fmt.Fscan(reader, &numbers[i])
		unique_numbers_map[numbers[i]] = 1
	}

	index_map := list_to_index_map(get_keys_sorted(unique_numbers_map))

	for _, value := range numbers {
		fmt.Fprintf(writer, "%d ",index_map[value])
	}
}

func get_keys_sorted(t map[int]int) []int {
	keys := make([]int, 0, len(t))

	for k := range t {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	return keys
}

func list_to_index_map(t []int) map[int]int {
	index_map := make(map[int]int, len(t))

	for index, value := range t {
		index_map[value] = index
	}

	return index_map
}
