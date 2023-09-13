package main

import (
	"fmt"
)

type Present struct {
	Value int
	Size  int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getValuesAndSizes(presents []Present) ([]int, []int) {
	var values []int
	var sizes []int

	for _, p := range presents {
		values = append(values, p.Value)
		sizes = append(sizes, p.Size)
	}

	return values, sizes
}

func getMemTable(presents []Present, n int) [][]int {
	values, sizes := getValuesAndSizes(presents)
	p := len(presents)

	var table = make([][]int, p+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}

	for i := range table {
		for j := range table[i] {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else if sizes[i-1] <= j {
				table[i][j] = max(values[i-1]+table[i-1][j-sizes[i-1]], table[i-1][j])
			} else {
				table[i][j] = table[i-1][j]
			}
		}
	}

	return table
}

func grabPresents(presents []Present, n int) []Present {
	var knapsack []Present

	presentsMap := make(map[int]Present)
	for _, present := range presents {
		presentsMap[present.Value] = present
	}

	table := getMemTable(presents, n)
	tab := table[len(presents)]
	sum := tab[len(table)-1]

	for i := len(tab) - 1; i >= 0; i-- {
		if sum == 0 {
			break
		}

		diff := sum - tab[i]
		if diff == 0 {
			continue
		}

		if v, ok := presentsMap[diff]; ok == false {
			continue
		} else {
			knapsack = append(knapsack, v)
			sum = sum - diff
		}
	}

	return knapsack
}

func main() {
	var n int = 5
	presents := []Present{
		Present{Value: 15, Size: 1},
		Present{Value: 4, Size: 2},
		Present{Value: 3, Size: 3},
		Present{Value: 5, Size: 4},
	}

	knapsack := grabPresents(presents, n)
	fmt.Println(knapsack)
}