package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"errors"
)

// readCSVFile() accepts a file and returns Contact struct slice (matches JSON schema) and logErrors struct slice (error validation)
// It uses the std go library's csv.Reader() to read the file and calls validateRecord() to make sure line-by-line validation is done
func readCSVFile(inputPath string) (contacts []Contact, logErrors []csv.ParseError, err error) {

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read()
	currentLine := 1

	for fields, err := reader.Read(); err != io.EOF; fields, err = reader.Read() {
		currentLine += 1
		if record, err := validateRecord(fields); err != nil {
			logErrors = append(logErrors, csv.ParseError{Line: currentLine, Err: err})
		} else {
			contacts = append(contacts, record)
		}
	}
	return contacts, logErrors, nil
}

func writeRecordsToJSON(filename string, records []Contact) error {
	asJSON, err := json.MarshalIndent(records,"","\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, asJSON, 0644)

}

func deleteCSVAfterProcessing(path string) error {
	if filepath.Ext(path) != ".csv" {
		return errors.New("non .csv file passed as argument")
	}
	return os.Remove(path)
}

func writeErrorsToFileAsCSV(filename string, errs []csv.ParseError) error {
	return ioutil.WriteFile(filename, errorsToCSV(errs), 0644)
}

func errorsToCSV(lineErrors []csv.ParseError) []byte {
	lines := make([]string, len(lineErrors)+1)
	if len(lineErrors) > 0 {
		lines[0] = "LINE_NUM,ERROR_MSG"
		for i, err := range lineErrors {
			lines[i+1] = fmt.Sprintf("%d,%s", err.Line, err.Err)
		}
	}
	return []byte(strings.Join(lines, "\n"))
}

