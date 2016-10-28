package main 

import (
   "log"
)

func logFileMessage(message string, filename string) {
   log.Printf(message + "\n", filename)
}

func logNERError(e error, s string) {
   JSON_PROCESSING_ERR := "Ignoring Line ERROR: Handling NER JSON:"   
   log.Println(JSON_PROCESSING_ERR, e)
   log.Println(s)
}

func logIntMessage(message string, count int) {
   log.Printf(message + "\n", count)
}   
