package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  log4batch [-d] DEBUG|INFO|WARN|ERROR Text to print and log.")
	fmt.Println("  Option -d: prints message with date and time.")
	fmt.Println("  export 'log4batch_debug' to print debug-messages.")
	os.Exit(0)
}

func main() {
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warningLog := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	debugLog := log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime)
	dPart := flag.String("d", "INFO|ERROR|WARNING|DEBUG", "Message-type.")
	help := flag.Bool("h", false, "print a little usage.")
	flag.Parse()

	if len(os.Args) == 1 {
		infoLog.Println("")
		os.Exit(0)
	}

	if *help {
		usage()
	}

	MESSAGE := os.Args[1:]
	myMessage := strings.Join(MESSAGE[1:], " ")
	onlyMessage := strings.Join(MESSAGE[0:], " ")
	if *dPart == "INFO" {
		newMessage, _ := strings.CutPrefix(myMessage, "INFO")
		infoLog.Println(newMessage)
	} else if *dPart == "DEBUG" {
		log4batch_debug := os.Getenv("log4batch_debug")
		if log4batch_debug != "" {
			newMessage, _ := strings.CutPrefix(myMessage, "DEBUG")
			debugLog.Println(newMessage)
		}
	} else if *dPart == "WARN" {
		newMessage, _ := strings.CutPrefix(myMessage, "WARN")
		warningLog.Println(newMessage)
	} else if *dPart == "ERROR" {
		newMessage, _ := strings.CutPrefix(myMessage, "ERROR")
		errorLog.Println(newMessage)
	} else {
		//newMessage, _ := strings.CutPrefix(myMessage, "INFO")
		fmt.Println(onlyMessage)
	}
}
