package main 

import (
  	"os"
)

func getFileContent(fp *os.File, fi os.FileInfo) string {
   getTikaRecursive(fi.Name(), fp, ACCEPT_MIME_JSON)
   content := fl_recursive_keys_values[TIKA_PLAIN_TEXT]
   return content.(string)
}

func getEntityData(content string, fname string) string {
   getNERData(content, fname)
   return "abc"
}

