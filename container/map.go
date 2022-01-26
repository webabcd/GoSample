package container

import "fmt"

func MapSample() {
	map_sample1()
}

func map_sample1() {
	var a map[int]string = map[int]string{0: "a", 1: "b"}
	a[0] = "x"
	a[2] = "y"
	fmt.Println(a, a[0], a[2])

	// b := make(map[string]float32, 100)
}
