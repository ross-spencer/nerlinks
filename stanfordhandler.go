package main 

import (
   "os"
   "fmt"
   "strings"
   "encoding/json"
)

//consts for the NER bits and pieces we want
const NER_PATTERN_TYPE = "ner"
const NER_PERSON   = "PERSON"
const NER_LOC      = "LOCATION"
const NER_DATE     = "DATE"
const NER_TIME     = "TIME"
const NER_ORD      = "ORDINAL"
const NER_MISC     = "MISC"

//entities we care about...
var ALL_ENTITIES = [...]string{NER_PERSON, NER_LOC, NER_DATE, NER_TIME, NER_ORD, NER_MISC}

//comparison values we want
const NER_CHAR_OFF_END = "characterOffsetEnd"
const NER_CHAR_OFF_BEGIN = "characterOffsetBegin"
const NER_TEXT_VALUE = "originalText"
const NER_INDEX_VALUE = "index"

//interfaces to store JSON results
var all_ner_values []map[string]interface{}
var ner_keys_values map[string]interface{}

//getNERData is used to build the data we're going to use per
//file given the system. 
func getNERData (content string, fname string) {
   preprocessed := nerPreprocess(content)
   resp := makeNERDataConnection(POST, ner_extract, preprocessed)
   readNERJson(resp, "")
   groupNERData()
}

func nerPreprocess (content string) string {
   //newlines
   return strings.Replace(content, "\n", ", ", -1)
}

func groupNERData() {

   //we're grouping split values to make whole names
   var indexes []float64   //scope for this var is important...
   for k, _ := range all_ner_values {
      if k+1 < len(all_ner_values) {
         combine := true
         var value string
         var nertype string
         for combine == true {
            val1 := all_ner_values[k][NER_CHAR_OFF_END].(float64)
            val2 := all_ner_values[k+1][NER_CHAR_OFF_BEGIN].(float64)
            if (val1 + 1) == val2 {
               name1 := all_ner_values[k][NER_TEXT_VALUE].(string)
               name2 := all_ner_values[k+1][NER_TEXT_VALUE].(string)
               idx1 := all_ner_values[k][NER_INDEX_VALUE].(float64)
               idx2 := all_ner_values[k+1][NER_INDEX_VALUE].(float64)
               indexes = ExtendIntSlice(indexes, idx1)
               indexes = ExtendIntSlice(indexes, idx2)
               value = value + name1 + " " + name2
               nertype = all_ner_values[k][NER_PATTERN_TYPE].(string)
               k = k+1
            } else {
               if value != "" {
                  fmt.Println(nertype, value)
                  break
               } else {
                  idx := false
                  for _, v := range indexes {
                     if all_ner_values[k][NER_INDEX_VALUE] == v {
                        idx = true
                     }
                     if all_ner_values[k+1][NER_INDEX_VALUE] == v {
                        idx = true
                     }
                  }
                  if idx == false { 
                     fmt.Println(all_ner_values[k][NER_PATTERN_TYPE], all_ner_values[k][NER_TEXT_VALUE])
                  }
               }
               combine = false
            }
         }
      }
   }

}

//readNERJson processes the JSON output by the Standord NER
//extractor. 
func readNERJson (output string, key string) {

   //value to hold the keys we extract
   var mdkeys []string

   //TODO: improve this to be more precise...
   trimmed := strings.Replace(output, "{\"sentences\":[", "", 1)
   trimmed = strings.Replace(trimmed, "]}]}", "", 1)

   //we can get multiple JSON sets from Stanford NER
   json_strings := strings.Split(trimmed, "},")
   
   for k, v := range json_strings {
      last := v[len(v)-1:]
      if last != "}" {
         json_strings[k] = v + "}"
      }
   }
   
   for _, v := range json_strings {
      var nermap map[string]interface{}
      if err := json.Unmarshal([]byte(v), &nermap); err != nil {
         fmt.Fprintf(os.Stderr, "Ignoring Line ERROR: Handling NER JSON: %v \"%v\"\n", err, v)
         continue
      }
      ner := getNERKeys(nermap, &mdkeys, NER_PATTERN_TYPE) 
      if ner {
         for _, v := range ALL_ENTITIES {
            if nermap[NER_PATTERN_TYPE] == v {
               all_ner_values = Extend(all_ner_values, nermap)
               break
            } else {
               //we can take a look at what other values NER has picked up
            }
         }
      }
   }
} 

//getNERKeys filters out only elements that have a named entity recognition
//elment associated with it for utilisation within the tool.
func getNERKeys (nermap map[string]interface{}, mdkeys *[]string, needle string) bool {   
   found := false
   keys := make([]string, len(nermap))
   i := 0
   for k := range nermap {
      if k == needle { found = true }
      keys[i] = k
      i++
   }
   *mdkeys = keys 
   return found
}

//Extend allows us to arbitrarily extend slices containing named entity
//recognition information.
func Extend(slice []map[string]interface{}, element map[string]interface{}) ([]map[string]interface{}) {
   n := len(slice)
   if n == cap(slice) {
      // Slice is full; must grow.kb
      // We double its size and add 1, so if the size is zero we still grow.
      newSlice := make([]map[string]interface{}, len(slice), 2*len(slice)+1)
      copy(newSlice, slice)
      slice = newSlice
   }
   slice = slice[0 : n+1]
   slice[n] = element
   return slice
}

func ExtendIntSlice(slice []float64, element float64) []float64 {
   n := len(slice)
   if n == cap(slice) {
      // Slice is full; must grow.kb
      // We double its size and add 1, so if the size is zero we still grow.
      newSlice := make([]float64, len(slice), 2*len(slice)+1)
      copy(newSlice, slice)
      slice = newSlice
   }
   slice = slice[0 : n+1]
   slice[n] = element
   return slice
}


