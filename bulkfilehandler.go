package main 

import (
  	"os"
   "fmt"
)

func getFileContent(fp *os.File, fi os.FileInfo) (string, error) {
   err := getTikaRecursive(fi.Name(), fp, ACCEPT_MIME_JSON)
   if err != nil {
      return "", err
   }
   content := ""
   if val, ok := fl_recursive_keys_values[TIKA_PLAIN_TEXT]; ok {
      content = val.(string)
   } else {
      return "", fmt.Errorf("No plain text data to analyse via NER")
   }
   return content, nil
}

func getEntityData(content string, fname string) []EntityData {
   return getNERData(content, fname)
}

