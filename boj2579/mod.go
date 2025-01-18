package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var steps []int
var lines int
var memoMap = map[int]int{}

func main() {
	var cur int
	var scoreSum = 0

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	lines, _ = strconv.Atoi(sc.Text())

	for i := 0; i < lines; i++ {
		sc.Scan()
		cur, _ = strconv.Atoi(sc.Text())

		scoreSum += cur
		steps = append(steps, cur)
		
		memoMap[i] = -1
	}
	
	fmt.Println(scoreSum - min(figureOpt(1) + steps[0], figureOpt(0)))

}

func r_figureOpt(p int) int {
	var leftSlices = lines - p

	if leftSlices > 4 {
		var case1 = steps[p+1] + figureOpt(p+2)
		var case2 = steps[p+2] + figureOpt(p+3)

		return min(case1, case2)
	}

	if leftSlices == 4 {
		return min(steps[p+1], steps[p+2])
	}

	if leftSlices == 3 {
		return steps[p+1]
	}

	return 0
}


func figureOpt(position int) int {
	if (memoMap[position]) != -1 {
		return memoMap[position]
	}

	var figuredValue = r_figureOpt(position)
	memoMap[position] = figuredValue

	return figuredValue
}
