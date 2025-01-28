package main

import (
	"bufio"
	"fmt"
	"os"
)

var connections_by_nodes map[int]([]int)

var nodes_quantity int
var edges_quantity int

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &nodes_quantity)
	fmt.Fscan(reader, &edges_quantity)

	connections_by_nodes = make(map[int][]int, nodes_quantity)

	for i := 1; i <= nodes_quantity; i++ {
		connections_by_nodes[i] = []int{}
	}

	var p1, p2 int
	
	for i := 0; i < edges_quantity; i++ {
		fmt.Fscan(reader, &p1)
		fmt.Fscan(reader, &p2)

		connections_by_nodes[p1] = append(connections_by_nodes[p1], p2)
		connections_by_nodes[p2] = append(connections_by_nodes[p2], p1)
	}

	groups := 0

	for {
		target_key := get_first_key()

		if target_key == -1 {
			break
		}

		loop(target_key)
		groups++
	}

	fmt.Fprintln(writer, groups)
}

func loop(target int) {
	target_connections := connections_by_nodes[target]
	delete(connections_by_nodes, target)

	if len(target_connections) == 0 {
		return
	}

	for _, connected_node := range target_connections {
		loop(connected_node)
	}
}

func get_first_key() int {
	for key := range connections_by_nodes {
		return key
	}

	return -1
}
