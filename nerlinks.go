package main

import (
  	"os"
	"fmt"
   "flag"
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

   findOpenConnections()

   filepath.Walk(file, readFile)
}
