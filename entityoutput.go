package main

import (
   "fmt"
)

//list to store the structs
var all_list []EntityData

func collateEntities(edat []EntityData) {
   fmt.Println("Collating statistics.")
   for _, v := range edat {
      all_list = ExtendEntitySlice(all_list, v)
   }
}

func allentityhandler() {
   fmt.Println()
   fmt.Println("rrrrrrrrrrrrrrr", len(all_list))
   fmt.Println()
   fmt.Println()
   for _, v := range all_list {
      fmt.Println(v.efile.fname)
   }
   
}
