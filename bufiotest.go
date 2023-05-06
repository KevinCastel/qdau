package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "./quadA.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erreur lors de l'exécution:", err)
		fmt.Println("Sortie d'erreur:")
		fmt.Println(string(output))
		return
	}

	/*
		La partie chatGPT en haut
		Ma partie en bas
	*/

	lineLength := -1
	countLine := 0
	isLineLengthFound := false

	for _, char := range output {
		if !isLineLengthFound {
			lineLength++
		}
		if char == '\n' {
			isLineLengthFound = true
			countLine++
		}
	}
	fmt.Println("longueur de la lignes :", lineLength)
	fmt.Println("Nombre de lignes:", countLine)
}
