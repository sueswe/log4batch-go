package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warningLog := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	dPart := flag.String("d", "INFO|ERROR|WARNING", "Message-type.")
	flag.Parse()

	if len(os.Args) == 1 {
		infoLog.Println("")
		os.Exit(0)
	}

	MESSAGE := os.Args[1:]
	myMessage := strings.Join(MESSAGE[1:], " ")

	if *dPart == "INFO" {
		newMessage, _ := strings.CutPrefix(myMessage, "INFO")
		infoLog.Println(newMessage)
	} else if *dPart == "WARN" {
		newMessage, _ := strings.CutPrefix(myMessage, "WARN")
		warningLog.Println(newMessage)
	} else if *dPart == "ERROR" {
		newMessage, _ := strings.CutPrefix(myMessage, "ERROR")
		errorLog.Println(newMessage)
	} else {
		//newMessage, _ := strings.CutPrefix(myMessage, "INFO")
		fmt.Println(MESSAGE)
	}
}
