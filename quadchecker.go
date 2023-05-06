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


func GetArgs(a[]string) map[string]int{
  /*
    Itere tous les arguments afin d'y
    recuperer les noms de fichier et les
    valeurs
  */
  mapQuadInformation := map[string]int{}
  isLastWasInt := false
  key := ""
  for _, i := range a{
    if IsInt(i){
      if isLastWasInt{
        isLastWasInt = false
      }
      key += string(i)
    } else {
      isLastWasInt = true
      
    }
    fmt.Println(i)
  }
  return mapQuadInformation
}

func IsInt(s string) bool{
  /*
    Appelé pour vérifier si
    la chaine de caractere
    contient des nombres.
    Si slle ne contient pas
    de nombres, c'est considere
    comme un nom de fichier
  */
  isInt := true
  for _,c := range s{
    if c < '0' || c > '9'{
      isInt = false
    }
  }
  return isInt
}

func quadchecker(){
  /*
    Exercice function
  */
}