package main

import (
   "os"
   "fmt"
   "bufio"
   "strings"
)

func responsehandler() {

   var input = true
   for input {
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Enter text: ")
      inputstr, _ := reader.ReadString('\n')
      input = checkquit(inputstr)
   }
}

func checkquit(inputstr string) bool {
   inputstr = strings.Replace(inputstr, "\n", "", -1)
    if inputstr == "false" || inputstr == "quit" || inputstr == "q" {
      return false
   }
   return true
}
