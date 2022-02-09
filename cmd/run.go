package main

import (
	"math/rand"

	. "github.com/jaroddev/evolugo/crossovers"
	"github.com/jaroddev/evolugo/genetic"
	. "github.com/jaroddev/evolugo/insertions"
	. "github.com/jaroddev/evolugo/mutations"
	. "github.com/jaroddev/evolugo/selections"

	"github.com/jaroddev/evotest/problems"
	"github.com/jaroddev/evotest/record"
)

const filename = "seeds.txt"

func configureAlgorithm(recorder genetic.Recorder) *genetic.GA {
	algorithm := &genetic.GA{}

	onemax := problems.NewBasicConfig()
	onemax.Attach(algorithm)

	// parameters
	algorithm.Selection = &Tournament{Participant: 40, Winner: 2}
	algorithm.Mutation = &Flip{Frequency: 2}
	algorithm.Crossover = &MonoPoint{ChildrenNumber: 2}
	algorithm.Insertion = &Elitist{}

	algorithm.Recorder = recorder

	return algorithm
}

func main() {
	lines := record.Read(filename)
	seeds := record.ConvertLinesToSeeds(lines)
	recorders := make([]record.Recorder, 0)

	for _, seed := range seeds {
		rand.Seed(seed)

		recorder := record.NewRecorder()
		algorithm := configureAlgorithm(&recorder)

		algorithm.Run()
		recorders = append(recorders, recorder)
	}

	mergedRecorder := record.Merge(recorders...)
	mergedRecorder.Save()
}
