package day01

import (
	"adventofcode2024/utils"
	"fmt"
)

func Run() {
    input := utils.ReadFileToArrayOfLines("./samples/sample.txt")
	for index, value := range input {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }
}
