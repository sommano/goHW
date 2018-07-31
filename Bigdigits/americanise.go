package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	inFilename, outFilename, err := filenamesFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inFile, outFile := os.Stdin, os.Stdout
	if inFilename != "" {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}

	if outFilename != "" {
		if outFile, err = os.Create(outFilename); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}
	if err = americanise(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}
