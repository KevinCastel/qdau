// go 1.10.4
package main

<<<<<<< HEAD
import ("fmt"
       "os"
       "strconv")

=======
import (
	"fmt"
	"os"
)
>>>>>>> ffbcd76 (Replacing for the arguments catch-up)

func main() {
	args := os.Args
	if len(args) >= 2 {
		m := GetArgs(args[1:])
		fmt.Println(m)
	} else {
		fmt.Println("Not a quad function")
	}
}

<<<<<<< HEAD

func GetArgs(a[]string) map[string]int{
  /*
    Itere tous les arguments afin d'y
    recuperer les noms de fichier et les
    valeurs
  */
  mapQuadInformation := map[string]int{}
  isLastWasInt := false
  key := ""
  intStr := ""
  for _, i := range a{
    if IsInt(i){
      if isLastWasInt{
        isLastWasInt = false
        fmt.Println("Ajout de la cle ", key)
        ass, err := strconv.Atoi(intStr)
        if err != nil{
          
        } else {
          fmt.Println("Une erreur a etait tdouve")
        }
        mapQuadInformation[key] = ass
        intStr = ""
        key = ""
      }
      key += string(i)
    } else {
      intStr += string(i)
      isLastWasInt = true
      
    }
    fmt.Println(i)
  }
  return mapQuadInformation
=======
func GetArgs(a []string) map[string][]int {
	/*
	   Itere tous les arguments afin d'y
	   recuperer les noms de fichier et les
	   valeurs
	*/
	mapQuadInformation := map[string][]int{}
	letters := ""
	numeric := ""
	for _, i := range a {
		for _, c := range []rune(i) {
			fmt.Println("c:'", string(c), "'")
			if c > 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
				letters += string(c)
			} else if c >= '0' && c <= '9' {
				numeric += string(c)
			}
		}
	}
	return mapQuadInformation
>>>>>>> ffbcd76 (Replacing for the arguments catch-up)
}

func IsInt(s string) bool {
	/*
	   Appelé pour vérifier si
	   la chaine de caractere
	   contient des nombres.
	   Si slle ne contient pas
	   de nombres, c'est considere
	   comme un nom de fichier
	*/
	isInt := true
	for _, c := range s {
		if c < '0' || c > '9' {
			isInt = false
		}
	}
	return isInt
}

func RemoveSpaceInString(s string) string {
	newString := ""
	a := []rune(s)
	for _, c := range a {
		if c >= '0' && c <= '9' {
			newString += string(c)
		}
	}
	return newString
}

func quadchecker() {
	/*
	   Exercice function
	*/
}
