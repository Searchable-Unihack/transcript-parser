package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Transcript struct {
	Created time.Time
	Id      string
	Updated time.Time
	Results []Results
	Status  string
}

type Results struct {
	ResultIndex int64
	Results     []Result
}

type Result struct {
	Final        bool
	Alternatives []Alternative
}

type Alternative struct {
	Transcript string
	Confidence float64
	Timestamps []Word
}

type Record struct {
	Time   float64 `json:"time"`
	Text   string  `json:"text"`
	Weight float64 `json:"weight"`
	Source string  `json:"source"`
}

type Word []interface{}

func main() {

	args := os.Args

	if len(args) != 3 {
		log.Fatalf("You must specify an input file and an output file (in that order).")
	}

	inputFile := args[1]
	outputFile := args[2]

	rawInput, err := ioutil.ReadFile(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	var transcript Transcript
	err = json.Unmarshal(rawInput, &transcript)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("results: %d", len(transcript.Results[0].Results))

	var records []Record

	for _, result := range transcript.Results[0].Results {
		alternative := result.Alternatives[0]
		// log.Printf("%d: %s (%0.2f)", i, alternative.Transcript, alternative.Timestamps[0][1].(float64))
		var record Record
		record.Time = alternative.Timestamps[0][1].(float64)
		record.Text = alternative.Transcript
		record.Weight = 1.0 + alternative.Confidence
		record.Source = "audio"
		records = append(records, record)
	}

	output, err := json.Marshal(&records)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(outputFile, output, 0644)

	if err != nil {
		panic(err)
	}

}
