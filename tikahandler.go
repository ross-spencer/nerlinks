package main

import (
  	"os"
	"fmt"
   "strings"
	"encoding/json"
)

//fl: denotes file level metadata keys
var fl_available_md_keys []string
var fl_keys_values map[string]interface{}

var fl_recursive_md_keys []string
var fl_recursive_keys_values map[string]interface{}

func getTikaId (fp *os.File) {
   resp := makeConnection(PUT, tika_path_detect, fp, "")
	fmt.Fprintln(os.Stdout, "RESPONSE:", resp)
}

func getTikaMetadataPUT (fp *os.File, accepttype string) string {
   fp.Seek(0,0)
   resp := makeConnection(PUT, tika_path_meta, fp, accepttype)
	return resp
}

func getTikaMetadataPOST (fname string, fp *os.File, accepttype string) {
   fp.Seek(0,0)
   resp := makeMultipartConnection(POST, tika_path_meta_form, fp, fname, accepttype)
   readTikaMetadataJson(resp, "", &fl_keys_values, &fl_available_md_keys)
}

func getTikaRecursive (fname string, fp *os.File, accepttype string) {
   fp.Seek(0,0)
   resp := makeMultipartConnection(POST, tika_path_meta_recursive, fp, fname, accepttype) 
   trimmed := strings.Trim(resp, "[ ]")
   readTikaMetadataJson(trimmed, "", &fl_recursive_keys_values, &fl_recursive_md_keys)
}

func readTikaMetadataJson (output string, key string, kv *map[string]interface{}, mdkeys *[]string) {

   //we can get multiple JSON sets from TIKA
   json_strings := strings.Split(output, "},")

   for k, v := range json_strings {
      last := v[len(v)-1:]
      if last != "}" {
         json_strings[k] = v + "}"
      }
   }

	var tikamap map[string]interface{}

   for _, v := range json_strings {
	   if err := json.Unmarshal([]byte(v), &tikamap); err != nil {
		   fmt.Fprintln(os.Stderr, "ERROR: Handling TIKA JSON,", err)
	   }
	   *kv = tikamap
	   getTikaKeys(tikamap, mdkeys) 
   }
} 

func getTikaKeys (tikamap map[string]interface{}, mdkeys *[]string) {	
	keys := make([]string, len(tikamap))
	i := 0
	for k := range tikamap {
		keys[i] = k
		i++
	}
	*mdkeys = keys    //alt: /meta/{field} TIKA URL
}

