package main

import "os"

//callback for walk needs to match the following:
//type WalkFunc func(path string, info os.FileInfo, err error) error
func readFile (path string, fi os.FileInfo, err error) error {   
   fp := openFile(path)
   switch mode := fi.Mode(); {
   case mode.IsRegular():
      logFileMessage("INFO: '%s' being processed.", fi.Name())
      content, err := getFileContent(fp, fi)
      if err != nil {
         logFileMessage("INFO: '%s' cannot be handled by Tika.", fi.Name())
      } else {
         edat := getEntityData(content, fi.Name()) 
         collateEntities(edat)         
      }
   case mode.IsDir():
      logFileMessage("INFO: '%s' is a directory.", fi.Name())         
   default: 
      logFileMessage("INFO: '%s' is something completely different.", fi.Name())   
   }
   return nil
}

