package main

import (
	"errors"
	"log"
	"fmt"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	switch {
	case len(arr) < 1:
		return 0, errors.New("empty line")
	case idx < 0:
		return 0, errors.New("invalid index\n")
	case idx > len(arr):
		return 0, errors.New("there is no such index\n")
	}
	start := unsafe.Pointer(&arr[0])
	size := unsafe.Sizeof(int(0))

	item := *(*int)(unsafe.Pointer(uintptr(start) + size*uintptr(idx)))
	return item, nil
}

func main () {
	array := []int{19, 30, 5, 691, 10, 2}

	res, err := getElement(array, 14)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
	fmt.Println("--------------------------")

	array = []int{19, 30, 5, 691, 10, 2}

	new, err := getElement(array, 3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(new)
	fmt.Println("--------------------------")

	array = []int{19, 30, 5, 691, 10, 2}

	tt, err := getElement(array, -2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tt)

	fmt.Println("--------------------------")

	array = []int{}

	ww, err := getElement(array, 8)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ww)

}