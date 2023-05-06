//go 1.10.4

package main

import ("fmt"
       "os")


func main() {
  args := os.Args
  if len(args) >= 2{
    m := GetArgs(args[1:])
    fmt.Println(m)
  } else {
    fmt.Println("Not a quad function")
  }
}


func GetArgs(a[]string) map[string]string{
  mapQuadInformation := map[string]string{}
  for _, i := range a{
    fmt.Println(i)
    
  }
  return mapQuadInformation
}

func IsInt(s string) bool{
  var b bool
  
  return b
}



func quadchecker(){
  /*
    Exercice function
  */
}