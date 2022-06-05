package main

import (
	_ "embed"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/src"
)

func main() {
	numberOfRows := 1e4
	src.Naive(uint64(numberOfRows))

}
