package main

//list to store the structs
var all_list []EntityData

type entitycat struct {
   etype string
   evalue string
   ecount int
}

var categories []entitycat

func collateEntities(edat []EntityData) {
   logIntMessage("Collating %d statistics.", len(edat))
   for _, v := range edat {
      all_list = ExtendEntitySlice(all_list, v)
   }
}

func allentityhandler() {
   for _, v1 := range all_list {
      loop := true
      var e1 entitycat
      if len(categories) > 0 {
         for _, c := range categories {
            if v1.etype == c.etype && v1.evalue == c.evalue {
               loop = false
            }
         }
      }
      if loop == true {
         e1.etype = v1.etype
         e1.evalue = v1.evalue
         for _, v2 := range all_list {
            if e1.evalue == v2.evalue && e1.etype == v2.etype {
               e1.ecount = e1.ecount +1
            }
         }
         categories = ExtendCategorySlice(categories, e1)
      }
   }
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
