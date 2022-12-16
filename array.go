package main

import "fmt"

func main(){

  var names [3]string

  names[0] = "SBY"
  names[1] = "JK"
  names[2] = "BDG"
  
  fmt.Println(names[0])
  fmt.Println(names[1])
  

  var values = [3]string{
    "JKT",
    "SBY",
    "BDG",
  }

  fmt.Println(values)
  fmt.Println(values[0])
}
