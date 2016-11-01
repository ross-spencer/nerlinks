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

   displaycategories()

   val, etype := checktype()
   fmt.Println("\nCategory selected:", etype)

   if val == true {
      displayvalues(etype)
      checkvalue()
   }

   if checkyesno() {
      responsehandler()
   } else {
      os.Exit(0)
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
      fmt.Printf("%d) %s: %d\r\n", k+1, x, count)
   }
}

func checktype() (bool, string) {
   fmt.Print("Please enter entity category value: ")
   reader := bufio.NewReader(os.Stdin)
   inputstr, err := reader.ReadString('\n')
   if err != nil {
      return checktype()
   }
   checkquit(inputstr)

   //else...
   inputstr = strings.Replace(inputstr, "\r", "", -1)
   inputstr = strings.Replace(inputstr, "\n", "", -1)

   i, err := strconv.Atoi(inputstr)
   if err != nil {
      return checktype()
   } else {
      //fix skew introduced for display purposes...
      i = i-1
   }
         
   if i < len(ALL_ENTITIES) {
      return true, ALL_ENTITIES[i]
   } else {
      fmt.Print("\r\nOption not found. Please try again.")
      return checktype()
   }
   return false, ""
}

func displayvalues(etype string) {
   //newline before input choices
   fmt.Println()
   fmt.Println("Values extracted from documents (option no. value, filecount): \r\n")

   cols := 3
   paging := (20 * cols)

   intpad := fmt.Sprintf("%d", len(strconv.Itoa(len(categories))))
   intpad = "(%" + intpad + "d) "
   
   colcount := 0
   pagecount := 0

   for _, x := range categories {
      if x.etype == etype {
         colcount = colcount + 1
         pagecount = pagecount + 1
         fmt.Printf(intpad + "%30s c.[%3d]   ", x.index, x.evalue, x.ecount)
         if colcount >= cols {
            colcount = 0
            fmt.Print("\r\n")
         } 
         if pagecount == paging {
            pagecount = checkpagecount()
         }
      }
   }
   //two newline before new input
   fmt.Println("\r\n")
}

func checkpagecount() int {
   fmt.Print("\r\nPress enter to page: ")
   reader := bufio.NewReader(os.Stdin)
   inputstr, err := reader.ReadString('\n')
   if err != nil {
      checkpagecount()
   }
   checkquit(inputstr) 
   fmt.Print("\r\n")
   return 0
}


func checkvalue() bool {
   fmt.Print("Enter Option: ")
   reader := bufio.NewReader(os.Stdin)
   inputstr, err := reader.ReadString('\n')
   if err != nil {
      return checkvalue()
   }
   checkquit(inputstr) 
   inputstr = strings.Replace(inputstr, "\r", "", -1)
   inputstr = strings.Replace(inputstr, "\n", "", -1)
   i, err := strconv.Atoi(inputstr)
   if err != nil {
      return checkvalue()
   }      
   res := getresult(i)
   if res {
      return true
   }
   return true
}

func getresult(i int) bool {
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
      fmt.Print("Entity entry not found, enter another option.\r\n")
      return false
   } else {
      fmt.Println("---")
      elapsed := time.Since(start)
      fmt.Printf("Catalogue query took %s\r\n", elapsed)
      fmt.Println()
   }
   return true
}

func checkquit(inputstr string) {
   inputstr = strings.Replace(inputstr, "\n", "", -1)
   inputstr = strings.Replace(inputstr, "\r", "", -1)   
   if inputstr == "false" || inputstr == "quit" || inputstr == "q" || inputstr == "n" {
      os.Exit(0)
   }
}

func checkyesno() bool {
   fmt.Print("Look for another value (y/n): ")
   reader := bufio.NewReader(os.Stdin)
   inputstr, err := reader.ReadString('\n')
   if err != nil {
      checkyesno()
   }
   inputstr = strings.Replace(inputstr, "\n", "", -1)
   inputstr = strings.Replace(inputstr, "\r", "", -1)
   checkquit(inputstr)
   if inputstr == "y" {
      return true
   }
   if inputstr == "n" {
      os.Exit(0)     //todo: send back up to main control look instead
   }
   return checkyesno()
}
