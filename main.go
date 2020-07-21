package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type RelevantPaths struct {
	inputDirPath string
	outputDirPath string
	errorDirPath string
}

type Contact struct {
	ID   int `json:"id"`
	Name struct {
		First  string `json:"first"`
		Middle string `json:"middle,omitempty"`
		Last   string `json:"last"`
	} `json:"name"`
	Phone string `json:"phone"`
}

func main() {
	p := new(RelevantPaths)
	flag.StringVar(&p.inputDirPath, "input-directory","./in","input dir to watch for new csv files")
	flag.StringVar(&p.outputDirPath,"output-directory","./out","output dir to store  JSON files")
	flag.StringVar(&p.errorDirPath,"error-directory","./err","error dir to store validation logs")
	flag.Parse()
	go func() {
		c := time.Tick(1 * time.Second)
		for range c {
			// Note this loop runs the function
			// in the same goroutine so we make sure there is
			// only ever one.
			p.watchInputDir()
		}
	}()
	// Blocking the main() from ending so the directory watcher is indefinitely active until interrupted:
	select {}

}

func (p RelevantPaths) watchInputDir() {
	seenFiles := make(map[string]int)
	currentDir, err := ioutil.ReadDir(p.inputDirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range currentDir {
		fileName := file.Name()
		if fileName != ".DS_Store" {
			if file.IsDir() || path.Ext(fileName) != ".csv"  {
				log.Printf("omitting %s, not a csv file \n", fileName)
				continue
			} else if _, repeatedFile := seenFiles[fileName]; repeatedFile {
				log.Printf("file %s has been previously processed", fileName)
			} else {
				// Safe to proceed to processing
				err := p.processCSV(fileName)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}
}

/**
processCSV() orchestrates the workflow of processing the the CSV file
(1) It calls readCSVFile() to first reads the input file
(2) It calls writeRecordsToJSON() to write the validated fields as output
(3) It calls to writeErrorsToFileAsCSV() the error log (if any) to CSV
(4) It finally calls deleteCSVAfterProcessing() to delete the input file from the input dir
 */
func (p RelevantPaths) processCSV(fileName string) error {
	records, parseErrs, err := readCSVFile(p.absoluteInputPath(fileName))
	if err != nil {
		return err
	}

	if err = writeRecordsToJSON(p.absoluteOutputPath(fileName), records); err != nil {
		return err
	}

	if err = writeErrorsToFileAsCSV(p.absoluteErrorPath(fileName), parseErrs); err != nil {
		return err
	}
	if err := deleteCSVAfterProcessing(p.absoluteInputPath(fileName)); err != nil {
		return err
	}
	return nil
}

func (p *RelevantPaths) absoluteInputPath(filename string) string {
	return filepath.Join(p.inputDirPath, filepath.Base(filename))
}

func (p *RelevantPaths) absoluteOutputPath(filename string) string {
	filename = strings.TrimSuffix(filename, ".csv") + ".json"
	return filepath.Join(p.outputDirPath, filepath.Base(filename))
}

func (p *RelevantPaths) absoluteErrorPath(filename string) string {
	return filepath.Join(p.errorDirPath, filepath.Base(filename))
}
