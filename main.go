package main

import (
	"fmt"

	"github.com/bitmaskit/go-insertion-sort/inserion"
)

func main() {
	arr := []int{4,3,2,1}
	insertion.Sort(arr)
	fmt.Println(arr)
}
