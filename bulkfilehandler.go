package main 

import (
  	"os"
   "fmt"
)

type contenterror struct {
   content string
   err error
   fname string
}  

func extractAndAnalyse(filepool []filedata) (bool, error) {
   
   //create a temporary list to buffer results coming out of NER...
   var tmp_list []EntityData

   //make channel run goroutine...
   ch := make(chan contenterror)
   for _, fi := range filepool {
      go getFileContent(fi, ch)
   }
   for range filepool{
      ce := <- ch
      if ce.err != nil {
         logFileMessage("INFO: '%s' cannot be handled by Tika.", ce.fname)
      } else {
         edat := getEntityData(ce.content, ce.fname) 
         tmp := collateEntities(edat)        
         //store our tmp value in the tmp_list...
         tmp_list = append(tmp_list, tmp...)          
      }
   }
   //once computations are complete add to all_list for global scope...
   all_list = append(all_list, tmp_list...)
   return false, nil
}

func openFile (path string) (*os.File, error) {
   fp, err := os.Open(path)
   if err != nil {
      return nil, err
   }
   return fp, nil
}

func getFileContent(fi filedata, ch chan contenterror) {
   //create empty struct to return...
   var ce contenterror
   ce.fname = fi.fname

   //what are we doing..?
   logFileMessage("INFO: '%s' being processed.", fi.fname)

   //process...
   fp, err := openFile(fi.fpath)
   defer fp.Close()
   if err != nil {
      ce.err = err
      ch <- ce
      return
   }

   _, fl_recursive_keys_values, err := getTikaRecursive(fi.fname, fp, ACCEPT_MIME_JSON)
   if err != nil {
      ce.err = err
      ch <- ce
      return
   }

   if val, ok := fl_recursive_keys_values[TIKA_PLAIN_TEXT]; ok {
      ce.content = val.(string)
      ch <- ce
      return
   } else {
      ce.err = fmt.Errorf("No plain text data to analyse via NER")
      ch <- ce
      return 
   }
}

func getEntityData(content string, fname string) []EntityData {
   return getNERData(content, fname)
}

