package main

import "strings"

//list to store the structs
var all_list []EntityData

type entitycat struct {
   index    int
   etype    string
   evalue   string
   ecount   int
}

var categories []entitycat

func collateEntities(edat []EntityData) []EntityData {

   var tmp_list []EntityData
   logIntMessage("Collating %d statistics.", len(edat))
   for _, v := range edat {
      tmp_list = ExtendEntitySlice(tmp_list, v)
   }
   return tmp_list
}

func allentityhandler() {
   index := 0
   for _, v1 := range all_list {
      loop := true
      var e1 entitycat
      if len(categories) > 0 {
         //if we've got the value we don't need to add it again
         for _, c := range categories {
            if entitycompare(c, v1) {
               loop = false
            }
         }
      }
      if loop == true {
         index = index + 1
         e1.index = index
         e1.etype = v1.etype
         e1.evalue = v1.evalue
         for _, v2 := range all_list {
            if (entitycompare(e1, v2)) {
               e1.ecount = e1.ecount +1
            }
         }
         categories = ExtendCategorySlice(categories, e1)
      }
   }   
}

func entitycompare(v1 entitycat, v2 EntityData) bool {
   return valuetypecompare(v1.evalue, v2.evalue, v1.etype, v2.etype)
}

func valuetypecompare(value1 string, value2 string, type1 string, type2 string) bool {
   value1 = strings.ToLower(value1)
   value2 = strings.ToLower(value2)
   if value1 == value2 && type1 == type2 {
      return true
   }
   return false
}

func ExtendCategorySlice(slice []entitycat, element entitycat) []entitycat {
   n := len(slice)
   if n == cap(slice) {
      // Slice is full; must grow.kb
      // We double its size and add 1, so if the size is zero we still grow.
      newSlice := make([]entitycat, len(slice), 2*len(slice)+1)
      copy(newSlice, slice)
      slice = newSlice
   }
   slice = slice[0 : n+1]
   slice[n] = element
   return slice
}

//delete entity from the slice so we can update dynamically
func deleteCatEntity(list []entitycat, index int) []entitycat {
   list[index] = list[len(list)-1] 
   list = list[:len(list)-1]
   return list
}
