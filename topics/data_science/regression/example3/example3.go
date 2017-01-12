// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to create training, test, and holdout data sets.
package main

import (
	"bytes"
	"io"
	"log"

	"github.com/kniren/gota/dataframe"
	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm on our localhost.  By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Get the diabetes dataset from Pachyderm's data
	// versioning at the latest commit.
	var b bytes.Buffer
	if err := c.GetFile("diabetes", "master", "diabetes.csv", 0, 0, "", false, nil, &b); err != nil {
		log.Fatal()
	}

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	diabetesDF := dataframe.ReadCSV(bytes.NewReader(b.Bytes()))

	// Calculate the number of elements in each set.
	trainingNum := diabetesDF.Nrow() / 2
	testNum := diabetesDF.Nrow() / 4
	holdoutNum := diabetesDF.Nrow() / 4
	if trainingNum+testNum+holdoutNum < diabetesDF.Nrow() {
		trainingNum++
	}

	// Create the subset indices.
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)
	holdoutIdx := make([]int, holdoutNum)

	// Enumerate the training indices.
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// Enumerate the test indices.
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// Enumerate the holdout indices.
	for i := 0; i < holdoutNum; i++ {
		holdoutIdx[i] = trainingNum + testNum + i
	}

	// Create the subset dataframes.
	trainingDF := diabetesDF.Subset(trainingIdx)
	testDF := diabetesDF.Subset(testIdx)
	holdoutDF := diabetesDF.Subset(holdoutIdx)

	// Create a map that will be used in writing the data
	// to files.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
		2: holdoutDF,
	}

	// Create a repo for our training, test, and holdout data.
	if err := c.CreateRepo("regression_split"); err != nil {
		log.Fatal(err)
	}

	// Create the respective files.
	for idx, setName := range []string{"training.csv", "test.csv", "holdout.csv"} {

		// Create a pipe to push data from the dataframes
		// into Pachyderm.
		r, w := io.Pipe()

		// Create a "regression_split" repo.
		commit, err := c.StartCommit("regression_split", "master")
		if err != nil {
			log.Fatal(err)
		}

		// Put the file into the repo.  Here we will utilize a go
		// routine, because putfile will only end after we close the
		// writer.  However, we can't close the writer until we
		// read it all out of the reader.
		go func() {
			if _, err := c.PutFile("regression_split", commit.ID, setName, r); err != nil {
				log.Fatal(err)
			}
			return
		}()

		// Write the dataframe out as a CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}

		// Close the writer or reading from the other end of the
		// pipe will never finish.
		if err := w.Close(); err != nil {
			log.Fatal(err)
		}

		// Finish the commit.
		if err := c.FinishCommit("regression_split", commit.ID); err != nil {
			log.Fatal(err)
		}

		// Close the pipe reader.
		if err := r.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
