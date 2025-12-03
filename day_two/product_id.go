package daytwo

import (
	utils "advent_of_code/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func formatInput(list []string) []string {
	formattedList := []string{}

	for _, line := range list {
		split := strings.Split(strings.TrimSpace(line), ",")
		formattedList = slices.Concat(formattedList, split)
	}

	return formattedList
}

func splitStringIntoChunks(str string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(str); i += chunkSize {
		end := i + chunkSize
		if end > len(str) {
			end = len(str)
		}
		chunks = append(chunks, str[i:end])
	}
	return chunks
}

func GiftShop() {
	file := utils.ReadFile("day_two/input.txt")
	input := formatInput(file)

	invalidIds := []int{}
	totalSum := 0

	for _, numRange := range input {

		if len(numRange) == 0 {
			continue
		}

		ranges := strings.Split(numRange, "-")
		leftStr := ranges[0]
		rightStr := ranges[1]

		left, err := strconv.Atoi(leftStr)

		if err != nil {
			panic(fmt.Sprint("Invalid range", ranges))
		}

		right, err := strconv.Atoi(rightStr)

		if err != nil {
			panic(fmt.Sprint("Invalid range", ranges))
		}

		for i := left; i <= right; i++ {
			strI := strconv.Itoa(i)
			if len(strI)%2 != 0 {
				continue
			}

			middle := len(strI) / 2

			firstHalf := strI[:middle]
			secondHalf := strI[middle:]

			if firstHalf == secondHalf {
				invalidIds = append(invalidIds, i)
			}

		}

	}

	for _, id := range invalidIds {
		totalSum = totalSum + id
	}

	fmt.Println("[Day_two] Total Sum: ", totalSum)

}

func GiftShopPartTwo() {
	file := utils.ReadFile("day_two/input.txt")
	input := formatInput(file)

	invalidIds := []int{}
	totalSum := 0

	for _, numRange := range input {

		if len(numRange) == 0 {
			continue
		}

		ranges := strings.Split(numRange, "-")
		leftStr := ranges[0]
		rightStr := ranges[1]

		left, err := strconv.Atoi(leftStr)

		if err != nil {
			panic(fmt.Sprint("Invalid range", ranges))
		}

		right, err := strconv.Atoi(rightStr)

		if err != nil {
			panic(fmt.Sprint("Invalid range", ranges))
		}

		for i := left; i <= right; i++ {
			strI := strconv.Itoa(i)

			middle := len(strI) / 2

			for j := 1; j <= middle; j++ {

				chunks := splitStringIntoChunks(strI, j)

				allRepeatedValues := true

				if len(chunks) == 2 {
					if chunks[0] != chunks[1] {
						allRepeatedValues = false
					}
				} else {
					for k := 1; k < len(chunks)-1; k++ {
						if chunks[k-1] != chunks[k] || chunks[k] != chunks[k+1] {
							allRepeatedValues = false
						}
					}
				}

				if allRepeatedValues && !slices.Contains(invalidIds, i) {
					invalidIds = append(invalidIds, i)
				}

			}
		}

	}

	for _, id := range invalidIds {
		totalSum = totalSum + id
	}

	fmt.Println("[Day_two] Total Sum: ", totalSum)
}
