package day_one

import (
	utils "advent_of_code/utils"
	"fmt"
	"strconv"
)

// from 0 - 99
const NUMBERS_IN_DIAL = 100

func CalculatePassword() {
	position := 50
	password := 0

	file := utils.ReadFile("day_one/input.txt")

	for _, line := range file {
		rotation := string(line[0])
		rest := line[1:]
		amount, err := strconv.Atoi(rest)

		if err != nil {
			panic(fmt.Sprint("Invalid line, expecting number got:", rest))
		}

		if rotation == "L" {
			position = position - amount

			for position < 0 {
				position = NUMBERS_IN_DIAL - utils.AbsInt(position)
			}
		}

		if rotation == "R" {
			position = position + amount

			for position >= 100 {
				position = position - NUMBERS_IN_DIAL
			}
		}

		if position == 0 {
			password = password + 1
		}

	}

	fmt.Println("[Day_One] Password:", password)
}

func CalculatePassword0x434C49434B() {
	position := 50
	password := 0

	file := utils.ReadFile("day_one/input.txt")

	for _, line := range file {
		rotation := string(line[0])
		rest := line[1:]
		amount, err := strconv.Atoi(rest)

		if err != nil {
			panic(fmt.Sprint("Invalid line, expecting number got:", rest))
		}

		if rotation == "L" {
			// weird, but I know that I counted one extra
			if position == 0 {
				password--
			}

			position = position - amount

			for position < 0 {
				password++
				position = NUMBERS_IN_DIAL - utils.AbsInt(position)
			}

			if position == 0 {
				password = password + 1
			}

		}

		if rotation == "R" {
			position = position + amount

			for position >= 100 {
				password++
				position = position - NUMBERS_IN_DIAL
			}

		}

	}

	fmt.Println("[Day_One] Password:", password)
}
