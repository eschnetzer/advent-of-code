package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sumArray(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func main() {
	var list1 []int
	var list2 []int
	var distances []int
	var sum int

	// Open the file
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read the file as a CSV
	reader := csv.NewReader(input)
	inputData, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Loop through the records
	for _, record := range inputData {
		for i := 0; i < len(record); i++ {
			line := strings.Split(record[i], "   ")
			i1, err := strconv.Atoi(line[0])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			list1 = append(list1, i1)

			i2, err := strconv.Atoi(line[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			list2 = append(list2, i2)
		}

	}

	// Sort the lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Calculate distance between values
	for i := 0; i < len(list1); i++ {
		distances = append(distances, int(math.Abs(float64(list2[i]-list1[i]))))
	}

	sum = sumArray(distances)

	fmt.Println(sum)
}
