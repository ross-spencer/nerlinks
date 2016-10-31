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
   fmt.Println("Please choose an entity category (category, count): \n")
   for k, x := range ALL_ENTITIES {
      count := 0
      for _, y := range categories {
         if x == y.etype {
            count = count + 1
         }
      }  
      fmt.Printf("%d) %s: %d\n", k, x, count)
   }
}

func checktype() (bool, string) {
   var input = true
   for input {

      reader := bufio.NewReader(os.Stdin)
      fmt.Print("\nPlease enter entity category value: ")
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
         fmt.Print("\nOption not found. Please try again.")
         checktype()
      }
   }
   return false, ""
}

func displayvalues(etype string) {
   //newline before input choices
   fmt.Println()
   fmt.Println("Values extracted from documents (option no. value, filecount): \n")
   cols := 3

   intpad := fmt.Sprintf("%d", len(strconv.Itoa(len(categories))))
   intpad = "%" + intpad + "d) "
   
   for _, x := range categories {
      if x.etype == etype {
         fmt.Printf(intpad + "%30s c.(%3d)   ", x.index, x.evalue, x.ecount)
         if x.index % cols == 0 {
            fmt.Print("\n")
         } 
      }
   }
   //two newline before new input
   fmt.Println("\n")
}

func checkvalue() bool {

   var input = true
   for input {

      reader := bufio.NewReader(os.Stdin)
      fmt.Print("\nEnter Option: ")
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
            fmt.Println("\nFiles listing this term:")
            typeout := false
            termout := false
            for _, y := range all_list {
               if x.evalue == y.evalue && x.etype == y.etype {
                  var padlen = 0
                  if typeout == false && termout == false {
                     typeout = true
                     termout = true
                     fmt.Printf("Type:  %s\n", y.etype)
                     if padlen < len(y.efile.fname) {
                        padlen = len(y.efile.fname)
                     }
                     fmt.Printf("Value: %s\n", y.evalue)
                     fmt.Println("---")
                  }
                  //todo make filename output more dynamic...
                  fmt.Printf("File name: %45s   Term count: %d\n", y.efile.fname, y.efile.ecount)
                  found = true
               }
            }
         }
      }

      if !found {
         fmt.Print("Option not found, enter another option.\n")
         checkvalue()
      }

      fmt.Println("---")
      elapsed := time.Since(start)
      fmt.Printf("Catalogue query took %s\n", elapsed)
      fmt.Println()
      input = false
   }
   return checkyesno()
}

func checkquit(inputstr string) {
   fmt.Println("quit")
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
