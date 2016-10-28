package main

import (
   "os"
   "fmt"
   "bufio"
   "strings"
   "strconv"
)

func responsehandler() {
   var input = true
   var inputstr string
   for input {
      displaycatoptions()
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Enter Option: ")
      var cat = false
      for !cat {
         inputstr, _ := reader.ReadString('\n')
         cat = checkcat(inputstr) 
      } 
      checkquit(inputstr)
   }
}

func displaycatoptions() {
   //newline before input choices
   fmt.Println()
   fmt.Println("Values extracted from documents (option no. value, count): \n")
   cols := 3
   for _, x := range categories {
      fmt.Printf("%2d) %25s (%d)     ", x.index, x.evalue, x.ecount)
      if x.index % cols == 0 {
         fmt.Print("\n")
      } 
   }
   //two newline before new input
   fmt.Println("\n")
}

func checkcat(inputstr string) bool {
   checkquit(inputstr) 
   //else...
   inputstr = strings.Replace(inputstr, "\n", "", -1)
   i, err := strconv.Atoi(inputstr)
   if err != nil {
      return false
   }

   found := false

   for _, x := range categories {
      if x.index == i {
         fmt.Println("\nFiles listing this term:")
         typeout := false
         termout := false
         for _, y := range all_list {
            if x.evalue == y.evalue && x.etype == y.etype {
               if typeout == false && termout == false {
                  typeout = true
                  termout = true
                  fmt.Printf("Type:  %s\n", y.etype)
                  fmt.Printf("Value: %s\n", y.evalue)
                  fmt.Println("---")
               }
               fmt.Printf("File name: %20s     Term count: %d\n", y.efile.fname, y.efile.ecount)
               found = true
            }
         }
      }
   }

   if !found {
      fmt.Print("Option not found, enter another option:")
      return false
   }

   fmt.Println("---")
   fmt.Println()
   return checkyesno()
}

func checkquit(inputstr string) {
   inputstr = strings.Replace(inputstr, "\n", "", -1)
    if inputstr == "false" || inputstr == "quit" || inputstr == "q" {
      os.Exit(0)
   }
}

func checkyesno() bool {
   var yesno = false
   reader := bufio.NewReader(os.Stdin)
   fmt.Print("Look for another value (y/n):")
   for !yesno {
      inputstr, _ := reader.ReadString('\n')
      inputstr = strings.Replace(inputstr, "\n", "", -1)
      checkquit(inputstr)
      if inputstr == "y" {
         return true
      }
      if inputstr == "n" {
         os.Exit(0)
      }
   }
   return false
}
