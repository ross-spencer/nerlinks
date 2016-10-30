package main

import (
  	"os"
	"fmt"
   "flag"
   "time"
   "path/filepath"
)

var file string
var vers bool

var fthrottle = 4

func init() {
   flag.StringVar(&file, "file", "false", "File to extract information from.")
   flag.BoolVar(&vers, "version", false, "[Optional] Output version of the tool.")
}

func processall(file string) {
   //check the services are up and running
   findOpenConnections()

   //make a listing of all the files we're going to process
   //efficient enough with memory?
   err := filepath.Walk(file, readFile)
   if err != nil {
      logStringError("%v", err)
      os.Exit(1)
   }

   //time how long it takes to prcess files and extract entities
   start := time.Now()

   //read each file into each server and collect results

   if len(allfiles) <= 0 {
      fmt.Fprintf(os.Stderr, "No files to process.\n")   
      os.Exit(1)
   }

   for x := 0; x < len(allfiles); x+=fthrottle {
      remain := min(x+fthrottle,len(allfiles))
      filepool := allfiles[x:remain]
      b := true
      for b {
         b, err = extractAndAnalyse(filepool)
         if err != nil {
            logStringError("%v", err)
            os.Exit(1)
         }
      }
   }   
   
   //output that time...
   elapsed := time.Since(start)
   fmt.Printf("\nNamed entity analysis took %s\n", elapsed)

   //collect all reaults together to enable interaction
   //fast, works on slices
   //allentityhandler()

   //user interaction...
   //responsehandler()
}

func main() {
   flag.Parse()
   if flag.NFlag() <= 0 {    // can access args w/ len(os.Args[1:]) too
      fmt.Fprintln(os.Stderr, "Usage:  links [-file ...]")
      fmt.Fprintln(os.Stderr, "               [Optional -version]")
      fmt.Fprintln(os.Stderr, "Output: [TBD]")
      flag.Usage()
      os.Exit(0)
   }
   if vers {
      fmt.Fprintln(os.Stdout, getVersion())
      os.Exit(1)
   }
   processall(file)
}

//math.Min uses float64, so let's not cast
func min(a, b int) int {
   if a < b { return a }
   return b
}
