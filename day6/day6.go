package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/blackdev1l/aoc2021/utils"
)

func cycle(list *[]int) {

	for i := 0; i < len((*list)); i++ {
		if (*list)[i] == 0 {
			(*list)[i] = 6
			(*list) = append((*list), 9)
		} else {
			(*list)[i]--
		}

	}

	//fmt.Println((*list))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	list := parseInput()
	L1 := list[0:50]
	L2 := list[50:100]
	L3 := list[100:150]
	L4 := list[150:200]
	L5 := list[200:300]
	runFor(256, &wg, L1)
	runFor(256, &wg, L2)
	runFor(256, &wg, L3)
	runFor(256, &wg, L4)
	runFor(256, &wg, L5)

	wg.Wait()
	fmt.Printf("fishes are %v\n", len(L1)+len(L2))
}

func runFor(days int, wg *sync.WaitGroup, list []int) {
	defer wg.Done()
	for i := 1; i < days+1; i++ {
		fmt.Printf("after %v days\n ", i)
		cycle(&list)
	}
}

func parseInput() []int {
	result := []int{}
	file := utils.ReadFIleAsStringLines("input.txt")
	for _, c := range strings.Split(file[0], ",") {
		i, err := strconv.Atoi(c)
		utils.Check(err)
		result = append(result, i)
	}

	return result

}
