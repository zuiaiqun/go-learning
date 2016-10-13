package test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func MarshalCensorWords(inFile, outFile string) {
	unifile, erri := os.Open(inFile)
	if erri != nil {
		fmt.Printf("open %s err", inFile)
		return
	}
	defer unifile.Close()
	ofile, erro := os.Create(outFile)
	if erro != nil {
		return
	}
	defer ofile.Close()
	wordsArray := make([]string, 0, 10000)
	uniLineReader := bufio.NewReaderSize(unifile, 400)
	line, _, bufErr := uniLineReader.ReadLine()
	for nil == bufErr {
		wordsArray = append(wordsArray, string(line))
		line, _, bufErr = uniLineReader.ReadLine()
	}
	byteArray, err := json.Marshal(wordsArray)
	if err != nil {
		fmt.Println("err")
		return
	}
	ofile.Write(byteArray)
	fmt.Println("ok")
}
