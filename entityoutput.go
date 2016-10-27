package main

import (
   "fmt"
)

//list to store the structs
var all_list []EntityData

func collateEntities(edat []EntityData) {
   fmt.Println("Collating statistics.", len(edat))
   for _, v := range edat {
      all_list = ExtendEntitySlice(all_list, v)
   }
}

func allentityhandler() {
   for _, v := range all_list {
      fmt.Println(v.efile.fname, v.etype, v.evalue, v.efile.ecount)
   }
   
}
