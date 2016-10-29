package main

import (
   "os"
   "fmt"
   "time"
   "bufio"
   "strings"
   "strconv"
)

func responsehandler() {
   var input = true
   var inputstr string
   for input {
      displaycatoptions()
      var cat = false
      for !cat {
         reader := bufio.NewReader(os.Stdin)
         fmt.Print("Enter Option: ")
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

   intpad := fmt.Sprintf("%d", len(strconv.Itoa(len(categories))))
   intpad = "%" + intpad + "d) "
   
   for _, x := range categories {
      fmt.Printf(intpad + "%30s (%d)     ", x.index, x.evalue, x.ecount)
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

   start := time.Now()
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
   elapsed := time.Since(start)
   fmt.Printf("Catalogue query took %s\n", elapsed)
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
   var yes = true
   reader := bufio.NewReader(os.Stdin)
   for yes {
      fmt.Print("Look for another value (y/n): ")
      inputstr, _ := reader.ReadString('\n')
      inputstr = strings.Replace(inputstr, "\n", "", -1)
      checkquit(inputstr)
      if inputstr == "y" {
         return true
      }
      if inputstr == "n" {
         yes = false
         os.Exit(0)     //todo: send back up to main control look instead
      }
   }
   return false
}
