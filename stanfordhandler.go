package main 

import (
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

//struct to return to other classes
type EntityData struct {
   etype       string
   evalue      string
   efile       FileData
}

type FileData struct {
   fname string
   ecount int
}

//list to store the structs
var entity_list []EntityData

//getNERData is used to build the data we're going to use per
//file given the system. 
func getNERData (content string, fname string) []EntityData {
   entity_list = nil
   all_ner_values = nil
   preprocessed := nerPreprocess(content)
   resp := makeNERDataConnection(POST, ner_extract, preprocessed)
   readNERJson(resp, "")
   groupNERData(fname)
   return entity_list
}

//pre-process lines of the output as some characters make it
//harder such as newlines where it interferes with word boundaries
func nerPreprocess (content string) string {
   //replace newlines
   return strings.Replace(content, "\n", ", ", -1)  //TODO: improve handling further down code
}

//group NER data for faceting... 
func groupNERData(fname string) {
   var indexes []float64
   var value string
   var name1 string
   var nertype string

   for k, _ := range all_ner_values {

      if k+1 < len(all_ner_values)-1 {
         value = ""
         nertype = ""
         combine := true
         

         for combine == true {
            val1 := all_ner_values[k][NER_CHAR_OFF_END].(float64)
            val2 := all_ner_values[k+1][NER_CHAR_OFF_BEGIN].(float64)
            name1 = all_ner_values[k][NER_TEXT_VALUE].(string)
            nertype = all_ner_values[k][NER_PATTERN_TYPE].(string)
           //name2 := all_ner_values[k+1][NER_TEXT_VALUE].(string)

            if (val1 + 1) == val2 {

               idx1 := all_ner_values[k][NER_INDEX_VALUE].(float64)
               idx2 := all_ner_values[k+1][NER_INDEX_VALUE].(float64)
               indexes = ExtendFloatSlice(indexes, idx1)
               indexes = ExtendFloatSlice(indexes, idx2)
               

               value = value + name1 + " "
               
               fmt.Println("yyy", value)

               if k+1 < len(all_ner_values)-1 {                  
                  k = k+1
               } else {
                  break
               }

            } else {

               if value != "" {
                  value = value + name1
                  addEntity(nertype, value, fname)
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
                  t1 := all_ner_values[k][NER_PATTERN_TYPE].(string)
                  v1 := all_ner_values[k][NER_TEXT_VALUE].(string)
                  if idx == false { 
                     addEntity(t1, v1, fname)
                  }
               }
               combine = false
            }
         }
      } else {
         name1 = all_ner_values[k][NER_TEXT_VALUE].(string)
      }
   }
   if value != "" {
      fmt.Println("ggg", value + name1, nertype)
      addEntity(value + name1, nertype, fname)
   }
}

//add entity to global entity list as we require
func addEntity (etype string, evalue string, fname string) {
   edata := initEntity(etype, evalue, fname)
   if len(entity_list) == 0 {
      entity_list = ExtendEntitySlice(entity_list, edata)
      return 
   } else {
      //see if this is a duplicate and handle accordingly
      for k, v := range entity_list {
         if etype == v.etype && evalue == v.evalue {
            //duplicate increment count
            edata = v
            edata.efile.ecount = edata.efile.ecount + 1
            entity_list = deleteEntity(entity_list, k)
            entity_list = ExtendEntitySlice(entity_list, edata)
            return
         } 
      }
      entity_list = ExtendEntitySlice(entity_list, edata) 
      return 
   }
}

//delete entity from the slice so we can update dynamically
func deleteEntity(list []EntityData, index int) []EntityData {
   list[index] = list[len(list)-1] 
   list = list[:len(list)-1]
   return list
}

//initialize an entity structure
func initEntity(etype string, evalue string, fname string) EntityData {
   var edata EntityData
   edata.etype = etype
   edata.evalue = evalue
   edata.efile.fname = fname
   edata.efile.ecount = 1
   return edata
}

//readNERJson processes the JSON output by the Standord NER
//extractor. 
func readNERJson (output string, key string) {
   //value to hold the keys we extract
   var mdkeys []string
   
   //process JSON and extract entities we care about
   json_strings, err := processNERJson(output)   
   if err != nil {
      return 
   }
   for _, v := range json_strings {
      if getNERKeys(v, &mdkeys, NER_PATTERN_TYPE) {
         for _, all := range ALL_ENTITIES {
            if v[NER_PATTERN_TYPE].(string) == all {
               all_ner_values = ExtendJSONSlice(all_ner_values, v)
            }
         }
      }
   }
} 

//JSON doesn't seem to work out of the box, so pre-process
func processNERJson(output string) ([]map[string]interface{}, error) {
   var wordjson []map[string]interface{}
   var nermap map[string]interface{}
   var err error

   output = fixupKnownErrors(output)

   err = json.Unmarshal([]byte(output), &nermap)
   if err != nil {
      logStringError("Issue unmarshalling JSON: %v", err)
      return wordjson, err
   }

   for _, n := range nermap {
      if rec1, ok1 := n.([]interface{}); ok1 {   
         for _, v1 := range rec1 {
            if rec2, ok2 := v1.(map[string]interface{}); ok2 {
               for k2, v2 := range rec2 { 
                  if k2 == "tokens" {
                     if rec3, ok3 := v2.([]interface{}); ok3 {
                        for _, v3 := range rec3 {                      
                           if rec4, ok4 := v3.(map[string]interface{}); ok4 {
                              wordjson = ExtendJSONSlice(wordjson, rec4)
                           }
                        }
                     }
                  }
               }
            }
         }
      }      
   }
   return wordjson, err
}

//Pre-process for known errors in JSON stream
func fixupKnownErrors(output string) string {
   //replace null characters seen in some streams
   out := strings.Replace(output, "\x00", " ", -1)
   return out
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
func ExtendJSONSlice(slice []map[string]interface{}, element map[string]interface{}) ([]map[string]interface{}) {
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

func ExtendFloatSlice(slice []float64, element float64) []float64 {
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

func ExtendEntitySlice(slice []EntityData, element EntityData) []EntityData {
   n := len(slice)
   if n == cap(slice) {
      // Slice is full; must grow.kb
      // We double its size and add 1, so if the size is zero we still grow.
      newSlice := make([]EntityData, len(slice), 2*len(slice)+1)
      copy(newSlice, slice)
      slice = newSlice
   }
   slice = slice[0 : n+1]
   slice[n] = element
   return slice
}


