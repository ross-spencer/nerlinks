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

func init() {
   flag.StringVar(&file, "file", "false", "File to extract information from.")
   flag.BoolVar(&vers, "version", false, "[Optional] Output version of the tool.")
}

func openFile (path string) *os.File {
   fp, err := os.Open(path)
   if err != nil {
      fmt.Fprintln(os.Stderr, "ERROR:", err)
      os.Exit(1)     //should only exit if root is null, consider no-exit
   }
   return fp
}

func processall(file string) {
   //check the services are up and running
   findOpenConnections()

   //time how long it takes to prcess files and extract entities
   start := time.Now()

   //read each file into each server and collect results
   //slow, works on server requests
   filepath.Walk(file, readFile)

   //output that time...
   elapsed := time.Since(start)
   fmt.Printf("\nNamed entity analysis took %s\n", elapsed)

   //collect all reaults together to enable interaction
   //fast, works on slices
   allentityhandler()

   //user interaction...
   responsehandler()
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
