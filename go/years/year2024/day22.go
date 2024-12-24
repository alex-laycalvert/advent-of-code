package year2024

import (
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

const BASE = 16777216

// ANSWERS:
//
// Part 1: 21147129593
//
// Part 2: 2445
type Day22 struct {
	Input []string
}

func (d Day22) Part1() string {
	answer := 0
	for _, line := range d.Input {
		secret, err := strconv.Atoi(line)
		utils.PanicErr(err)
		for range 2000 {
			secret, _ = next(secret)
		}
		answer += secret
	}

	return strconv.Itoa(answer)
}

func (d Day22) Part2() string {
	numSecrets := 2000
	maxValue := 0
	maxPrices := make(map[int]int)
	seen := make(map[int]int)

	for _, line := range d.Input {
		clear(seen)
		value, err := strconv.Atoi(line)
		utils.PanicErr(err)

		// get the first 4 secrets/prices
		value, price1 := next(value)
		value, price2 := next(value)
		value, price3 := next(value)
		value, price4 := next(value)

		for range numSecrets - 4 {
			// get the next secret/price
			nextValue, nextPrice := next(value)
			value = nextValue
			// this emulates a sliding window of the list of all changes in price
			key := getKey(price2-price1, price3-price2, price4-price3, nextPrice-price4)
			if _, ok := seen[key]; !ok {
				seen[key] = key
				maxPrices[key] += nextPrice
				if maxPrices[key] > maxValue {
					maxValue = maxPrices[key]
				}
			}

			// shift all prices forward to get next window
			price1, price2, price3, price4 = price2, price3, price4, nextPrice
		}
	}
	return strconv.Itoa(maxValue)
}

type Secret struct {
	Value  int
	Price  int
	Change int
}

func nextSecret(secret Secret) Secret {
	value := (secret.Value ^ (secret.Value * 64)) % BASE
	value = (value ^ (value / 32)) % BASE
	value = (value ^ (value * 2048)) % BASE
	price := value % 10
	return Secret{
		Value:  value,
		Price:  price,
		Change: price - secret.Price,
	}
}

func next(secret int) (value int, price int) {
	value = (secret ^ (secret * 64)) % BASE
	value = (value ^ (value / 32)) % BASE
	value = (value ^ (value * 2048)) % BASE
	return value, value % 10
}

func getKey(n1, n2, n3, n4 int) int {
	return (n1+9)*6859 + (n2+9)*361 + (n3+9)*19 + (n4 + 9)
}
