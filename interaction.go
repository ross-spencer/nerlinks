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
   displaycategories()
   for input {
      val, etype := checktype()
      if val == true {
         input = false
         displayvalues(etype)
         for val {
            more := checkvalue()
            if more == true {
               val = false
               responsehandler()  
            }
         }
      }
   }
}

func displaycategories() {
   fmt.Println()
   fmt.Println("Please choose an entity category (category, count): \r\n")
   for k, x := range ALL_ENTITIES {
      count := 0
      for _, y := range categories {
         if x == y.etype {
            count = count + 1
         }
      }  
      fmt.Printf("%d) %s: %d\r\n", k, x, count)
   }
}

func checktype() (bool, string) {
   var input = true
   for input {

      reader := bufio.NewReader(os.Stdin)
      fmt.Print("\r\nPlease enter entity category value: ")
      inputstr, _ := reader.ReadString('\n')

      checkquit(inputstr) 
      //else...
      inputstr = strings.Replace(inputstr, "\r", "", -1)
      inputstr = strings.Replace(inputstr, "\n", "", -1)
      i, _ := strconv.Atoi(inputstr)
      if i < len(ALL_ENTITIES) {
         input = false
         return true, ALL_ENTITIES[i]
      } else {
         fmt.Print("\r\nOption not found. Please try again.")
         checktype()
      }
   }
   return false, ""
}

func displayvalues(etype string) {
   //newline before input choices
   fmt.Println()
   fmt.Println("Values extracted from documents (option no. value, filecount): \r\n")
   cols := 3

   intpad := fmt.Sprintf("%d", len(strconv.Itoa(len(categories))))
   intpad = "%" + intpad + "|d) "
   
   colcount := 0
   for _, x := range categories {
      if x.etype == etype {
         colcount = colcount + 1
         fmt.Printf(intpad + "%30s c.[%3d]   ", x.index, x.evalue, x.ecount)
         if colcount >= cols {
            colcount = 0
            fmt.Print("\r\n")
         } 
      }
   }
   //two newline before new input
   fmt.Println("\r\n")
}

func checkvalue() bool {

   var input = true
   for input {

      reader := bufio.NewReader(os.Stdin)
      fmt.Print("\r\nEnter Option: ")
      inputstr, _ := reader.ReadString('\n')

      checkquit(inputstr) 
      //else...
      inputstr = strings.Replace(inputstr, "\r", "", -1)
      inputstr = strings.Replace(inputstr, "\n", "", -1)
      i, _ := strconv.Atoi(inputstr)

      found := false

      start := time.Now()
      for _, x := range categories {
         if x.index == i {
            fmt.Println("\r\nFiles listing this term:")
            typeout := false
            termout := false
            for _, y := range all_list {
               if x.evalue == y.evalue && x.etype == y.etype {
                  var padlen = 0
                  if typeout == false && termout == false {
                     typeout = true
                     termout = true
                     fmt.Printf("Type:  %s\r\n", y.etype)
                     if padlen < len(y.efile.fname) {
                        padlen = len(y.efile.fname)
                     }
                     fmt.Printf("Value: %s\r\n", y.evalue)
                     fmt.Println("---")
                  }
                  //todo make filename output more dynamic...
                  fmt.Printf("File name: %45s   Term count: %d\r\n", y.efile.fname, y.efile.ecount)
                  found = true
               }
            }
         }
      }

      if !found {
         fmt.Print("Option not found, enter another option.\r\n")
         checkvalue()
      }

      fmt.Println("---")
      elapsed := time.Since(start)
      fmt.Printf("Catalogue query took %s\r\n", elapsed)
      fmt.Println()
      input = false
   }
   return checkyesno()
}

func checkquit(inputstr string) {
   inputstr = strings.Replace(inputstr, "\n", "", -1)
   inputstr = strings.Replace(inputstr, "\r", "", -1)   
   fmt.Println(inputstr)
    if inputstr == "false" || inputstr == "quit" || inputstr == "q" || inputstr == "n" {
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
      inputstr = strings.Replace(inputstr, "\r", "", -1)
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
