package record

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Row struct {
	Generation  int
	BestFitness float64
}

type Recorder struct {
	Rows []Row
}

func NewRecorder() Recorder {
	return Recorder{
		Rows: make([]Row, 0),
	}
}

func Merge(recorders ...Recorder) Recorder {
	recorder := NewRecorder()
	maxLength := 0

	for index := range recorders {
		if maxLength < len(recorders[index].Rows) {
			maxLength = len(recorders[index].Rows)
		}
	}

	for i := 0; i < maxLength; i++ {
		newRow := Row{}
		average := 0

		for index := range recorders {
			if len(recorders[index].Rows) > i {
				newRow.BestFitness += recorders[index].Rows[i].BestFitness
				newRow.Generation = recorders[0].Rows[i].Generation
				average++
			}
		}

		newRow.BestFitness /= float64(average)
		recorder.Rows = append(recorder.Rows, newRow)
	}

	return recorder
}

func (recorder *Recorder) Record(Generation int, BestFitness float64) {
	newRow := &Row{Generation, BestFitness}
	recorder.Rows = append(recorder.Rows, *newRow)
}

func (r *Recorder) Save() {
	records := [][]string{
		{"generation", "fitness"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, row := range r.Rows {
		records = append(records, []string{
			fmt.Sprint(row.Generation), fmt.Sprint(row.BestFitness),
		})
	}

	w.WriteAll(records) // calls Flush internally
	if err := w.Error(); err != nil {
		fmt.Println("whouah, got an error man")
	}
}
