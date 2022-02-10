package problems

import "math/rand"

var whole float64 = 1.0

func CreateProbaVectorFromOperatorList(operatorList []interface{}) (proba []float64) {
	if len(operatorList) == 0 {
		panic("pls initialize the operator list")
	}

	// Share to all operator, the same initial probability
	share := whole / float64(len(operatorList))
	proba = make([]float64, 0)

	for range operatorList {
		proba = append(proba, share)
	}

	return
}

func Pick(proba []float64) int {
	n := rand.Float64()
	sum := 0.0
	for index := range proba {
		sum += proba[index]
		if n <= sum {
			return index
		}
	}

	return 0 // or 1
}

func Update(proba []float64, pmin float64, pos int, improvement, rewardMultiplyer float64) {

	viableOperatorCount := getViableOperatorCount(proba, pmin)

	if viableOperatorCount == 0 {
		panic("tmp measure")
	}

	reseted := ResetPosFromProba(proba, pos)

	if improvement > 0 {
		if viableOperatorCount > 1 {
			addReward(proba, reseted, pmin, pos, improvement, rewardMultiplyer)
		}
	} else {
		applyPenalty(proba, reseted, pmin, pos, rewardMultiplyer)
	}

	checkError(proba, pos)
}

func getViableOperatorCount(proba []float64, pmin float64) (viable int) {
	for index := range proba {
		if proba[index] > pmin {
			viable++
		}
	}

	return
}

func ResetPosFromProba(proba []float64, pos int) (ratios []float64) {
	// copy of the proba vector
	ratios = make([]float64, len(proba))
	copy(ratios, proba)

	total := 0.0

	for index := range proba {
		if index != pos {
			total += whole - proba[index]
		}
	}

	for index := range proba {
		if index != pos {
			ratios[index] = (whole - proba[index]) / total
		}
	}

	ratios[pos] = 0.0
	return
}

func addReward(proba []float64, reseted []float64, pmin float64, pos int, improvement, rewardMultiplyer float64) {
	reward := improvement * rewardMultiplyer
	gap := 0.0
	for index := range proba {
		if index != pos {
			proba[index] = proba[index] - reseted[index]*reward
			if proba[index] < pmin {
				gap += pmin - proba[index]
				proba[index] = pmin
			}
		}
	}
	proba[pos] += reward - gap
}

func applyPenalty(proba []float64, reseted []float64, pmin float64, pos int, rewardMultiplyer float64) {
	var reward float64
	penalty := 2 * rewardMultiplyer * float64((len(proba) - 1))

	if proba[pos]-penalty >= pmin {
		reward = 2 * rewardMultiplyer
	} else {
		reward = (2 * rewardMultiplyer) - (pmin - (proba[pos] - penalty))
	}

	for index := range proba {
		if index == pos {
			proba[index] -= reward * float64(len(proba)-1)
		} else {
			proba[index] += reward * ((1.0 - reseted[index]) / float64(len(proba)-2))
		}
	}
}

func checkError(proba []float64, pos int) {
	sum := 0.0
	for index := range proba {
		sum += proba[index]
	}

	if sum > 100 {
		proba[pos] += 100.0 - sum
	}

}
