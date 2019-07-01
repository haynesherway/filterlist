package main

import (
	"flag"
	"fmt"

	"github.com/haynesherway/filterlist"
)

var fruitList filterlist.FilterList

func main() {
	flag.Var(&fruitList, "fruits", "fruits to search for")
	flag.Parse()

	fmt.Println("Searching for the following fruits:")
	fmt.Println(fruitList)
}
