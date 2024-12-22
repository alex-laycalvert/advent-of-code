package year2024

import (
	"fmt"
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
		secretValue, err := strconv.Atoi(line)
		utils.PanicErr(err)
		secret := Secret{secretValue, 0, 0}
		for range 2000 {
			secret = nextSecret(secret)
		}
		answer += secret.Value
	}

	return strconv.Itoa(answer)
}

func (d Day22) Part2() string {
	secrets := make([][]Secret, len(d.Input))
	numSecrets := 2000
	for i, line := range d.Input {
		secrets[i] = make([]Secret, numSecrets)
		secretValue, err := strconv.Atoi(line)
		utils.PanicErr(err)
		secret := Secret{
			Value:  secretValue,
			Change: 0,
			Price:  secretValue % 10,
		}
		for j := range numSecrets {
			secret = nextSecret(secret)
			secrets[i][j] = secret
		}
	}

	uniqueChangeKeys := make(map[string]string)
	allSequences := make([]map[string]int, len(secrets))
	for i, secrets := range secrets {
		allSequences[i] = make(map[string]int, 0)
		for j := 0; j < len(secrets)-3; j++ {
			changeKey := fmt.Sprintf("%d,%d,%d,%d", secrets[j].Change, secrets[j+1].Change, secrets[j+2].Change, secrets[j+3].Change)
			if _, ok := allSequences[i][changeKey]; ok {
				continue
			}
			allSequences[i][changeKey] = secrets[j+3].Price
			if _, ok := uniqueChangeKeys[changeKey]; ok {
				continue
			}
			uniqueChangeKeys[changeKey] = changeKey
		}
	}

	maxPrice := 0
	for changeKey := range uniqueChangeKeys {
		sum := 0
		for _, changes := range allSequences {
			if price, ok := changes[changeKey]; ok {
				sum += price
			}
		}
		if sum > maxPrice {
			maxPrice = sum
		}
	}

	return strconv.Itoa(maxPrice)
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
