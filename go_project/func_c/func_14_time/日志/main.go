package main

import (
	"fmt"
	"log"
	"os"
)

func logsOutput() {
	for i := 0; i < 10; i++ {
		log.Println("this is logs")
	}

}

func outputFile() {
	logFile, err := os.OpenFile("./access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(logFile)

}

func main() {
	logsOutput()
}
