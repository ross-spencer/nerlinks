package main 

import (
   "log"
)

const JSON_UNEXPECTED_END = "unexpected end of JSON input"

func logFileMessage(message string, filename string) {
   log.Printf(message + "\n", filename)
}

func logNERError(e error, s string) {
   JSON_PROCESSING_ERR := "Ignoring Line ERROR: Handling NER JSON:"   
   log.Println(JSON_PROCESSING_ERR, e, "x")
   log.Println(s)
}

func logStringError(message string, e error) {
   log.Printf(message + "\n", e)
}

func logIntMessage(message string, count int) {
   log.Printf(message + "\n", count)
}   

func logBasicMessage(message string) {
   log.Printf(message + "\n")
}