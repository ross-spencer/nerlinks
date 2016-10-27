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

func findOpenConnections() {
   var tika string = "http://127.0.0.1:9998/"
   var resp int8

   resp = testConnection(tika)
   if resp == CONN_BAD {
      fmt.Fprintln(os.Stdout, "INFO: Tika connection not available to connect to. Check localhost:9998.")
      os.Exit(1)
   }

   var ner string = "http://127.0.0.1:9000/"
   resp = testConnection(ner)
   if resp == CONN_BAD {
      fmt.Fprintln(os.Stdout, "INFO: Stanford NER connection not available to connect to. Check localhost:9000.")
      os.Exit(1)
   }   
}

func openFile (path string) *os.File {
   fp, err := os.Open(path)
   if err != nil {
      fmt.Fprintln(os.Stderr, "ERROR:", err)
      os.Exit(1)     //should only exit if root is null, consider no-exit
   }
   return fp
}

//callback for walk needs to match the following:
//type WalkFunc func(path string, info os.FileInfo, err error) error
func readFile (path string, fi os.FileInfo, err error) error {   
   fp := openFile(path)
   switch mode := fi.Mode(); {
   case mode.IsRegular():

      fmt.Println()
      fmt.Println(fi.Name())
      fmt.Println()


      content, err := getFileContent(fp, fi)
      if err != nil {
         return err
      }     
      getEntityData(content, fi.Name()) 
   case mode.IsDir():
      fmt.Fprintln(os.Stderr, "INFO:", fi.Name(), "is a directory.")      
   default: 
      fmt.Fprintln(os.Stderr, "INFO: Something completely different.")
   }
   return nil
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
