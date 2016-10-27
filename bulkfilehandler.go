package main 

import (
  	"os"
)

func getFileContent(fp *os.File, fi os.FileInfo) string {
   getTikaRecursive(fi.Name(), fp, ACCEPT_MIME_JSON)
   return fl_recursive_keys_values[TIKA_PLAIN_TEXT].(string)
}

func getEntityData(content string, fname string) []EntityData {
   return getNERData(content, fname)
}

