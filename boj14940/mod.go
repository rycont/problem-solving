package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var width, height int

var cellmap [][]int
var distmap [][]int

var target_x, target_y int
var NOT_VISITED = -1

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var task_q list.List

type Task struct {
	x    int
	y    int
	prev int
}

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &height)
	fmt.Fscan(reader, &width)

	cellmap = make([][]int, height)
	distmap = make([][]int, height)

	for y := 0; y < height; y++ {
		cellmap[y] = make([]int, width)
		distmap[y] = make([]int, width)

		for x := 0; x < width; x++ {
			fmt.Fscan(reader, &cellmap[y][x])

			if cellmap[y][x] == 2 {
				target_x = x
				target_y = y
			}

			if cellmap[y][x] == 0 {
				distmap[y][x] = 0
			} else {
				distmap[y][x] = NOT_VISITED
			}
		}
	}

	task_q.PushBack(Task{
		x:    target_x,
		y:    target_y,
		prev: -1,
	})

	for task_q.Len() > 0 {
		q_item := task_q.Front()
		task_q.Remove(q_item)
		task := q_item.Value.(Task)

		loop(task)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Fprintf(writer, "%d ", distmap[y][x])
		}
		fmt.Fprint(writer, "\n")
	}
}

func loop(task Task) {
	x := task.x
	y := task.y
	prev := task.prev

	if x < 0 || y < 0 || y == height || x == width {
		return
	}

	if cellmap[y][x] == 0 {
		distmap[y][x] = 0
		return
	}

	this_dist := prev + 1

	if distmap[y][x] <= this_dist && distmap[y][x] != NOT_VISITED {
		return
	}

	distmap[y][x] = this_dist

	task_q.PushBack(Task{
		x:    x,
		y:    y + 1,
		prev: this_dist,
	})

	task_q.PushBack(Task{
		x:    x + 1,
		y:    y,
		prev: this_dist,
	})

	task_q.PushBack(Task{
		x:    x,
		y:    y - 1,
		prev: this_dist,
	})

	task_q.PushBack(Task{
		x:    x - 1,
		y:    y,
		prev: this_dist,
	})
}
