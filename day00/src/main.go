package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() []int {
	arr := make([]int, 0)

	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a sequence of integers in range [-100000 100000], separated by a newline:")
	for scan.Scan() {
		numStr := scan.Text()
		numNum, err := strconv.Atoi(strings.Trim(numStr, "\n\t\r "))
		if err != nil || numNum < -100000 || numNum > 100000 {
			fmt.Println("wrong input value")
			continue
		}
		arr = append(arr, numNum)
	}
	return arr
}

func countOccurrence(arr []int) map[int]int{
	arrMap := make(map[int]int, 0)

	for _, num := range arr {
		if _, exist := arrMap[num]; exist == false {
			arrMap[num] = 0
		}
		arrMap[num] += 1
	}
	return arrMap
}

func findMode(arr []int) int {
	arrMap := countOccurrence(arr)
	var smallest, count int

	for key, value := range arrMap {
		if value > count {
			smallest = key
			count = value
		} else if value == count && key < smallest{
			smallest = key
		}
	}
	return smallest
}

func findSumQuantity(arr []int) (int, int) {
	var sum int
	var i int
	for _, num := range arr {
		sum += num
		i++
	}
	return sum, i
}

func findMedian(arr []int, sum int, quantity int) float64 {
	sort.Ints(arr)
	middle := quantity / 2

	if quantity % 2 == 0 {
		return float64(arr[middle - 1] + arr[middle]) / 2
	}
	return float64(arr[middle])
}

func findSD(arr []int, mean float64, quantity int) float64 {
	var sdSum float64
	for _, num := range arr {
		sd := math.Pow(float64(num) - mean, 2)
		sdSum += sd
	}
	return math.Sqrt(sdSum / float64(quantity))
}

func main() {
	arr := readInput()
	if len(arr) == 0 {
		log.Fatalln("No integers were entered.")
	}
	sum, quantity := findSumQuantity(arr)
	mode := findMode(arr)
	mean := float64(sum) / float64(quantity)
	median := findMedian(arr, sum, quantity)
	sd := findSD(arr, mean, quantity)

	flagMode := flag.Bool("Mode", false, "display mode value")
	flagMean := flag.Bool("Mean", false, "display mean value")
	flagMedian := flag.Bool("Median", false, "display median value")
	flagSD := flag.Bool("SD", false, "display standard deviation")
	flag.Parse()

	noFlags := !*flagMode && !*flagMean && !*flagMedian && !*flagSD

	if *flagMean || noFlags {
		fmt.Printf("Mean: %.2f\n", mean)
	}
	if *flagMedian || noFlags {
		fmt.Printf("Median: %.2f\n", median)
	}
	if *flagMode || noFlags {
		fmt.Println("Mode: ", mode)
	}
	if *flagSD || noFlags {
		fmt.Printf("SD: %.2f\n", sd)
	}
}