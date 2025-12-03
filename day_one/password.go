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

	fmt.Println("Password:", password)
}
