package main

import (
	"math/rand"
	"time"

	. "github.com/jaroddev/evolugo/crossovers"
	"github.com/jaroddev/evolugo/genetic"
	. "github.com/jaroddev/evolugo/insertions"
	. "github.com/jaroddev/evolugo/mutations"
	. "github.com/jaroddev/evolugo/selections"

	"github.com/jaroddev/evotest/problems"
	"github.com/jaroddev/evotest/record"
)

// basic config is best-three-clone-elitist.csv

func main() {

	// prepare 30 seeds
	rand.Seed(time.Now().UnixNano())

	algorithm := &genetic.GA{}

	onemax := problems.NewBasicConfig()
	onemax.Attach(algorithm)

	// parameters
	algorithm.Mutation = &Flip{Frequency: 3}
	algorithm.Selection = &SelectBest{ParentNumber: 2}
	algorithm.Crossover = &Clone{ChildrenNumber: 2}
	algorithm.Insertion = &Elitist{}

	recorder := record.NewRecorder()
	algorithm.Recorder = &recorder

	algorithm.Run()
	recorder.Save()
}
